package main

import (
	"errors"
	"fmt"
	"strings"
)

// struct
type person struct {
	name    string
	age     int
	favFood []string
}

/********************BankAccount Project********************/

// Account struct 생성
type Account struct {
	owner   string // 소문자로 표시 -> private
	balance int    // 소문자로 표시 -> private
}

// error 설정
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// Go에는 Exception이 없고 직접 error코드를 직접 써야됨
// error에는 두가지 value가 있음 - error, nil
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Go에서 자동으로 호출해주는 Method인 String
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

/******************** BankAccount Project - END ********************/

/********************Dictionary Project********************/
// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errCantUpdate = errors.New("Cant update non-existing word")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

/******************** Dictionary Project - END ********************/

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
} // GO는 return값을 여러개 받을 수 있다.

func repeatMe(words ...string) {
	fmt.Println(words)
} // Type앞에 ... 입력해주면 값을 무제한으로 받을 수 있다.

func lenAndUppercase(name string) (lenght int, uppercase string) {
	// defer는 function이 값을 return하고 나면 실행된다.
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
} // return 값을 미리 정해주면 return 값 뒤에 따로 명시할 필요는 없다. -> naked return

// 반복문 for
func superAdd(numbers ...int) int {
	/*
		for index, number := range numbers {
			fmt.Println(index, number)
		}
		return 1
	*/ // range는 index값을 의미한다. range는 for안에서만 사용가능

	/*
		for i := 0; i < len(numbers); i++ {
			fmt.Println(numbers[i])
		}
		return 1
	*/

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// if
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true // else부분을 굳이 쓰지 않고 return 할수 있고, Go는 if안에 다른변수를 넣어서 조건을 만들수 있다.
}

// Switch
func canIDrive(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false // switch도 if와 마찬가지로 변수를 넣어서 조건을 지정할수 있다.
	// if~elseif... elseif... -> switch와 같은 의미
}

func main() {
	fmt.Println("Hello Golang!!!")
	/*
		math.Phi(0)
		something.SayHello() // 대문자로 시작하면 export 가능
		something.sayBye()
	*/

	// 상수
	const name string = "nico"
	// name = "Lynn" 에러발생, 상수는 값을 변경할 수 없다.
	fmt.Println(name)

	// 변수
	nick := "zico" // var nick string = "zico" 이거와 같은 표현이다. 축약형은 func안에서만 가능하다.
	nick = "Lean"
	fmt.Println(nick)

	fmt.Println(multiply(2, 2))

	totalLenght, upperName := lenAndUpper("song")
	fmt.Println(totalLenght, upperName) // multiple value를 반환하는 방법

	// totalLenght만 값을 받고 싶은 경우, _로 적으면 값을 무시한다는 뜻
	// totalLenght, _ := lenAndUpper("song")  뒤에 upperName은 _이기 때문에 값을 무시
	// fmt.Println(totalLenght)

	repeatMe("song", "zico", "nico", "cjh")

	// naked return
	totalLenght, up := lenAndUppercase("nico")
	fmt.Println(totalLenght, up)

	// superAdd(1, 2, 3, 4, 5, 6)

	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result)

	// if
	fmt.Println(canIDrink(16))

	// switch
	fmt.Println(canIDrive(16))

	// pointers
	a := 2
	b := &a // &를 붙이면 메모리 주소를 확인할 수 있음. b에 a의 주소를 저장
	a = 5
	*b = 20
	//fmt.Println(*b) *를 붙이면 훝어본다는 느낌, b의 값은 a의 메모리주소이지만, *붙어서 a의값 2가 출력됨
	fmt.Println(a) // a와b는 연결되어있기 때문에, *b의 값을 변경해주면 a의 값도 변경됨(20)
	// & -> 메모리 주소확인
	// * -> 주소를 살펴보고 주소에 담긴 값을 확인, 주소에 *를 써서 주소에 담긴 값도 변경할 수 있다

	// Array(배열)
	//names := [5]string{"nico", "anna", "coke"}
	//names[3] = "jisu"
	//names[4] = "rose"
	//fmt.Println(names)

	// Go에서 Array 크기에 제한을 주고 싶거나, 제한이 없게 하고 싶으면 []로 표시 -> Slice기능
	names := []string{"nico", "anna", "coke"}
	names = append(names, "jisu") // Slice한 배열에 append로 names의 배열에 값을 추가,수정할수 있다.
	fmt.Println(names)

	// map
	/*
		nico := map[string]string{"name": "nico", "age": "12"} // key, value
		for key, value := range nico {
			fmt.Println(key, value)
		}
	*/

	// struct
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)

	// BankAccount Project
	account := NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	// Dictionary Project

	// Search for a word
	dictionary := Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// Add a word to the dictionary
	word := "hello"
	definition2 := "Greeting"
	err2 := dictionary.Add(word, definition2)
	if err2 != nil {
		fmt.Println(err2)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)
	err3 := dictionary.Add(word, definition2)
	if err3 != nil {
		fmt.Println(err3)
	}

	// Update a word
	baseWord := "hi"
	dictionary.Add(baseWord, "First")
	err4 := dictionary.Update(baseWord, "Second")
	if err4 != nil {
		fmt.Println(err4)
	}
	word1, _ := dictionary.Search(baseWord)
	fmt.Println(word1)

	// Delete a word
	baseWord2 := "bye"
	dictionary.Add(baseWord2, "First")
	dictionary.Search(baseWord2)
	dictionary.Delete(baseWord2)
	word, err5 := dictionary.Search(baseWord2)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println(word)
	}

}
