# Goland로 Golang시작하기

vscode에서 작업환경을 만들어서 작업해 본적은 있으나 Goland를 사용해 본적은 없다. 그래서 작성하는 포스트

## 프로젝트 생성

![image](https://user-images.githubusercontent.com/113662725/226096584-d51bf41c-7b98-4e04-80d7-2fc0faa9443a.png)

New Project 선택

![image](https://user-images.githubusercontent.com/113662725/226096829-fb547388-40c6-41df-9929-84d422e760b4.png)
- Location: 프로젝트 생성 폴더인데 main으로 폴더명을 만들어야 go run이 가능하다.!
- GOROOT: GOROOT는 Go언어가 설치되어 있는 위치를 말한다. GOROOT오른쪽 +버튼을 누르면 다른버전의 SDK설치할 수 있다. 나는 가장 최근 버전을 설치했다.

Environment는 뭘까?? 나중에 알아봐야겠다.

그리고 create버튼을 눌러 생성하자!

![image](https://user-images.githubusercontent.com/113662725/226096881-350a6c5a-4c16-40ad-af72-e9f83b16d502.png)
- 프로젝트가 정상적으로 생성되었다.

## Hello World!

Hello Wolrd!를 출력할 파일을 go파일을 한개 만들어보자.!

![image](https://user-images.githubusercontent.com/113662725/226096937-b6cf11a8-3876-48cb-ab4b-c9be7e5aef45.png)

Hello.go 파일을 하나 생성한다.

![image](https://user-images.githubusercontent.com/113662725/226096997-54ffc96f-ac3e-4106-b233-3a83ff0a6697.png)

아래 이미지와 같이 코드를 짜준다. fmt는 Println을 사용하기 위해 추가했다.

![image](https://user-images.githubusercontent.com/113662725/226097141-af5ee086-0a76-4294-9d13-a9b0ea93b832.png)

ctrl + shift + f10을 누르면 실행해주고 콘솔에 다음과 같이 출력된다.

![image](https://user-images.githubusercontent.com/113662725/226097218-dc36f550-93da-4fe1-8249-08d05eb8fa1a.png)


코드에디터 윗부분에 GOPATH is empty라고 나오는데 GOPATH를 따로 잡아주지 않아서 그렇다. GOLAND에서는 프로젝트별로 GOPATH를 잡을 수 있는거같은데 다음번에 알아보자.!

GOLAND의 기능이 좋아서 설정에 손이 별로 안가고 손쉽게 HelloWolrd를 출력할 수 있다.~~!