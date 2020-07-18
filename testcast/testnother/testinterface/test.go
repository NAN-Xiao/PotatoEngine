package testinterface

type Ispace interface {
	CC()

}

type BaseSpace struct {

	Name string
}
func(this *BaseSpace)DD(){

}

type MyClass struct {
	BaseSpace
	Ispace
}

func Test()  {

}