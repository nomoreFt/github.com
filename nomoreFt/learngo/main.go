package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}
	return result
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:

		return false
	}
	return false
}

func main() {
	name := "nico"
	name = "lynn"
	fmt.Println(name)
	//변수 return 무시
	totalLenght, _ := lenAndUpper("nico")
	fmt.Println(totalLenght)

	repeatMe("a", "b", "c", "d")
	fmt.Println(totalLenght)

	//func - defer, multy return, return 변수 생성
	fmt.Println(lenAndUpper2("hyunwoo"))

	//for - range
	result := superAdd(1, 2, 3, 4, 5, 6, 7, 7)
	fmt.Println(result)

	//if else variable Expression
	fmt.Println(canIDrink(13))

	//switch
	fmt.Println(canIDrink2(16))

	//pointer 높은 수준 코드로 low-level-programming 가능
	a := 2
	b := &a
	*b = 202020

	fmt.Println(a, *b, &a, b, &b)
	//* 메모리의 값, & 해당 메모리주소
	//메모리 주소를 할당받은 pointer는 *으로 해당 메모리 주소의 값을 바꿀 수 있음

	//Arrays
	naming := [5]string{"hi", "my", "name"}
	naming[4] = "is"

	fmt.Println(naming)
	//slices
	naming2 := []string{"hi", "my", "name"}
	naming2 = append(naming2, "is")
	naming2 = append(naming2, "hyunwoo")

	fmt.Println(naming2)

	//maps
	hyunwoo := map[string]string{"name": "hw", "age": "12"}
	fmt.Println(hyunwoo)
	for _, value := range hyunwoo {
		fmt.Println(value)
	}

	//struct

	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)

}

type person struct {
	name    string
	age     int
	favFood []string
}
