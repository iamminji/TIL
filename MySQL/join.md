# join

## 차집합

`minus` 사용해서 차집합
```
SELECT select_list1
FROM table_name1
MINUS
SELECT select_list2
FROM table_name2;
```

1) `join` 사용해서 차집합
```
SELECT
    select_list
FROM
    table1
LEFT JOIN table2
    ON join_predicate
WHERE
    table2.column_name IS NULL;
```

2) `join` 사용해서 차집합
```
SELECT
    id
FROM
    t1
LEFT JOIN
    t2 USING (id)
WHERE
    t2.id IS NULL;
```
