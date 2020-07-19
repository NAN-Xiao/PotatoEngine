package main

import "reflect"

type Ispace interface {
	CC()
}

type BaseSpace struct {
	Name string
}
func(this *BaseSpace)DD(){
 println("dd")
}

type MyClass struct {
	BaseSpace
	Ispace
}

func main()  {
	m:=new(MyClass)
   TestIN(m)
}

func TestIN(in Ispace){
	t:=reflect.ValueOf(in)

	fn:=t.MethodByName("DD")
	fn.Call(nil)

}