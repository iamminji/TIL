# Chapter07. 내장 제어 구문 (Built in Control Structure)

## match 표현식
스칼라의 `match` 표현식은 여타 언어의 `switch` 문과 유사하다.
자바와의 다른 점이 있다면 암묵적으로 `case` 문에 `break` 가 있고, 표현식의 결과가 값이라는 점이다.

<pre><code>val firstArg = if (!args.isEmpty) args(0) else ""
val friend =
  firstArg match {
    case "salt" => "pepper"
    case "chips" => "salsa"
    case "eggs" => "bacon"
    case _ => "huh?"
  }
</code></pre>
