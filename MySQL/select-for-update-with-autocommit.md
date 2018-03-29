### SELECT FOR UPDATE와 autocommit

MySQL 에서 ```SELECT ... FOR UPDATE``` 구문을 써서 100개의 row를 가져왔다고 했을 때.

100 개의 작업을 다 끝낸 후에 ```commit```을 임의로 하곤 했다.

만약에 ```autocommit```을 활성화 하고 1개의 작업만 (update) 하면 바로 락이 풀리나? 가 궁금했다.

보니까 해당 구문을 사용할땐 반드시 disabled 하라는 구절이 도큐먼트에 있었다. 활성화 하면 잠금이 걸리지 않는다고 한다.

<pre>
Locking of rows for update using SELECT FOR UPDATE only applies when autocommit is disabled (either by beginning transaction with START TRANSACTION or by setting autocommit to 0. If autocommit is enabled, the rows matching the specification are not locked.
</pre>

* https://dev.mysql.com/doc/refman/5.7/en/innodb-locking-reads.html
