
WordCount
======================

###### 최초 작성일: 2018-09-29

[WordCount.java](https://github.com/iamminji/TIL/blob/master/hadoop/code/WordCount.java) 실행 방법

<pre><code>echo “Hello World Hello Hadoop” > file01.txt
echo “Goodbye Hadoop Bye Hadoop” > file02.txt
hadoop fs -mkdir input
hadoop fs -put file* input/

# 이렇게 하고 그 다음 jar 실행
hadoop jar MyJar.jar input output

# output 이미 존재하면 에러 날 수 있는데 아래 처럼 삭제하면 됨
hadoop fs -rm -r output

# 결과 파일 확인
hadoop fs -cat output/part-r-00000

# 결과 파일 명이 다르면 아래 처럼 현재 파일 명 확인
hadoop fs -ls output/
</code></pre>

코드 참고
https://hadoop.apache.org/docs/stable/hadoop-mapreduce-client/hadoop-mapreduce-client-core/MapReduceTutorial.html
