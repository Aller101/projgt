package singleton

import "sync"

var (
	privateSingleton *singleton
	mtx              sync.Mutex
)

type singleton struct {
	name string
}

func (s *singleton) GetName() string {
	return s.name
}

func NewSingleton(name string) *singleton {

	mtx.Lock()
	defer mtx.Unlock()

	if privateSingleton != nil {
		return privateSingleton
	}

	privateSingleton = &singleton{name: name}

	return privateSingleton
}
