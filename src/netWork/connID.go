package netWork

type ConnID int32

func (this ConnID)Set(id int32)  {
	this= ConnID(id)
}
func (this *ConnID)Get() int32 {
	return  int32(*this)
}
