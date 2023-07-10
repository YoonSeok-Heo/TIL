# Serialization and Deserialization

직렬화(Serialization)는 객체를 바이트 스트림으로 변환하는 작업이고, 
역질렬화(Deserialization)는 바이트 스트림을 java객체로 다시 생성하는 역방향 프로세스이다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/cc56da6d-a426-469f-80a1-83d020b6b64d)

생성된 바이트 스트림은 플랫폼에 독립적이다.
따라서 한 플랫폼에서 직렬화된 객체를 다른 플랫폼에서 역직렬화할 수 있다.

java객체를 직렬화할 수 있게 하려면 java.io.Serializable 인터페이스를 구현해야한다.
ObjectOutputStream 클래스에는 객체를 직렬화하기 위한 writeObject() 메서드가 포함되어 있다.

```java
public final void writeObject(Object obj) throws IOException
```

---

ObjectInputStream 클래스에는 객체를 역직렬화하기 위한 readObject() 메서드가 포함되어 있습니다.

```java
public final Object readObject() throws IOException, ClassNotFoundException
```

---

## 직렬화의 장점

1. 오브젝트의 상태를 저장/유지할 수 있다.
2. 네트워크를 통해 오브젝트를 보내고 받을 수 있다.

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/3019d3a5-9692-4c22-bb4e-1e04f60f366d)

java.io.Serializable 인터페이스를 구현하는 클래스의 객체만 직렬화할 수 있다.
Selializable은 마커 인터페이스(데이터 멤버와 메서드가 없음)이며, 
이 인터페이스는 자바 클래스를 'mark'하여 해당 클래스의 객체가 특정 기능을 사용할 수 있도록 하는데 사용된다.

### 마커 인터페이스의 다른 예:
- Cloneable
- Remote


---

### 중요한점.

1. 부모 클래스가 Serializable인터페이스를 구현한 경우 자식 클래스는 이를 구현할 필요가 없지만, 반대의 경우는 해당하지 않는다.
2. 정적(static)이 아닌 데이터 멤버만 직렬화 프로세스를 통해 저장된다.
3. 정적 데이터 멤버와 일시적 데이터 멤버는 직렬화 프로세스를 통해 저장되지 않는다.
따라서 정적이 아닌 데이터 멤버의 값을 저장하지 않으려면 일시적인 데이터 멤버로 만들어야한다.
4. 객체가 역직렬화될 때 객체의 생성자는 호출되지 않는다.
5. 연관된 객체는 직렬화 가능한 인터페이스를 구현해야 한다.

```java
// 5번의 예
class A implements Serializable{

    // B also implements Serializable
    // interface.
    B ob=new B();
}
```

---

**SerialVersionUID**직렬화 런타임은 각 직렬화 가능 클래스에 버전 번호를 연결하는데. 
이 버전 번호는 직렬화된 객체의 발신자와 수신자가 직렬화와 관련하여 호환되는 해당 객체에 대한 클래스를 로드했는지 확인하기 위해 역직렬화 중에 사용된다.
수신자가 해당 발신자의 클래스와 다른 UID를 가진 객체에 대한 클래스를 로드한 경우 역직렬화 시 `InvalidClassException`이 발생한다.

> 내가 이해한 내용으로 설명하자면 `SerialVersionUID`는 키값과 같이 이해하고 있다.
> 
> `SerialVersionUID`를 명시하지 않으면 컴파일러가 임의의 값(컴파일러마다 키값이 다르다는걸 의미)을 부여하는데, 
> 만약 `SerialVersionUID'명시하지 않고 다른 컴파일러로 직렬화한 데이터를 역질렬화 하게 된다면 키값이 달라서 `InvalidClassException`이 발생한다고 볼 수 있다는 것이다.
> 그러니 다른 컴파일러에서 역직렬화하기 위해선 'SerialVersionUID'를 동일하게 해줘야할 의무가 있다.


---

## Example

### 1.
```java
// Java code for serialization and deserialization
// of a Java object
import java.io.*;

class Demo implements java.io.Serializable
{
	public int a;
	public String b;

	// Default constructor
	public Demo(int a, String b)
	{
		this.a = a;
		this.b = b;
	}

}

