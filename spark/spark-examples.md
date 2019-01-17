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
