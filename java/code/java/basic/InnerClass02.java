package java.basic;

class OuterClass02 {
    private int num = 175;

    public class InnerClass {
        public int getNum() {
            System.out.println("This is the getnum method of the inner class");
            return num;
        }
    }
}


public class InnerClass02 {
    public static void main(String[] args) {
        OuterClass02 outerClass02 = new OuterClass02();

        OuterClass02.InnerClass innerClass = outerClass02.new InnerClass();
        System.out.println(innerClass.getNum());
    }

}
