# Zookeeper

## 주키퍼 API

주키퍼에는 다음과 같은 API 가 있다.

data 를 담고 있는 /path 라는 이름의 znode 생성
```
create /path data
```

/path 라는 znode 를 제거
```
delete /path
```

/path 라는 znode 가 존재하는지 확인 (Stat)
```
exists /path
```

/path 라는 znode 에 data를 저장
```
setData /path data
```

/path 라는 znode의 데이터를 가져옴
```
getData /path
```

/path 라는 znode의 모든 자식 목록을 반환
```
getChildren /path
```

### 주키퍼의 노드 종류

주키퍼에는 총 4가지 종류의 노드가 있다. (3.5 기준으로 TTL, container 가 생겼다.)

- persistent
영구적인 노드로 `delete` 를 이용해 제거가 가능하다.
- ephemeral
임시 노드로 주키퍼 클라이언트와 세션이 끊기기 전까지 살아있다가, 끊기면 사라지는 노드이다. 혹은 세션을 맺은 클라이언트가 삭제해야 제거되는 노드이다.
- ttl
- container

각 노드에는 순차적 (sequential) 노드도 존재한다. 단순히 노드가 생성되면 뒤에 sequential number 가 붙는다고 생각하면 된다. 즉 tasks 라는 zNode가 있을 때 해당 옵션(-s)으로 생성하면 tasks-(숫자) 가 붙는 것이다.
