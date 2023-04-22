# Iterator vs foreach

---

## Iterator

- Iterator는 Collection framwork안에 들어있는 인터페이스이다.
- 컬렉션에서 탐색, 데이터 엑세스, 데이터 제거를 할 수 있다.

### Iterator method

java.util.Iterator의 3개 메소드

1. boolean hasNext(): Iterator에 반복할 요소가 더 있으면 true를 반환한다.
2. Object next(): hasNext()가 참을 반환하면 컬렉션의 다음 요소를 반환한다. 만약 다음 요소가 없는 경우 'NoSuchElementException'을 반환한다.
3. void remove(): 컬렉션의 현재 요소를 제거한다. 이 메서드는 next()가 호출되기 전에 이 함수가 호출되면 'IllegalStateException'을 던진다.

```java
import java.util.ArrayList;
import java.util.Iterator;

public class Sample {

    public static void main(String[] args)
    {
        ArrayList<String> list = new ArrayList<String>();

        list.add("A");
        list.add("B");
        list.add("C");
        list.add("D");
        list.add("E");

        // Iterator to traverse the list
        Iterator iterator = list.iterator();

        System.out.println("List elements : ");

        while (iterator.hasNext()) {
            Object element = iterator.next();
            System.out.print((String) element + " ");
            // iterator.remove()
            if ("B".equals((String) element)){
                iterator.remove();
            }
        }
        System.out.println(); // 개행
        System.out.println("다음 요소 여부: " + iterator.hasNext());
        System.out.println("B 삭제 확인: " + list);
    }
}
```
```console
출력확인

List elements : 
A B C D E 
다음 요소 여부: false
B 삭제 확인: [A, C, D, E]
```
hasNext()를 이용해 다음 요소가 있는지 확인하고 next를 이용해 요소를 가져온다.

remove는 next를 이용해 가져온 값을 리스트에서 삭제한다.

---
## for each

- for each문은 자바 1.5에 추가되었다. for each라는 키워드가 있는건 가니고 for문의 조건식 부분이 다르다.
- 컬렉션의 아이템을 처음부터 끝까지 탐색하거나, 에외가 발생할 때까지 탐색을 수행한다.
- 즉, 1 ~ 10까지의 정수값을 가진 List에서 for-each문으로 1 ~ 3까지만 탐색하는건 기본동작에 없고, if문과 같은 조건문으로 반복문을 빠져나와야한다.


## 기존 for문

```java
String[] alpha = {"A", "B", "C"};
for(int i = 0; i < numbers.length; i++) {
    System.out.print(alpha[i] + " ");
}
```
```console
출력확인

A B C
```

## Enhanced for 문(for each)

```java
String[] alpha = {"A", "B", "C"};
for(String s : alpha) {
    System.out.print(s + " ");
}
```
```console
출력확인

A B C
```

기존의 for문과 foreach문은 같은 동작을 하지만 조건부가 훨씬 간단한 동작을 하는것을 알 수 있다.

단, for eachn문은 따로 반복회수를 명시적으로 주는 것이 불가능해 1스탭씩 순차적으로 반복할 때만 사용가능하다.

---
## 차이점
1. for each문에서는 크기 확인이 필요하지 않다. 하지만 Iterator에서는 hasNext()를 올바르게 사용해서 NoSuchElementException이 발생하는걸 막아줘야한다.
2. for each문에서는 객체가 변경되거나 삭제되면 ConcurrentModificationException이 발생할 수 있다.
3. 성능은 비슷하다.

### 속도
2억번의 덧셈을 실행해보았다. 속도는 항상 비슷하다.
```console
for each 시간차이(ms) : 312
for each 시간차이(ms) : 302
for each 시간차이(ms) : 238
for each 시간차이(ms) : 238
for each 시간차이(ms) : 247
for each 시간차이(ms) : 257
for each 시간차이(ms) : 242
for each 시간차이(ms) : 238
for each 시간차이(ms) : 250
for each 시간차이(ms) : 240

iter 시간차이(ms) : 298
iter 시간차이(ms) : 280
iter 시간차이(ms) : 247
iter 시간차이(ms) : 245
iter 시간차이(ms) : 244
iter 시간차이(ms) : 252
iter 시간차이(ms) : 254
iter 시간차이(ms) : 247
iter 시간차이(ms) : 239
iter 시간차이(ms) : 244
```

---
## 결론

### for each를 주로 사용하자
- 현재 자바에서는 될 수 있으면 JDK 1.5부터 추가된 Enhanced for 문을 사용하도록 권장하고 있다.
- 같은 성능을 유지하면서도 코드의 명확성이 확보된다.
- Iterator의 경우 boilerplate code가 많이 들어간다.(코드가 간결해짐)

### Iterator를 사용할 경우
- 배열이나 리스트내부의 객체를 수정 또는 제거할 경우 Iterator를 사용하자.

