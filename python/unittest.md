### 파이썬 단위 테스트

파이썬에서 단위 테스트 모듈 이름은 ```unittest``` 이다.

최근에 TDD를 공부하면서, 클래스 단위로 테스트를 진행하고 싶은데 해당 인스턴스가 매번 테스트 할 때마다 생성되고 있는 것이 마음에 들지 않았다.

즉, unittest 를 상속하고 setUp을 오버라이드 하여, 해당 메소드 안에 테스트 하고 싶은 클래스 인스턴스를 생성하면 각 테스트가 실행될 때 마다 setUp이 다시 실행된다는 것이다.

이게 무슨말이냐 하면 예를 들어서 코드가 아래와 같다.

<pre>
import uniitest

class MyTest(unittest.TestCase):
  def setUp(self):
    self.my_class = MyClass()

  def test_add(self):
    self.assertEqual(self.my_class.add(1,3), 4)

  def test_mul(self):
    self.assertEqual(self.my_class.mul(3,1), 3)
</pre>


이러면 나는 self.my_class 가 한번 생성되고 해당 인스턴스를 재활용할 거라고 생각했는데, 실제로는 각 테스트가 실행될 때 마다 다시 setUp이 실행되었다.

이때 사용할 수 있는 것이 setUpClass 이다. (파이썬2.7 부터 가능하다!)
해당 메서드를 classmethod 데코레이터를 사용하면, 원했던 기능대로 동작한다.

setUp 부분을 아래처럼 수정해주면 된다!

<pre>
  @classmethod
  def setUpClass(cls):
    cls.my_class = MyClass()
</pre>

마찬가지로 테스트가 끝나고 실행될 tearDown 역시 tearDownClass를 활용하면 된다!

아래의 링크들에서 다양한 테스트 예제들을 살펴볼 수 있다.
* http://pythontesting.net/framework/unittest/unittest-fixtures/
