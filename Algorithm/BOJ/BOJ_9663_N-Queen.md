# BOJ 9663 N-Queen
#프루트포스 #백트래킹

[BOJ 9663 N-Queen](https://www.acmicpc.net/problem/9663)

맨처음에 2차원 배열로 해결하려 했으나. 한줄에 퀸이 한개만 들어간다고 생각하면 n개의 배열만 있으면 된다.
그리고 퀸의 인덱스를 넣어주면 된다. 그럼 모든 위치를 n개의 배열로 나타낼 수 있다. 뿐만아니라 위치 계산도 O(1)으로 가능해진다.



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
