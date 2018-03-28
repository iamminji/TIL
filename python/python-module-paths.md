### 파이썬 모듈 경로

site-packages 에 user module 을 만들어서, 해당 경로를 추가하였다.
그런데 user module 이 실제로 있는 경로 안에서 import 하면 어떻게 되는지 궁금해서 모듈을 import 할 떄 어떤 것이 우선시 되는지 찾아보았다.

기본적으로 파이썬의 import module 우선 순위는 다음과 같다.
1. 빌트인 모듈
2. sys.path 로 출력되는 경로 순서

여기서 sys.path 는 아래의 경로들을 의미한다.
1. 현재 디렉토리 (sys.path 를 출력하면 가장 먼저 빈 문자열이 들어가 있는데 이를 의미)
2. PYTHONPATH (shell 에서 지정한 경로)
3. 설치할때 지정한 경로

대신에 현재 경로에 심볼릭 링크가 있다고 해서, 심볼릭 링크가 sys.path에 포함 되지는 않는다. (그렇다고 import 못하는 건 아니다. 되긴 됨)

* https://docs.python.org/3/tutorial/modules.html#the-module-search-path
