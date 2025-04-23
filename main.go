// 패키지 선언
package main

// 단일 패키지 가져오기
// import "fmt"

// 복수 패키지 가져오기
import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"strconv"

	myfirstpkg "github.com/tlswlsdntks/go-lang-myfirstpkg"
)

func plus(x int, y int) int {
	return x + y
}

func minus(x int, y int) (val int) {
	val = x - y
	return
}

func calc(x int, y int) (int, int) {
	return x + y, x - y
}

func myName() string {
	name := "신진우"
	return name
}

func addOne(num int) {
	num += 1
}

func addTwo(num *int) {
	*num += 1
}

type Calculator struct {
	X int
}

// 구조체의 복사본을 전달받아 사용
func (c Calculator) add(x int) int {
	c.X += x
	return c.X
}

// 구조체의 주소를 전달받아 사용
func (c *Calculator) add2(x int) int {
	/*
		c *Calculator 는 포인터 변수이다.
		*c 는 포인터가 가리키는 구조체 값을 의미하며, 이 값을 통해 필드에 접근하려면 (*c).X 라고 써야 한다.
		Go 에서는 c.X 가 (*c).X 와 동일하게 처리되기 때문에, 대부분 c.X 라고만 사용한다.
	*/
	// (*c).X += x
	c.X += x
	return c.X
}

type Ducky interface {
	DuckSound() string
	DuckWalk() string
	isSwim() string
}

type Bird struct {
	name string
}

func (b Bird) DuckSound() string {
	return "꽥꽥"
}

func (b Bird) DuckWalk() string {
	return "뒤뚱뒤뚱"
}

func (b Bird) isSwim() string {
	if b.name == "오리" {
		return "있다"
	} else {
		return "없다"
	}
}

func typeCheck(i interface{}) {
	fmt.Printf("%T \n", i)
}

// 지연된 함수의 인수는 defer 문이 평가될 때, 평가된다.
func deferA() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 지연된 함수의 호출은 주변 함수가 반환된 후, 후입선출 순서로 실행된다.
func deferB() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// 지연된 함수는 반환하는 함수의 명명된 반환 값을 읽고 할당할 수 있다.
func deferC() (i int) {
	defer func() {
		i++
		fmt.Println(i)
	}() // (): 익명함수 즉시 실행
	return 1
}

// 가변 인자: 매개변수타입은 슬라이스 타입으로 지정
func numbers(nums ...int) {
	fmt.Println(nums)
	fmt.Printf("%T \n", nums)
}

func strings(str ...string) {
	fmt.Println(str)
	fmt.Printf("%T \n", str)
}

func devide(x int, y int) int {
	return x / y
}

func devide2(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("Error: divide by zero")
	}
	return x / y, nil
}

