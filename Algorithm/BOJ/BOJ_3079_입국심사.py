from sys import stdin

input = stdin.readline

n, m = map(int, input().split())
arr = [int(input()) for i in range(n)]

# arr.sort()

l, r = min(arr), max(arr) * m
answer = r

while l <= r:
    total = 0
    mid = (l + r) // 2

    for i in range(n):
        total += mid // arr[i]

    if total >= m:
        r = mid -1
        answer = min(mid, answer)

    else:
        l = mid + 1

print(answer)
