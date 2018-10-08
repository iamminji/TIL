
Chapter02. 스파크 맛보기
==
  
_사전에 Mac Mojave에 스파크와 스칼라를 설치하였고 예제 대부분을 스칼라로 진행할 예정이다._
  
  
##### 예제. 스칼라로 줄 세기
  
1. 터미널에서 스파크 셸 실행
<code> > spark-shell </code>

그러면 아래와 같은 텍스트가 뜬다.
<pre><code>Welcome to
     ____              __
    / __/__  ___ _____/ /__
   _\ \/ _ \/ _ `/ __/  '_/
  /___/ .__/\_,_/_/ /_/\_\   version 2.3.1
     /_/

Using Scala version 2.11.8 (Java HotSpot(TM) 64-Bit Server VM, Java 1.8.0_181)
Type in expressions to have them evaluated.
Type :help for more information.
</code></pre>
  
그러렴 콘솔에서 아래 처럼 진행한다. 사전에 test.txt 에는 4줄의 텍스트를 적어두었다.
  
<pre><code>scala> var lines = sc.textFile("/파일/경로/test.txt")
lines: org.apache.spark.rdd.RDD[String] = /파일/경로/test.txt MapPartitionsRDD[1] at textFile at <console>:24

scala> lines.count()
res0: Long = 4

scala> lines.first()
res1: String = Hello

scala> :quit
</code></pre>
  
#### 스파크의 핵심 개념 소개
넓게 보면 모든 스파크 애플리케이션은 클러스터에서 다양한 병렬 연산을 수행하는 드라이버 프로그램으로 구성된다. 드라이버 프로그램은 당신이 만든 애플리케이션의
 main 함수를 갖고 있으며 클러스터의 분산 데이터세트를 정의하고 그 데이터세트에 연산 작업을 수행한다.
  
드라이버 프로그램들은 연산 클러스터에 연결을 나타내는 <code>SparkContext</code> 객체를 통해 스파크에 접속한다. 
셸에서는 이 <code>SparkContext</code> 객체는 자동적으로 sc라는 변수에 만들어진다.

<pre><code>scala> sc
res1: org.apache.spark.SparkContext = org.apache.spark.SparkContext@31006a75
</code></pre>
  
<code>SparkContext</code> 객체를 하나 만들었다면 그것으로 __RDD__ 를 만들어 낼 수 있다. 예제에서는 텍스트 파일의 각 라인을 표현하는 
__RDD__ 를 만들기 위해 <code>sc.textFile()</code>을 호출 했었다. 그러고 나면 그 라인에 <code>count()</code> 같은 다양한 연산을 수행해 볼 수 있다.
  
이런 연산들을 수행하기 위해 드라이버 프로그램들은 보통 익스큐터(executor)라 불리는 다수의 노드를 관리한다.
  
#### SparkContext 초기화하기
한 애플리케이션이 스파크이 연동되려면 우선 프로그램 내에서 관련 스파크 패키지들을 임포트(import)하고 <code>SparkContext</code> 객체를 생성해야 한다.
  
<pre><code>import org.apache.spark.SparkConf
import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._

val conf = new SparkConf().setMaster("local").setAppName("My App")
val sc = new SparkContext(conf)
</code></pre>
  
위의 예제는 <code>SparkContext</code>를 초기화하는 가장 간단한 형태를 보여 주고 있다. 이 때 다음의 두 가지 인자를 전달해 주어야 한다.

* 클러스터 URL
    - 위 예제에서 local이라 쓰인 부분, 스파크에게 어떤 식으로 클러스터에세 접속할 지 알려 준다. local은 한 개의 스레드(thread)나 단일의 로컬 
    머신에서 돌 때 따로 접속할 필요가 없음을 알려 주는 특수한 값이다.
* 애플리케이션 이름
    - 위 예제에서 My App 부분. 클러스터에 접속한다면 클러스터 UI에서 저 이름으로 애플리케이션을 구분할 수 있다.