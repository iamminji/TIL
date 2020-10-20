# HBase CLI command 정리

## hbase 테이블 별 용량 확인하기

path 는 기본적으로 zookeeper base path 이다. (변경되었다면 hbase master status 웹 UI 에서 확인할 수 있고 임의로 변경했다면 사용자가 당연히 알겠지?) 기본은 `/hbase/data/default` 이다.

```
hdfs dfs -du -h  /hbase/data/default/
```
