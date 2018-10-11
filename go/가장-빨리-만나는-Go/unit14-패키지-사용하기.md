
unit14. 패키지 사용하기
==
  
  
패키지를 사용하려면 __import__ 키워드를 사용한다.
  
+ import "패키지"
  
14.1 여러 패키지 사용하기
--
import 키워드를 여러 개 사용하여 패키지를 가져오면 된다. 너무 패키지가 많으면 괄호로 묶는다.
  
<pre><code>package main
import "fmt"
import "runtime"
</code></pre>
  
<pre><code>package main
import (
    "fmt"
    "runtime"
)
</code></pre>
  
14.2 전역 패키지로 사용하기
--
<pre><code>package main
import . "fmt"

func main() {
    Println("Hello, World!")
}
</code></pre>
  
__fmt__ 키워드 없이 전역에서 쓸 수 있지만 중복 네이밍이 있을 수 있으니까 왠만하면 쓰지 말자.
  
14.3 패키지 별칭 사용하기
--
<pre><code>package main
import f "fmt"

func main() {
    f.Println("Hello, World!")
}
</code></pre>
  
사용하고 안쓰려면 언더바!
