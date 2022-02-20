# 프로세스

리눅스에서는 두 가지 목적으로 프로세스를 생성한다.

1. 같은 프로그램의 처리를 여러 개의 프로세스가 나눠서 처리 (`fork()`)
2. 전혀 다른 프로그램을 생성 (`execve()`)

## 프로세스 생성

### fork() 함수
같은 프로그램의 처리를 위해 여러 개의 프로세스가 나눠서 처리 하는 경우, `fork()` 를 사용한다.
프로세스가 `fork()` 를 호출하면 자식 프로세스가 생성 되고 자기 자신은 부모 프로세스가 된다. `fork()` 가 일어나는 순서는 다음과 같다.

1. 자식 프로세스 용 메모리 영역을 작성하고 거기에 부모 프로세스의 메모리를 복사
2. `fork()` 함수의 리턴 값이 각기 다른 것을 이용하여 부모/자식 프로세스가 서로 다른 일을 하도록 분기

부모 프로세스의 메모리를 복사 하여 자식 프로세스가 생성 되지만 자식 프로세스가 모든 프로퍼티(데이터)들을 상속 받는 것은 아니다.

#### 부모 프로세스로 부터 상속 받는 데이터
부모 프로세스의 메모리가 복사 되어 자식 프로세스가 생성 된다. 상속 받는 데이터(프로퍼티) 는 다음과 같다.

```
1. real user ID, real group ID, effective user ID, effective group ID
2. supplementary group IDs
3. process group ID
4. session ID
5. controlling terminal
6. set-user-ID flag and set-group-ID flag
7. current working directory
8. root directory
9. file mode creation mask
10. signal mask and dispositions
11. the close-on-exec flag for any open file descriptors
12. environment
13. attached shared memory segments
14. resource limits
15. Memory mappings
```

#### 부모 프로세스와 자식 프로세스의 차이점

```
1. `fork()` 시에 리턴되는 integer value 값이 다르다. (부모 프로세스는 자식 프로세스의 ID 를 출력하게 된다. 자식 프로세스는 0 을 리턴한다.)
2. 프로세스 ID 가 다르다.
3. the two processes have different parent process IDs—the parent process ID of the child is the parent; the parent process ID of the parent doesn't change
4. the child's values for tms_utime, tms_stime, tms_cutime, and tms_ustime are set to 0
5. file locks set by the parent are not inherited by the child
6. pending alarms are cleared for the child
7. the set of pending signals for the child is set to the empty set
```

### execve() 함수
전혀 다른 프로그램을 생성할 때에는 `execve()` 함수를 사용한다.

1. 실행 파일을 읽은 다음 프로세스의 메모리 맵에 필요한 정보를 읽어 들인다.
2. 현재 프로세스의 메모리를 새로운 프로세스의 데이터로 덮어쓴다.
3. 새로운 프로세스의 첫 번째 명령부터 실행한다.

이는 프로세스의 수가 증가 되는 것이 아니라 기존 프로세스의 메모리를 새로운 프로세스 인 것 처럼 바꿔 치기 하는 것이라고 보면 된다.

### fork and exec
전혀 다른 프로그램을 (`execve()`) 를 새로 생성 (`fork()`) 할 때는 `fork and exec` 라고 한다. 사실 실제 순서는 언급한 것과 다르다. 정확히는 새로 생성(fork) 하고 다른 프로그램으로 덮어씌우는(exec) 것이다.

bash 를 예로 들자. bash 에서 새로운 프로세스 echo 를 실행한다고 보면, bash 의 자식 프로세스를 만들고 이 자식 프로세스를 echo 메모리로 덮어씌우는 것이다.

### 종료
종료할 때는 `exit()` 함수를 사용하면 메모리를 전부 회수한다.

### 고아 프로세스
고아 프로세스는 부모 프로세스가 `wait` 을 호출하는 대신 종료된 경우이다.
이 경우 해당 프로세스의 부모로 `init` (pid=1) 인 프로세스가 새 부모 프로세스가 되어서 처리 될 수 있다.


고아 프로세스를 만드는 방법 예시
```c
int main() {
    int pid = fork();

    if (pid > 0) {
        // this is parent
        // terminate parent process, no wait
        exit(0);
    }

    if (pid == 0) {
        // this is child
        sleep(60);
    }
    return 0;
}
```

### 좀비 프로세스
좀비 프로세스는 종료 되었지만 부모 프로세스가 `wait` 을 호출하지 않은 프로세스이다.
좀비 프로세스는 메모리나 CPU 를 사용하지는 않지만(메모리에는 남아 있다.)  process table 의 공간을 차지한다. 
process table 은 유한한 리소스이기 때문에 좀비 프로세스가 많이 발생하면 다른 프로세스를 시작할 수 없다.

#### 좀비 프로세스가 오래 살고 있다면?
부모 프로세스는 어떤 이유에서든 계속 존재하고 있다는 것이다.

#### 그러면 좀비는 어떻게 죽일 수 있을까?
사실 좀비는 이미 죽은 프로세스이다. 그래서 단순히 `kill` 할 수가 없다. 그래서 해당 좀비 프로세스의 부모 프로세스를 `kill` 하는 방법으로 해결할 수 있다.

```shell
$ kill -s SIGCHLD <Parent PID>
```

> SIGCHLD 는 자식 프로세스가 종료될 때 부모 프로세스에게 보내는 시그널이다. 이 시그널은 부모 프로세스에게 `wait` 을 호출하고 좀비 자식을 정리하라는 의미이다.


만약 부모 프로세스가 죽었다면 고아 프로세스가 되어서 `init` 에게 다시 입양 된다.


좀비 프로세스를 만드는 방법 예시
```c
int main() {
    int pid = fork();

    if (pid > 0) {
        // this is parent
        slee(60);
    }

    if (pid == 0) {
        // this is child
        // exit but parent doesn't wait for child
        exit(0);
    }
    return 0;
}
```

## 프로세스간 통신
프로세스간 통신 (협력) 을 하는 경우
- 여러 사용자가 동일한 정보를 사용하는 경우
- 특정 task 를 sub 로 나누어 빠르게 처리하고자 하는 경우 (aka. 병렬실행)

이런 경우를 위해 프로세스들은 데이터와 정보를 교환할 수 있는 프로세스간 통신 (IPC) 기법을 사용한다. 이 때 두 가지 모델이 바로 `message passing`, `shared memory` 이다.

많은 처리 코어를 가진 시스템 상에서는 메시지 전달이 공유 메모리 보다 더 나은 성능을 보인다.

### 메시지 전달 (message passing)
- 메시지 전달 모델은 충돌을 회피할 필요가 없어서 적은 양의 데이터를 교환하는데 유용하다.
- 메시지 전달 모델은 분산 시스템에서 공유 메모리보다 구현하기 쉽다.

### 공유 메모리 (shared memory)
- 공유 메모리 영역을 구축할 때만 시스템 호출을 사용하고 이후에는 일반적인 메모리 접근으로 취급되어 커널의 도움이 필요 없다.

### IPC 기법

- 소켓
- RPC
- Pipe
- Named Pipe

