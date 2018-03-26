### 숫자공간

파이썬은 이름공간과 값공간이 따로 있는데, 어떤 변수에 값을 할당하는 개념이 아니라 값은 이미 존재하고 그 값과 이름 공간만 연결해주는 개념이라고 보면 된다.

Cpython 에서는 ```-5 부터 256 까지``` 메모리에 미리 할당해두고 이름 공간에만 매치 시켜주게 했는데 그래서 되게 재미난 결과가 나오기도 한다.

<pre>
>>> a = 256
>>> b = 256
>>> a is b
True
>>> a = 257
>>> b = 257
>>> a is b
False
>>> 257 is 257
True
</pre>

256 까지는 이미 메모리에 올라가 있어서(이미 존재하는 객체이기 때문에) is로 타입 비교를 해도 같은거고 257은 서로 다르기 때문에 a is b가 False가 나온다. 대신 값 자체를 비교하면 True 가 나온다. 넘 재밌당

* https://stackoverflow.com/questions/306313/is-operator-behaves-unexpectedly-with-integers
* https://docs.python.org/2/c-api/int.html