func main() {
	// 01. 첫 golang 프로그램 실행
	/*
		실행 순서
			go build ./main.go
			./main.exe
			go run main.go

		go-path
			go env
			go env | grep GOPATH
			C:\Users\tlswl\go 를 VSCode 에서 확인
			bin: 소스 파일(패키지)를 컴파일하여 실행 파일(바이너리)이 생성되는 디렉터리이다.
			pkg:
				패키지를 컴파일하여 라이브러리 파일이 생성되는 디렉터리이고,
				pkg 디렉터리 아래에는 <운영체제>_<아키텍쳐> 형식으로 디렉터리가 생성된다.
				64비트 리눅스라면 linux_amd64 디렉터리 아래에 라이브러리 파일이 생성된다.
			src: 내가 작성한 소스 파일과 인터넷에서 자동으로 받아온 소스 파일이 저장되는 디렉터리이다.
	*/
	fmt.Println("Hello, World")

	// 02. 변수와 상수
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
		uint16: unsigned int(양수)
		uint32
		uint64
		float32
		float64
		int
		uint
		uintptr
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

	// 03. 연산자
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
	// &^: 각 비트 단위로 첫 번째 자릿값이 1이고, 두 번째 자리값이 0인 경우에만 1을 반환한다.
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

	// 04. 배열
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

	// 05. 슬라이스: 동적 배열
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

	// 06. 맵: 키와 값 형태이며 순서를 보장하지 않는다.
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

	// 07. 조건문
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
		fmt.Println("Hello World")
	}
	if num0 := 0; num0 < 1 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World")
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

	// 08. 반복문
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

	// 09. 구조체
	/*
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

	// 10. 포인터
	numb := 1
	var numpointer *int = &numb
	fmt.Println(&numb)
	fmt.Println(*numpointer)

	// 포인터로 값 변경
	*numpointer = 2
	fmt.Println(numb)

	// geth 코드 - 구조체 사용 예시
	// go-ethereum\core\types\state_account.go, line: 31

	// 11. 함수
	/*
		func 함수명(매개변수 매개변수타입) 리턴타입 {
			실행문
			return 반환값
		}
	*/
	fmt.Println(plus(42, 13))
	fmt.Println(minus(42, 13))
	fmt.Println(calc(42, 13))
	// fmt.Println(name)

	num01 := 1
	addOne(num01)
	fmt.Println(num01) // 1

	num02 := 1
	addTwo(&num02)
	fmt.Println(num02) // 2

	// geth 코드 - 함수 사용 예시
	// go-ethereum\core\types\block.go, line: 243

	// 12. 메서드: 구조체에 속한 함수
	/*
		func (리시버: 구조체) 메서드명(매개변수 매개변수타입) 반환타입 {
			실행문
			return 반환값
		}
	*/
	// 구조체를 통한 호출
	cal := Calculator{X: 10}
	fmt.Println(cal.add(11))
	fmt.Println(cal)

	fmt.Println(cal.add2(11))
	fmt.Println(cal)

	// 13. 인터페이스: 덕 타이핑
	// 사람이 오리처럼 행동하면 오리로 봐도 무방하다라는게 덕 타이핑(Duck Typing)이다.
	// '객체를 미리 판단하지 않고 변수와 메서드가 사용되는 때에 객체를 판단하겠다'라는 의미이다.
	/*
		type 인터페이스명 interface {
			메서드 집합
		}
	*/
	duck := Bird{"오리"}
	fmt.Printf("%s는 %s하고 울며, %s하고 걸으며 수영을 할 수 %s. \n", duck.name, duck.DuckSound(), duck.DuckWalk(), duck.isSwim())

	typeCheck(int8(1))
	typeCheck(int16(2))
	typeCheck(int32(3))
	typeCheck(int64(4))
	typeCheck(uint8(5))
	typeCheck(uint16(6))
	typeCheck(uint32(7))
	typeCheck(uint64(8))

	// geth 코드 - 인터페이스 사용 예시
	// go-ethereum\consensus\consensus.go, line: 33
	// go-ethereum\core\headerchain.go, line: 404

	// 14. 패키지
	/*
		기본 제공 패키지
			https://pkg.go.dev/std

		패키지 목록
			go env | grep GOROOT
			C:\Program Files\Go\src

		go mod: 모듈 관리 도구
			프로젝트의 의존성(외부 라이브러리나 패키지)을 쉽게 관리할 수 있다.

		외부 패키지 생성
			go mod init github.com/{깃허브 아이디}/go-lang-myfirstpkg

		원격 저장소 연동
			git remote add origin git@github.com:tlswlsdntks/go-lang-myfirstpkg.git
			git push -u origin main

		외부 패키지 다운로드
			go mod init go-lang-practice
			go get github.com/{깃허브 아이디}/go-lang-myfirstpkg

		외부 모듈 가져오기
			import "github.com/{깃허브 아이디}/go-lang-myfirstpkg"
	*/
	myfirstpkg.MyFirstFunc()

	// 15. 함수의 기능
	// defer: stack 과 같음, 후입선출(LIFO)
	defer fmt.Println("마지막으로 실행")
	defer fmt.Println("세 번째로 실행")
	defer fmt.Println("두 번째로 실행")
	fmt.Println("첫 번째로 실행")

	// 지연된 함수의 인수의 평가
	deferA()

	// 지연된 함수의 호출
	deferB()

	// 지연된 함수는 반환값 할당
	deferC()

	// 가변 인자
	numbers(1, 2, 3, 4, 5)
	strings("a", "b", "c")

	// 함수 리터럴
	/*
		익명 함수
		func () {
			코드
		}

		익명 함수를 변수에 대입
		f := func () {
			코드
		}
		f(): 함수 호출

		직접 호출
		func () {
			코드
		}()
	*/
	func() {
		fmt.Println("Hello World")
	}()

	// 캡쳐: 익명 함수는 외부 변수를 참조값으로 가져옴
	i = 0
	func() {
		i++
	}()
	fmt.Println(i)

	// 혹은 포인터를 이용해야함
	addTwo(&i)
	fmt.Println(i)

	// geth 코드 - 함수의 기능 사용 예시
	// go-ethereum\core\state\statedb.go, line: 551

	// 16. 에러 처리
	val := devide(10, 0) // panic: runtime error: integer divide by zero, src\runtime\panic.go, line: 236, 239
	fmt.Println(val)
	val, err := fmt.Println(devide2(10, 0))
	fmt.Println(val, err)

	// 에러 생성
	fmt.Println(errors.New("error message"))
	fmt.Println(fmt.Errorf("error message"))

	// 패닉 발생: 시스템 중단
	val, err = devide2(10, 0)
	if err != nil {
		panic("Error: Something went wrong")
	} else {
		fmt.Println(val)
	}

	// defer 함수를 사용하여 패닉을 반환하고, 복구
	defer func() {
		r := recover()
		fmt.Println(r) // Error: Something went wrong
		fmt.Println("Recovered from panic")
	}()
	panic("Error: Something went wrong")

	// geth 코드 - 에러 처리 사용 예시
	// go-ethereum\cmd\geth\consolecmd.go, line: 70

}