class Test
{
	public static void main(String[] args)
	{
		Demo object = new Demo(1, "geeksforgeeks");
		String filename = "file.ser";
		
		// Serialization
		try
		{
			//Saving of object in a file
			FileOutputStream file = new FileOutputStream(filename);
			ObjectOutputStream out = new ObjectOutputStream(file);
			
			// Method for serialization of object
			out.writeObject(object);
			
			out.close();
			file.close();
			
			System.out.println("Object has been serialized");

		}
		
		catch(IOException ex)
		{
			System.out.println("IOException is caught");
		}


		Demo object1 = null;

		// Deserialization
		try
		{
			// Reading the object from a file
			FileInputStream file = new FileInputStream(filename);
			ObjectInputStream in = new ObjectInputStream(file);
			
			// Method for deserialization of object
			object1 = (Demo)in.readObject();
			
			in.close();
			file.close();
			
			System.out.println("Object has been deserialized ");
			System.out.println("a = " + object1.a);
			System.out.println("b = " + object1.b);
		}
		
		catch(IOException ex)
		{
			System.out.println("IOException is caught");
		}
		
		catch(ClassNotFoundException ex)
		{
			System.out.println("ClassNotFoundException is caught");
		}

	}
}
```

#### output
```
Object has been serialized
Object has been deserialized 
a = 1
b = geeksforgeeks
```

---

### 2. 
```java
// Java code for serialization and deserialization
// of a Java object
import java.io.*;

class Emp implements Serializable {
private static final long serialversionUID =
								129348938L;
	transient int a;
	static int b;
	String name;
	int age;

	// Default constructor
public Emp(String name, int age, int a, int b)
	{
		this.name = name;
		this.age = age;
		this.a = a;
		this.b = b;
	}

}

public class SerialExample {
public static void printdata(Emp object1)
	{

		System.out.println("name = " + object1.name);
		System.out.println("age = " + object1.age);
		System.out.println("a = " + object1.a);
		System.out.println("b = " + object1.b);
	}

public static void main(String[] args)
	{
		Emp object = new Emp("ab", 20, 2, 1000);
		String filename = "shubham.txt";

		// Serialization
		try {

			// Saving of object in a file
			FileOutputStream file = new FileOutputStream
										(filename);
			ObjectOutputStream out = new ObjectOutputStream
										(file);

			// Method for serialization of object
			out.writeObject(object);

			out.close();
			file.close();

			System.out.println("Object has been serialized\n"
							+ "Data before Deserialization.");
			printdata(object);

			// value of static variable changed
			object.b = 2000;
		}

		catch (IOException ex) {
			System.out.println("IOException is caught");
		}

		object = null;

		// Deserialization
		try {

			// Reading the object from a file
			FileInputStream file = new FileInputStream
										(filename);
			ObjectInputStream in = new ObjectInputStream
										(file);

			// Method for deserialization of object
			object = (Emp)in.readObject();

			in.close();
			file.close();
			System.out.println("Object has been deserialized\n"
								+ "Data after Deserialization.");
			printdata(object);

			// System.out.println("z = " + object1.z);
		}

		catch (IOException ex) {
			System.out.println("IOException is caught");
		}

		catch (ClassNotFoundException ex) {
			System.out.println("ClassNotFoundException" +
								" is caught");
		}
	}
}
```

### output
```
Object has been serialized
Data before Deserialization.
name = ab
age = 20
a = 2
b = 1000
Object has been deserialized
Data after Deserialization.
name = ab
age = 20
a = 0
b = 2000
```


> 두번째 예제에서 객체를 역직렬화하는 동안 a와 b의 값이 변경되었다.
> 이는 a는 `transient`변수이고 b는 `static`변수이기 때문이다.
> 
> `transient`변수는 직렬화 과정에서 직렬화가 되지 않으며, 역직렬화 과정에서 기본값으로 초기화 된다.
> 
> `static`변수는 역직렬화 중에 클래스에 정의된 현재 값으로 로드된다.

---

### 3.
```java
//java code for final with transient

import java.io.*;


class Dog implements Serializable{
		int i=10;
		transient final int j=20;
}
class GFG {
	public static void main (String[] args)throws IOException,ClassNotFoundException
	{
		Dog d1=new Dog();
		//Serialization started
		System.out.println("serialization started");
		FileOutputStream fos= new FileOutputStream("abc.ser");
		ObjectOutputStream oos=new ObjectOutputStream(fos);
		oos.writeObject(d1);
		System.out.println("Serialization ended");
	
		//Deserialization started
		System.out.println("Deserialization started");
		FileInputStream fis=new FileInputStream("abc.ser");
		ObjectInputStream ois=new ObjectInputStream(fis);
		Dog d2=(Dog) ois.readObject();
		System.out.println("Deserialization ended");
		System.out.println("Dog object data");
		//final result
		System.out.println(d2.i+"\t" +d2.j);
	}
}
```

### output
```
serialization started
Serialization ended
Deserialization started
Deserialization ended
Dog object data
10    20
```