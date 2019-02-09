# Python Tips

최근 파이썬을 하도 안 써서, 많이 까먹었다. 그래서내가 까먹기 싫어서 정리하는 파이썬 3.X (어쩌면 2.7도) 에서 사용하는 팁들을 정리해본다.

## 리스트 list

### 문자열 뒤집기

```
>>> a = [1,2,3]
>>> list(reversed(a)) # 1
[3, 2, 1]
>>> a[::-1] # 2
[3, 2, 1]
```

파이썬에서 문자열을 뒤집을 수 있는 방법은 (일단 지금 생각나는) 두 가지다.
첫 째는 [reversed](https://docs.python.org/3/library/functions.html#reversed) 를 사용하는 것이다. 해당 함수는 파이썬 빌트인 함수로, 결과 값으로 iterator 를 리턴하기 때문에 list 로 감싸주어야 한다.

두 번째는 파이썬 리스트의 성질을 이용하는 방법이다. 파이썬의 리스트는 총 세 개의 인덱스를 받을 수 있는데 첫번째는 start 고 두번째가 end, 마지막이 jump 다. 처음과 두번째를 비우고 마지막 부분을 -1 로 지정하면 거꾸로 간다는 의미이므로 뒤집혀진 리스트를 리턴해주게 된다.

## 딕셔너리 dictionary

파이썬의 딕셔너리는 자바의 map 과 같다. 다만 다른게 있다면 값 타입을 지정하지 않아도 된다는 것이다. (스칼라의 Any나 Go 의 interface 같은)

그러나 파이썬의 강력한 기능은 그냥 딕셔너리가 아니라, `collections` 에 있는 딕셔너리 서브 클래스들이다. 해당 빌트인 라이브러리에 들어있는 딕셔너리 들은 파이썬의 코드를 더욱 짧게! 더욱 __pythonic__ 하게 바꿔준다.

### 딕셔너리 한 줄로 값 추가하기

이게 뭔 말인가 싶을 수 있다. `int` 나 `str` 같은 타입이 아니라 `list` 나 `dict` 같은 자료구조를 딕셔너리의 값으로 사용하고 싶을 때가 있다.

그러면, 파이썬에서 보통 딕셔너리를 사용할 때 키가 없으면 값을 아래와 같이 초기화 해준다.

```
>>> d = {}
>>> if 1 not in d:
...     d[1] = list("hey!")
... else:
...     d[1].append("hey!")
...
```

그런데 이게 굉장히 길어보이고, 번거롭게 느껴질 때가 있다. 그럴 경우엔 `collections` 의 `defaultdict` 를 사용해 보자. 위의 코드를 한 줄로 작성할 수 있다.

```
>>> from collections import defaultdict
>>> d = defaultdict(list)
>>> d[1].append("hey")
>>> d
defaultdict(<class 'list'>, {1: ['hey']})
```

### 딕셔너리 순서 보장하기
딕셔너리는 또 당연하게도 기본적으로 순서가 보장되지 않는다. 이 때문에 다른 언어에서는 순서를 보장하기 위한 자료구조를 직접 만들지만 파이썬엔 [OrderedDict](https://docs.python.org/3/library/collections.html#collections.OrderedDict) 라는 자료구조도 있다!

```
>>> from collections import OrderedDict
>>> a = OrderedDict()
>>> a[1] = 1
>>> a[2] = 2
>>> a[3] = 3
>>> a
OrderedDict([(1, 1), (2, 2), (3, 3)])
```

해당 자료구조엔 심지어 `pop` 할 때 마지막 것 부터 하기 위한 boolean 매개변수도 존재한다.

```
>>> a
OrderedDict([(1, 1), (2, 2), (3, 3)])
>>> a.popitem(last=True)
(3, 3)
>>> a.popitem(last=False)
(1, 1)
>>> a
OrderedDict([(2, 2)])
```

### 스위치 switch-case
파이썬에선 switch-case 구문이 없다. 그래서 보통은 if-else 를 사용하지만 dictionary 를 사용하면 switch-case 와 비슷한 효과를 나타낼 수 있다.

``` 
def f(x):
    return {'a': '1', 'b': '2'}.get(x, '3')
```

파이썬의 dictionary 를 그대로 리턴 하되, switch 의 default 구문 처럼 쓰기 위해 get 을 마지막에 추가 하였다. 이러면 없는 키에 대해선 (default) '3' 을 리턴하게 된다.

## 기타 등등

### (iterator가 있는, 리스트나 튜플 등의) 자료구조 참/거짓 한 줄로 검증하기

#### any
파이썬의 빌트인엔 `any` 라는 함수도 있다. 이 함수는 `iterator` 가 구현된 자료구조를 내부에 받아서, 값이 하나라도 `True` 면 `True` 를 리턴하고 아니면 `False` 를 리턴한다.

#### all
`any` 가 있으니 반대로 `all` 도 있다. 전부 `True` 여야 `True` 를 리턴한다.

> 알고리즘 짤 때 값 검증을 굳이 loop 돌지 않고 쓸 수 있어 좋다. 당연하게도 전체 iter 를 돌기 때문에 시간 복잡도는 O(n) 이다.

```
>>> all([True,False,True,True])
False
>>> any([True,False,True,True])
True
```

