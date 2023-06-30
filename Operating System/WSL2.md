
## WSL2 설치 조건

- Windows10 2004 이상(빌드 19041) 버전(버전이 낮으면 윈도우 업데이트를 해야한다.)
- CPU가 가상화를 지원해야한다.

## 환경 만들기

1. CPU가 가상화를 지원해야한다.(보충설명 해야하는데 나중에.)
2. WSL2는 windows10 hyper-v가 실행되어야 한다.
- Hyper-V 실행[(클릭)][DISM hyper-v setting]




## 설치

1.  PowerShell 또는 Windows 명령 프롬프트에서 아래 명령을 입력합니다.(관리자 권한 실행)
```
wsl --install
```
- 위 명령어는 최신 Linux커널을 다운로드하고, WSL2, Ubuntu를 기본값으로 설치합니다.
- 설치 완료 시 아래와 같이 재부팅 하라는 말이 나옵니다. 재부팅을 해줍니다.
  ![image](https://user-images.githubusercontent.com/51642448/152667565-190a8950-4fb1-45c4-adbe-792c2e4743ab.png)



2. Microsoft Store 에서 사용하고자 하는 ubuntu를 설치 실행합니다.

![image](https://user-images.githubusercontent.com/51642448/152667645-7ed50060-2cfd-4f4e-a4ea-570d65c45387.png)

![image](https://user-images.githubusercontent.com/51642448/152668244-6dd50e84-bf38-4a05-88d6-ac7a1ddd8770.png)

Ubuntu가 정상동작 합니다.

WSL2가 처음 나왔을땐 이것저것 많이 해줬어야했는데 지금은 많이 단조로워져서 설치하기 쉬웠다.

[DISM hyper-v setting]: https://docs.microsoft.com/ko-kr/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v