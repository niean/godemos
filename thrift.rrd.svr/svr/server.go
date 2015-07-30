package svr

import (
	"../g"
	"../rrd"
	"fmt"
	"github.com/niean/thrift/lib/go/thrift"
	"log"
	"time"
)

func Start() {
	cfg := g.Config()

	if !cfg.Rrd.Enable {
		log.Println("rrd.Start warning, not enable")
		return
	}

	// get rrd svr config
	protocol := cfg.Rrd.Protocol
	buffered := cfg.Rrd.Buffered
	framed := cfg.Rrd.Framed
	addr := cfg.Rrd.Listen
	cliTimeout := cfg.Rrd.CallTimeout

	// protocal factory
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		log.Fatalln("rrd.Start error, invalid protocol", protocol)
	}

	// transport factory
	var transportFactory thrift.TTransportFactory
	if buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(10240) // 10K bytes
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	// run server
	go func() {
		if err := RunServer(transportFactory, protocolFactory, addr, cliTimeout); err != nil {
			fmt.Println("running server error,", err)
		}
	}()
}

// 之所以把RunServer独立为函数,是因为这个函数之前的、建立thrift通路的 代码, 是可以和RunCilent复用的
func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, cliTimeout int64) error {
	// transport
	var transport thrift.TServerTransport
	var err error
	if cliTimeout <= 0 {
		transport, err = thrift.NewTServerSocket(addr)
	} else {
		transport, err = thrift.NewTServerSocketTimeout(addr, time.Millisecond*time.Duration(cliTimeout))
	}
	if err != nil {
		return err
	}

	// server
	handler := NewRRDHBaseBackendHandler()
	processor := rrd.NewRRDHBaseBackendProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	// start server
	log.Println("rrd.RunServer ok, listening on", addr)
	return server.Serve()
}
