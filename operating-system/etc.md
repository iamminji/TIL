## 심볼릭 링크와 하드링크
심볼릭 링크로 걸면 원본 파일을 삭제하면 링크가 깨진다.
하드 링크를 걸면 원본 파일을 삭제해도 깨지지 않는다.

파일은 inode 로 관리 되는데 하드 링크로 걸 때는 같은 inode 를 가르키는 다른 파일을 생성한 것이여서 원본 파일 즉 최초의 파일을 삭제 하여도 신규로 건 (하드링크로 건) 파일은 살아 있기 때문에 깨지지 않는 것이다.

inode 는 자신에게 참조된 링크를 카운트 하는데, 이 값이 0 일 때만 파일이 실제로 삭제가 된다.

(심볼릭 링크는 바로 가기와 같은 것이다.)

