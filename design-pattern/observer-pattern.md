### 옵저버 패턴

옵저버 패턴
- 한 객체의 상태가 바뀌면 그 객체에 의존하는 다른 객체들한테 연락이 가고 자동으로 내용이 갱신되는 방식으로 일대다(one-to-many) 의존성을 정의합니다.


옵저버 패턴은 뉴스를 구독하는 것과 비슷합니다.
1. The subscriber opens subscription for the newspaper
2. You subscribe to the newspaper
3. Somebody else subscribes to the newspaper
4. When there's a new newspaper, you and that somebody else get a new
newspaper
5. If you don't want to receive the newspaper anymore, you cancel your
subscription and you will not receive next newspaper (but others will)

옵저버 패턴으로 얻는 장점은 아래와 같습니다.
1. Subject와 Observer 사이에 느슨한 결합을 유지합니다. Subject는 오직 Obeservers 들(인터페이스)에 대해서만 압니다.
2. Subject와 Observer 사이에 브로드 캐스트 매시징 기능이 있습니다.
3. 다수의 Obeservers 는 실행 중에 변경될 수 있습니다.
4. Subject는 원하는 수의 Observers 들을 유지할 수 있습니다.


옵저버 패턴이 주로 쓰이는 곳은 아래와 같습니다.
- 외부에서 발생한 이벤트(사용자 입력같은)에 대한 응답.
- 객체의 속성 값 변화에 따른 응답.

파이썬의 Observer pattern 구현 예제입니다.

<code>Subject</code>클래스
```
import time


class Subject(object):
  def __init__(self):
    self.observers = []
    self.cur_time = None

  def register_observer(self, observer):
    if observer in self.observers:
      print observer, 'already in subscribed observers'
    else:
      self.observers.append(observer)

  def unregister_observer(self, observer):
    try:
      self.observers.remove(observer)
    except ValueError:
      print 'No such observer in subject'

  def notify_observers(self):
    self.cur_time = time.time()

  for observer in self.observers:
    observer.notify(self.cur_time)
```

파이썬에선 인터페이스 개념이 없습니다.
예제는 대신 추상 클래스를 구현하기 위해 ABCMeta, abstractmethod 를 사용하고 있습니다.

<code>Observer</code>클래스
```
from abc import ABCMeta, abstractmethod
import datetime


class Observer(object):
  """Abstract class for observers, provides notify method as
    interface for subjects."""
  __metaclass__ = ABCMeta

  @abstractmethod
  def notify(self, unix_timestamp):
    pass
```

그리고 이 Observer 추상 클래스를 아래와 같이 상속 받습니다.

```
class EUTimeObserver(Observer):
  def __init__(self, name):
    self.name = name

  def notify(self, unix_timestamp):
    time = datetime.datetime.fromtimestamp(int(unix_timestamp)).strftime('%Y
%m-%d %H:%M:%S')
    print 'Observer', self.name, 'says:', time
```


참고
* http://www.acornpub.co.kr/book/python-design-patterns
* https://ko.wikipedia.org/wiki/%EC%98%B5%EC%84%9C%EB%B2%84_%ED%8C%A8%ED%84%B4
