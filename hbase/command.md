# HBase CLI command 정리

## HBase 테이블 이름 변경
HBase 는 테이블 `rename` 개념이 없고, 테이블 명을 변경해주기 위해선 `snapshot`을 뜨고 해당 스냅샷의 이름을 바꿔야 한다. 

뜬 `snapshot`으로 다시 테이블 생성 `restore` 하면 기존 이름으로 생성되므로 주의한다.

```
hbase shell> disable 'tableName'
hbase shell> snapshot 'tableName', 'tableSnapshot'
hbase shell> clone_snapshot 'tableSnapshot', 'newTableName'
hbase shell> delete_snapshot 'tableSnapshot'
hbase shell> drop 'tableName'
```


# hbase 테이블 별 용량 확인하기
__사실 hbase cli 는 아니고 hdfs 명령어이긴 하다__

path 는 기본적으로 zookeeper base path 이다. (변경되었다면 hbase master status 웹 UI 에서 확인할 수 있고 임의로 변경했다면 사용자가 당연히 알겠지?) 기본은 `/hbase/data/default` 이다.

```
hdfs dfs -du -h  /hbase/data/default/
```
