# Callable과 Future
#Callable #JAVA #멀티쓰레드 #Future

Callable은 JAVA5에서 Runnable의 한계점(결과 반환, Exception 발생)을 보완해서 나온 인터페이스이다.

Callable은 Runnable과 같이 멀티쓰레드 동작을 구현한다는 공통점이 있지만 다음과 같은 차이점이 있다.

[Runnable에 대하여 알고싶다면 클릭!! 이 문서를 읽기전에 읽는걸 추천](https://github.com/YoonSeok-Heo/TIL/blob/main/JAVA/Thread/Runnable.md)
- Runnable
  - 어떤 객체도 리턴하지 않는다
  - Exception을 발생시키지 않는다.
- Callable
  - 특정 타입의 객체를 리턴
  - Exception을 발생시킬 수 있다.
