package runtime

const (
	_EINTR     = -0x7ffffff6
	_EFAULT    = -0x7fffecff
	_EAGAIN    = -0x7ffffff5
	_ENOMEM    = -0x80000000
	_ETIMEDOUT = -0x7ffffff7
	_EACCES	   = -0x7ffffffe

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x8
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x4

/*	_MADV_DONTNEED   = 0x4
	_MADV_HUGEPAGE   = 0xe
	_MADV_NOHUGEPAGE = 0xf
*/
	_SA_RESTART  = 0x10
	_SA_ONSTACK  = 0x20
	//_SA_RESTORER = 0x4000000
	_SA_SIGINFO  = 0x40

	_SIGHUP    = 0x1
	_SIGINT    = 0x2
	_SIGQUIT   = 0x3
	_SIGILL    = 0x4
	_SIGTRAP   = 0x16
	_SIGABRT   = 0x6
	_SIGBUS    = 0x1e
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	//_SIGUSR1   = 0xa
	_SIGSEGV   = 0xb
	//_SIGUSR2   = 0xc
	_SIGPIPE   = 0x7
	_SIGALRM   = 0xe
	//_SIGSTKFLT = 0x10
	_SIGCHLD   = 0x5
	_SIGCONT   = 0xc
	_SIGSTOP   = 0xa
	_SIGTSTP   = 0xd
	_SIGTTIN   = 0x10
	_SIGTTOU   = 0x11
	_SIGURG    = 0x1a
	_SIGXCPU   = 0x1c
	_SIGXFSZ   = 0x1d
	_SIGVTALRM = 0x1b
	_SIGPROF   = 0x18
	_SIGWINCH  = 0x14
	//_SIGIO     = 0x1d
	//_SIGPWR    = 0x1e
	_SIGSYS    = 0x19
	_SIGUSR1   = 0x12
	_SIGUSR2   = 0x13

	_FPE_INTDIV = 0x14
	_FPE_INTOVF = 0x15
	_FPE_FLTDIV = 0x16
	_FPE_FLTOVF = 0x17
	_FPE_FLTUND = 0x18
	_FPE_FLTRES = 0x19
	_FPE_FLTINV = 0x1a
	_FPE_FLTSUB = 0x1b

	_BUS_ADRALN = 0x28
	_BUS_ADRERR = 0x29
	_BUS_OBJERR = 0x2a

	_SEGV_MAPERR = 0x1e
	_SEGV_ACCERR = 0x1f

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_PTHREAD_CREATE_DETACHED = 0x1

	_HOST_NAME_MAX = 0xff
	_O_NONBLOCK = 0x80
	_FD_CLOEXEC = 0x1
	_F_GETFL = 0x8
	_F_SETFL = 0x10
	_F_SETFD = 0x4
	_SC_NPROCESSORS_ONLN = 0x23
	_B_PAGE_SIZE = 0x1000
	_B_ERROR = -0x1

/*	_EPOLLIN       = 0x1
	_EPOLLOUT      = 0x4
	_EPOLLERR      = 0x8
	_EPOLLHUP      = 0x10
	_EPOLLRDHUP    = 0x2000
	_EPOLLET       = 0x80000000
	_EPOLL_CLOEXEC = 0x80000
	_EPOLL_CTL_ADD = 0x1
	_EPOLL_CTL_DEL = 0x2
	_EPOLL_CTL_MOD = 0x3

	_AF_UNIX    = 0x1
	_F_SETFL    = 0x4
	_SOCK_DGRAM = 0x2
*/
)

type Sigset uint64

type fpxreg struct {
	significand [4]uint16
	exponent    uint16
	padding     [3]uint16
}

type sigactiont struct {
	sa_handler  uintptr
	sa_flags    uint64
	sa_restorer uintptr
	sa_mask     uint64
}

type siginfo struct {
	si_signo int32
	si_errno int32
	si_code  int32
	si_status int32
	si_addr uintptr
}

type xmmreg struct {
	element [4]uint32
}

type fpstate struct {
	cwd       uint16
	swd       uint16
	ftw       uint16
	fop       uint16
	rip       uint64
	rdp       uint64
	mxcsr     uint32
	mxcr_mask uint32
	_st       [8]fpxreg
	_xmm      [16]xmmreg
	padding   [24]uint32
}

type stackt struct {
	ss_sp     *byte
	ss_flags  int32
	pad_cgo_0 [4]byte
	ss_size   uint64
}

type semt struct {
	id int32
	_padding [3]int32
}

type sigcontext struct {
	r8          uint64
	r9          uint64
	r10         uint64
	r11         uint64
	r12         uint64
	r13         uint64
	r14         uint64
	r15         uint64
	rdi         uint64
	rsi         uint64
	rbp         uint64
	rbx         uint64
	rdx         uint64
	rax         uint64
	rcx         uint64
	rsp         uint64
	rip         uint64
	eflags      uint64
	cs          uint16
	gs          uint16
	fs          uint16
	err         uint64
	trapno      uint64
	oldmask     uint64
	cr2         uint64
	fpstate     *fpstate
	__reserved1 [8]uint64
}