# BOJ 1149 RGB거리
#DP

[BOJ 1149 RGB거리](https://www.acmicpc.net/problem/1149)

DP기초 문제이다.

## 문제 해석
- 집의 개수와 집마다 RGB색상으로 칠하는데 비용을 입력받는다.
- 옆집은 같은 색으로 칠할 수 없을 경우 모든 집을 색칠하는데 최소비용은?

## 문제 풀이

![image](https://user-images.githubusercontent.com/113662725/218497420-50fa603a-fe06-4299-b570-766b88233ca2.png)
- 입력값

![image](https://user-images.githubusercontent.com/113662725/218497262-c6c56cd4-4325-4cff-87a3-6fd0e31c23eb.png)
- DP행렬을 만들어준다. 소스상에서는 0으로 초기화된 DP행렬을 만들었다.

![image](https://user-images.githubusercontent.com/113662725/218498511-99cdd152-3906-47e0-8e4a-30431d0a1417.png)
- DP테이블에서 0, 1, 2열 순서대로 RGB의 색상이다. 
- R의 값을 입력 받았으면 전 집의 G와 B를 비교해야한다.
- R행열의 칠해져있는 부분에서 (1, 0)은 입력된 값과 전 값인(0, 1), (0, 2)값과 비교해서 더 작은 값을 인풋으로 입력된 R(26)과 더해서 넣어준다. 
```python
dp[i][0] = min(dp[i - 1][1], dp[i - 1][2]) + R
```
- G행열의 칠해져있는 부분에서 (1, 1)은 입력된 값과 전 값인(0, 0), (0, 2)값과 비교해서 더 작은 값을 인풋으로 입력된 G(40)과 더해서 넣어준다. 
```python
dp[i][1] = min(dp[i - 1][0], dp[i - 1][2]) + G
```
- G행열의 칠해져있는 부분에서 (1, 2)은 입력된 값과 전 값인(0, 0), (0, 1)값과 비교해서 더 작은 값을 인풋으로 입력된 G(83)과 더해서 넣어준다. 
```python
dp[i][2] = min(dp[i - 1][0], dp[i - 1][1]) + B
```

![image](https://user-images.githubusercontent.com/113662725/218501442-ee9f3c5c-988d-48e5-9937-c31593c7bfc5.png)
- 위의 과정을 반복하게 되면 DP행렬을 위 이미지와 같이 바뀐다. 
- 마지막 행은 총 소모하게 될 금액을 의미하며 96이 가장 적으므로 96이 답이된다.


## Python

```python
n = int(input())

dp = [[0] * 3 for i in range(n + 1)]

for i in range(1, n + 1):
    R, G, B = map(int, input().split())
    
    dp[i][0] = min(dp[i - 1][1], dp[i - 1][2]) + R
    dp[i][1] = min(dp[i - 1][0], dp[i - 1][2]) + G
    dp[i][2] = min(dp[i - 1][0], dp[i - 1][1]) + B
    
print(min(dp[-1]))
```
최근에 다시 풀면서 진행한 방식으로 RGB값을 받을때마다 계산을 실행했다.

```python
n = int(input())
dp = []
for i in range(n):
    dp.append(list(map(int, input().split())))

for i in range(1, len(dp)):
    dp[i][0] = min(dp[i - 1][1], dp[i - 1][2]) + dp[i][0]
    dp[i][1] = min(dp[i - 1][0], dp[i - 1][2]) + dp[i][1]
    dp[i][2] = min(dp[i - 1][0], dp[i - 1][1]) + dp[i][2]

print(min(dp[n - 1][0], dp[n - 1][1], dp[n - 1][2]))
```
예전에 풀어둔 코드인데 RGB값을 모두 다 받고 한번에 DP계산을 진행

두개 코드의 가장 큰 차이점은 for문을 1번실행하냐? 2번실행 하느냐이다. 물론 전자가 한번 실행하기에 속도가 더빠르다.
![image](https://user-images.githubusercontent.com/113662725/218502468-39e6794b-d728-4661-b896-bf032548812d.png)
