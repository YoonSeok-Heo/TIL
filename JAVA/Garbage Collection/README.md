# Garbage Collection

솔루션에서 G1GC를 사용하는데 GC과정에서 서비스 지연이 발생해 GC를 공부하기로 마음먹었다...
GC를 구글링하다보면 대부분의 포스트에서 GC튜닝은 최후의 보루로 사용하라고해서 지금 들어가있는 솔루션의 GC를 튜닝할 생각은없다.
그래도 알아두면 좋겠지 마음에 공부해본다.

---

## GC란?

JVM은 Garbage Collector가 존재한다.
가비지 컬렉터는 동적으로 할당했던 메모리 영역중 더 이상 참조되지 않는 불필요한 메모리를 알아서 정리하는 역할을 한다.
가비지 컬렉션은 메모리 누수를 방지하기 위해 주기적으로 불필요한 메모리를 정리해 주는 기능이다.

> 동적 할당 메모리 영역은 런타임시에 사용되는 Heap영역 메모리를 뜻한다.

c나 c++과 같은 언어에서는 개발자가 직접 메모리 할당과 해제를 컨트롤한다.(ex. malloc, free와 같이)
메모리를 관리한다는건 생각보다 디테일하고 복잡한 과정인데, 
Garbage Collector는 이를 상황에 맞게 해준다.

> 물론 개발자가 원하지 않게 동작 할 수도 있을것이다. 나처럼...

---

## Garbage Collection(GC) 용어

### - stop the world(STW):
  - GC를 실행하기 위해 JVM이 애플리케이션을 멈춘다.
    STW가 발생하면 **GC를 실행하는 스레드를 제외한 나머지 스레드는 모두 작업을 멈춘다.**
### - Mark: 
  - 사용되는 메모리와 **사용되지 않는 메모리를 식별**한다.
### - Sweep: 
  - Mark 단계에서 사**용되지 않음으로 식별된 메모리를 헤제**한다.
### - Young 영역(young region, young generation): 
  - 새롭게 생성된 객체가 대부분이 할당되는 영역
  - 대부분의 객체가 금방 접근불가 상태가 되므로 많은 객체가 Young영역에 생성되었다 사라진다.
  - **Young영역애 대한 GC를 Minor GC라고 부른다.**
### - Old 영역(old region, old generation):
  - Young영역에서 접근가능한 상태를 유지해 살아남은 객체가 복사 또는 이동하는 영역
  - 복사되는 과정에서 대부분 Young영역보다 크게 할당되며, 크기가 큰 만큼 GC는 적게 발생한다.
  - **Old 영역에 대한 GC를 Major GC 또는 Full GC라고 부른다.**
  > 영역의 크기가 GC와 무슨 상관이 있는지 정확히 모르겠다. 이건 좀 더 알아봐야겠다.

---

## Young 영역

객체가 제일 먼저 생성되는 Young영역부터 알아보자.
Young영역은 3개의 영역으로 나누어진다.

- Eden 영역
- Survivor 영역(2개)

Survivor 영역은 2개로 나누어지므로, 총 3개의 영역으로 나누어진다고 볼 수 있다. 

각 영역의 처리절차는 다음과 같다.

1. 새로 생성된 대부분의 개체는 Eden영역에 위치한다.
2. Eden 영역에서 GC가 한 번 발생한 후 **살아남은 객체는 Survivor영역 중 하나**로 이동된다.
3. Eden 영역에서 GC가 발생하면 이미 살아남은 객체가 존재하는 Suvivor영역으로 계속 쌓인다.
4. 하나의 Survivor 영역이 가득 차게 되면 그 중 살아남은 객체를 **다른 Survivor영역으로 이동**하고, 가득찬 Suvivor영역은 아무 데이터도 없는 상태로 된다.
5. 이 **과정을 반복해도 계속해서 살아남아 있는 객체를 Old 영역으로 이동**시키게 된다.

위 절차에서 알겠지만 Survivoer 영역 중 하나는 반드시 비어 있는 상태로 남아있어야 한다. 
만약 두 Survivor영역에 모두 데이터가 존재하거나, 
두 영역 모두 사용량이 0이라면 정상적인 상황이 아니라고 생각해도 된다.

---

## Old 영역

Old 영역은 기본적으로 데이터가 가득 차면 GC를 실행한다. 
GC방식에 따라 처리 절차가 다르며, GC방식은 다양하게 있다.

1. Serial GC
2. Parallel GC
3. Parallel Old GC
4. Concurrent Mark & Sweep GC(CMS)
5. G1(Garbage Frist) GC

요즘은 G1GC를 많이 사용하며 JAVA9 부터는 JVM의 기본 GC이다.

각각의 GC의 설명은 위 목록을 눌러서 확인.

---


### Minor GC를 통해 Old영역까지 데이터가 쌓인 것을 표현한 이미지 이다.

[![이미지 클릭 시 원본으로 이동](https://github.com/YoonSeok-Heo/TIL/assets/113662725/c9958110-8733-4857-8c26-177b4547429f)](https://d2.naver.com/helloworld/1329)
<br>이미지 클릭 시 출처로 이동

---

## 기본적인 GC의 동작 과정

Young영역과 Old영역은 서로 다른 메모리 구조로 되어 있기 때문에 세부적인 동작 방식은 다르지만, 
기본적으로 가비지 컬렉션이 실행되면 2가지 공통적인 단계를 따르게 된다.

1. Stop The World
2. Mark And Sweep

### 1. Stop The World(STW)

Stop The World는 가비지 컬렉션을 실행하기 위해 **JVM이 애플리케이션 실행을 멈춘다.** 
GC가 실행될 때는 **GC를 실행하는 쓰레드를 제외한 모든 쓰레드의 작업이 중단되고, 
GC가 완료되면 모든 쓰레드의 작업이 재개된다.**

GC를 실행하는 쓰레드 외의 모든 쓰레드의 작업이 중단되면서 애플리케이션이 멈추기 때문에, 
GC의 성능 개선을 위해 튜닝한다는 말은 보통 STW의 시간을 줄이는 작업을 의미한다.

> 나도 STW 시간이 3초이상 나오는 경우가 있어 실시간으로 처리하는 과정에서 작업들이 쌓이는 현상이 발생하고 있다.

### 2. Mark and Sweep 

- Mark: 사용되는 메모리와 사용되지 않는 메모리를 식별하는 작업
- Sweep: Mark 단계에서 사용되지 않음으로 식별된 메모리를 헤제하는 작업이다.

STW를 통해 모든 작업을 중단시키면, 
GC는 스택의 모든 변수 또는 Reachable객체를 스캔하면서 각각이 어떤 객체를 참조하는지 탐색한다.
**사용되고 있는 메모리를 식별**하는데, 이러한 과정을 Mark라고 한다.
이후에 **Mark가 되지 않은 객체들을 메모리에서 제거**하는데,
이 과정을 Sweep이라고 한다.


