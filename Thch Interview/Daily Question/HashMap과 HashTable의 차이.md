# HashMap과 HashTable의 차이

## HashMap

- HashMap은 동기화를 지원하지 않는다.

  - 단일스레드 환경에서 사용하기 좋은 자료구조

IDE의 구현부를 보면, 아래 Hashtable과는 다르게 synchronized 키워드가 없음

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

또, HashMap은 key 값이나 value 값에 null이 들어갈 수 있음
```java
Map<Integer, String> map = new HashMap<>();
map.put(null, "null");

System.out.println(map.get(null)); // null
```
