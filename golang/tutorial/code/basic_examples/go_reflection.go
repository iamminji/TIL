package basic_examples

import (
	"fmt"
	"reflect"
)

type ReflectionStruct struct {

	/* tag 붙여 써야됨! */
	name string `tag1:"1"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func reflectionExample() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(t.Name())
	fmt.Println(t.Size())

	fmt.Println(t.Kind() == reflect.Float64)
	fmt.Println(t.Kind() == reflect.Int64)

	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Kind() == reflect.Int64)
	fmt.Println(v.Float())

}

func reflectionExample02() {
	r := ReflectionStruct{}
	name, ok := reflect.TypeOf(r).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(r).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))
}

func main() {
	reflectionExample()
	reflectionExample02()
}
