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

## Thread 클래스
Thread는 쓰레드 생성을 위해 Java에서 미리 구현해둔 클래스이다. Thread는 기본적으로 다음과 같은 메소드들을 제공한다.

- sleep
  - 현재 쓰레드 멈추기
  - 자원을 놓아주지는 않고, 제어권을 넘겨주므로 데드락이 발생할 수 있음
- interupt
  - 다른 쓰레드를 깨워서 interruptedException 을 발생시킴
  - Interupt가 발생한 쓰레드는 예외를 catch하여 다른 작업을 할 수 있음
- join
  - 다른 쓰레드의 작업이 끝날 때 까지 기다리게 함
  - 쓰레드의 순서를 제어할 때 사용할 수 있음

Thread클래스로 쓰레드를 구현하려면 이를 상속받는 클래스를 만들고 run메소드를 구현해야한다. 그 다음 start메소드를 호출하면 구현한run메소드가 실행된다.
