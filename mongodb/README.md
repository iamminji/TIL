# 몽고 DB
## Overview
- 데이터의 기본 단위 Document (aka. Row)
- Collection (RDBMS 로 따지면 Table)

## 저장
### insertMany
#### ordered 옵션
`ordered` 옵션을 `false`로 주면 저장 시 충돌이 나는 것은 제외하고 나머지 것들은 다 저장됨

`ordered` 가 몽고디비에 도큐먼트 저장 시, 요청 순서를 정렬하여 저장할 것인지에 대한 여부이기 때문에 옵션을 `false`로 주면 실패난 것은 패스하는 것이다.

```shell
db.movies.insertMany([
    {"_id": 5, "title": "!"},
    {"_id": 6, "title": "!"},
    {"_id": 6, "title": "!"}, // 이것만 저장 안됨
    {"_id": 7, "title": "!"},
], {"ordered": false})
```

기본 값은 `true` 이다.

#### Message size limit
몽고디비는 48MB보다 큰 메시지 사이즈를 허용하지 않아서  `insertMany`시 이 사이즈가 넘어가면 쪼개져서 저장된다.

## 갱신
### updateOne
#### upsert 옵션
`upsert` 옵션을 줘서 매치되는게 없으면 도큐먼트를 추가하게 만들 수도 있다. 이 경우 삽입될 때만 특정 필드의 값을 주고 싶다면 (갱신은 제외) `$setOnInsert` 라는 옵션을 사용하면 된다.

### findOneAndUpdate
`findOndAndDelete`, `findOneAndReplace` 등도 있다.

기본적으로 이 3가지 연산은 수정된 도큐먼트의 값을 원자적으로 얻어 스레드 간 race condition 을 피하는, 심플한 코드를 작성할 수 있게 해준다.

#### retuenNewDocument
해당 옵션을 추가하면 변경된 도큐먼트를 리턴값으로 주게 된다. (없으면 변경 전 걸 줌)

## 정렬
### 복합 인덱스 설계하기
1. 동등 필터에 대한 키를 맨 앞에 표시해야 한다.
2. 정렬에 사용되는 키는 multivalued field 앞에 표시해야 한다.
3. multivalued 필터에 대한 키는 마지막에 표시해야 한다.

