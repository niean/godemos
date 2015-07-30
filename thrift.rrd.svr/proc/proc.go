package proc

import (
	nproc "github.com/toolkits/proc"
	"log"
)

var (
	PingCnt = nproc.NewSCounterQps("PingCnt")

	SendCnt      = nproc.NewSCounterQps("SendCnt")
	SendItemsCnt = nproc.NewSCounterQps("SendItemsCnt")

	QueryCnt      = nproc.NewSCounterQps("QueryCnt")
	QueryItemsCnt = nproc.NewSCounterQps("QueryItemsCnt")

	LastCnt      = nproc.NewSCounterQps("LastCnt")
	LastItemsCnt = nproc.NewSCounterQps("LastItemsCnt")
)

func Start() {
	log.Println("proc.Start ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	ret = append(ret, PingCnt.Get())

	ret = append(ret, SendCnt.Get())
	ret = append(ret, SendItemsCnt.Get())

	ret = append(ret, QueryCnt.Get())
	ret = append(ret, QueryItemsCnt.Get())

	ret = append(ret, LastCnt.Get())
	ret = append(ret, LastItemsCnt.Get())

	return ret
}
