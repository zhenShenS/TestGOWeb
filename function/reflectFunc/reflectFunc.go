package reflectFunc

import (
	"fmt"
	"reflect"
)

func TestReflectFunc() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
