## About redux
리덕스는 컴포넌트 밖에서 관리 로직을 처리하는 라이브러리다. 리덕스를 사용하면 스토어라는 객체 내부에 상태를 담게 된다.

_상태는 state 이다._

상태에 어떤 변화를 일으키고 싶을 때 아래와 같이 진행한다. 이 과정을 dispatch 라고 한다.
<pre><code>컴포넌트 -> (action) -> 스토어
</code></pre>

스토어가 action 을 받으면 reducer가 상태를 어떻게 변경시켜야 할지 정하고 처리되면 새 상태를 스토어에 저장한다.

스토어 안에 상태가 바뀌면 스토어를 구독하고 있는 컴포넌트에 바로 전달한다. 부모 컴포넌트로 props를 전달하는 작업은 생략하며, 리덕스에 연결하는 함수를 사용하여 컴포넌트를 스토어에 구독시킨다.

* store 스토어 : 애플리케이션의 상태 값들을 내장
* action 액션 : 상태 변화를 일으킬 때 참조하는 객체
* dispatch 디스패치 : 액션을 스토어에 전달하는 것을 의미
* reducer 리듀서 : 상태를 변화시키는 로직이 있는 함수
* subscribe 구독 : 스토어 값이 필요한 컴포넌트는 스토어를 구독한다.

#### action
액션 객체 안에서 type 값은 필수고 나머지는 선택이다.
<pre><code>const INCREMENT = 'INCREMENT';
const increment = (diff) => ({
    type: INCREMENT,
    diff: diff
})
</code></pre>

#### reducer
리듀서는 파라미터 두 개를 받는데, 첫번째는 현재 상태고 두 번쨰는 액션 객체다. 리덕스에선 상태를 업데이트 할 때는 컴포넌트의 상태를 다루는 것 처럼 직접 업데이트 하는 것이 아니라 새로운 객체를 만든 후에 그 안에 상태를 정의해야 한다.

Object.assign 을 써도 되지만 전개 연산자가 더 깔끔하니까 전개 연산자를 쓰도록 하자.

<pre><code>function reducer(state = {}, action) {
  switch(action.type) {
    case INCREMENT:
      return {
        ...state,
        number: state.number + action.diff
      };
      // 또는
      //  return Object.assign({}, state, {
      //    number: state.number + action.diff
      //  });
      //
  }
}
</code></pre>

#### store
스토어를 생성할 때는 `createStore` 함수를 사용한다. 파라미터로는 리듀서 함수가 들어가고, 두 번째 파라미터를 설정하면 해당 값을 스토어으 ㅣ기본 값으로 사용한다.

<pre><code>const { createStore } = Redux;
const store = createStore(reducer);
</code></pre>

#### subscribe
리덕스 에서 구독 함수를 직접 실행하진 않고 `connect` 가 대신해준다.

#### dispatch
스토어에 액션을 넣을 때는 `store.dispatch` 함수를 사용한다.


### redux의 세 가지 원칙

#### 스토어는 단 한개
모든 상태는 하나의 스토어에서 관리 된다.

#### 상태는 읽기 전용
상태를 변화시키는 유일한 방법은 무슨 일이 벌어나는지를 묘사하는 액션 객체를 전달하는 방법 뿐이다.

#### 변화는 순수 함수로 구성
리듀서는 이전 상태와 액션을 받아 다음 상태를 반환하는 순수 함수이다.

참고
- 리액트를 다루는 기술
- [https://dobbit.github.io/redux/introduction/ThreePrinciples.html](https://dobbit.github.io/redux/introduction/ThreePrinciples.html)
