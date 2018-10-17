채널 Channel
==
  
  
채널
--
채널은 두 고루친이 서로 통신하고 실행흐름을 동기화하는 수단을 제공한다. 채널에 데이터를 전달하고 받을 땐 <code><-</code> 연산자를 사용한다. 

<pre><code>c <- "ping"</code></pre>
<code>ping</code> 이라는 string 문자열을 채널 c에 전달한다.
  
<pre><code>msg := <- c</code></pre>
<code>msg</code> 라는 변수에 채널 c의 값을 빼 저장한다. 채널에 값이 없으면 들어올 때 가지 기다린다.
  

### 채널 버퍼링
채널에 버퍼를 1개 이상 설정하면 __asynchronous channel__ 이 생성된다.
  
### select
Go 언어는 <code>switch</code> 구문 대신 분기를 위해 <code>select</code> 를 사용한다.
