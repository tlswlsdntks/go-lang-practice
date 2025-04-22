package main

import (
	"fmt"
	"math/big"
	"strconv"
)

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
	// 왈루스 기호
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
	fmt.Printf("형 변환: %T \n", int(float1))

	// 타입 생성
	type MyInt int
	fmt.Printf("타입 생성: %T \n", MyInt(10))

	// geth 코드 - 타입 사용 예시
	// consensus\ethash\consensus.go
	// *ethash: 현재 사용하지 않는 합의 알고리즘

	// 산술 연산자
	a, b := 10, 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	c := a % b
	fmt.Println(c)

	// 비교 연산자
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

	// 대입 연산자
	a = b
	fmt.Println(a)

	// 논리 연산자
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(!true)

	// 증감 연산자
	a, b = 10, 20
	a++
	b--
	b--
	b--
	b--
	fmt.Println(a, b)

	// 비트 연산자
	a = 10             // 0000 1010
	b = 30             // 0001 1110
	fmt.Println(a & b) // 10: 0000 1010
	fmt.Println(a | b) // 30: 0001 1110
	fmt.Println(a ^ b) // 20: 0001 0100
	fmt.Println(^b)    // 225: 1110 0001
	// &^: 각 비트 단위로 첫 번째 자릿값이 1이고, 두번째 자리값이 0인 경우에만 1을 반환한다.
	fmt.Println(a &^ b) // 0: 0000 0000
	fmt.Println(b &^ a) // 20: 0001 0100

	// 이진수 변환
	fmt.Println(strconv.FormatInt(int64(a^b), 2)) // 20: 0001 0100

	// 시프트 연산자
	a = 10              // 0000 1010
	fmt.Println(a << 3) // 80: 0101 0000
	fmt.Println(a >> 3) // 1: 0000 0001

	// 연산자 우선순위
	/**
		precedence  Operator
		5			* / % << >> & &^
		4			+ - | ^
		3			== != < <= > >=
		2			&&
		1			||
	**/

	// geth 코드 - 연산자 사용 예시
	// common\math\big.go
	fmt.Println(strconv.FormatUint(uint64(^big.Word(0))>>63, 2))           // 1: 0000 0001
	fmt.Println(strconv.FormatUint(32, 2))                                 // 32: 0010 0000
	fmt.Println(strconv.FormatUint(32<<(uint64(^big.Word(0))>>63), 2))     // 64: 0100 0000
	fmt.Println(strconv.FormatUint((32<<(uint64(^big.Word(0))>>63))/8, 2)) // 8: 0000 1000

	// 배열 선언
	// var arr [3]int = [3]int{1, 2, 3}
	arr := [3]int{1, 2, 3}

	// 배열 타입/길이
	fmt.Printf("%T \n", arr)
	fmt.Println(len(arr))

	// 배열에 값 저장
	arr[0] = 1

	// 배열 읽기
	fmt.Println(arr)
	fmt.Println(arr[0]) // 1

	// 다중 배열
	// var arr2 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr2 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(arr2[1][1]) // 5
	arr2[2][2] = 10
	fmt.Println(arr2)

	// geth 코드 - 배열 사용 예시
	// crypto\blake2b\blake2b_generic.go
	// var precomputed = [10][16]byte{}
}
