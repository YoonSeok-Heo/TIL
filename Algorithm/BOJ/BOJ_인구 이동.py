from sys import stdin

input = stdin.readline

dx, dy = [0, 1, 0, -1], [1, 0, -1, 0]

n, l, r = map(int, input().split())
arr = [list(map(int, input().split())) for i in range(n)]

for i in range(n):
    for j in range(n):
        visit = [[0] * n for _ in range(n)]
        
        for a in visit:
            print(a)
        print()


for i in arr:
    print(i)