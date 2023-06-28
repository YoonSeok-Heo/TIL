# Garbage Collection

솔루션에서 G1GC를 사용하는데 GC과정에서 서비스 지연이 발생해 GC를 공부하기로 마음먹었다...
GC를 구글링하다보면 대부분의 포스트에서 GC튜닝은 최후의 보루로 사용하라고해서 지금 들어가있는 솔루션의 GC를 튜닝할 생각은없다.
그래도 알아두면 좋겠지 마음에 공부해본다.

---

## GC란?

JVM은 Garbage Collector가 존재한다.
가비지 컬렉터는 동적으로 할당했던 메모리 영역중 더 이상 참조되지 않는 불필요한 메모리를 알아서 정리하는 역할을 한다.
가비지 컬렉터가 메모리 누수를 방지하기 위해 주기적으로 불필요한 메모리를 정리해 주는 기능이다.

> 동적 할당 메모리 영역은 런타임시에 사용되는 Heap영역 메모리를 뜻한다.

c나 c++과 같은 언어에서는 개발자가 직접 메모리 할당과 해제를 컨트롤한다.(ex. malloc, free와 같이)
메모리를 관리한다는건 생각보다 디테일하고 복잡한 과정인데. 
Garbage Collector는 이를 상황에 맞게 해준다.

---

## Garbage Collection(GC) 용어

- stop the world(STW): GC를 실행하기 위해 JVM이 애플리케이션을 멈춘다.
STW가 발생하면 GC를 실행하는 스레드를 제외한 나머지 스레드는 모두 작업을 멈춘다.
- Young 영역(young region, young generation): 
  - 새롭게 생성된 객체가 할당되는 영역
  - 대부분의 객체가 금방 Unreachable 상태가 되므로 많은 객체가 Young영역에 생성되었다 사라진다.
  - **Young영역애 대한 GC를 Minor GC라고 부른다.**
- Old 영역(old region, old generation):
  - Young영역에서 Reachable 상태를 유지해 살아남은 객체가 복사 또는 이동하는 영역
  - 복사되는 과정에서 대부분 Young영역보다 크게 할당되며, 크기가 큰 만큼 가비지는 적게 발생한다.
  - **Old 영역에 대한 GC를 Major GC 또는 Full GC라고 부른다.**

---

##