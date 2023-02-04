package singleton

import "sync"

var (
	instance *db
)

var lock = &sync.Mutex{}

type db struct {
}

func GetDbInstance() *db {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &db{}
	}
	return instance
}
