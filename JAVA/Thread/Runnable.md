# Runnable
#Runnable #Thread 

보통 쓰레드 객체를 만들 때 Thread 클래스를 상속하여 만들기도 하지만 Runnable 인터페이스를 구현하는 방법을 주로 쓴다. 왜냐하면 Thread클래스를 상속하면 다른 클래슬르 상속할 수 없기 때문이다.

[Thread에 대하여 알고싶다면 클릭!! 이 문서를 읽기전에 읽는걸 추천](https://github.com/YoonSeok-Heo/TIL/blob/main/JAVA/Thread/Thread)

위 링크에 Thread를 이용한 코드이다.
```java
public class Sample extends Thread{
    int seq;
    public Sample(int seq){
        this.seq = seq;
    }

    public void run(){
        System.out.println(this.seq + " thread start. ");   // 쓰레드 시작
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
        }
        System.out.println(this.seq + " thread end.");  // 쓰레드 종료
    }
    public static void main(String[] args){
        ArrayList<Thread> threads = new ArrayList<>();

        IntStream.range(0, 10).forEach(i -> {
                    Thread t = new Sample(i);
                    t.start();
                    threads.add(t);
        });

        // 이것도 스트림을 사용해보았다.
        threads.stream().forEach(thread ->{
            try {
                thread.join();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        System.out.println("main end.");  // main 메서드 종료
    }
}
```

위 코드를 Runnable인터페이스를 사용하는 방식으로 변경해보자

```java
public class Sample implements Runnable{ //여기수정
    int seq;
    public Sample(int seq){
        this.seq = seq;
    }

    public void run(){
        System.out.println(this.seq + " thread start. "); 
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
        }
        System.out.println(this.seq + " thread end."); 
    }
    public static void main(String[] args){
        ArrayList<Thread> threads = new ArrayList<>();

        IntStream.range(0, 10).forEach(i -> {
                    Thread t = new Thread(new Sample(i)); //여기수정!
                    t.start();
                    threads.add(t);
        });

  
        threads.stream().forEach(thread ->{
            try {
                thread.join();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        System.out.println("main end.");
    }
}
```
Thread를 extends하던 것에서 Runnable을 Implements하도록 변경되었다.

**※ Runnable인터페이스는 run메서드 구현을 강제한다.

Thread의 생성자로 Runnable 인터페이스를 구현한 객체를 넘길 수 있는데 이 방법을 사용한 것이다. 이렇게 변경된 코드는 이 전에 만들었던 예제와 완전히 동일하게 동작한다. 다만 인터페이스를 이용했으니 상속 및 기타 등등에서 좀 더 유연한 프로그램으로 발전했다고 볼 수 있겠다.




