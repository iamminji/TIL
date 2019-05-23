# vim

## 특정 패턴을 포함한 라인을 삭제하고 싶을 때

```
:g/pattern/d
```

ref 
- [https://vim.fandom.com/wiki/Delete_all_lines_containing_a_pattern](https://vim.fandom.com/wiki/Delete_all_lines_containing_a_pattern)

## 모든 라인 밑에 빈 줄을 (new line) 을 넣고 싶을때

숫자만 바꿔서 1 대신에 3을 넣으면 every 3 line 후에 new line 이 들어간다.

```
:%s/\v(.*\n){1}/&\r
```

ref
- [https://stackoverflow.com/questions/10413906/how-to-add-a-line-after-every-few-lines-in-vim](https://stackoverflow.com/questions/10413906/how-to-add-a-line-after-every-few-lines-in-vim)
- [https://vim.fandom.com/wiki/Add_a_newline_after_given_patterns](https://vim.fandom.com/wiki/Add_a_newline_after_given_patterns)


## custom macro 만들기

ref
- [https://superuser.com/questions/93492/how-to-add-a-command-in-vim-editor](https://superuser.com/questions/93492/how-to-add-a-command-in-vim-editor)
- [http://vimdoc.sourceforge.net/htmldoc/usr_40.html#40.2](http://vimdoc.sourceforge.net/htmldoc/usr_40.html#40.2)

