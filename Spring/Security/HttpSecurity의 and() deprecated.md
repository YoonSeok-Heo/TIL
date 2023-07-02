# HttpSecurity에서 and()가 비활성화 되다.

작성일자인 2023.07.01 기준으로 불과 얼마전에 security를 사용했을 땐 분명 and()를 잘 사용했다.
그리고 오늘 해보니 deprecated된걸 확인할 수 있었다.

> Security가 변화하는게 많은것 같다. 사실 내 생각이다. 

### 직접 사용했던 코드를 확인해보자.

올해 초까지 사용하던 코드의 일부이다. 여기선 아무 문제가 없었다. spring-security-core 6.0.2 버전이다.
and()를 사용하기도하고 lambda를 사용해서 코드를 만들기도 했다.

```java
public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http
        .httpBasic().disable()
        .csrf().disable()
        .cors().disable()
        .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS)
        .and()
        .authorizeHttpRequests((authz) -> authz
            .requestMatchers(HttpMethod.POST, "/user").permitAll()
            .requestMatchers("/accessDeniedPage").permitAll()
            .requestMatchers("/admin/**").hasRole("ADMIN")

            .anyRequest().authenticated()
                )
        );
        return http.build();
}
```

그런데 오늘 위와 비슷한 형식의 코드를 작성하니 Intellij에서 에러로 잡았다....ㅜㅜ
어떤에러지 싶어서 에러를 확인해보니 Intellij에서 deprecated된 코드를 에러로 잡고있었다. 
Intellij 오류로 잡는게 아니고 기본 정책으로 정해져있다.

> 몇달전만 해도 사용했던 코드가 deprecated되다니. 
> 
> 나는 개발자이기때문에 회사에선 스프링 6버전대를 사용하지 않더라도 혼자 공부할땐 deprecated된 코드를 절대 사용할 수 없다.
> 
> 그러면서 해당 내용을 알아보기 시작했다.


---

## 왜 없어졌을까?

관련 내용을 찾다가 HttpSecurity로 필터내용을 정의할 때 lambda로 작성하는 것을 권장한다고 나와 있었다.
**그 문서를 어디서 봤는지 기억이 안난다...ㅠㅠ**

여튼 그래서 내 코드중에서 lambda로 작성한 코드는 Intellij가 잡지 않았던 것으로 생각된다.

---

## 변경 방법

말그대로 lambda형식으로 바꾸면 된다.

```java
public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http
        .httpBasic((basic) -> basic
                .disable()
        )
        .csrf((c) -> c
                .disable()
        )
        .cors((c) -> c
            .disable()
        )
        .sessionManagement((session) -> session
            .sessionCreationPolicy(SessionCreationPolicy.STATELESS)
        )
        
        .authorizeHttpRequests((authz) -> authz
                .requestMatchers(HttpMethod.POST, "/api/v1/user").permitAll()
                .anyRequest().authenticated()
        );

        return http.build();
    }
```
나는 이런식으로 해결 했다.

### !! 그런데 여기서 궁금한점이 생겼다

CustomDsl을 어떻게 등록할 것인가?

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/d463c27b-13ee-4da1-80d4-075ed3d7ecca)

and를 빼면?

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/e51513c2-0daf-40a5-94e0-a6f2ebe35c50)

무지한 내가 뭔가 뺴먹은게 있을 것이다. 

### 나의 해결 방법

정답은 아닐 것이다..

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/3abcc0f1-a14c-40c7-bf7e-96f72b24a3f2)

그냥 http를 한번 더 ㅋㅋㅋㅋㅋ ... 아직 모든 테스트를 해보진 않았따. 



어카지?

