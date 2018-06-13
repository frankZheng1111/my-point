package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type SnowFlake struct {
	Sm             sync.Mutex
	StartTimestamp int64 // 41位
	CurrentId      int64 // 12位  0 ~ 4095
	DatabaseId     string
	WorkerId       string
	LastTimestamp  string
}

func FillInBinaryString(bstring string, length int) string {
	result := bstring
	for len(result) < length {
		result = "0" + result
	}
	return result
}

func (snowFlake *SnowFlake) CurrentTimestamp() string {
	return FillInBinaryString(fmt.Sprintf("%b", time.Now().UnixNano()/1000000-snowFlake.StartTimestamp), 41)
}

func (snowFlake *SnowFlake) GenerateId() interface{} {
	IdStr := "0"
	snowFlake.Sm.Lock()
	currentTimestamp := snowFlake.CurrentTimestamp()
	if snowFlake.LastTimestamp != currentTimestamp {
		snowFlake.LastTimestamp = currentTimestamp
		snowFlake.CurrentId = 0
	}
	IdStr += fmt.Sprintf("%s%s%s%s", currentTimestamp, snowFlake.DatabaseId, snowFlake.WorkerId, FillInBinaryString(strconv.FormatInt(snowFlake.CurrentId, 2), 12))
	snowFlake.CurrentId++
	snowFlake.Sm.Unlock()
	fmt.Println(IdStr)
	Id, _ := strconv.ParseInt(IdStr, 2, 64)
	return Id
}

func main() {
	snowFlake := new(SnowFlake)
	snowFlake.StartTimestamp = time.Now().UnixNano() / 1000000
	snowFlake.DatabaseId = "00000"
	snowFlake.WorkerId = "00000"
	for i := 0; i < 50; i++ {
		fmt.Println(snowFlake.GenerateId())
	}
}
