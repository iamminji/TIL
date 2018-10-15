Go Coding Convention
==
  
  
Package Structure
--
<pre><code>bin/
   hello                          # command executable
   outyet                         # command executable
src/
   github.com/golang/example/
       .git/                      # Git repository metadata
hello/
    hello.go               # command source
outyet/
    main.go                # command source
    main_test.go           # test source
stringutil/
    reverse.go             # package source
    reverse_test.go        # test source
   golang.org/x/image/
       .git/                      # Git repository metadata
bmp/
    reader.go              # package source
    writer.go              # package source
   ... (many more repositories and packages omitted) ...
</code></pre>
  
Naming Conventions
--
패키지 외부에서도 사용하려면 파일 명의 맨 앞글자를 __대문자__ 로 지정하고, 내부에서만 사용할 경우 __소문자__ 로 지정한다.