package lock

import (
	"errors"
	"github.com/gogf/gf/v2/os/gmlock"
)

// Lock 加锁
func Lock(lockKey string, f func()) error {
	lock := gmlock.TryLock(lockKey)
	if !lock {
		return errors.New("系统繁忙，请重试")
	}
	defer func() {
		gmlock.Unlock(lockKey)
		gmlock.Remove(lockKey)
	}()
	f()
	return nil
}
