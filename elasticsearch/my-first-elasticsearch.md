## Elasticsearch

Elasticsearch 를 mac os 에 설치해보고 튜토리얼을 따라해보았다.


- elasticsearch 설치

```
brew install elasticsearch
```

- elasticsearch가 정상적으로 설치 되었는지 확인한다. 아래 명령어 입력 후 브라우저 상에서 <code>localhost:9200</code> 으로 접근한다.

```
> elasticsearch
```

- json 으로 결과 값이 나오는걸 확인했다면, python 패키지를 설치해준다.

```
    pip install elasticsearch
```

- python 코드로 아래와 같이 입력한다. (이 때, 2번이 실행중이어야 한다.)

```
from datetime import datetime
from elasticsearch import Elasticsearch
es = Elasticsearch()

doc = {
    'author': 'kimchy',
    'text': 'Elasticsearch: cool. bonsai cool.',
    'timestamp': datetime.now(),
}
res = es.get(index="test-index", doc_type='tweet', id=1)
```

- 브라우저에 <code>http://localhost:9200/_search?q=author:kimchy</code> 로 접속한다.

결과가 아래의 형태처럼 나오면 성공!

```
{
    "took": 38,
    "timed_out": false,
    "_shards": {
        "total": 5,
        "successful": 5,
        "skipped": 0,
        "failed": 0
    },
"hits": {
    "total": 1,
    "max_score": 0.2876821,
    "hits": [
        {
            "_index": "test-index",
            "_type": "tweet",
            "_id": "1",
            "_score": 0.2876821,
            "_source": {
                "text": "Elasticsearch: cool. bonsai cool.",
                "timestamp": "2018-04-22T14:32:47.336112",
                "author": "kimchy"
                }
            }
        ]
    }
}
```
