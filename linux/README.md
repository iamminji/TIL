# 리눅스

## 커널
### 시스템 콜
프로세스는 프로세스의 생성이나 하드웨어 조작 등 커널의 도움이 필요할 경우 시스템 콜을 통해 처리를 요청한다. 시스템 콜의 종류는 아래와 같다.

- 프로세스 생성, 삭제
- 메모리 확보, 해제
- 프로세스 간 통신 (IPC)
- 네트워크
- 파일 시스템 다루기
- 파일 다루기


프로세스는 사용자 모드로 주로 실행되지만 시스템 콜에 의해 커널 모드로 실행되고는 한다. (시스템 콜을 호출하면 CPU 에서 인터럽트 이벤트가 발생하여 이 인터럽트 이벤트를 통해 사용자 모드에서 커널 모드로 변경 된다.)


## 프로세스 관리
리눅스에서는 두 가지 목적으로 프로세스를 생성한다.

1. 같은 프로그램의 처리를 여러 개의 프로세스가 나눠서 처리 (`fork()`)
2. 전혀 다른 프로그램을 생성 (`execve()`)

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

## 메모리 관리
각 프로세스도 메모리를 사용하고, 커널 자체도 메모리를 사용한다.

#### 메모리 정보
`free` 명령어로 확인할 수 있다.

#### 메모리 부족
메모리 사용량이 증가하면 free 한 메모리 영역이 줄어든다. 그리고 이 영역을 다 사용하게 된다면 메모리 관리 시스템은 커널 내부의 해제 가능한 메모리 영역을 해제한다.
그래도 메모리가 부족하다면 시스템은 `Out Of Memory` 상태가 된다.

이러한 경우 메모리 관리 시스템은 적절한 프로세스를 찾아 강제 종료 `kill` 을 한다. 이를 __OOM Killer__ 라고 한다.

##### 죽일 (...) 프로세스를 찾는 방법
메모리 관리 시스템은 죽이기 가장 좋은 프로세스를 선택한다. 여기서 _좋다_ 라는 의미는 죽였을 때 가장 큰 메모리를 확보할 수 있고 시스템에 큰 영향이 가지 않을 수 있다는 것에 대한 의미이다.

이를 위해 리눅스는 `oom_score` 라는 값을 각 프로세스마다 갖고 있다. 이런식으로 확인 가능하다.

```
$ cat /proc/<pid>/oom_score
```

이 값이 크면 클 수록 죽을 확률이 높다.

[스택오버플로우](https://unix.stackexchange.com/questions/153585/how-does-the-oom-killer-decide-which-process-to-kill-first), [linux_mm](https://linux-mm.org/OOM_Killer)  를 참고하였다.

만약 아주 중요한 프로세스를 메모리 관리 시스템이 kill 하게 된다면, 서비스는 큰 문제가 발생할 것이다. 그래서 리눅스 커널 파라미터 (vm.panic_on_oom) 의 기본값을 변경해 프로세스가 아니라 시스템을 종료하게 만들 수도 있다.


#### 메모리 할당
커널이 프로세스에 메모리를 할당하는 일은 크게 두 가지 타이밍에서 벌어진다.

1. 프로세스를 생성할 때
2. 프로세스를 생성한 뒤 추가로 동적 메모리를 할당할 때

프로세스가 생성된 뒤 추가로 메모리가 더 필요하면 프로세스는 커널에 메모리 확보용 시스템 콜을 호출해서 메모리 할당을 요청한다.
이 때 발생하는 문제점은 다음과 같이 있다.

1. 메모리 단편화
2. 다른 용도의 메모리에 접근 가능
3. 여러 프로세스를 다루기 곤란

이러한 점들 때문에 가상 메모리가 등장했다.

### 가상 메모리
메모리를 프로세스가 직접 접근하지 않고, 가상 주소라는 주소를 사용하여 간접적으로 접근하도록 하는 방식이다.

#### 가상 메모리 응용

- 파일 맵
- 디맨드 페이징
- Copy On Write 방식의 고속 프로세스 생성
- 스왑 (Swap)
- 계층형 페이지 테이블
- Huge Page


## 심볼릭 링크와 하드링크
심볼릭 링크로 걸면 원본 파일을 삭제하면 링크가 깨진다.
하드 링크를 걸면 원본 파일을 삭제해도 깨지지 않는다.

파일은 inode 로 관리 되는데 하드 링크로 걸 때는 같은 inode 를 가르키는 다른 파일을 생성한 것이여서 원본 파일 즉 최초의 파일을 삭제 하여도 신규로 건 (하드링크로 건) 파일은 살아 있기 때문에 깨지지 않는 것이다.

inode 는 자신에게 참조된 링크를 카운트 하는데, 이 값이 0 일 때만 파일이 실제로 삭제가 된다.

(심볼릭 링크는 바로 가기와 같은 것이다.)

## 참고
- 실습과 그림으로 배우는 리눅스 구조
- https://stackoverflow.com/questions/21220107/what-parent-process-stuff-gets-shared-in-newly-created-child-process-in-linux
- https://medium.com/pocs/%EB%A6%AC%EB%88%85%EC%8A%A4-%EC%BB%A4%EB%84%90-%EC%9A%B4%EC%98%81%EC%B2%B4%EC%A0%9C-%EA%B0%95%EC%9D%98%EB%85%B8%ED%8A%B8-2-78406a13c5c9
