# Git 유용한 명령어 정리


## 푸시 강제
```
git push --force-with-lease origin master
```
이미 커밋 푸시한 내용을 `rebase` 로 합쳐버렸다면 커밋은 되어도 푸시할 때 에러가 난다.
왜냐하면 커밋 로그가 변경되었기 때문이다. 

그때 푸시를 강제로 하는 명령어이다.


## 머지 충돌 해결 후
머지 충돌을 해결한 후에 커밋할 때 `i` 옵션으로 해 준다.

```
git commit -i
```

## 커밋 바꾸기
커밋 로그를 하나로 합칠 때 사용한다.

`rebase` 를 써서 본문의 내용을 `pick` 에서 `squash` 로만 바꿔 주면 된다.

이미 푸쉬된 커밋 로그 4개를 합칠 때
```
git rebase -i origin/master~4 master
git push origin +master
```

## pull 강제
로컬 git 에서 pull 하는데 merge conflict 가 나는데 local history 는 없어도 상관 없을 때 사용!

```
git fetch --all
git reset --hard origin/master.
```

## tag force push

```
git tag -f <tag>
git push -f origin --tag
```

#### 참고
- https://stackoverflow.com/questions/1125968/how-do-i-force-git-pull-to-overwrite-local-files
