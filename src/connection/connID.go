package connection

type ConnID struct {
	id int32
}
func (this *ConnID)Set(id int32)  {
	this.id=id
}
func (this *ConnID)Get() int32 {
	return this.id
}
