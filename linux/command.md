# 자주 쓰는 linux command 정리

## ls

- ls 에서 정렬해서 보고 싶을때 `-t`
- ls 에서 사이즈 확인해보고 싶을 때 `-h`

## df
- 서버 디스크 별 용량 보고 싶을 때 `df -h`

## shell

```
PATH="$(perl -e 'print join(":", grep { not $seen{$_}++ } split(/:/, $ENV{PATH}))')"
```
path export 할 때 실수로 중복된 path 목록들이 생성된 경우가 있다. 이러한 path 목록에서 unique 한 것만 남기고 정렬까지 해 주는 커맨드

```
tail -f 로그명 | perl -nle 's/\\x(..)/pack("C",hex($1))/eg;print $_'
```
로그 볼 때 한글 문자열이 human readable 한 게 아니라 \x 로 시작되는 경우가 있다(특히 파이썬에서). 이를 한글 표기로 변경해주는 옵션 커맨드


## diff

두 파일에서 첫번째 컬럼만 값이 다른 경우를 보고 싶었다. 
```
awk 'NR==FNR{c[$1]++;next};c[$1] == 0' file1.txt file2.txt > diff_file.txt
```

참고
[https://unix.stackexchange.com/questions/174599/using-diff-on-a-specific-column-in-a-file](https://unix.stackexchange.com/questions/174599/using-diff-on-a-specific-column-in-a-file

