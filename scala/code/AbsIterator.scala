




object TargetTest2 extends App {
    def loop(body: => Unit): LoopUnlessCond = new LoopUnlessCond(body)
    protected class LoopUnlessCond(body: => Unit) {
        def unless(cond: => Boolean) {
            body
            if (!cond) unless(cond)
        }
    }
    var i = 10
    loop {
        i -= 1
        println("i = " + i)
    } unless(i == 0)
}
