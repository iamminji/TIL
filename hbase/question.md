# 궁금한 점
궁금한 내용을 목록화 하고, 도큐먼트와 코드를 살펴본 후 하나씩 정리해보도록 하자

- [ ] RegionServer 가 죽으면 담당하는 Region 이 다른 서버로 이동 되는데, 이 때 Region 이 disk 에 쓰여지지 않았다면 (HFile 로 저장되지 않았다면), 즉 메모리 상태에 있었다면 RegionServer 가 죽으면 해당 Region 은 날라가는 거 아닐까? (-> 뭔 소리지 replica 3 까지 되는데 정신차려....)

WAL 때문에 Memstore 로 복구 될 수 있는것인건가?

- [ ] RegionServer 가 죽으면 Zookeeper 랑 통신이 끊기게 되는데 그 이후에 해당 Region 은 어떤 룰(?) 에 의해서 다른 RegionServer 로 옮겨지는 걸까?

