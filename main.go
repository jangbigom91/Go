package main

import (
	"fmt"
	"strings"
)

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
}
