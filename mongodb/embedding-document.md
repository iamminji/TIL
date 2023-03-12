# Embedding MongoDB
몽고디비는 도큐먼트 안에 내장 도큐먼트 (Embedded document, nested doucment) 를 둘 수 있다.

전형적인 RDBMS 에서는 타입이 다른 데이터는 서로 다른 테이블에 저장하여 JOIN 을 사용하지만 몽고 디비는 도큐먼트 안에 도큐먼트를 둘 수 있다는 이야기이다.

예를 들면 `user` 라는 컬렉션 안에 데이터에 `address` 라는 도큐먼트를 만들고 `address` 관련 데이터를 포함하게 한다.

즉 아래와 같은 구조에서
```shell
> db.user.findOne()
{
    _id: 111111,
    email: “email@example.com”,
    name: {given: “Jane”, family: “Han”},
}

> db.address.find({user_id: 111111})
{
    _id: 121212,
    street: “111 Elm Street”,
    city: “Springfield”,
    state: “Ohio”,
    country: “US”,
    zip: “00000”
}
```
이렇게 포함시키게 한다.

```shell
> db.user.findOne({"_id": 111111})
{
    _id: 111111,
    email: "email@example.com",
    name: {given: "Jane", family: "Han"},
    address: {
        street: "11 Elm Street",
        city: "Springfield",
        state: "Ohio",
        country: "US",
        zip: "00000",
    }
}
```
접근할 때는 dot 을 사용한다. (`user.address`)

## 언제 Embedding document 를 사용하는게 좋을까?

내장 도큐먼트는 관계성을 갖는 데이터를 저장하는데 효율적이고 깔끔한 방법중 하나이다. 특히 도큐먼트 (데이터) 를 함께 조회할 때 좋다. (user 의 address 를 가져와야하는 경우)
몽고디비의 스키마는 디자인하는 일반적인 방법은 기본적으로 내장 도큐먼트를 사용하는 것이다. 그리고 어플리케이션이나 데이터베이스 측면에서의 Join 은 특별한 이유가 있어야 한다.


## The Embedded Document Pattern
함께 쓰는 것은 함께 저장하도록 한다!

## The Embedded Subset Pattern
데이터의 일부분만 필요한 경우 일부분만을 저장할 수도 있다. 예를 들어`movie` 와 `review` 컬렉션이 있다고 하자. 영화를 조회할 때 항상 최근 리뷰가 필요하다면 `review` 의 일부분만을 저장할 수도 있다.

## The Extended Reference Pattern
Subset Pattern 과 유사한데 데이터를 저장하는 게 아니라 다른 도큐먼트 (컬렉션) 에 접근할 수 있는 레퍼런스(와 일부 필트) 를 저장하는 것이다.

레퍼런스 도큐먼트의 일부 필드를 정기적으로 가져와야 한다면 그 필드들도 함께 저장한다!

## Unbounded Lists
리스트 사이즈 (Embedding document 의 타입) 가 계속 커질 예정이면 저장하지 않는 것도 방법 중 하나이다. 몽고디비는 단일 도큐먼트의 사이즈 제한이 있고, 메모리 사용량에 부정적 영향을 끼칠 수도 있기 때문이다.

따라서 리스트가 무한으로 커지면 그냥 다른 컬렉션에 저장하도록 하자!

## Independent Access
각 컬렉션에 따로 접근해야할 요구사항이 있으면 분리해서 사용하도록 한다.

독립적으로 접근하거나 만들어야 할 경우 서로 다른 컬렉션에 저장하도록 하자!


# Reference
- https://www.mongodb.com/basics/embedded-mongodb
