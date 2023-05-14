# JVM 메모리 구조

다른 언어를 접하면서 그 언어만의 메모리 관리를 알아보다가 기존에 사용하던 JAVA도 정리해야겠다는 생각이 들어 정리해본다.

## JVM이란?

- Java Virtual Machine의 약자이며, 자바 가상 머신이라고 부른다. 
- 자바 바이트코드를 직접 실행하는 주체이다.
- CPU나 OS에 관계없이 실행이 가능하도록 도와준다.
- 메모리 관리를 자동으로 해준다(Garbage Collection)

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/8ed11c3c-6c81-4869-8239-490c98d826a5)

- **Java Compiler**: 사용자가 작성한 코드(.java)를 **Byte code(.class)로 변환**시킨다.
- **Class Loader**: 변환된 Byte code(.class)를 **JVM이 할당받은 메모리 영역으로 적재하는 역할**을 한다.
- **Execution Engine**: Class Loader를 통해 JVM에 적재된 Byte code를 **명령어 단위로 실행**시킨다.
- **Garbage Collector**: 어플리케이션이 생성한 객체의 생존 여부를 판단하여 더이상 참조되지 않거나 null인 **객체의 메모리를 반납**한다.
- **Runtime Data Area**: JVM의 메모리 영역으로 자바 어플리케이션을 실행할 때 사용되는 **데이터들을 적재하는 영역**이다.

> Java Compiler, Execution Engine, Garbage Collector는 따로 포스트를 작성해야겠다. 이번 포스트에서는 Runtime Data Area에 대해서 자세히 알아본다.

--- 

## Runtime Data Area

JVM의 메모리 영역이며, 애플리케이션을 실행할 때 사용하는 데이터들이 적제된다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/256054f1-5c07-4cf4-8687-495045655539)

### 1. Method Area(메소드 영역)

정적 영역(Static Area)이라고도 하며 Byte code가 로드되는 곳이다. 

- JVM이 무언가를 실행하려면 해당 바이트 코드들이 메모리 공간에 적재되어야 한다. 
- 클래스와 인터페이스의 메타 데이터들을 저장된다.
  - Type정보(Interface인지 class인지 등)
  - Constant Pool(상수 풀 : 문자 상수, 타입, 필드, 객체 참조가 저장됨)
  - static 변수, final class 변수 등

### 2. Heap Area(힙 영역)

new 키워드로 생성된 객체와 배열이 생성되는 영역이다.(new: 동적으로 생성된 인스턴스)

- Heap Area에 생성된 객체들은 다른 객체 필드 또는 스택에 존재하는 다른 메서드에 의해 참조될 수 있다.
- 메소드가 실행되면 Stack Area에는 참조값만 저장하고 Heap Area에 객체 데이터를 저장한다.
- 메소드 영역에 로드된 클래스만 생성이 가능하고 Garbage Collector가 참조되지 않는 메모리를 확인하고 제거하는 영역이다.

> ex. 배열이나 리스트를 선언하게 되면 힙영역으로 적재된다.

### 3. Stack Area(스택 영역)

Stack Area는 메서드가 호출될 시 할당되는 영역이다.

- 지역 변수, 파라미터, 리턴 값 등 임시로 사용하는 값들이 생성되는 영역이다.
- Thread를 사용하여 프로그래밍 할 경우 각 Thread는 한 개의 Stack영역을 할당 받는다. 즉, 스레드는 각자의 메모리 공간을 가지고 메서드를 수행한다.

### 4. PC Register(PC 레지스터)

Thread가 생성될 때마다 생성되는 영역으로 Program Counter 즉, 현재 쓰레드가 실행되는 부분의 주소와 명령을 저장하고 있는 영역이다. (*CPU의 레지스터와 다름)

### 5. Native Method Stack

자바 외 언어로 작성된 네이티브 코드를 위한 메모리 영역입니다.