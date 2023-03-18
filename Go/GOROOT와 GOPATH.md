# GOROOT와 GOPATH

## GOROOT?

golang을 설치했다면 설치한 golang의 폴더가 위치한곳이 goroot이다.

나는 goland를 통해서 golang을 설치했다.

나에게 GOROOT는 "C:\tool\language\Go\sdk\go1.20.2"인것이다.

![image](https://user-images.githubusercontent.com/113662725/226098801-161b05ae-dce4-4614-92c1-c26d84b514a4.png)

### GOROOT 환경변수

1. 시스템 환경 변수 편집

    ![image](https://user-images.githubusercontent.com/113662725/226099737-c52513a6-8857-44fa-bf0b-b7f99bd9e7e1.png)<br><br>

2. 환경 변수

    ![image](https://user-images.githubusercontent.com/113662725/226099773-d4f36665-12b2-48f1-8ffd-430e3ac932f9.png)<br><br>

3. 시스템 변수 -> 새로 만들기 -> 변수 이름(GOROOT), 변수 값을 golang이 설치된 디렉터리로 적는다.

    ![image](https://user-images.githubusercontent.com/113662725/226099846-e3337d0f-89ef-4a81-b36c-e36523456aa7.png)<br><br>



## GOPATH?

GOPATH는 workspace라고 보면될것같다. 

GOPATH는 GO프로젝트의 import위치 소스코드 모듈을 다운로드 하는 위치도 모두 결정하며 기본 경로는 **C:\Users\{USER}\go**(windows기준)이다.
모든 작업은 위 경로에서 진행되며 하위 디렉터리로는 **src, bin, pkg**가 있어야한다.

### 그렇다면 여러개의 프로젝트를 진행한다면..?
의문이 들것이다. workspace를 한개만 정할 수 있다면, 여러개의 프로젝트를 진행 할 순 없을까?


![image](https://user-images.githubusercontent.com/113662725/226099993-66aff15e-3539-4cbd-9222-f06ba6a478a7.png)

goland를 쓰자.. ㅋㅋㅋ 지원 잘해주더라 프로젝트별로 gopath 설정이 가능하다.

그리고 1.11?버전부터 module을 지원해서 프로젝트별로 작업이 가능하다고 한다!! 이것도 알아보자
