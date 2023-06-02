## Tree(트리)



**트리**는 그래프의 일종으로 특수한 조건을 만족하는 그래프를 말한다.

트리는 생김새가 나뭇가지와 비슷하여 트리라는 이름으로 불리운다.



### 1. 용어정리

- 루트

  : 루트 노드는 부모노드가 없는 트리 구조의 최상의 노드를 말한다. 경우에 따라서 모든 노드를 루트노드로 만들 수 있다.



![image](https://user-images.githubusercontent.com/51642448/132092340-d1f8f0dc-08d4-486e-a5a8-eac645c62c93.png)

​

​	위와 같은 노드는 부모가 없는 1번노드를 부모 노드라고 생각 할 수 있다.

​

![image](https://user-images.githubusercontent.com/51642448/132092390-47d8e1bd-541a-4482-9b23-970d4d5d7c6f.png)



​	하지만 위와같이 노드를 적절히 회전시켜 8번 노드가 루트인 트리도 만들 수 있다.

​	트리와 관련된 문제는 루트를 임의로 선택해야 할 수 있다는걸 알아두자.





**1번노드 기준으로 파란색 노드가 해당노드이다.**

- 부모(parent)노드

  : 부모 노드는 연결되어 있는 노드 중 **자신보다 상위에 있는 노드**를 의미

  : 1번 노드 기준 부모노드는 2번노드이다.



![image](https://user-images.githubusercontent.com/51642448/132093821-f4e2af0a-a93e-4a3d-a731-2584c992d399.png)

- 자식(child)노드

  : 자식 노드는 연결되어 있는 노드 중 **자신보다 하위에 있는 노드**를 의미

  : 1번 기준 자식노드는 3번노드이다.



![image](https://user-images.githubusercontent.com/51642448/132093832-44aa5510-c545-4198-8b7e-6f7b63424009.png)

- 단말(leaf)노드

  : 단말 노드는 **자식이 없는 노드**를 의미

  : 자식이 없는 노드는 4, 6, 7, 9이다.



![image](https://user-images.githubusercontent.com/51642448/132093739-5b985905-56e6-4980-acd8-30c52fd9025b.png)

- 형제(siblin)노드

  : 형제 노드는 **같은 부모를 갖는 노드**를 의미

  : 1번노드의 형제노드는 7, 4번이다.



![image](https://user-images.githubusercontent.com/51642448/132093783-12d60aef-e10c-4355-bd64-f849addec94e.png)



- 조상(ancestor)노드

  : 부모 노드와 부모 노드의 부모들을 의미한다. 즉 해당 **노드에서 루트 노드까지 경로에 있는 모든 노드**를 의미

  : 1번노드와 루트노드 그사이의 노드 모두 조상 노드이다.



![image](https://user-images.githubusercontent.com/51642448/132093796-0a65fe87-0b04-4154-809f-5fb9941d1bba.png)

- 자손(descendant)노드

  :  조상 노드와 반대로 **자신의 모든 자식 노드**

  : 1번 노드의 모든 자식노드



![image](https://user-images.githubusercontent.com/51642448/132093808-7537f542-b43a-40e3-9b55-1d0562960b06.png)



- 깊이(depth)

  : 루트 노드에서 해당 노드까지의 거리를 의미. **깊이가 같은 노드를 레벨이 같다**고 한다.

  : 1번의 깊이는 3이다

![image](https://user-images.githubusercontent.com/51642448/132093907-e53be607-268e-4744-9bda-24e662c2b9c4.png)

- 높이(height)

  : **트리에서 가장 깊이 **있는 노드의 깊이를 높이라고 한다.

  : 아래 트리의 높이는 6이다.



![image](https://user-images.githubusercontent.com/51642448/132093891-95108848-7ea9-495b-aa9c-9da9f4660546.png)







트리의 특징

1. 사이클이 존재하지 않는다.
2. **어떤 노드(정점)에서 다른 노드로 가는경로가 유일** 위의 트리에서 1번노드에서 어떤노드를 가더라도 경로는 하나뿐
3. 노드(정점)의 총 갯수를 **V**개라고 했을 때 간선의 개수는 **V - 1**개이다.
4. 루트 노드를 제외한 모든 노드의 입력 차수는 1이다. 각각의 노드는 단 하나의 부모를 가질 수 있다.
5. **트리는 재귀적 속성**을 갖는다. 트리에 있는 한 노드와 그 노드 자손들은 모두 모은 것이 트리



트리는 그래프의 일종이므로 다양한 방법으로 표현할 수 있다. 인접 행렬, 인접 리스트를 이용하여 표현. 이뿐만 아니라 트리가 갖는 몇가지 특성으로 특정 노드의 부모 노드만을 저장하여 표현하는 방법, 1차원 배열을 이용하여 저장하는 방법도 사용한다. 부모 노드만을 저장하는 방법의 경우 유니온 파인드라는 자료구조를 다룰 때, 1차원 배열을 이용하여 저장하는 방법은 주로 최대 2개의 자식 노드를 가질 수 있는 이진 트리를 표현할 때 사용.



트리를 순회하는 방법

- DFS
- BFS
- 전위
- 중위
- 후의

### 2. DFS(Deptg Frist Search)를 알아보자

DFS(Deptg Frist Search), 깊이 우선 탐색이라고 한다.

말그대로 깊이를 우선 탐색한다.

![image](https://user-images.githubusercontent.com/51642448/132092390-47d8e1bd-541a-4482-9b23-970d4d5d7c6f.png)

위의 트리를 깊이우선 탐색해보자

8번 노드가 루트 노드이다.

![image](https://user-images.githubusercontent.com/51642448/132126787-06671f9e-e3b3-4eed-be20-a536c48bca0c.png)

**8 > 2 > 7 > 4 > 1 > 3 > 5 > 9 > 6** 의 순서대로 탐색하게 된다.

이는 항상 이렇게 순회한다는건 아니다 어떻게 코드를 짜느냐에 따라 다르다.



![image](https://user-images.githubusercontent.com/51642448/132126872-554cc54a-3174-4ea8-9a02-f14839b8fe6e.png)

위와 같은 방법으로도 탐색할 수 있다.

**8 > 2 > 1 > 3 > 6 > 5 > 9 > 4 > 7** 위의 방법과는 다르지만 둘다 자식노드를 먼저 찾아서 순회 하는걸 알 수 있다.



#### 간단하게 순회 방법을 알아보았다. 코드 내에선 어떻게 동작하나 알아보자

깊이 우선 탐색은 **스택**을 사용한다.





![image](https://user-images.githubusercontent.com/51642448/132127454-db6f90e1-22ab-4154-9cdb-5863c91daa40.png)

1. 처음에 시작할 노드를 스택에 넣어준다.



**여기서 순회가 끝날 때까지 반복작업이 이루어진다.**

2. 스택의 값을 팝해준다. 그럼 스택의 특징상 마지막 값이 나오게 된다.

   팝해준 값은 8

   8은 이제 방문한 노드다.



![image](https://user-images.githubusercontent.com/51642448/132127629-8f8f246b-0b92-46ed-a326-b966a1bcca68.png)



3. 방금 팝한 8의 **자식 노드**를 다시 스택에 넣어준다.

    - 넣어주면서 이미 방문한 노드는 스택에 넣지 않는 작업이 필요로하다.
    - 이를 해결하기 위해 visit리스트를 사용하였다. 방문한 노드를 visit 리스트에 넣어주었다.



![image](https://user-images.githubusercontent.com/51642448/132127641-cd0c257b-7f6c-4336-b49e-617fc230cd72.png)

4. 8의 자식 노드 2를 넣어주게 된다.

   ![image](https://user-images.githubusercontent.com/51642448/132127616-458e6276-3731-415d-af36-102303f73b0e.png)



**위의 2 ~ 4의 작업을 반복하게 된다.**



![image](https://user-images.githubusercontent.com/51642448/132128006-753b3120-9899-4fa1-b4df-440420bd39c7.png)

2를 pop하고 2의 자식노드 7, 4, 1을 스택에 넣어준다.

여기서 다음 pop의 대상은 stack의 특징상 1이 된다.

뒤에서 계속 pop을 하게 되면 깊이 우선탐색이 완성된다.

다음은  DFS python 코드이다.

```python
def DFS(tree: list, start_node: int):
    stack = []
    visit = []
    stack.append(start_node)
    
    while stack:
        node = stack.pop()
        print(node)
        visit.append(node)
        for n in tree[node]:
            if n not in visit:
                stack.append(n)
```



