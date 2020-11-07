# with-recursive statement

single column 에 대해 선형으로 증가하는 컬럼 값을 만들고 싶을 때 사용한다.

예)
```
WITH RECURSIVE cte (n) AS
(
  SELECT 1
  UNION ALL
  SELECT n + 1 FROM cte WHERE n < 5
)
SELECT * FROM cte;
```

결과
```
+------+
| n    |
+------+
|    1 |
|    2 |
|    3 |
|    4 |
|    5 |
+------+
```

#### 참고
https://dev.mysql.com/doc/refman/8.0/en/with.html#common-table-expressions-recursive
