// +build ignore

package main

import "syscall"

func lockReg(fd, cmd, typ, off, whence, ln int) error {
	var lock syscall.Flock_t
	lock.Type = int16(typ)
	lock.Start = int64(off)
	lock.Whence = int16(whence)
	lock.Len = int64(ln)

	return syscall.FcntlFlock(uintptr(fd), cmd, &lock)
}

func ReadLock(fd, off, whence, ln int) error {
	return lockReg(fd, syscall.F_SETLK, syscall.F_RDLCK, off, whence, ln)
}

func ReadWLock(fd, off, whence, ln int) error {
	return lockReg(fd, syscall.F_SETLKW, syscall.F_RDLCK, off, whence, ln)
}

func WriteLock(fd, off, whence, ln int) error {
	return lockReg(fd, syscall.F_SETLK, syscall.F_WRLCK, off, whence, ln)
}

func WriteWLock(fd, off, whence, ln int) error {
	return lockReg(fd, syscall.F_SETLKW, syscall.F_WRLCK, off, whence, ln)
}

func Unlock(fd, off, whence, ln int) error {
	return lockReg(fd, syscall.F_SETLK, syscall.F_UNLCK, off, whence, ln)
}

func lockTest(fd, typ, off, whence, ln int) (int32, error) {
	var lock syscall.Flock_t
	lock.Type = int16(typ)
	lock.Start = int64(off)
	lock.Whence = int16(whence)
	lock.Len = int64(ln)

	if err := syscall.FcntlFlock(uintptr(fd), syscall.F_GETLK, &lock); err != nil {
		return 0, err // 检查失败, 返回错误
	}
	if lock.Type == syscall.F_UNLCK {
		return 0, nil // 当前未上锁
	}
	return lock.Pid, syscall.EWOULDBLOCK // 当前已上锁
}

func IsReadLockable(fd, off, whence, ln int) bool {
	var pid, err = lockTest(fd, syscall.F_RDLCK, off, whence, ln)
	return pid == 0 && err == nil
}

func IsWriteLockable(fd, off, whence, ln int) bool {
	var pid, err = lockTest(fd, syscall.F_WRLCK, off, whence, ln)
	return pid == 0 && err == nil
}

func main() {

}
