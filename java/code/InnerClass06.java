
public class InnerClass06 {
    static class NestedClass {
        public void myMethod() {
            System.out.println("This is my nested class");
        }
    }

    public static void main(String[] args) {
        InnerClass06.NestedClass innerClass06 = new InnerClass06.NestedClass();
        innerClass06.myMethod();
    }
}
