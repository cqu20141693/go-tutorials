package gsync

import (
	"sync"
	"testing"
)

func TestRWMutex(t *testing.T) {
	mutex := &sync.RWMutex{}
	useLock(mutex)

	useRLock(mutex)

}

func useRLock(mutex *sync.RWMutex) {
	defer mutex.RUnlock()
	mutex.RLock()
	// Read 共享变量
}

func useLock(mutex *sync.RWMutex) {
	defer mutex.Unlock()
	// Update 共享变量
	mutex.Lock()
}
