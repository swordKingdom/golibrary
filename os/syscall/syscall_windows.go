package syscall

import (
	"syscall"
)

const (
	SIGHUP  = syscall.Signal(0x1)
	SIGINT  = syscall.Signal(0x2)
	SIGQUIT = syscall.Signal(0x3)
	SIGILL  = syscall.Signal(0x4)
	SIGTRAP = syscall.Signal(0x5)
	SIGABRT = syscall.Signal(0x6)
	SIGBUS  = syscall.Signal(0x7)
	SIGFPE  = syscall.Signal(0x8)
	SIGKILL = syscall.Signal(0x9)
	SIGSEGV = syscall.Signal(0xb)
	SIGPIPE = syscall.Signal(0xd)
	SIGALRM = syscall.Signal(0xe)
	SIGTERM = syscall.Signal(0xf)
)
