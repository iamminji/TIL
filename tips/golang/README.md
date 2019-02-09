# Golang Tips

## 문자열 처리

### 유니코드 정규화 unicode normalization

```
package main

import (
    "fmt"
    "unicode"

    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

func main() {
    t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
    result, _, _ := transform.String(t, "žůžo")
    fmt.Println(result)
}
```
