package creational.abstract_factory.buttons;

public class MacOSButton implements Button{

    @Override
    public void paint() {
        System.out.println("Yon have created MacOSButton");
    }
}
