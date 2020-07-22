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

////
type A interface {
	Get()
}

type B struct {
	ID int32
}

func (this *B)Get()  {
	
}

func TTest(a A)  {
	b,ok:=a.(*B)
	if ok{
		print(b.ID)
	}

}





