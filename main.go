package main

import "fmt"

func main() {
	/**
		실행 순서
		go build ./main.go
		./main.exe
		go run main.go

		GO-PATH
		go env
		go env | grep GOPATH
		C:\Users\tlswl\go 구성
		bin: 빌드 후 실행파일 저장
		src: 작성 공간
		pkg: 프로그램 빌드할 때 생성된 파일?
	**/
	fmt.Println("Hello, World!")

	var num int
	fmt.Println(num)

	var num1 int = 10
	fmt.Println(num1)

	var num2 = 20
	num2 = 20
	fmt.Println(num2)

	num3 := 30
	fmt.Println(num3)

	// var num4, num5 int = 40, 50
	num4, num5 := 40, 50
	fmt.Println(num4, num5)

	var (
		num6 int = 60
		num7 int = 70
	)
	fmt.Println(num6, num7)

	const PI = 3.14
	// PI = 3.1415
	fmt.Println(PI)

	const (
		iota1 = iota
		iota2
		iota3
	)
	fmt.Println(iota1, iota2, iota3) // 0 1 2

	/**
		자료형
		bool
		int8
		int16
		int32
		int64
		unit16: unsigned(양수)
		unit32
		unit64
		float32
		float64
		int
		unit
		unitptr
		byte
		rune
		string
		complex64
		complex128
	**/

	// 형 변환
	var float1 float64 = 3.14
	fmt.Printf("%T", int(float1))

	// 타입 생성
	type MyInt int
	fmt.Println(MyInt(10))
	fmt.Printf("%T", MyInt(10))

	// 타입 사용 예시
	// C:\Users\tlswl\OneDrive\바탕 화면\이더리움(EVM) 기반 블록체인 개발 A to Z (feat. Solidity & Go) 초격차 패키지 Online\프로젝트\go-ethereum\consensus\ethash\consensus.go
	// *ethash: 현재 사용하지 않는 합의 알고리즘
}
