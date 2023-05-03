# Stream의 foreach와 for-loop는 다르다.

Java 8부터는 Stream과 lambda를 제공하는데 반복적인 일을 처리하기 위해 Stream을 공부하고 있다.

나의 궁금점은 Collections.forEach문과 Stream.forEach문의 차이에서부터 왔다.

```java
// Collections
Arrays.asList(files).forEach(file -> {
        ...
});

// Stream
Arrays.asList(files).stream().forEach(file -> {
        ...
});
```
두개의 차이점이다. 이 부분은 나중에 알아보도록 하고 오늘은 Stream과 for-loop에 대해서 알아보도록하자.

### ※모든 반복문을 stream으로 구현하려고 했다면 그러지 않는게 좋을것이다.!

## 모든 반복문에 Stream을 사용하면 안되는 이유??

함수형 프로그래밍이 가독성도 좋아진다. 다음 예를 보자

```java
//기존의 for-loop
for (int i = 0; i < list.size(); i++) {
  System.out.println(list.get(i));
}

//향상된 for-Each
for (String item : list) {
  System.out.println(item);
}

//stream.forEach()
list.stream().forEach(System.out::println);
```
위 코드만 보아도 stream을 사용한 코드는 한줄로 간편하고 가독성도 좋아졌다. 

하지만 모든 요소를 순회하는 stream.forEach() 사용에 대해서는 신중하게 생각헤 볼 필요가있다. 여러 페이지를 구글링 해봐도 stream().forEach()는 출력문을 사용할 때가 아니면 사용하지 말라고한다.

## 로직이 있는 Stream.forEach

Stream내부에 로직이 있는 경우 잘 활용하고 있는지 확인해보자.

### 종료로직이 있는 Stream.forEach

stream에서 반복문을 강제적으로 종료시키는 방법이 없다. 종료 조건이 있는 반복문을 stream으로 구현한다면 필요하지 않는 구간까지 반복하는 비효율이 발생할 것이다.

```java
//for-loop로 짠 경우
for (int i = 0; i < 100; i++) {
    if (i > 50) {
      break;
      //50번 돌고 반복을 종료한다.
    }
    System.out.println(i);
}

IntStream.range(1, 100).forEach(i -> {
    if (i > 50) {
      return;
      //각 수행에 대해 다음 수행을 막을 뿐, 100번 모두 조건을 확인한 후에야 종료한다.
    }
    System.out.println(i);
});
```

> Stream.forEach()문에서 break와 비슷한 효과를 얻을 수 있는 방법이 있긴하다 lambda식 내에서 Exception을 날리는 방법이다.!

### Stream.forEach의 사용법?? 

아래 코드와 같이 Filter나 map과 같이 다른 연산을 이용하고 마지막으로 forEach로 넘어간뒤 데이터를 확인하는 정도로만 사용하는것이 좋다
```java
IntStream.range(1, 100)
        .filter(i -> i <= 50)
        .forEach(System.out::println);
```

## 결론

Stream은 직관적이면서 코드 라인수도 줄일 수 있고 lambda와 함께 사용하며, 쉽게 병렬처리도 가능하다는 장점이 있지만 단순히 for, while문의 대체재가 아니다.

상황에 맞게 적절히 선택해서 사용해야한다.