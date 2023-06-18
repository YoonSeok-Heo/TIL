## Intro

음수가 있는 최단거리를 구하는 방법이 벨만포드이다 벨만 포드를 알아보겠다.


## Bellman-ford

- 벨만 포드 알고리즘은 한 노드에서 다른 노드까지의 최단 거리를 구하는 알고리즘이다.
- 다익스트라 알고리즘과 가장 큰 차이점은 간선이 **음수**일때 사용할 수 있다.
- 하지만 효율성이 O(VE) 느리다...
 


### 1. 다익스트라 알고리즘 아이디어

매 단계마다 모든 간선을 전부 확인해 노드간의 최단 거리를 구한다.

#### 알고리즘 동작 원리

1. 모든 간선을 전부 확인하면서 노드간의 최단거리를 구한다.
2. 음수 간선의 순환을 감지

1 - 2의 과정을 모든 노드에서 반복한다. 그래서 속도가 O(VE)이다.

### 2. 구현 방법

최단거리 테이블 최대값으로 초기화

1. 출발 노드 설정
2. 최단거리 테이블 초기화
3. 다음 과정 반복
   1. 모든 간선을 확인
   2. 각 간선을 거쳐 다른 노드로 가는 비용을 계산하여 최단 거리 테이블 갱신

- 음수 간선 순환이 발생하는지 체크하고 싶다면 3번 과정을 한번 더 수행한다. 이때 테이블이 갱신된다면 음수의 순환이 발생하고 있다는 것이다.

아래 소스는 백준 11657 문제를 풀이한 내역이다.

```python
import sys

INF = int(1e9)

input = sys.stdin.readline

v, e = map(input().split())

edges = []

cost = [INF] * (v + 1)

for _ in range(e):
  a, b, c = map(int, input().split())

  edges.append((a, b, c))

def bellman_ford(start):

  cost[start] = 0

  for i in range(v):
    for j in range(e):

      s = edges[j][0]
      e = edges[j][1]
      c = edges[j][2]

      if cost[s] != INF and cost[e] > cost[s] + c:
        cost[e] = cost[s] + c

        if i == v - 1:
          return True

  return False

cycle = bellman_ford(1)

if cycle:
  print("-1")
else:
  for i in range(2, v + 1):
    if cost[i] == INF:
      print("-1")
    else:
      print(cost[i])
```