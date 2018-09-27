# akka 를 이용한 통신 예제

###### 최초 작성일: 2018-09-27


Messages.scala

<pre><code>package progscala2.introscala.shapes

object Messages {

  object Exit

  object Finished

  case class Response(message: String)

}
</code></pre>

ShapesDrawingActor.scala

<pre><code>package progscala2.introscala.shapes

import akka.actor.Actor

class ShapesDrawingActor extends Actor {

  import Messages._

  override def receive: Receive = {
    case s: Shape =>
      s.draw(str => println(s"ShapeDrawingActor:$str"))
      sender ! Response(s"ShapeDrawingActor: $s drawn")
    case Exit =>
      println(s"ShapesDrawingActor: exiting...")
      sender ! Finished
    case unexpected =>
      val response = Response(s"ERROR: Unknown message:$unexpected")
      println(s"ShapesDrawingActor:$response")
      sender ! response
  }

}
</code></pre>

Shapes.scala

<pre><code>package progscala2.introscala.shapes

case class Point(x: Double = 0.0, y: Double = 0.0)


abstract class Shape {
  def draw(f: String => Unit): Unit = f(s"draw: ${this.toString}")

}

case class Circle(center: Point, radius: Double) extends Shape

case class Rectangle(lowerLeft: Point, height: Double, width: Double) extends Shape

case class Triangle(point1: Point, point2: Point, point3: Point) extends Shape
</code></pre>

ShapesDrawingDriver.scala

<pre><code>package progscala2.introscala.shapes

import akka.actor.{Props, Actor, ActorRef, ActorSystem}
import com.typesafe.config.ConfigFactory


private object Start

object ShapesDrawingDriver {
  def main(args: Array[String]): Unit = {
    val system = ActorSystem("DrawingActorSystem", ConfigFactory.load())
    val drawer = system.actorOf(Props(new ShapesDrawingActor), "drawingActor")
    val driver = system.actorOf(Props(new ShapesDrawingDriver(drawer)), "drawingService")
    driver ! Start
  }
}


class ShapesDrawingDriver(drawerActor: ActorRef) extends Actor {

  import Messages._

  override def receive: Receive = {
    case Start =>
      drawerActor ! Circle(Point(0.0, 0.0), 1.0)
      drawerActor ! Rectangle(Point(0.0, 0.0), 2, 5)
      drawerActor ! 3.14159
      drawerActor ! Triangle(Point(0.0, 0.0), Point(2.0, 0.0), Point(1.0, 2.0))
      drawerActor ! Exit
    case Finished =>
      println(s"ShapesDrawingDriver: cleaning up...")
      context.system.terminate()
    case response: Response =>
      println("ShapesDrawingDriver: Response = " + response)
    case unexpected =>
      println("ShapesDrawingDriver: ERROR: Received an unexpected message = " + unexpected)
  }
}
</code></pre>
