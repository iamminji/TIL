
class OuterClass01 {
    int num;

    private class InnerClass {
        public void print() {
            System.out.println("This is an inner class");
        }
    }

    void displayInner() {
        InnerClass innerClass = new InnerClass();
        innerClass.print();
    }
}


public class InnerClass01 {
    public static void main(String[] args) {
        OuterClass01 outerClass = new OuterClass01();
        outerClass.displayInner();
    }

}
