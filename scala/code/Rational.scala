

object Rational {
    def main(args: Array[String]): Unit = {
        val rational = new Rational(1, 2)
        println(rational)

        val rational2 = new Rational(2, 3)
        println(rational add rational2)
    }
}

class Rational(n: Int, d: Int) {
    require(d != 0)
    val number: Int = n
    val denom: Int = d

    override def toString = n + "/" + d

    def add(that: Rational): Rational = new Rational(number * that.denom + that.number * denom, denom * that.denom)

    def lessThan(that: Rational) = this.number * that.denom < that.number * this.denom

    def max(that: Rational): Rational = {
        if (lessThan(that)) that else this

    }
}