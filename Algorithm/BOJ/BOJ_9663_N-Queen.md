# BOJ 9663 N-Queen
#프루트포스 #백트래킹

[BOJ 9663 N-Queen](https://www.acmicpc.net/problem/9663)

## 문제 해석
- N x N인 체스판 위에 N개의 퀸들이 서로 공격할 수 없도록 놓는 경우의 수를 구하는 문제이다.

### 서로 공격할 수 없는 조건
1. 자신을 기준으로 일직선상에 놓을 수 없다.
2. 자신을 기준으로 대각선상에 놓을 수 없다.

### 해결 방안

맨 처음에 2차원 배열로 해결하려 했으나. 한 줄에 퀸이 한 개만 들어간다고 생각하면 N개의 배열만 있으면 된다.
그리고 퀸의 인덱스를 넣어주면 된다. 그럼 모든 위치를 N개의 배열로 나타낼 수 있다. 그뿐만 아니라 위치 계산도 O(1)으로 가능해진다.

![image](https://user-images.githubusercontent.com/113662725/218292591-bcc45910-1918-43d6-a641-abf511ce8f1f.png)

위와같이 퀸의 위치를 배열로 넣어주면 된다.

### 풀이
2개의 함수를 작성했다. DFS형식으로 작동하는 함수와 퀸을 놓을 수 있는지 없는지 체크(check(x))하는 함수!

DFS형식의 재귀함수는 설명 생략 보면 안다.!

체크하는 함수를 설명해보겠다.
```python 
def check(x):
    for i in range(x):
        if rows[x] == rows[i] or abs(rows[x] - rows[i]) == x - i:
            return False
    return True
```


![image](https://user-images.githubusercontent.com/113662725/218293208-c419348f-c039-4220-80db-ae4865ca79db.png)
- check에서 파라미터로 들어오는 x값은 행이다. (위 표에서 색상 별로 x값이라고 생각하면 된다.


![image](https://user-images.githubusercontent.com/113662725/218293358-d224a69a-7314-47b1-bbc6-3744844adc79.png)
- for문에서 x만큼 반복하게 되는데. x까지 라인에서 같은 열에 같은 값이 있는지 확인한다.

#### rows[x] == rows[i]
같은 행 열을 체크하는 수식.
- 퀸이 초록색이라고 생각하면 rows[2] == rows[0]은 둘다 0값으로 False로 리턴되는것을 확인할 수 있다.

#### abs(rows[x] - rows[i]) == x - i
대각선상에 있는 퀸을 체크하는 수식

![image](https://user-images.githubusercontent.com/113662725/218293455-85f1c2be-867b-4b1b-8068-741cabc48458.png)
- 예시
    - x = 2, i = 0
    - abs(rows[x] - rows[i]) = abs(2 - 0) = 2
    - x - i = 2 - 0 = 2
- 설명
    - ![image](https://user-images.githubusercontent.com/113662725/218293680-f6adcea6-cb51-4cef-9e61-b4a3746c5b2a.png)
    - 0, 0에서는 대각선 위치는 삼각형으로 봤을때 (1, 1), (2, 2) 처럼 가로 세로의 길이가 같아야한다.


## Python
```python
import sys
input = sys.stdin.readline

def check(x):
    for i in range(x):
        if rows[x] == rows[i] or abs(rows[x] - rows[i]) == x - i:
            return False
    return True

def dfs(x):
    global cnt

    if x == n:
        cnt += 1
        return

    for i in range(n):
        if visited[i]: 
            continue

        rows[x] = i
        if check(x):
            visited[i] = True
            dfs(x+1)
            visited[i] = False

n = int(input())
rows = [0] * n
visited = [False] * n
cnt = 0

dfs(0)
print(cnt)
```
