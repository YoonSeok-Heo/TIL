from sys import stdin
import heapq

input = stdin.readline

n = int(input())

answer = 0
arr = [int(input()) for i in range(n)]

heapq.heapify(arr)

for i in range(n - 1):
    a = heapq.heappop(arr)
    b = heapq.heappop(arr)
    sum_cards = a + b
    answer += sum_cards
    heapq.heappush(arr, sum_cards)
    
print(answer)