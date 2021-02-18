# Apache HBase Performance Tuning

https://hbase.apache.org/book.html#performance

## Writing to HBase

### Batch Loading
Use the bulk load tool if you can. See Bulk Loading. Otherwise, pay attention to the below.

### Table Creatioin: Pre-Creating Regions
HBase 는 디폴트로 한개의 region 이 생성 된다. 이 의미는 모든 클라이언트가 region 이 split 될 정도로 
커지지 않을 때 까지 동일한 region 에 write 할 것이라는 의미이다.


그래서 HBase 에서는 좋은 튜닝 방법 중 하나로, 빈 region 을 사전에 생성하기도 한다. 
(단 너무 많은 region 을 생성하지는 말자. 성능 하락의 주범이다.)


사전에 region 을 생성하는 방법으로는 2 가지가 있다.

#### Admin 을 사용
```
byte[] startKey = ...;      // your lowest key
byte[] endKey = ...;        // your highest key
int numberOfRegions = ...;  // # of regions to create
admin.createTable(table, startKey, endKey, numberOfRegions);
```

#### 직접 split 하기
```
byte[][] splits = ...;   // create your own splits
admin.createTable(table, splits);
```

비슷한 동작을 shell 에서도 할 수 있다.
```
# create table with specific split points
hbase>create 't1','f1',SPLITS => ['\x10\x00', '\x20\x00', '\x30\x00', '\x40\x00']

# create table with four regions based on random bytes keys
hbase>create 't2','f1', { NUMREGIONS => 4 , SPLITALGO => 'UniformSplit' }

# create table with five regions based on hex keys
create 't3','f1', { NUMREGIONS => 5, SPLITALGO => 'HexStringSplit' }
```

### Table Creation: Deferred Log Flush
