package util

import "sync"

// type Single interface {
// 	Single()
// }

type singleton struct {
	Id   int
	Name string
}

// func (s *singleton) Single() {

// }

var instance *singleton
var instanceOnce sync.Once

func GetSingleInstance() *singleton {
	instanceOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
