# Queue

큐는 스택과 다르게 **선입선출**의 구조를 가진다. 번호표와 같은 구조라고 생각하면 된다. 먼저 번호표를 뽑은사람부터 업무가 처리된다. 이러한 구조가 선입선출의 큐이다.

## 큐(Queue)의 동작

- Front: 가장 앞에 있는 원소
- Rear: 가장 뒤에 있는 원소

- enQueue(): 큐에 데이터를 넣는다
  - Rear값을 +1 해주고 Rear가 가르키는 자리에 데이터를 넣는다.
- deQueue(): 큐에서 데이터를 뺀다.
  - Front가 가르키는 자리의 값을 빼고 Front의 값을 +1 해준다.
- isEmpty(): 큐가 비어있는지 확인한다.
- isFull(): 큐가 꽉 차 있는지 확인한다.
- peek(): 앞에있는 원소를 삭제하지 않고 반환한다.

## 구현

### Golang

slice로 구현되어 있어서 Front, Rear개념이 딱히 필요없어서 간단하게만 만들었다.

```go
package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enQueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) deQueue() int {
	retVal := q.items[0]
	q.items = q.items[1:]
	return retVal
}

func (q *Queue) isEmpyty() bool {
	return len(q.items) == 0
}

// 큐의 배열을 slice로 구현했기 때문에 isFull은 구현하지 않았다.
//func (q *Queue) isFull() bool{
//	return true
//}

func (q *Queue) peek() int {
	return q.items[0]
}

func main() {

	fmt.Println("Queue")
	queue := Queue{make([]int, 0)}
	fmt.Println("enQueue: 1")
	(&queue).enQueue(1)
	fmt.Println("enQueue: 2")
	(&queue).enQueue(2)
	fmt.Println("enQueue: 3")
	(&queue).enQueue(3)
	fmt.Println("queue의 원소들:", (&queue).items)
	fmt.Println("peek:", (&queue).peek())
	fmt.Println("deQueue:", (&queue).deQueue())
	fmt.Println("peek:", (&queue).peek())
	fmt.Println("deQueue:", (&queue).deQueue())
	fmt.Println("deQueue:", (&queue).deQueue())
	fmt.Println("IsEmpty:", (&queue).isEmpyty())

}
```


## 선형 큐의 단점

- 큐에 데이터가 꽉 차있으면 데이터를 더 추가할 수 없다.
  - 큐의 사이즈를 늘리는 방법이 있으나 새로운 큐를 생성하고 기존 데이터를 복사하는 과정이 오래걸리므로 시간 효율 측면에서 좋지 못하다.
- 단순 배열로 구현한다면 공간 효율성이 좋지 못하다.
  - enQueue, deQueue 과정에서 front가 계속 뒤로 밀리게 되고 앞쪽의 공간이 남게된다.
  - 데이터를 앞쪽으로 당겨오게 된다면 모든 원소를 옮겨줘야하므로 속도 측면에서 상당히 느리다.