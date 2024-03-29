# Exception in Java


## Exception?

Java에서의 **예외 처리(Exception Handling)**는 애플리케이션의 정상적인 흐름을 유지할 수 있도록 런타임 오류를 처리하는 효과적인 수단이다.
ClassNotFoundExceptio, IOException, SQLException, RemoteException등과 같은 런타임 오류를 처리하는 메커니즘이다.

Exception은 프로그램 실행 중, 즉 런타임에 발생하는 원치 않거나 예기치 않은 이벤트로 프로그램 명령의 정상적인 흐름을 방해한다.
Exception은 프로그램에서 포착하여 처리할 수 있으며, 메서드 내에서 예외가 발생하면 예외는 객체를 생성한다.
이 객체를 Exception object라고 한다.
Exception object에는 Exception 의 이름과 설명, Exception이 발생했을 때의 프로그램 상태 등 Exception에 대한 정보가 포함된다.

## 예외가 발생하는 주요 이유

- 사용자의 잘못된 입력(int변수에 문자를 집어넣는다.)
- 장치 오류(장치와의 연결이 헤제된다?)
- 네트워크 연결 오류
- 물리적 제한(디스크 메모리 부족)
- 코드 오류
- 사용할 수 없는 파일 열기(암호화 파일, 확장자가 다른 파일?)


## Error?

Errors는 메모리 부족, 메모리 누수, 스택 오버플로 오류, 라이브러리 비호환성, 무한 재귀 등과 같이 **복구할 수 없는 상태**를 나타낸다.
Error는 일반적으로 프로그래머의 통제 범위를 벗어나는 것이므로 **Error를 처리하려고 하지말자.**

> Error는 런타임중에서 잡기는 힘들고 코드 수정이나, 호환성 있는 라이브러리를 사용한다거나 해야한다.


## Error와 Exception의 차이

- Error는 애플리캐이션에서 포착하려고 시도해서는 안되는 심각한 문제를 나타낸다.
- Exception은 애플리케이션에서 포착하려고 시도할 수 있는 조건을 나타낸다.

---

## Exception의 계층 구조

모든 Exception과 Error유형은 Throwable클래스의 서브클래스이다. 

2개의 분기 중 한 개 분기는 Exception 클래스이고 예외 조건에 사용된다. 예를 들면 NullPointerException이 Exception의 예이다.
또 다른 분기는 Error인데 Java 런타임 시스템(JVM)에서 런타임 환경 자체(JRE)와 관련된 오류를 나타내는 데 사용된다.
StackOverflowError가 이러한 오류의 예이다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/0c80144f-cc79-44df-ae4d-d728c00014b6)

### Exception의 유형

Java는 다양한 클래스 라이브러리와 관련된 여러 유형의 예외를 정의하고있다.
또한 Java는 사용자가 직접 예외를 정의할 수 있도록 허용한다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/3ce0925e-65a7-4833-ab52-969b43db05c8)


### 예외의 두 가지 방식으로 분류할 수 있다.

1. 기본 제공 예외
    - Checked Exception
    - Unchecked Exception
2. 사용자 정의 예외

---

## A. 기본 제공 예외

기본적으로 제공되는 예외는 Java라이브러리에서 사용할 수 있는 예외이다.
이러한 예외는 특정 오류 상황을 설명하는 적합하다.

#### Checked Exception:

RuntimeException을 상속하지 않는 예외들을 말하는데, Checked Exception이 발생할 수 있는 메소드를 사용할 경우,
복구가 가능한 예외들이기 때문에 반드시 예외처리하는 코드를 작성해야한다.
이때 해결하지 않으면 컴파일 시 Checked Exception이 발생한다.(컴파일 시 발생하는 예외라 Compile Time Exception이라고도 한다.)

> chtch문으로 예외를 잡거나 throws로 자신을 호출한 클래스로 예외를 던지는 방법으로 해결한다.
> 
> 대표적인 Exception: IOException, SQLException

#### Unchecked Exception:

컴파일러는 컴파일 시 이러한 예외를 검사하지 않는다. 
간단히 말해, 프로그램이 체크되지 않은 예외를 던졌을 때 이를 처리하거나 선언하지 않더라도 컴파일 오류가 발생하지 않는다.
프로그램에 오류가 있을 때 발생하도록 의도된 것이다.

> 대표적인 Exception: NullPointerException, IllegalArgumentException


---

## B. 사용자 정의 예외

Java의 기본 제공 예외로는 특정 상황을 설명할 수 없는 경우가 있다.
이러한 경우 사용자가 'User-Defined Exception'이라고 하는 예외를 생성할 수 있다.

---

## Java에서 예외 처리의 장점

1. 프로그램 실행을 완료하기 위한 조항 제공
2. 프로그램 코드 및 오류 처리 코드의 손쉬운 식별
3. 오류 전파
4. 의미 있는 오류 보고
5. 예외 정보를 인쇄하는 방법

---

## 예외 정보를 출력하는 메서드

### 1. printStackTrace() 

Exception이 어디서 발생했는지 추적해서 알려준다.
내가 주로 사용하는 방법이다.

```java
//program to print the exception information using printStackTrace() method

import java.io.*;

class GFG {
	public static void main (String[] args) {
	int a=5;
	int b=0;
		try{
		    System.out.println(a/b);
		}
	    catch(ArithmeticException e){
		    e.printStackTrace();
	    }
	}
}
```

