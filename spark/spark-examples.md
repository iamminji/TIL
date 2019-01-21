# 스파크


## 글자 수 세기 예제

##### WordCount.scala
```
package com.my.package

import org.apache.spark.SparkContext
import org.apache.spark.SparkContext._
import org.apache.spark.SparkConf


object WordCount {
    def main(args: Array[String]) {
        val infile = args(0)
        val conf = new SparkConf().setAppName("word count")
        val sc = new SparkContext(conf)
        val input = sc.textFile(infile)
        val words = input.flatMap(line => line.split(" "))
        val counts = words.map(word => (word, 1)).reduceByKey{case(x, y) => x + y}
        counts.saveAsTextFile(args(1))
    }
}
```

##### build.sbt
```
name := "spark-tutorial-test"

version := "0.1"

scalaVersion := "2.11.0"


libraryDependencies ++= Seq(
    "org.apache.spark" %% "spark-sql" % "2.3.1",
    "org.apache.spark" %% "spark-core" % "2.3.1",
    "org.apache.spark" %% "spark-streaming" % "2.3.1"
)
```

여기서 spark 버젼 2.4.X 를 쓰고 스칼라 2.12.X 를 썼는데 인텔리제이에선 정상 동작하고, 터미널에선 에러가 아래처럼 났다.

```
Exception in thread "main" java.lang.BootstrapMethodError: java.lang.NoClassDefFoundError: scala/runtime/java8/JFunction2$mcIII$sp
```

알고보니 인텔리제이의 스칼라/스파크 버젼이랑 터미널에서 심볼릭 링크로 걸려있는 스칼라/스파크 버젼이 달라서 그랬다.

설치한 스파크는 버젼 2.3.X 인데 스칼라는 2.12.X 여서 호환이 안되었던 정말 어처구니 없는 문제였다.

##### 실행

```
sbt clean
sbt compile
sbt package
```

후에 그 다음 스파크 실행

```
spark-submit --class com.my.package.WordCount /path/to/spark-tutorial-test.jar /path/to/inputfile /path/to/outputfile
```

하면 정상 동작

## 스파크 명령어

### run-example
스파크 배포판엔 기본적으로 예제 코드가 있다.

```
run-example SparkPi
```

위의 명령어를 치면 pi 가 출력된다.

### spark-submit
스파크 실행 명령어로, 어떤 특별한 로직을 수행하는건 아니고 자바의 리플렉션 API 를 이용해 사용자가 지정한 클래스의 메인함수를 실행하는 스파크의 런처 모듈 중 하나다.

>스파크 애플리케이션을 실행하기 위해서는 메인 함수를 가진 애플리케이션을 작성하고, 스파크에서 제공하는 `spark-submit` 스크립트를 이용해 실행한다.

### 스파크 잡 죽이기
```
yarn application -kill <application_id>
```
