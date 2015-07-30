package svr

import (
	"../g"
	"../proc"
	"../rrd"
	"log"
	"sync"
)

type RRDHBaseBackendHandler struct {
	sync.RWMutex
	RecvCnt int64
}

func NewRRDHBaseBackendHandler() *RRDHBaseBackendHandler {
	return &RRDHBaseBackendHandler{RecvCnt: 0}
}

func (this *RRDHBaseBackendHandler) Ping() (err error) {
	// statistics
	proc.PingCnt.Incr()
	return nil
}

func (this *RRDHBaseBackendHandler) Send(items []*rrd.GraphItem) (r string, err error) {
	// statistics
	proc.SendCnt.Incr()
	proc.SendItemsCnt.IncrBy(int64(len(items)))

	if g.Config().Debug {
		log.Printf("recv: %v", items)
	}

	this.Lock()
	this.RecvCnt += 1
	this.Unlock()

	return "OK", nil
}

func (this *RRDHBaseBackendHandler) Query(requests []*rrd.QueryRequest) (r []*rrd.QueryResponse, err error) {
	// statistics
	proc.QueryCnt.Incr()
    log.Println(requests)

	ret := make([]*rrd.QueryResponse, 0)
	return ret, nil
}

func (this *RRDHBaseBackendHandler) Last(requests []*rrd.LastRequest) (r []*rrd.LastResponse, err error) {
	// statistics
	proc.LastCnt.Incr()
    log.Println(requests)

	ret := make([]*rrd.LastResponse, 0)
	return ret, nil
}
