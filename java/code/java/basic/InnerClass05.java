package java.basic;

interface Message {
    String greet();
}


public class InnerClass05 {
    public void displayMessage(Message m) {
        System.out.println(m.greet() + ", This is an example of anonymous inner class as an argument");
    }

    public static void main(String[] args) {
        InnerClass05 innerClass05 = new InnerClass05();

        innerClass05.displayMessage(new Message() {
            @Override
            public String greet() {
                return "Hello";
            }
        });

        /* Java8
        * Anonymous Class To Lambda
        * */
        innerClass05.displayMessage(() -> "Hello");
    }

}
