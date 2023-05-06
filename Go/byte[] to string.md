# Byte[] to string

byte배열로 받은 데이터를 문자열로 바꾸는 방법은 여러가지가 있다.

---
## 사용법

### 1. string
```go
byteData := []byte(testString)
_ = string(byteData)
```

### 2. fmt.Sprinf
```go
byteData := []byte(testString)
_ = fmt.Sprintf("%s", byteData)
```

### 3. bytes.NewBuffer
```go
byteData := []byte(testString)
_ = bytes.NewBuffer(byteData).String()
```

### 4. reflect
```go
byteData := []byte(testString)
bh := (*reflect.SliceHeader)(unsafe.Pointer(&byteData))
sh := reflect.StringHeader{bh.Data, bh.Len}
_ = *(*string)(unsafe.Pointer(&sh))
```

---

## Benchmark

> ## 변환 크기당 벤치마크
> 
> 1억회 반복
> 
> ### 12bytes benchmark
> 
> "test string~" 를 변환하는 벤치마크
> 
> ```console
> BenchmarkBytesToStringUseString-6               100000000                4.183 ns/op
> BenchmarkBytesToStringUseSprintf-6              100000000              111.0 ns/op
> BenchmarkBytesToStringUserNewBuffer-6           100000000                4.175 ns/op
> BenchmarkBytesToStringUseReflect-6              100000000                0.2522 ns/op
> ```
> 
> Reflect를 사용하는 방법이 월등히 빠른걸 확인 할 수 있다.
> 
> ### 1200bytes benchmark
> 
> "test string~" * 100 변환
> 
> ```console
> BenchmarkBytesToStringUseString-6               100000000              306.3 ns/op
> BenchmarkBytesToStringUseSprintf-6              100000000              409.6 ns/op
> BenchmarkBytesToStringUserNewBuffer-6           100000000              306.7 ns/op
> BenchmarkBytesToStringUseReflect-6              100000000                0.2668 ns/op
> ```
> 
> ### 12000bytes benchmark
> 
> ```console
> BenchmarkBytesToStringUseString-6               100000000             2212 ns/op
> BenchmarkBytesToStringUseSprintf-6              100000000             2776 ns/op
> BenchmarkBytesToStringUserNewBuffer-6           *** Test killed: ran too long (11m0s).
> ```
> 실행 도중에 죽었다.. 시간설정 없이 시간이 너무 지체되어서 그런것 같다. 2개만 다시
> ```console
> BenchmarkBytesToStringUserNewBuffer-6           100000000             2242 ns/op
> BenchmarkBytesToStringUseReflect-6              100000000                0.2793 ns/op
> ```
> 
> ### 결과 
> 
> reflect를 이용한 변환은 문자열의 길이와 상관없이 비슷한 시간이 걸리는것으로 확인된다. 이유가 뭘까? 알아봐야겠다.

--- 

> ## 반복수에 따른 
> 
> 1200bytes의 데이터 를 변환
> 
> ### 100회
> ```console
> BenchmarkBytesToStringUseString-6                    100
> BenchmarkBytesToStringUseSprintf-6                   100
> BenchmarkBytesToStringUserNewBuffer-6                100
> BenchmarkBytesToStringUseReflect-6                   100
> ```
> 안나옴...ㅡ.ㅡ
> 
> ### 10,000회
> ```console
> BenchmarkBytesToStringUseString-6                  10000               474.4 ns/op
> BenchmarkBytesToStringUseSprintf-6                 10000               307.4 ns/op
> BenchmarkBytesToStringUserNewBuffer-6              10000               204.9 ns/op
> BenchmarkBytesToStringUseReflect-6                 10000
> ```
> reflect는 안나온다.~ 
> 
> ### 1,000,000회
> ```console
> BenchmarkBytesToStringUseString-6                1000000               283.6 ns/op
> BenchmarkBytesToStringUseSprintf-6               1000000               414.2 ns/op
> BenchmarkBytesToStringUserNewBuffer-6            1000000               315.0 ns/op
> BenchmarkBytesToStringUseReflect-6               1000000                 0.09450 ns/op
> ```
> 
> ### 100,000,000회
> ```console
> BenchmarkBytesToStringUseString-6               100000000              302.6 ns/op
> BenchmarkBytesToStringUseSprintf-6              100000000              410.0 ns/op
> BenchmarkBytesToStringUserNewBuffer-6           100000000              306.2 ns/op
> BenchmarkBytesToStringUseReflect-6              100000000                0.2782 ns/op
> ```

> ### 반복수에 따른결과가 이상해서 다시한번 테스트 해보겠다.
> 
> 위에서 테스트한 방법은 **golang에서 test**를 사용해 해보았고 이번엔 직접 for문을 돌려서 테스트 해보겠다.
> 
> ### 100회
> ```console
> 1200bytes
> string: 0s   
> Sprintf: 0s  
> NewBuffer: 0s
> reflect: 0s  
> ```
> 
> ### 10,000회
> ```console
> 1200bytes
> string: 5.1087ms
> Sprintf: 4.3256ms
> NewBuffer: 3.8797ms
> reflect: 0s
> 
> ```
> 
> ### 1,000,000회
> ```console
> 1200bytes
> string: 252.8599ms
> Sprintf: 345.7148ms
> NewBuffer: 226.1223ms
> reflect: 514.2µs
> ```
> 
> ### 100,000,000회
> ```console
> 1200bytes
> string: 24.7054633s
> Sprintf: 35.0372343s
> NewBuffer: 25.297152s
> reflect: 52.8281ms
> ```
> 
> 다른 결과가 나왔다. 이 결과가 좀더 정상적인것같다. 이걸 참고하자.


> ### 결과
> 
> **reflect >>>>>>> Newbuffer > string >> Springf**
> 
> reflect가 월등히 빠르다.
