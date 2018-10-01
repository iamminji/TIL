
abstract class AnonymousClass {
    public abstract void myMethod();
}

public class InnerClass04 {
    public static void main(String[] args) {
        AnonymousClass inner = new AnonymousClass() {
            @Override
            public void myMethod() {
                System.out.println("This is an example of anonymous class");
            }
        };
        inner.myMethod();
    }
}
