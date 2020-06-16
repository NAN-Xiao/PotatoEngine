package main

func main() {
	LaunchServer()
	sp := NewLoginSpace("Login")
	RegistSpace("Login", sp)
	Serv()
	select {}
}
