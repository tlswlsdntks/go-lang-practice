package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"runtime"
	"strconv"
)

func main() {
	/*
		실행 순서
			go build ./main.go
			./main.exe
			go run main.go

		go-path
			go env
			go env | grep GOPATH
			C:\Users\tlswl\go 구성
			bin: 빌드 후 실행파일 저장
			src: 작성 공간
			pkg: 프로그램 빌드할 때 생성된 파일?
	*/
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
	// 왈러스 기호
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

	/*
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
	*/

	// 형 변환
	var float1 float64 = 3.14
	fmt.Printf("형 변환: %T \n", int(float1))

	// 타입 생성
	type MyInt int
	fmt.Printf("타입 생성: %T \n", MyInt(10))

	// geth 코드 - 타입 사용 예시
	// go-ethereum\consensus\ethash\consensus.go, line: 42
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
	// go-ethereum\common\math\big.go, line: 32
	fmt.Println(strconv.FormatUint(uint64(^big.Word(0))>>63, 2))           // 1: 0000 0001
	fmt.Println(strconv.FormatUint(32, 2))                                 // 32: 0010 0000
	fmt.Println(strconv.FormatUint(32<<(uint64(^big.Word(0))>>63), 2))     // 64: 0100 0000
	fmt.Println(strconv.FormatUint((32<<(uint64(^big.Word(0))>>63))/8, 2)) // 8: 0000 1000

	// 배열 선언 및 초기화
	// var arr [3]int
	arr := [3]int{1, 2, 3}

	// 배열 타입/길이
	fmt.Printf("%T \n", arr)
	fmt.Println(len(arr))

	// 다중 배열
	// var arr2 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr2 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	arr2[2][2] = 10
	fmt.Println(arr2[1][1])
	fmt.Println(arr2)

	// geth 코드 - 배열 사용 예시
	// go-ethereum\crypto\blake2b\blake2b_generic.go, line: 15

	// 슬라이스: 동적 배열
	// 슬라이스 선언 및 초기화
	// var s []int
	s := []int{1, 2, 3}

	// make(): 0 초기화
	// s := make([]int, 3)

	// 슬라이스 타입/길이
	fmt.Printf("%T \n", s)
	fmt.Println(len(s))

	// 슬라이스에 값 추가
	s = append(s, 4)
	fmt.Println(s)

	// 슬라이스 용량 지정
	s2 := make([]int, 5, 5)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// 슬라이스 값 복사
	dst := make([]int, 10)
	cp := []int{1, 2, 3, 4, 5}
	fmt.Println(dst)
	copy(dst, cp)
	fmt.Println(dst)

	// 배열 슬라이싱: 슬라이스 변환
	arr = [3]int{1, 2, 3}
	slice := arr[0:3]
	// fmt.Println(reflect.TypeOf(slice))
	fmt.Printf("%T", slice)

	// 슬라이스 슬라이싱
	fmt.Println(slice[1:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[:])
	fmt.Println(slice[:3])
	fmt.Println(slice[0:1:3]) // [str, stop, cap]

	// geth 코드 - 슬라이스 사용 예시
	// go-ethereum\common\types.go, line: 319

	// 맵: 키와 값 형태이며 순서를 보장하지 않는다.
	// 맵 선언 및 초기화
	var numMap = make(map[string]int)
	numMap["one"] = 1
	numMap["two"] = 2
	numMap["three"] = 3
	numMap["four"] = 4

	numMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	fmt.Println(numMap)

	// 요소 추가
	numMap["five"] = 5

	// 요소 제거
	delete(numMap, "five")

	// 요소 접근
	fmt.Println(numMap["one"])

	// 요소 존재 여부 확인
	value, isKey := numMap["one"]
	fmt.Println(isKey, value)
	value, isKey = numMap["five"]
	fmt.Println(isKey, value)

	// geth 코드 - 맵 사용 예시
	// go-ethereum\eth\api_backend.go, line: 333

	// if 조건문
	/*
		if 초기문; 조건문 {
			실행문
		}

		if 조건문{
			실행문
		} else if {
			실행문
		} else {
			실행문
		}
	*/
	if true {
		fmt.Println("Hello World!")
	}
	if num0 := 0; num0 < 1 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World!")
	}

	// switch 조건문
	switch os := runtime.GOOS; os {
	case "windows.":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}

	// geth 코드 - 조건문 사용 예시
	// go-ethereum\node\defaults.go, line: 80

	// 반복문
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1_000 {
		sum += sum
	}
	fmt.Println(sum)

	i := 1
	for {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i, v := range [5]int{1, 2, 3, 4, 5} {
		fmt.Printf("index: %d, value: %d \n", i, v)
	}

	// geth 코드 - 조건문 사용 예시
	// go-ethereum\cmd\geth\main.go, line: 372

	/*	구조체 정의
		type 타입명 struct {
			필드명 타입
			필드명 타입
		}
	*/
	type Vertex struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	// 구조체 변수 선언 및 초기화
	var v1 Vertex = Vertex{1, 2}
	var v2 Vertex = Vertex{X: 1, Y: 2}
	var v3 Vertex = Vertex{X: 1}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3) // {1 0}

	// 구조체 필드 접근
	v := Vertex{X: 10}
	v.X = 10
	v.Y = 20
	fmt.Printf("X: %d, Y: %d \n", v.X, v.Y)

	// 구조체의 json 형태 제공
	data, _ := json.Marshal(v) // _(underscore): key 를 사용하지 않겠다는 의미
	fmt.Println(string(data))  // {"x":10,"y":20}

	// geth 코드 - 구조체 사용 예시
	// go-ethereum\core\types\block.go, line: 206

}
