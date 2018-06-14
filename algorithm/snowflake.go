package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const MAX_CURRENT_ID = 4095
const TIME_UNIT = 1 * time.Millisecond // 时间戳最小单位

type SnowFlake struct {
	Sm             sync.Mutex
	StartTimestamp int64 // 41位
	CurrentId      int64 // 12位  0 ~ 4095
	DatabaseId     string
	WorkerId       string
	LastTimestamp  string
}

func FixedLengthBinaryString(num int64, length int) string {
	result := strconv.FormatInt(num, 2)
	for len(result) < length {
		result = "0" + result
	}
	return result
}

func (snowFlake *SnowFlake) CurrentTimestamp() string {
	return FixedLengthBinaryString(time.Now().UnixNano()/int64(TIME_UNIT/time.Nanosecond)-snowFlake.StartTimestamp, 41)
}

func (snowFlake *SnowFlake) GenerateId() int64 {
	IdStr := "0"
	snowFlake.Sm.Lock()
	currentTimestamp := snowFlake.CurrentTimestamp()
	if snowFlake.LastTimestamp != currentTimestamp {
		snowFlake.LastTimestamp = currentTimestamp
		snowFlake.CurrentId = 0
	}
	if snowFlake.CurrentId > MAX_CURRENT_ID {
		snowFlake.Sm.Unlock()
		time.Sleep(TIME_UNIT)
		return snowFlake.GenerateId()
	}
	IdStr += fmt.Sprintf("%s%s%s%s", currentTimestamp, snowFlake.DatabaseId, snowFlake.WorkerId, FixedLengthBinaryString(snowFlake.CurrentId, 12))
	snowFlake.CurrentId++
	snowFlake.Sm.Unlock()
	Id, _ := strconv.ParseInt(IdStr, 2, 64)
	return Id
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	snowFlake := new(SnowFlake)
	snowFlake.StartTimestamp = time.Now().UnixNano() / int64(TIME_UNIT/time.Nanosecond)
	snowFlake.DatabaseId = "00000"
	snowFlake.WorkerId = "00000"
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			fmt.Printf("%d ", snowFlake.GenerateId())
			wg.Done()
		}()
	}
	wg.Wait()
}
