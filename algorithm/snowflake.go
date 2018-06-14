package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const TIME_UNIT = 1 * time.Millisecond // 时间戳最小单位
const CURRENT_ID_MAX_BIT = 12
const DATA_ID_MAX_BIT = 5
const WORKER_ID_MAX_BIT = 5
const MAX_CURRENT_ID = 1<<CURRENT_ID_MAX_BIT - 1

type SnowFlake struct {
	Sm             sync.Mutex
	StartTimestamp int64 // 41位
	CurrentId      int64 // 12位  0 ~ 4095
	DataCenterId   int64
	WorkerId       int64
	LastTimestamp  int64
}

func FixedLengthBinaryString(num int64, length int) string {
	result := strconv.FormatInt(num, 2)
	for len(result) < length {
		result = "0" + result
	}
	return result
}

func (snowFlake *SnowFlake) ParseToTimestamp(times time.Time) int64 {
	return times.UnixNano() / int64(TIME_UNIT/time.Nanosecond)
}

func (snowFlake *SnowFlake) GenerateId() int64 {
	snowFlake.Sm.Lock()
	currentTimestamp := snowFlake.ParseToTimestamp(time.Now()) - snowFlake.StartTimestamp
	if snowFlake.LastTimestamp != currentTimestamp {
		snowFlake.LastTimestamp = currentTimestamp
		snowFlake.CurrentId = 0
	}
	if snowFlake.CurrentId > MAX_CURRENT_ID {
		time.Sleep(TIME_UNIT)
		snowFlake.Sm.Unlock()
		return snowFlake.GenerateId()
	}
	Id := currentTimestamp<<(CURRENT_ID_MAX_BIT+WORKER_ID_MAX_BIT+DATA_ID_MAX_BIT) |
		snowFlake.DataCenterId<<(CURRENT_ID_MAX_BIT+WORKER_ID_MAX_BIT) |
		snowFlake.WorkerId<<CURRENT_ID_MAX_BIT |
		snowFlake.CurrentId
	snowFlake.CurrentId++
	snowFlake.Sm.Unlock()
	return Id
}

func main() {
	timeStr := "2018-06-14T9:09:26.371Z"
	times, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(times)
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	snowFlake := new(SnowFlake)
	snowFlake.StartTimestamp = snowFlake.ParseToTimestamp(times)
	snowFlake.DataCenterId = 0
	snowFlake.WorkerId = 0
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			fmt.Printf("%d ", snowFlake.GenerateId())
			wg.Done()
		}()
	}
	wg.Wait()
}
