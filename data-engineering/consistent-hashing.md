# Consistent Hashing

## Hasing
Consistent Hasing이란 분산 시스템에서 데이터를 일관성있게 분산하는 알고리즘이다. 

일반적으로 modular 연산을 사용하는 해싱의 경우, 노드가 추가/삭제 될 때 전체 아이템(키)에 대한 연산을 다시 진행해 데이터를 분산해야 한다. 이러한  (rehasing) 재분배 과정은 분산시스템의 성능과 확장성을 저하시킬 수 있다.

Consistent Hashing(일관성 해싱)은 이를 해결하기 위한 알고리즘으로 (기본적으로) 원형/링 형식의 해시공간을 사용한다.

> 좀 찾아보니 ring이 아니라 jump 개념도 있긴 했다.

일반적인 해싱은 거의 모든 키가 재매핑되어야 했지만, 일관성 해싱은 평균적으로 K/n의 키만 재매핑된다고 한다. (n은 전체 노드의 개수, K는 item의 개수)

## Consistent Hashing

Consistent Hashing은 아래와 같은 구조를 가진다.
- Requestor (작은 색깔의 원) / Server node (또는 bucket)로 구성된 Ring 구조가 있다.
- 수가 정해져있지 않은 Ring이라고 가정한다. (node 또는 bucket들은 늘어나거 줄수 있음)
- Requestor들은 동일 해시 함수를 사용한다.


[!image](https://ik.imagekit.io/ably/ghost/prod/2022/01/hashing@2x.png?tr=w-1520,q-50)


요청은 시계 방향으로 매핑해서 (가장 가까운 Server Node가) 처리한다.
[!image](https://ik.imagekit.io/ably/ghost/prod/2022/07/mapping-in-the-hashing@2x.png?tr=w-1520,q-50)


매핑된 노드에 장애가 나면, 그 다음 가까운 노드에게 요청이 할당된다.
[!image](https://ik.imagekit.io/ably/ghost/prod/2022/01/failure-of-node3@2x.png?tr=w-1520,q-50)


> 노드들은 bucket 이라고도 불리며, 노드 추가시 (virtual) vnode, vbucket 이라고도 한다. 가상노드는 실제 노드의 Replica로 분산과 밸런싱을 위해서 사용된다.


### 일관성 해싱 최악의 성능
- 링의 노드 변화가 크면 재분배해야할 부하가 많아, 성능에 악영향을 끼친다.
- 일단 해싱함수를 사용하기 때문에 잘못된 함수 사용 시, 핫 스팟이 발생할 여지가 있다. (특정 노드에 부하가 집중 됨)


## 참고
- https://www.joinc.co.kr/w/man/12/hash/consistent
- https://ably.com/blog/implementing-efficient-consistent-hashing
- https://charsyam.wordpress.com/2016/10/02/%EC%9E%85-%EA%B0%9C%EB%B0%9C-consistent-hashing-%EC%97%90-%EB%8C%80%ED%95%9C-%EA%B8%B0%EC%B4%88/
