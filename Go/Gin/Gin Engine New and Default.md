# Gin Engine New and Default
#Go #Gin-Engine #Gin #Gin-Engine-New #Gin-Engine-Default

Gin공부 중 미들웨어를 추가하고 싶어서 찾아보고 있는데 gin.New()로 생성하는 부분을 확인했다. 

지금까지 gin.Default()로 생성하던 것과 달리 New()는 어떤 동작을 하고 무었인지 알아봐야겠다. 


## Default()

![image](https://user-images.githubusercontent.com/113662725/226365452-fdadcfae-6853-4ba7-8803-9681ea9eb9a8.png)

설명에 보면 기본값으로 로거랑 복구 미들웨어가 들어있는 엔진 인스턴스를 반환한다고 써있는거 같다.(영어를 잘 못함)

지금까지 계속 생성하던 Defualt인스턴스는 기본 설정이 된 가장 기본적인 인스턴스로 보면 될 것 같다.


## New()

![image](https://user-images.githubusercontent.com/113662725/226366256-184e275a-8a7b-4a60-ac70-d72d4f8cb723.png)

New는 미들웨어가 첨부되지 않은 빈 인스턴스를 반환한다고 적혀있다.

**기본 구성**
- RedirectTrailingSlash: true
- RedirectFixedPath: false 
- HandleMethodNotAllowed: false 
- ForwardedByClientIP: true 
- UseRawPath: false 
- UnescapePathValues: true

기본 구성이라고는 적혀있으나 정확히 어떤기능을 하는지 잘 모른다. 확인해 봐야한다. 


### 결론!

Default에는 기본적인 미들웨어가 들어가있다. 그럼 추가가 안되는것인가??? 이건 잘모르겠다.

New 인스턴스에는 미들웨어가 들어있지 않은 빈 인스턴스가 반환되기 때문에 인스턴스에 Use함수를 이용해 미들웨어를 추가해 사용할 수 있다.

그냥 Default는 기본적인 엔진인스턴스 New는 아무것도 없는 빈 엔진 인스턴스라고 보면 될 것 같다. 

