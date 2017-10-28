package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func listener() (l net.Listener, err error) {
	var fd uintptr
	if _, err = fmt.Sscan(os.Getenv("TEST_FD"), &fd); nil != err {
		fmt.Printf("Sscan ret: %v\n", err)
		return
	}
	l, err = net.FileListener(os.NewFile(fd, os.Getenv("TEST_NAME")))
	if nil != err {
		return
	}
	switch l.(type) {
	case *net.TCPListener, *net.UnixListener:
	default:
		err = fmt.Errorf(
			"file descriptor is %T not *net.TCPListener or *net.UnixListener",
			l,
		)
		return
	}
	if err = syscall.Close(int(fd)); nil != err {
		return
	}
	return
}

func main() {
	l1, err1 := listener()
	l2, err2 := listener()
	fmt.Printf("l : %v\n", l1)
	fmt.Printf("err : %v\n", err1)
	fmt.Printf("ctrListener: %v\n", l2)
	fmt.Printf("goAgainErr: %v\n", err2)
}
