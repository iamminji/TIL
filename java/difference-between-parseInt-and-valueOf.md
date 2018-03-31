### 자바에서 parseInt와 valueOf 의 차이점

parseInt는 primitive int를 반환하고, valueOf(String) 은 new Integer()를 반환한다.

그래서 Integer.valueOf(s) 는 new Integer(Integer.parseInt(s)) 이것과 동일하다고 한다.

보면서, 자바는 그러면 왜 integer를 primitive와 object로 둘다 두어서 사람을 헷갈리게 하는건지? 언제 object를 써야하고 primitive를 써야하는건지? 잘 모르겠다.

자바를 안 써서 피부로 와닿지 않는 것 같기도 하다.

* https://softwareengineering.stackexchange.com/questions/203970/when-to-use-primitive-vs-class-in-java
