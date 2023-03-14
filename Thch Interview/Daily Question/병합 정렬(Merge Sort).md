# 병합 정렬(Merge Sort)

합병 정렬이라고도 부르며, 분할 정복 방법을 통해 구현한다.!

빠른 정렬로 분류되며, 퀵소트와 함께 많이 언급되는 정렬방식이며 퀵소트와는 반대로 **안정 정렬**에 속한다.

## 시간 복잡도
|평균|최선|최악|
|---|---|---|
|O(nlogn)|O(nlogn)|O(nlogn)|

모든 요소를 쪼갠 후, 다시 병합시키면서 정렬하기 때문에 속도가 같다.

## 코드

### Python 
```python
def merge_sort(arr):
    if len(arr) < 2:
        return arr

    mid = len(arr) // 2
    low_arr = merge_sort(arr[:mid])
    high_arr = merge_sort(arr[mid:])

    merged_arr = []
    l = h = 0
    while l < len(low_arr) and h < len(high_arr):
        if low_arr[l] < high_arr[h]:
            merged_arr.append(low_arr[l])
            l += 1
        else:
            merged_arr.append(high_arr[h])
            h += 1
    merged_arr += low_arr[l:]
    merged_arr += high_arr[h:]
    return merged_arr
```
