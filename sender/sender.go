package sender

import (
	"math/rand"
	"time"
)

func GetMessage(i int) int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return i
}
