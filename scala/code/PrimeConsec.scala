
import scala.collection.mutable


object PrimeConsec {

    val cache = new mutable.HashMap[Int, Boolean]

    def isPrime(num: Int): Boolean = {

        if (cache.contains(num)) {
            return cache(num)
        }

        var n = num - 1

        while (n > 1) {
            if (num % n == 0) {
                cache put(num, false)
                return false
            }
            n -= 1
        }
        cache put(num, true)
        true
    }

    def countKprimes(k: Int, start: Long, nd: Long): String = {
        // your code

        cache put(0, false)
        cache put(1, false)

        val primeList = mutable.HashMap[Int, List[Int]]()

        for (n <- start.until(nd)) {
            for (i <- 0 to n.toInt) {
                if (isPrime(i)) {
                    if (primeList.contains(n.toInt)) {
                        primeList.put(n.toInt, i :: primeList(n.toInt))
                    } else {
                        primeList.put(n.toInt, List(i))
                    }
                }
            }
            var count = 0
//            println(primeList.get(n.toInt).foreach(v => v))
        }
        ""
    }


    def main(args: Array[String]): Unit = {

        //        countKprimes(2, 0, 100)
        countKprimes(2, 0, 4)

    }
}
