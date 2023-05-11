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

---

Front, Rear 사용해서 구현. 테스트하려고 enqueue, dequeue만 구현했다.

```go
type PointQueue struct {
	items []int
	front int
	rear  int
}

func (q *PointQueue) enQueue(val int) {
	q.items[q.front] = val
	q.front++
}

func (q *PointQueue) deQueue() int {
	retVal := q.items[q.rear]
	q.rear++
	return retVal
}
```

## 속도 테스트를 해봐야지

> ### Slice, Front/Rear
> 
> 그냥 slice를 사용한 큐와 front, rear를 사용한 큐랑 비교해봐야겠다.
> 
> 슬라이스의 용량(Capacity)은 둘다 같은 상태로 시작한다.(데이터가 늘어남에 따라 슬라이스 복사가 발생하면 속도가 확 늘어나기 때문에)
> 
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"time"
> )
> 
> type SliceQueue struct {
> 	items []int
> }
> 
> func (q *SliceQueue) enQueue(val int) {
> 	q.items = append(q.items, val)
> }
> 
> func (q *SliceQueue) deQueue() int {
> 	retVal := q.items[0]
> 	q.items = q.items[1:]
> 	return retVal
> }
> 
> type PointQueue struct {
> 	items []int
> 	front int
> 	rear  int
> }
> 
> func (q *PointQueue) enQueue(val int) {
> 	q.items[q.front] = val
> 	q.front++
> }
> 
> func (q *PointQueue) deQueue() int {
> 	retVal := q.items[q.rear]
> 	q.rear++
> 	return retVal
> }
> 
> func main() {
> 	
> 	startTime := time.Now()
> 	queue := SliceQueue{make([]int, 0, 500000000)}
> 	for i := 0; i < 500000000; i++ {
> 		(&queue).enQueue(i)
> 	}
> 	
> 	for i := 0; i < 500000000; i++ {
> 		(&queue).deQueue()
> 	}
> 	fmt.Println("Slice 실행시간:", time.Since(startTime))
>     fmt.Println("슬라이스 길이:", len((&queue).items), "/ 슬라이스 용량:", cap((&queue).items))
> 
> 	startTime = time.Now()
> 	// 사용할 큐의 크기만큼 미리 할당을 해줘야한다.
> 	pointerQueue := PointQueue{make([]int, 500000000, 500000000), 0, 0}
> 	for i := 0; i < 500000000; i++ {
> 		(&pointerQueue).enQueue(i)
> 	}
> 	for i := 0; i < 500000000; i++ {
> 		(&pointerQueue).deQueue()
> 	}
> 	fmt.Println("Front/Rear 실행시간:", time.Since(startTime))
>     fmt.Println("슬라이스 길이:", len((&pointerQueue).items), "/ 슬라이스 용량:", cap((&pointerQueue).items))
> 
> 
> }
> ```
> 
> 
>> #### 결과
>> ```console
>> Slice 실행시간: 2.3534738s
>> 슬라이스 길이: 0 / 슬라이스 용량: 0
>> Front/Rear 실행시간: 2.2083191s
>> 슬라이스 길이: 500000000 / 슬라이스 용량: 500000000
>> ```
>> **여러번 반복한 결과 후자가 약간 빠르다.(약 5퍼센트)**
> 
>> ### 상황별 장단점.
>> 
>> 전자의 경우 모든 데이터들이 deQueue가 된다면 슬라이스 용량이 0으로 줄어든다. 그말은 즉 다시 데이터를 넣을때 용량을 초과하는 데이터가 들어가면 슬라이스 복사가 이루어지면서 속도가 느려지게된다.
>> 
>> 후자의 경우 큐를 생성할때 길이만큼만 enQueue가 가능하다. 1회성이라면 좋은 선택일것이다. 이점을 보완하기 위해 **환형 큐**를 사용한다.


> ### 용량에 따른 속도 차이
> #### 용량 설정 O
> ```go
> queue := SliceQueue{make([]int, 0, 500000000)}
> for i := 0; i < 500000000; i++ {
> 	(&queue).enQueue(i)
> }
> ```
> #### 용량 설정 X
> ```go
> queue := SliceQueue{make([]int, 0)}
> for i := 0; i < 500000000; i++ {
> 	(&queue).enQueue(i)
> }
> ```
>
> 
> ```console
> 용량 설정 O 실행시간: 2.3951141s
> 용량 설정 X 실행시간: 7.4095498s
> ```
> 
> 용량을 초과해서 슬라이스 복사가 이루어질때 많은 시간이 들어간다. 이런상황이 없도록 코딩해야한다.

## 선형 큐의 단점

- 큐에 데이터가 꽉 차있으면 데이터를 더 추가할 수 없다.
  - 큐의 사이즈를 늘리는 방법이 있으나 새로운 큐를 생성하고 기존 데이터를 복사하는 과정이 오래걸리므로 시간 효율 측면에서 좋지 못하다.
- 단순 배열로 구현한다면 공간 효율성이 좋지 못하다.
  - enQueue, deQueue 과정에서 front가 계속 뒤로 밀리게 되고 앞쪽의 공간이 남게된다.
  - 데이터를 앞쪽으로 당겨오게 된다면 모든 원소를 옮겨줘야하므로 속도 측면에서 상당히 느리다.