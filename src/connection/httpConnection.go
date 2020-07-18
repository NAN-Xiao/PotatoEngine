package connection

type HttpConnect struct {
	connID ConnID
}

func (this *HttpConnect) Read() (l int, err error) {
	return 0, nil
}
func (this *HttpConnect) Write(data []byte) {

}
func (this *HttpConnect) Close() bool {
	return false
}
func (this *HttpConnect) Listen() {

}
