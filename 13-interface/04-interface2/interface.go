package main

//一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其它类型或结构体来实现
type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (g *Logger) Log(l string) {

}

type GameService struct {
	Logger
}

func (g GameService) Start() {
}

func main() {
	var s Service = new(GameService)

	s.Start()
	s.Log("hello")
}
