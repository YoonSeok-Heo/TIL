# ArrayList vs LinkedList

자주 사용하는 List인 ArrayList와 LinkedList 두 개를 비교해 보자

ArrayList와 LinkedList는 모두 List 인터페이스를 구현하고 삽입 순서를 유지하며, 둘 다 동기화되지 않은 클래스이다.

## ArrayList와 LinkedList의 차이점

|   | `ArrayList` | `LinkedList` |
|---|-------------|--------------|
| 내부 구조 | 동적 배열 | 이중 연결 리스트 |
| 데이터 접근 시간 | O(1) | O(n) |
| 데이터 삽입/삭제 시간 (리스트의 시작점) | O(n) | O(1) |
| 데이터 삽입/삭제 시간 (리스트의 중간 또는 끝) | O(n) | O(1) |
| 메모리 사용량 | 적은 추가 공간 필요 | 추가 메모리 할당 필요 없음 |
| 데이터 접근 패턴 | 읽기 작업이 많은 경우 | 삽입/삭제 작업이 많은 경우 |
| 반복 작업 성능 | 빠름 | 느림 |
| 구현 인터페이스 | `List`, `RandomAccess` | `List`, `Queue` |

위 표는 일반적인 시나리오에서의 `ArrayList`와 `LinkedList`의 특징을 보여주고 있다.
이는 모든 경우에 해당하는 절대적인 규칙은 아니며, 
작업의 종류와 크기, 
데이터 구조의 특성에 따라 성능이 달라질 수 있다. 

> 삽입/삭제가 많은경우 LinkedList를 조회가 많은경우 ArrayList를 
> 
> 위 속도가 다른 이유는 내부 구조에 따라 데이터에 접근하는 방식이 다르기 때문에 속도 차이를 보이게 된다.

**따라서 어떤 데이터 구조를 선택할지는 사용 사례와 요구 사항에 따라 결정해야 한다.**

## 사용 예제

```java
import java.util.*;    
class TestArrayLinked{ 
    public static void main(String args[]){    
        List<String> al=new ArrayList<String>();//creating arraylist
        al.add("Ravi");//adding object in arraylist    
        al.add("Vijay");    
        al.add("Ravi");    
        al.add("Ajay");    
        
        List<String> al2=new LinkedList<String>();//creating linkedlist    
        al2.add("James");//adding object in linkedlist    
        al2.add("Serena");    
        al2.add("Swati");    
        al2.add("Junaid");    
    
        System.out.println("arraylist: "+al);  
        System.out.println("linkedlist: "+al2);  
    }    
}  
```