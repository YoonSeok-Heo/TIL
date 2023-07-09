# HashSet


`HashSet`은 Java 컬렉션 프레임워크에서 제공하는 클래스 중 하나이다. 
HashSet은 중복 요소를 허용하지 않고, 요소의 순서를 보장하지 않는 데이터 구조이다. 
`HashSet`은 해시 테이블을 기반으로 구현되어 있으며, 매우 빠른 검색, 삽입 및 삭제 작업을 제공합니다.

## HashSet의 특징

### 1. 중복이 허용되지 않는다.
- `HashSet`은 동일한 요소를 여러번 포함할 수 없다.
- 동일한 객체를 두 번 이상 추가하려고 하면 무시된다.

### 2. 요소의 순서를 보장하지 않는다.
- `HashSet`은 요소의 순서를 유지하지 않는다.
- 요소가 추가된 순서와는 관계없이 요소를 반복할 때 항상 다른 순서로 반환된다.

### 3. 빠른 검색 속도
- `HashSet`은 해시 테이블로 구현되어 있으므로 내부적으로 각 요소에 대한 해시 코드를 사용하여 요소를 저장하고 검색한다.
- 해시 코드를 이용하기 때문에 매우 빠른 검색 속도를 제공한다.

### 4. Null요소 허용
- `HashSet`은 null요소를 허용한다.
- 즉, null을 `HashSet`에 추가할 수 있고, null요소를 포함한 요소 집합을 만들 수 있다.
- null 또한 중복 추가가 불가능하다.

---

## 사용법 

```java
HashSet<String> set = new HashSet<>();

// 요소 추가
set.add("Apple");
set.add("Banana");
set.add("Orange");

// 요소 검색
boolean containsApple = set.contains("Apple"); // true

// 요소 삭제
set.remove("Banana");

// 크기 확인
int size = set.size(); // 2

// 모든 요소 반복
for (String element : set) {
    System.out.println(element);
}

// HashSet이 비어 있는지 확인
boolean isEmpty = set.isEmpty(); // false
```

> 순서가 보장되지 않은 상태로 출력한다.