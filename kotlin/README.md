# 코틀린

안드로이드 공식 언어로 채택 되면서, 자바와의 호환성도 좋고 스프링 부트와 같은 인기(?) 프레임 워크에서도
지원하고 있는 언어가 바로 코틀린이다.

> 최근에 스프링 부트를 자주 쓰는데, 코틀린 지원도 된다는게 나름 매력적으로 느껴졌다.
스칼라에도 플레이 프레임워크라는 웹 프레임워크가 있지만 문법도 어렵고 나에겐 진입 장벽이 너무 높다.
(그 밖에도 빌드 시간이 오래 걸린다던가 하는 단점도 있긴 하다.)
요새 레거시를 제외하고 신규 프로젝트는 코틀린으로도 많이 한다고 하니 궁금하기도 해서 튜토리얼만 빠르게 보기로 했다.

문법 자체는 크게 어렵지 않다.
label 은 golang 하고 비슷한 것 같고, switch 문에서 arrow operation 은 자바 14에 생긴 operation 과 같았다.
var 과 val 은 스칼라에도 있는 거기도 하고

type alias 는 마음에 든다. 자바를 하면서 동일한 이름의 클래스에 불필요하게 긴 패키지 명을 적어야 한다는 것이
마음에 안들었기 때문이었다. 그래서 파이썬의 as 가 있었으면 얼마나 좋았을까? 라는 생각을 하곤 했는데 코틀린에 비슷한게 있다니!

전체적으로 자바랑 비슷하지만 한 가지 다른 점은 static 이라는 키워드가 없다는 것이다. 정확히는 Companion object 가 static 을 대신한다. (스칼라에도 Object 라는 키워드가 있는데 이는 싱글턴을 보장하는 녀석이다.)

[이런 글](https://softwareengineering.stackexchange.com/questions/356415/why-is-there-no-static-keyword-in-kotlin) 도 있는데 static 이라는 녀석은 다른 멤버와는 다르게 행동할 가능성이 있고, companion object 를 사용함으로써 static 사용의 단점을 보완할 수 있다나 뭐라나.

[이 글](https://tourspace.tistory.com/109) 에서 static 을 대신하여 Object 키워드를 사용하는 방법이 자세히 나와 있다.

프로젝트를 직접 해본 것은 아니고 튜토리얼만 읽었을 뿐이라 사실 문법에 익숙하지는 않는다.
코틀린으로 스프링 부트를 만져보면서 공부해야겠다.

## 참고
- [https://www.tutorialspoint.com/kotlin/index.htm](https://www.tutorialspoint.com/kotlin/index.htm)
