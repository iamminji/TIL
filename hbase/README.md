# HBase




## 데이터 처리
HBase 테이블의 모든 로우는 rowkey 라는 이름의 유일한 식별자를 가지고 있다. HBase 테이블에서 데이터의 위치를 표시하는 데 사용되는 다른 coordinate 들도 있지만 rowkey는 가장 기본적인 수단이다. 관계형 데이터베이스와 같이 rowkey는 unique 하다.


#### HBase 테이블 이름 변경
HBase 는 테이블 `rename` 개념이 없고, 테이블 명을 변경해주기 위해선 `snapshot`을 뜨고 해당 스냅샷의 이름을 바꿔야 한다.

뜬 `snapshot`으로 다시 테이블 생성 `restore` 하면 기존 이름으로 생성되므로 주의한다.

```
hbase shell> disable 'tableName'
hbase shell> snapshot 'tableName', 'tableSnapshot'
hbase shell> clone_snapshot 'tableSnapshot', 'newTableName'
hbase shell> delete_snapshot 'tableSnapshot'
hbase shell> drop 'tableName'
```
