package singleton

// memastikan suatu object cuma bisa diintiate 1x saja
// untuk memastikan objek diinitiate 1x bisa pakai lock
// jadi fungsi akan ke lock sampai selesai baru bisa dieksekusi lagi
// atau pakai once. jadi fungsi instance = &db{} akan dijalankan cuma 1x

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
