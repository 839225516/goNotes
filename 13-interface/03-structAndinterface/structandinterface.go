package main

import "io"

type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

func usingCloser(closer io.Closer) {
	closer.Close()
}

func main() {
	// 实例化Socket
	s := new(Socket)

	// s实现了 io.Writer 和 io.Closer接口
	// ioW = s
	// usingWriter(ioW)
	usingWriter(s)

	// ioC = s
	// usingCloser(ioC)
	usingCloser(s)
}
