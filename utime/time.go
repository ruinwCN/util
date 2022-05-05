package utime

import (
	"time"
)

func Timestamp() int64 {
	//return time.Now().UnixNano() / int64(time.Millisecond)
	return time.Now().UnixMilli()
}
