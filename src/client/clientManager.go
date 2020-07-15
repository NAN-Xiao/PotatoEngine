package client

type ClientMap struct {
	Clients []Client
	ClientIndex int32
}
//添加一个客户端到全局
func (this *ClientMap)AddClient(cl Client)  {
	this.ClientIndex+=1
	cl.ClientID=this.ClientIndex
	this.Clients = append(this.Clients, cl)

}
//从全局删除客户端
func (this *ClientMap)RemoveClient(cl Client)  {
	if this.Clients==nil|| len(this.Clients)<=0{
		return
	}
	for i:=range this.Clients{
		if this.Clients[i]==cl{
			this.Clients=append(this.Clients[:i],this.Clients[i+1:]...)
		}
	}
}