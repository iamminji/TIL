# 도커

## ENTRYPOINT 와 CMD 의 차이

## ADD 와 COPY 의 차이

## 도커 자주 쓰는 명령어

`docker run` 는 pull, create, start, attach (-i, -t 옵션을 사용했을 떄)ㅁ 와 같ㅇ다.

`docker pull` 리모트에서 도커 이미지를 가져온다.

`docker run -p p1:p2` 로컬 포트 p1과 프로세스 포트 p2를 매핑해서 실행시킨다.

`docker images` 로컬에 있는 도커 이미지들을 볼 수 있다.

`docker exec -it <container id> sh` 도커 컨테이너 내부로 들어간다. (기본 셸 sh)

`docker logs -f <container id>` 도커 컨테이너의 표준 출력 로그 (?) 를 볼 수 있다.
