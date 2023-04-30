# Functions vs methods

ent공부를 하다보니 내가 알지못한 방법의 문법이 나오는거다...? 이건 뭐지 하고 찾아봤는데 golang에서는 function과 method가 약간 다르다.

개인적으로 method는 Java에서 객체를 사용하는 느낌이 들었다.

전반적인 내용은 [Golang functions vs methods 👯‍♀️](https://www.sohamkamani.com/golang/functions-vs-methods/)보고 작성한 내용이다.

## Syntax 

### 선언

#### Function

function은 인수의 유형, 반환 값, 함수 본문으로 구성되어 있다.(내가 기본적으로 알고있던 사용법)

```go
type Person struct {
  Name string
  Age int
}

// This function returns a new instance of `Person`
func NewPerson(name string, age int) *Person {
  return &Person{
    Name: name,
    Age: age,
  }
}
```

---

#### Method

method는 추가적으로 receiver부분(객체지향에서 클래스에 해당)을 명시해준다.

```go
// The `Person` pointer type is the receiver of the `isAdult` method
func (p *Person) isAdult() bool {
  return p.Age > 18
}
```

메서드 이름 앞에 (p *Person)이 추가되었으며, *Person클래스에 isAdult()라는 메서드를 선언했다고 보면된다.

---

### 실행

함수는 지정된 인수를 사용하여 독립적으로 호출되고 메서드는 receiver유형에 따라 호출된다.
```go
p := NewPerson("John", 21)

fmt.Println(p.isAdult())
// true
```

NewPerson()은 독립적으로 실행되었으며 p.isAdult()는 p에서 실행된 것을 확인할 수 있다.

---

## 교환성??

function과 method는 서로 바꿀 수 있다. 예를 들어 isAdult메서드를 함수로 바꾸고 NewPerson함수를 메서드로 사용할 수 있다.

```go
type PersonFactory struct {}

// `NewPerson` is now a method of a `PersonFactory` struct
func (p *PersonFactory) NewPerson(name string, age int) *Person {
  return &Person{
    Name: name,
    Age: age,
  }
}

// `isAdult` is now a function where we're passing the `Person` as an argument
// instead of a receiver
func isAdult(p *Person) bool {
  return p.Age > 18
}
```

이 경우 실행 구문이 이상해 보인다.

```go
factory := &PersonFactory{}

p := factory.NewPerson("John", 21)

fmt.Println(isAdult(p))
// true
```

위 코드는 불필요하고 복잡하게 짜여 있다. 이는 메서드와 함수의 차이는 대부분 구문상의 차이이며, 사용 사례에 적절한 추상화를 사용해야 한다는걸 의미한다.

---

## 사용 예

### Method chaining

메서드의 유용한 속성 중 하나가 코드를 깔끔하게 유지하면서 메서드를 서로 연결할 수 있다는 점이다. 메서드 체이닝을 사용하여 Person의 몇 가지 속성을 설정하는 예를 보자.

```go
type Person struct {
	Name string
	Age  int
}

func (p *Person) withName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) withAge(age int) *Person {
	p.Age = age
	return p
}

func main() {
	p := &Person{}

	p = p.withName("John").withAge(21)
	
  fmt.Println(*p)
  // {John 21}
}
```

### Stateful vs stateless execution

- 교환성?? 예제에서 새로운 Person인스턴스를 생성하기위해 PersonFactory 객체를 사용하는 보았을 것이다. 결과적으로 이것은 안티패턴(anti-pattern)이므로 피해야한다.

- stateless를 실행하기 위해서는 앞서 선언한 NewPerson함수와 같이 function을 사용하는것이 훨씬 낫다.

> stateless란 동일한 입력에 대해 항상 동일한 출력을 반환하는 모든 코드를 의미한다.

**※결론은 function이 특정 타입의 값을 많이 읽거나 수정하는 경우 해당 타입의 method여야 한다.**

---

### Semantics

시멘틱은 코드를 읽는 방법에 대한 것이다. 아래 소리내어 읽으면 어느쪽이 이해하기 쉬울까?

```go
customer := NewPerson("John", 21)

// Method
customer.isAdult()

// Function
isAdult(customer)
```

```customer.isAdult()```와 ```isAdult(customer)``` 두개의 코드 중 "Is the customer an adult?"라고 질문할 때 더 잘 읽힌다. 또한 "is x and adult?"라는 질문에 항상 사람 x의 컨텍스트에서 질문하게 된다.


## 결론

Go에서 함수와 메서드의 주요 차이점과 사용 사례 몇 가지를 살펴보았지만 항상 예외는 존재한다.! 이러한 규칙을 정석으로 받아들이지 않는 것이 중요하다.

결국 함수와 메서드의 차이점은 결과 코드가 어떻게 읽히는지에 있으며, 한 가지 방식이 다른 방식보다 더 잘 읽힌다고 느낀다면, 그것이 바로 올바른 추상화 일것이다.Go에서 함수와 메서드의 주요 차이점과 사용 사례 몇 가지를 살펴보았지만 항상 예외는 존재합니다! 이러한 규칙을 정석으로 받아들이지 않는 것이 중요합니다.

결국 함수와 메서드의 차이점은 결과 코드가 어떻게 읽히는지에 있습니다. 여러분이나 여러분의 팀이 한 가지 방식이 다른 방식보다 더 잘 읽힌다고 느낀다면, 그것이 바로 올바른 추상화입니다!