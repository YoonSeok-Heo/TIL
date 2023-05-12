# Circular Queue

환형 큐, 원형 큐로 부르며 [**<선형 큐>**](https://github.com/YoonSeok-Heo/TIL/blob/main/Data%20Structure/Queue.md)의 단점을 보완하기 위한 자료구조이다.

선형큐의 단점은 [**<클릭>**](https://github.com/YoonSeok-Heo/TIL/blob/main/Data%20Structure/Queue.md#%EC%84%A0%ED%98%95-%ED%81%90%EC%9D%98-%EB%8B%A8%EC%A0%90)에서 확인할 수 있다.

## 동작 과정

선형 큐에서 deQueue로 인해 발생한 배열의 빈 공간을 활용하기 위해 원형 큐가 생겼다.

---

![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/92adff71-372e-419a-b9cc-94fd4526b9dc)
그림 .1

![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/a1c5c6b0-e524-47ab-94c8-1865484f7b39)
그림 .2

그림 .1 이 기존 선형큐라면 그림.2는 환형 큐이다.

선형큐에서 앞부분과 뒷부분을 연결해준 형태로 보면 된다.

--- 

> ### enQueue()
> 
> #### 1. 초기상태
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/5ce0728e-8f3e-43c1-aa31-e25219ab501b)
> 
> 초기상태이다. front/rear 포인터가 같은 지점을 보고 있다. (즉, front/rear가 같은 지점을 보고있으면 큐가 비어있다고 볼 수 있다.)
> 
> #### 2. rear가 가리키는 값에 데이터 삽입.
> 
> **데이터를 넣기전에 큐가 꽉차있는지 확인한다.** 확인 방법은 front 1칸뒤에 rear가 있다면 꽉찬것이라고 볼 수 있다.
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/16928a86-7d16-4b00-90d3-9106d40fda47)
> 
> #### 3. rear = (rear + 1) % 8
> 
> rear의 값을 1 증가시킨다. 모듈러 8을 해주는 이유는 배열의 최대 크기 8을 초과하면 다시 1로 돌아가야하기 때문이다.(0부터 시작해야하는데 이미지를 잘못만들어서...)
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/a2bb0e50-d9ed-40cf-bb9a-35aadd580932)

---

> ### deQueue()
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/c227c67d-67e0-449a-b1ac-5e188d1e3bdd)
> 
> 위 그림과 같은 데이터가 들어있다고 가정한다.
> 
> #### 1. front가 가리키는 값 내보내기
> 
> 내보내기 전 큐 안에 데이터가 있는지 확인해야한다. 
> 
> front와 rear가 같은 위치를 가리키고 있다면 큐는 비어있으므로 deQueue를 실행할 수 없다.
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/e8e721a2-6422-4d45-b4e6-15c0a5bd76f6)
> 
> #### 2. front = (front + 1) % 8
> 
> front + 1 해주고 모듈러로 배열의 최대값만큼 나누어 나머지값을 구한다.(그림이 0부터 시작해야하는데..)
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/153fb622-e30d-413f-9ffd-2c6acd2838ce)

---

> ### isEmplty()
> 
> rear와 front가 같은 값을 가리키고 있으면 큐가 비어있는 상태이다.
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/d5ee0ad6-a9a1-41e5-b38f-8a53269f50bb)
> 
> ### isFull()
> 
> rear값이 front보다 한칸 전에 있으면 Full상태이다. Full상태를 확인하기 위해서는 배열의 한칸은 사용하지 비워둬야한다.
> 
> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/ed6e073e-451f-4be4-855d-19d60d367c58)
> 
>> 만약. 모든 데이터를 가득 채우게 된다면. front/rear포인터로 비어있는 상태와 구분할 수 없는 상태가 된다.
>> 
>> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/451e25fe-15a3-48e7-9d71-a0e04c523405)
>> ![image](https://github.com/YoonSeok-Heo/Algorithm/assets/113662725/d5ee0ad6-a9a1-41e5-b38f-8a53269f50bb)


