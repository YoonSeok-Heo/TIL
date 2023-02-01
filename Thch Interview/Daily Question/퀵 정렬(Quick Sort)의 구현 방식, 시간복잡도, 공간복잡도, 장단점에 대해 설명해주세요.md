# 퀵 정렬(Quick Sort)

- 퀵정렬은 **불안정 정렬**에 속한다.
- 다른 원소와의 비교만으로 정렬을 수행하는 **비교 정렬**에 속한다.
- 분할 정복 알고리즘의 하나로, 평균적으로 **매우 빠른 수행 속도**를 자랑하는 정렬 방법
- 분할 정복(divide and conquer)방법
  - 문제를 작은 2개의 문제로 분리하고 각각을 해결한 다음, 결과를 모아서 원래의 문제를 해결하는 전략.
  - 분할 정복은 대부분 순환 호출을 이용하여 구현
- 병합 정렬(merge sort)과 달리 퀵 정렬은 리스트를 **비균등하게** 분할한다.

## 정렬 과정

1. 리스트의 요소 중 한개를 선택한다. 이렇게 고른 원소를 피벗(pivot)이라고 한다.
2. 피벗을 기준으로 피벗보다 작은 요소는 모두 피벗의 왼쪽으로, 피벗보다 큰 요소는 피벗의 오른쪽으로 옮겨진다.
3. 피벗을 기준으로 왼쪽 리스트와 오른쪽 리스트를 다시 정렬한다.
    1. 분할된 부분 리스트에 대하여 순환 호출을 이용하여 1번과 2번과정을 반복
    
4. 부분 리스트들이 더이상 분할이 불가능할 때까지 반복(리스트의 크기가 0또는 1이 될때까지 반복)


## 퀵 정렬(quick sort)알고리즘 예제 partitioning

![image](https://user-images.githubusercontent.com/113662725/216049467-f5dcf22d-da7c-4440-b780-acc35c6cba48.png)
- Pivot값을 먼저 설정해준다. 피벗값을 설정하는 방법은 다양하다. 나는 맨앞에 있는 값을 설정
- 2개의 포인터를 준비 Left(L) Right(R)

![image](https://user-images.githubusercontent.com/113662725/216049922-5b67c2a1-002c-4b7a-a2cf-57edf5bd9e1a.png)
- L포인터 먼저 피벗값과 비교하여 작다면 다음 위치로 이동 크다면 잠시 고정해둔다.(L포인터는 피벗보다 크면 고정 작으면 다음으로 이동)

![image](https://user-images.githubusercontent.com/113662725/216050286-5c06432f-26d9-4911-8f06-803b8edcf5ba.png)
- L포인터를 고정하고 R포인터를 피벗값과 비교 피벗값보다 작을경우 고정 클경우 다음 위치로 이동한다.(R포인터는 피벗보다 작으면 고정 크면 다음으로 이동)

![image](https://user-images.githubusercontent.com/113662725/216050481-38779572-62ff-448b-823c-e64beb0c15ea.png)
- L포인터와 R포인터의 값이 둘다 고정되어 있으면 값을 서로 바꿔준다.

![image](https://user-images.githubusercontent.com/113662725/216050707-ce8c66af-08ef-4ddf-8c5f-c0461140cb9d.png)
- 값을 서로 바꾸고 L포인터와 R포인터의 값을 한칸씩 이동한다.

![image](https://user-images.githubusercontent.com/113662725/216050935-d06ad261-5eb4-47b5-8aea-c7d9d6976745.png)
- L포인터가 피벗값보다 크기 때문에 포인터 고정한다.

![image](https://user-images.githubusercontent.com/113662725/216051196-a406b5f9-93fd-47f6-b211-f77949ab8fac.png)
- R포인터는 피벗값보다 크기 때문에 다음 포인터로 이동한다.

![image](https://user-images.githubusercontent.com/113662725/216051572-b84d950f-16e3-4929-b0e1-a7528e94069b.png)
- R포인터가 피벗값보다 작으므로 고정한다.

![image](https://user-images.githubusercontent.com/113662725/216051735-2843420b-1442-49c0-82cd-c90a22691e7e.png)
- L포인터와 R포인터의 값이 둘다 고정되어 있으면 값을 서로 바꿔준다.

![image](https://user-images.githubusercontent.com/113662725/216051921-594a1ffa-602c-46d8-845e-c12cb335d8d5.png)
- L포인터가 피벗값보다 크기 때문에 포인터 고정한다.

![image](https://user-images.githubusercontent.com/113662725/216052129-8e76c994-64d2-49b1-8ac5-5c6b0b9bdc24.png)
- R포인터는 피벗값보다 크기 때문에 다음 포인터로 이동한다.

![image](https://user-images.githubusercontent.com/113662725/216052181-7f71c375-aa1a-48c4-ad5b-679fe0a8e918.png)
- R포인터는 피벗값보다 크기 때문에 다음 포인터로 이동한다.

![image](https://user-images.githubusercontent.com/113662725/216052379-ac244ac5-069b-42bb-b962-dc7d5669a413.png)
- R포인터가 L포인터보다 왼쪽에 있을경우 피벗값과 교환한다.

<br>

![image](https://user-images.githubusercontent.com/113662725/216052453-a8e066bd-8a5b-480d-a90d-f3b966dbc081.png)
- 교환한 피벗값보다 왼쪽에 있는 값은 작은 값들만 모이게 되고 오른쪽의 값들은 큰값들이 모이게된다
- Smaller값들과 Bigger의 값들을 방금 했던 작업들을 반복해주면 된다.(분할 정복)






