package main
import (
	"io"
	"os"
	"fmt"
	"reflect"
)

func main() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File

	fmt.Println(reflect.ValueOf(3))
	fmt.Println(reflect.ValueOf(3).Type())

	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Println("%d\n", i)
}
