# Redux
리덕스는 컴포넌트 밖에서 관리 로직을 처리하는 라이브러리다. 리덕스를 사용하면 스토어라는 객체 내부에 상태를 담게 된다.

## 왜 리덕스인가?
리액트의 컴포넌트 트리는 부모, 자식, 그리고 많은 형제 (...) 로 이루어진다. 여기서 상위 컴포넌트를 `App` 이라고 부른다고 한다면, 이 `App` 의 자손 컴포넌트들은
필연적으로 `App` 의 상태에 영향을 받을 수 밖에 없다.

`App` 에서 `state` 를 업데이트 하면 `App` 컴포넌트가 리렌더링 되고, 리액트 특성상 하위 컴포넌트 들도
모두 리랜더링 된다. 혹은 하위 컴포넌트 중에 하나만 상위에서 영향을 받은 `state` 로 업데이트 하고 싶어도, 다른 형제들까지 모두 리렌더링 할 수 밖에 없는 구조가 현재 리액트의 구조이다.

즉, 리덕스는 __상태 전달 문제__ 를 해결하기 위해 나타났다.

리덕스는 컴포넌트 밖에서 관리 로직을 처리하는 라이브러리이다.

컴포넌트 밖에 `store` 와 `reducer` 를 두고 각 컴포넌트들은 `store` 에서 `state` 를 `subscribe` 하고, 행동할 `action` 을 `reducer` 로 `dispatch` 하는 형태다.

쉽게 얘기하면 `store` 라는 구조를 따로 두고 여기에 모든 상태 `state` 를 저장해서,
컴포넌트가 가져다가 알아서 행동 `action` 한다고 보면 될 것 같다.

용어를 정리하자면 아래와 같다.

###### store 상태
애플리케이션의 상태 값들을 내장
###### action 액션
상태 변화를 일으킬 때 참조하는 객체
###### dispatch 디스패치
액션을 스토어에 전달하는 것을 의미
###### reducer 리듀서
상태를 변화시키는 로직이 있는 함수
###### subscribe 구독
스토어 값이 필요한 컴포넌트는 스토어를 구독한다.

각 용어들의 예제를 하단에 적도록 하겠다.

### Understanding how Redux Work

우선 `state` 에 어떤 변화를 일으키고 싶을 때는 아래와 같이 진행한다.

<pre><code>컴포넌트 -> (action) -> store</code></pre>

`store` 가 `action` 을 받으면 `reducer` 가 상태를 어떻게 변경시켜야 할지 정하고 처리되면 새 상태를 `store` 에 저장한다.

`store` 안에 `state` 가 바뀌면 `store` 를 `subscribe` 하고 있는 컴포넌트에 바로 전달한다. 부모 컴포넌트로 `props` 를 전달하는 작업은 생략하며, 리덕스에 연결하는 함수를 사용하여 컴포넌트를 `store` 에 `subscribe` 시킨다.

> subscribe 대신에 connect 함수를 쓰면 좀 더 유연하게 컴포넌트를 store 에 구독시킬 수 있다. connect 는 redux 패키지가 아니고 react-redux 패키지 안에 있으므로 따로 설치 해야 한다.

#### action

액션에 있어야할 것은 크게 두 가지로 나누어진다. 하나는 액션 타입이고 나머지 액션 생성자다.

액션 타입이란 컴포넌트가 행동할 액션의 종류를 상수 형태로 정의한 것이고 (type이라고 한다.) 액션 생정자는 해당 액션 타입을 이용해 진짜로 행동할 액션을 담는 객체다.

###### action type
<pre><code>const INCREMENT = 'INCREMENT';
</code></pre>

###### action creator
액션 객체 안에서 type 값은 __필수__ 고 나머지는 선택이다.

<pre><code>const increment = (diff) => ({
    type: INCREMENT,
    diff: diff
})
</code></pre>

#### reducer
리듀서는 파라미터 두 개를 받는데, 첫번째는 현재 상태고 두 번쨰는 액션 객체다. 리덕스에선 상태를 업데이트 할 때는 컴포넌트의 상태를 다루는 것 처럼 직접 업데이트 하는 것이 아니라 새로운 객체를 만든 후에 그 안에 상태를 정의해야 한다.

Object.assign 을 써도 되지만 전개 연산자가 더 깔끔하니까 (ES6 부터 지원)전개 연산자를 쓰도록 하자.

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

리듀서에서 주의 할 점은 리듀서는 __항상 같은 결과를 내놓아야 한다__ 는 것이다. 이 말인 즉슨, API 호출 처럼 결과 값이 변경되는 것이 아니라
동일한 상태를 받으면 그에 맞는 결과 상태는 늘 똑같아야 한다는 것이다!

#### store
스토어를 생성할 때는 `createStore` 함수를 사용한다. 파라미터로는 리듀서 함수가 들어가고, 두 번째 파라미터를 설정하면 해당 값을 스토어으 ㅣ기본 값으로 사용한다.

<pre><code>const { createStore } = Redux;
const store = createStore(reducer);
</code></pre>

#### subscribe
리덕스 에서 구독 함수를 직접 실행하진 않고 `connect` 가 대신해준다.

#### dispatch
스토어에 액션을 넣을 때는 `store.dispatch` 함수를 사용한다.

#### connect
connect 함수를 사용하여 컴포넌트를 스토어에 연결 시킬 수 있다. connect 함수에는 3가지 파라미터가 들어간다.
각각의 함수들은 없어도 되고, 있어도 된다. (필수가 아니다)

각각의 함수들은 컴포넌트에서 사용할 `props` 를 반환한다.

- mapStateToProps : store.getState() 결과 값인 state를 파라미터로 받아 컴포넌트의 props로 사용할 객체를 반환한다.
- mapDispatchToProps : dispatch를 파라미터로 받아 액션을 디스패치 하는 함수들을 객체 안에 넣어서 반환한다.
- mergeProps : state와 dispatch 가 동시에 필요한 함수를 props로 전달해야 할 때 사용하는데, 일반적으로는 잘 사용하지 않는다.

### redux의 세 가지 원칙

#### 스토어는 단 한개
모든 상태는 하나의 스토어에서 관리 된다.

#### 상태는 읽기 전용
상태를 변화시키는 유일한 방법은 무슨 일이 벌어나는지를 묘사하는 액션 객체를 전달하는 방법 뿐이다.

#### 변화는 순수 함수로 구성
리듀서는 이전 상태와 액션을 받아 다음 상태를 반환하는 순수 함수이다.

### 심화
#### 비동기 액션
비동기 API 를 호출할 때 최소 3가지 액션이 있다.

- 리듀서에게 요청이 시작되었음을 알리는 액션.
- 리듀서에게 요청이 성공적으로 완료되었다고 알리는 액션.
- 리듀서에게 요청이 실패했음을 알리는 액션.

#### 비동기 흐름
`redux-thunk` 나 `redux-promise` 같은 비동기 미들웨어는 store 의 `dispatch` 를 감싸서 액션이 아니라, 함수나 약속(?) 같은
다른 것들을 보낼 수 있게 해준다. 미들웨어는 보내는 것을 받아서 해석한 다음, 다음 미들웨어로 액션을 넘긴다.

참고
- 리액트를 다루는 기술
- [https://dobbit.github.io/redux/introduction/ThreePrinciples.html](https://dobbit.github.io/redux/introduction/ThreePrinciples.html)
- [https://redux.js.org/introduction/motivation](https://redux.js.org/introduction/motivation)