#### OutPut:
```console
java.lang.ArithmeticException: / by zero
at GFG.main(File.java:10)
```

### 2. toString()

설명 형식으로 보여준다.

```java
//program to print the exception information using printStackTrace() method

import java.io.*;

class GFG1 {
	public static void main (String[] args) {
	int a=5;
	int b=0;
		try{
		    System.out.println(a/b);
		}
	    catch(ArithmeticException e){
            System.out.println(e.toString());
	    }
	}
}
```

#### OutPut:
```console
java.lang.ArithmeticException: / by zero
```

### 3. getMessage()

예외에 대한 정보와 설명을 보여준다.

```java
//program to print the exception information using printStackTrace() method

import java.io.*;

class GFG1 {
	public static void main (String[] args) {
	int a=5;
	int b=0;
		try{
		    System.out.println(a/b);
		}
	    catch(ArithmeticException e){
            System.out.println(e.getMessage());
	    }
	}
}
```

#### OutPut:
```console
/ by zero
```

---

## JVM의 Exception처리 방법

### 기본 예외 처리

메서드 내에서 예외가 발생하면 메서드는 예외 객체라는 객체를 생성하여 런타임 시스템(JVM)에 넘긴다.
예외 객체에는 예외의 이름과 설명, 예외가 발생한 프로그램의 현재 상태가 포함된다. 
예외 객체를 생성하고 런타임 시스템에서 이를 처리하는 것을 예외를 던진다라고 표현한다.
예외가 발생한 메서드에 도달하기 위해 호출된 메서드 목록도 있다.
이렇게 순서대로 나열된 메소드 목록을 Call Stack이라고 한다.

### Exception은 다음과 같은 절차를 따른다.

1. 런타임 시스템은 호출 스택을 검색하여, 예외가 발생한 코드 블록이 포함된 메서드를 찾는다. 이 코드 블록을 예외 처리기라고 한다.
2. 런타임 시스템은 Exception이 발생한 메서드부터 검색을 시작하여 메서드가 호출된 역순으로 메서드를 스택을 타고 올라간다.
3. 적절한 핸들러를 찾으면 Exception을 적절한 핸들러로 전달한다. (적절한 핸들러란 Exception을 처리할 수 있는 객체와 유형이 일치하는 것을 의미함)
4. 런타임 시스템이 호출 스택의 모든 메서드를 검색해도 적절한 처리기를 찾을 수 없을 경우 런타임 시스템은 예외 객체를 런타임 시스템의 일부인 default Exception handler로 넘긴다.
5. 위 Handler에서 예외 정보를 출력하고 프로그램을 비정상 종료한다.

아래 이미지는 호출 스택의 흐름이다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/ac52f0eb-c0f1-4c3d-aacb-4ab2d1821aab

```java
// Java Program to Demonstrate How Exception Is Thrown

// Class
// ThrowsExecp
class GFG {

	// Main driver method
	public static void main(String args[])
	{
		// Taking an empty string
		String str = null;
		// Getting length of a string
		System.out.println(str.length());
	}
}
```

Output:

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/cf2a1d28-df34-4723-a683-34beb08e3b6e)


---

### 호출 스택 예외처리 방법

```java
// Java Program to Demonstrate Exception is Thrown
// How the runTime System Searches Call-Stack
// to Find Appropriate Exception Handler

// Class
// ExceptionThrown
class GFG {

	// Method 1
	// It throws the Exception(ArithmeticException).
	// Appropriate Exception handler is not found
	// within this method.
	static int divideByZero(int a, int b)
	{

		// this statement will cause ArithmeticException
		// (/by zero)
		int i = a / b;

		return i;
	}

	// The runTime System searches the appropriate
	// Exception handler in method also but couldn't have
	// found. So looking forward on the call stack
	static int computeDivision(int a, int b)
	{

		int res = 0;

		// Try block to check for exceptions
		try {

			res = divideByZero(a, b);
		}

		// Catch block to handle NumberFormatException
		// exception Doesn't matches with
		// ArithmeticException
		catch (NumberFormatException ex) {
			// Display message when exception occurs
			System.out.println(
				"NumberFormatException is occurred");
		}
		return res;
	}

	// Method 2
	// Found appropriate Exception handler.
	// i.e. matching catch block.
	public static void main(String args[])
	{

		int a = 1;
		int b = 0;

		// Try block to check for exceptions
		try {
			int i = computeDivision(a, b);
		}

		// Catch block to handle ArithmeticException
		// exceptions
		catch (ArithmeticException ex) {

			// getMessage() will print description
			// of exception(here / by zero)
			System.out.println(ex.getMessage());
		}
	}
}
```

### Output
```
/ by zero
```

---

## 프로그래머의 예외처리 방법

### 사용자 정의 예외처리:

Java의 예외처리는 try, catch, throw, throws 와 finally로 관리된다.
예외를 발생시킬 수 있다고 생각되는 프로그램은 try블록 내에 포함하면 되고, 
try블록 내에서 예외가 발생하면 예외가 던져진다. 
예외는 catch블록에서 포착하여 알맞은 방식으로 처리하게 된다.
또한 예외를 수동으로 던지려면 throw키워드를 사용한다.
finally블록은 Exception이 발생하고 알맞은 처리까지 진행 후 마지막으로 진행할 과정을 명시한다.

> finally는 JDK 18에서 사라지는것으로 알고 있다.

try-catch문을 살펴보자.

















