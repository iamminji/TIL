
public class InnerClass03 {
    void myMethod() {
        int num = 23;

        class MethodInnerClass {
            public void print() {
                System.out.println("This is method inner class: num => " + num);
            }
        }

        MethodInnerClass inner = new MethodInnerClass();
        inner.print();
    }

    public static void main(String[] args) {
        InnerClass03 inner = new InnerClass03();
        inner.myMethod();
    }
}
