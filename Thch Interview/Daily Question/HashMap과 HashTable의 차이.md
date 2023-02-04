# HashMap과 HashTable의 차이

## HashMap

- key나 value에 null값이 들어갈 수 있다.
- HashMap은 동기화를 지원하지 않는다.
  - 단일스레드 환경에서 사용하기 좋은 자료구조

구현부를 보면, 아래 Hashtable과는 다르게 synchronized 키워드가 없음

```java
// get method
public V get(Object key) {
    Node<K,V> e;
    return (e = getNode(hash(key), key)) == null ? null : e.value;
}
    
// put method
public V put(K key, V value) {
    return putVal(hash(key), key, value, false, true);
}
```


## HashTable

Hashtable 클래스는 컬렉션 프레임웍이 만들어지기 이전부터 존재하던 것이기 때문에 컬렉션 프레임워의 명명법을 따르지 않는다.

Vector나 Hashtable과 같은 기존의 컬렉션 클래스들은 호환을 위해, 설계를 변경해서 남겨두었지만 가능하면 사용하지 않는 것이 좋다. (대신 ArrayList와 HashMap을 사용하는 것이 좋다.)

- key나 value에 null값이 들어갈 수 없다.
- Hashtable은 동기화를 지원하여, thread-safe하다.
  - 멀티스레드 환경에서 사용하기 좋은 자료구조
  - 하지만, HashMap에 비해 느리다. (다른 스레드가 block되고 unblock 되는 대기 시간을 기다리기 때문)

구현부를 보면 synchronized 키워드가 붙어있다.
```java
// get method
public synchronized V get(Object key) {
    // ... 중략 ...
}
    
// put method
public synchronized V put(K key, V value) {
    // ... 중략 ...
}
```






