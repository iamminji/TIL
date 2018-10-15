package main

import (
	"fmt"
	"reflect"
)

func add3(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func makeAdd3(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), add3)

	// fn에 함수 정보 v를 설정하여 함수를 연결
	fn.Set(v)
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeAdd3(&intSum)
	makeAdd3(&floatSum)
	makeAdd3(&stringSum)

	fmt.Println(intSum(1, 2))
	fmt.Println(floatSum(2.1, 3.5))
	fmt.Println(stringSum("Hello, ", "World!"))

}
