package main

// 使用new 或 & 构造的类型实例是类型的指针
// ins := new(T)  // T为类型，可以是结构体、整型、字符串等

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

// 在go中，对结构体进行 & 取地址操作时，视为对该类型进行一次new的实例化操作
type Command struct {
	Name    string
	Var     *int
	Comment string
}

// 取地址实例化是最广泛的一种结构体实例化方式。可以使用函数封装上面的初始化过程
func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	tank := new(Player)

	// new实例化的结构体在成员赋值上和var实例化的写法一致
	tank.Name = "Xiaoming"
	tank.HealthPoint = 200

	var version int = 1

	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"

	cmd = newCommand(
		"version",
		&version,
		"show version",
	)
}
