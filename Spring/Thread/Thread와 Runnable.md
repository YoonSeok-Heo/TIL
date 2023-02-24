# Thread와 Runnable
#JAVA #멀티쓰레드 #Thread #Runnable

자바 초기부터 동시성 프로그래밍을 위해 만들어졌던 Thread와 Runnable를 살펴보도록 하자.

## 쓰레드와 멀티 쓰레드
쓰레드는 프로그램 실행의 가장 작은 단위이다. 일반적인 자바 어플리케이션은 한개의 메인 쓰레드에서 프로그램이 실행되기 때문에 동시에 여러 작업이 불가능하다.
여러개의 작업을 동시에 처리하기 위해 자바는 멀티 쓰레드 기반의 동시성 프로그래밍을 계속해서 발전시켜왔다.

|Version|Function|
|---|---|
|Java5 이전|Runnable과 Thread|
|Java5|Callable, Future 및 Executor, ExcutorService, Executors|
|Java7|Fork/Join 및 Recursive Task|
|Java8|CompletableFuture|
|Java9|\*Flow|

**\*Flow는 reactive프로그래밍에서 사용하나보다?**

## Thread
쓰레드 역시 예제를 통해서 알아보자. 쓰레드는 가장 간단하게 다음과 같이 만들 수 있다.

**Sample.java**
```java
public class Sample extends Thread{

    public void run(){
        System.out.println("thread run. ");
    }
    public static void main(String[] args){
        Sample sample = new Sample();
        sample.start();
    }
}
```
Sample클래스는 Thread클래스를 상속받고, Thread클래스에 run메서드를 구현하면 sample.start()실행시 sample객체의 run메서드가 수행된다.

> ※Thread클래스를 extends했기때문에 start메서드 실행시 run메서드가 수행된다. Thread클래스는 start실행 시 run메서드가 수행되도록 내부적으로 동작한다.

위 Sample예제 실행시 쓰래드가 한개이기때문에 어떻게 동작하는지 알기가 힘들다. 다음 예제를 보자
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
    // Stream 연습중이라 사용해봤다.
    public static void main(String[] args){
        IntStream.range(0, 10).forEach(i -> {
                    Thread t = new Sample(i);
                    t.start();
        });
        System.out.println("main end.");  // main 메서드 종료
    }
}
```
총 10개의 쓰레드를 실행시키는 예제로 어떤 쓰레드인지 확인하기 위해 쓰레드마다 순번을 부여했다. 쓰레드 시작과 끝을 알리는 출력을 하고 시작과 종료 사이에 1초간의 sleep을 주었다.

```console
main end.  <<- main메서드가 먼저 
5 thread start. 
1 thread start. 
3 thread start. 
6 thread start. 
8 thread start. 
2 thread start. 
4 thread start. 
0 thread start. 
9 thread start. 
7 thread start. 
1 thread end.
4 thread end.
5 thread end.
0 thread end.
8 thread end.
6 thread end.
3 thread end.
2 thread end.
7 thread end.
9 thread end.
```
대략 위와 같은 결과가 나오는데 결과에서 보이듯이 for문을 0부터 9까지 실행했는데 순서가 제각각으로 나오는 것을 확인할 수 있다.

Thread는 쓰레드 생성을 위해 Java에서 미리 구현해둔 클래스이다. Thread는 기본적으로 다음과 같은 메소드들을 제공한다. 순서에 상관없이 동시에 실행된다는 뜻이다. 가장 마지막에 실행되는 "main end."의 경우 쓰레드가 실행되기도 전에 출력된다. 쓰레드가 종료되기 전에 main메서드가 종료되었다는 사실이다.

## join
위 예제에서 쓰레드가 모두 수행되고 종료되기 전에 main메서드가 먼저 종료된다. 그렇다면 모든 쓰레드가 종료된 후 main메서드를 종료시키고 싶은경우에는 어떻게 해야할까?
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
생성된 쓰레드를 담기 위해 ArrayList객체인 threads를 만들고 쓰레드 생성시 생성된 객체를 threads에 저장. 그리고 main메서드가 종료되기 전에 threads에 담긴 각각의 thread에 join메서드를 호출하여 쓰레드가 종료될 때까지 대기하도록 했다. 쓰레드의 join메서드는 쓰레드가 종료될 때까지 기다리게 하는 메서드이다. 

위와같이 join을 추가한 후 수행하면 다음과 비슷한 결과가 나온다.
```console
7 thread start. 
5 thread start. 
2 thread start. 
6 thread start. 
4 thread start. 
1 thread start. 
3 thread start. 
0 thread start. 
8 thread start. 
9 thread start. 
2 thread end.
7 thread end.
6 thread end.
3 thread end.
4 thread end.
5 thread end.
0 thread end.
1 thread end.
8 thread end.
9 thread end.
main end. <<- 스레드가 모두 종료되고 실행된다.
```
"main end."가 가장 마지막에 출력되는 것을 확인 할 수 있다.

쓰레드 프로그래밍시 가장 많이 실수하는 부분이 쓰레드가 종료되지 않았는데 다음 로직을 수행하게 만드는 일이다.!! join 메서드를 기억하자.
