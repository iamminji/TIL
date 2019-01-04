## 커밋 바꾸기

커밋 바꾸기

### squash
커밋 로그를 하나로 합칠 때 사용한다.

`rebase` 를 써서 본문의 내용을 `pick` 에서 `squash` 로만 바꿔 주면 된다.

이미 푸쉬된 커밋 로그 4개를 합칠 때
```
git rebase -i origin/master~4 master
git push origin +master
```
