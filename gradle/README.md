# Gradle


## tutorial 

gradle 5.X 이후엔 leftShift가 없어서 책(Gradle 철저입문) 의 예제를 그대로 따라하면 `Could not find method leftShift() for arguments` 과 같은 에러가 난다.

그러므로 예제를 아래처럼 변경해준다.

```
task hello {
    doLast {
        println 'Hello Gradle world!'
    }
}
```

실행
```
$ gradle hello
$ gradle -q hello # 결과 값만 보여줌 (로그 없이)
```

ref
- [https://docs.gradle.org/current/userguide/tutorial_using_tasks.html](https://docs.gradle.org/current/userguide/tutorial_using_tasks.html)


## 다른 gradle 파일 명 사용하기

b 옵션을 붙이면 다른 이름을 가진 gradle task 파일 이름도 사용할 수 있다.

```
gradle -b sample.gradle task_name
```

