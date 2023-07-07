# Collections

Java 컬렉션은 객체 그룹을 저장하고 조작하는 아키텍처를 제공하는 프레임워크이다.

컬렉션 프레임워크에서는 검색, 정렬, 조작, 삭제 등 데이터에 대해 수행하는 모든 작업을 수행할 수 있으며,
다양한 인터페이스(Set, List, Queue, Deque)와 클래스(ArrayList, Vector, HashSet 등)을 제공한다.

 
## 컬렉션 프레임워크의 계층 구조

java.util 패키지에는 Collection프레임워크에 대한 모든 클래스와 인터페이스가 포함되어 있다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/6dad965f-a96e-4924-b3bc-59a813d43b73)


## Methods of Collections interface 

java.util.Collections 인터페이스는 많은 메서드가 있다. 확인해보자

| 메소드                                      | 접근 제어자 | 설명                                       |
| ---------------------------------------- | -------- |------------------------------------------|
| `boolean add(E element)`                 | public   | 지정된 요소를 컬렉션에 추가한다.                       |
| `boolean addAll(Collection<? extends E> collection)` | public   | 지정된 컬렉션의 모든 요소를 현재 컬렉션에 추가한다.            |
| `void clear()`                            | public   | 컬렉션에서 모든 요소를 제거한다.                       |
| `boolean contains(Object element)`        | public   | 컬렉션이 지정된 요소를 포함하고 있는지 여부를 반환한다.          |
| `boolean containsAll(Collection<?> collection)`    | public   | 컬렉션이 지정된 컬렉션의 모든 요소를 포함하고 있는지 여부를 반환한다.  |
| `boolean equals(Object object)`           | public   | 지정된 객체와 컬렉션이 동일한지 여부를 반환한다.              |
| `int hashCode()`                          | public   | 컬렉션의 해시 코드를 반환한다.                        |
| `boolean isEmpty()`                       | public   | 컬렉션이 비어 있는지 여부를 반환한다.                    |
| `Iterator<E> iterator()`                  | public   | 컬렉션을 반복하는 데 사용되는 이터레이터를 반환한다.            |
| `boolean remove(Object element)`          | public   | 컬렉션에서 지정된 요소를 제거한다.                      |
| `boolean removeAll(Collection<?> collection)`      | public   | 컬렉션에서 지정된 컬렉션의 모든 요소를 제거한다.              |
| `boolean retainAll(Collection<?> collection)`       | public   | 컬렉션에서 지정된 컬렉션에 포함된 요소만 남기고 나머지 요소를 제거한다. |
| `int size()`                              | public   | 컬렉션의 크기(요소의 개수)를 반환한다.                   |
| `Object[] toArray()`                      | public   | 컬렉션의 모든 요소를 포함하는 배열을 반환한다.               |
| `<T> T[] toArray(T[] array)`               | public   | 컬렉션의 모든 요소를 지정된 배열에 저장하고 해당 배열을 반환한다.    |

`Collections` 인터페이스의 메소드들은 모두 `public`접근 제어자를 가지고 있습니다. 이는 해당 메소드들이 인터페이스를 구현하는 클래스에서 모두 접근 가능하다는 것을 의미한다.

## Methods of Collection interface

위에서 제공한 표에는 `Collections` 인터페이스에만 포함된 메소드들이 포함되어 있으며, `stream()` 및 `spliterator()`와 같은 메소드는 `Collection` 인터페이스의 하위 인터페이스인 `java.util.Collection`에서 제공된다.

| 메소드                                      | 접근 제어자 | 설명                          |
| ---------------------------------------- | -------- |-----------------------------|
| `Stream<E> stream()`                      | default  | 컬렉션을 스트림으로 변환하여 반환한다.       |
| `default Spliterator<E> spliterator()`    | default  | 컬렉션의 Spliterator를 반환한다.     |
| `default boolean removeIf(Predicate<? super E> filter)` | default  | 지정된 조건에 맞는 요소들을 컬렉션에서 제거한다. |
| `default Stream<E> parallelStream()`      | default  | 컬렉션을 병렬 스트림으로 변환하여 반환한다.    |

`stream()` 메소드는 기본(default) 접근 제어자를 가지며, `spliterator()`, `removeIf()`, `parallelStream()` 메소드들 또한 마찬가지이다. 
이들 메소드들은 `Collection` 인터페이스의 디폴트 메소드로 정의되어 있으므로, 
구현 클래스에서 직접 호출할 수 있다.

참고로, `Stream` 및 `Spliterator`는 Java 8에서 추가된 기능으로, 
컬렉션 요소들을 효과적으로 처리하기 위한 강력한 기능을 제공한다.

## Methods of Iterator interface

아래에는 Java의 `Iterator` 인터페이스에 정의된 메소드들을 표로 나타낸 것이다.

| 메소드                 | 설명                         |
|-----------------------|----------------------------|
| `boolean hasNext()`   | 다음 요소가 있는지 여부를 확인한다.       |
| `E next()`            | 다음 요소를 반환한다.               |
| `void remove()`       | 컬렉션에서 현재 요소를 제거한다.         |
| `default void forEachRemaining(Consumer<? super E> action)` | 남은 모든 요소에 대해 지정된 작업을 수행한다. |

> 참고로, `forEachRemaining(Consumer<? super E> action)` 메소드는 Java 8에서 추가된 기능으로, 
> `Iterator`에 남아 있는 모든 요소에 대해 지정된 작업을 수행한다. 
> 이를 통해 반복문을 통해 요소를 하나씩 처리하는 대신, 더 간단하고 간결한 코드를 작성할 수 있다.

Iterable 인터페이스는 모든 컬렉션 클래스의 루트 인터페이스이다.
Collection 인터페이스는 Interable인터페이를 확장하므로 Collection인터페이스의 모든 서브 클래스도 Iterable인터페이스를 구현한다.

