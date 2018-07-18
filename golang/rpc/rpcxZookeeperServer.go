// go run -tags zookeeper rpcxZookeeperServer.go
package main

import (
	"context"
	"flag"
	"fmt"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"github.com/smallnest/rpcx/share"
	"log"
	"time"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Println(reply.C, " = ", args.A, " * ", args.B)
	return nil
}

func (t *Arith) PrintMetaData(ctx context.Context, args string, reply *string) error {
	*reply = "success"
	fmt.Println("req.Meta: ", ctx.Value(share.ReqMetaDataKey).(map[string]string)["reqMeta"])
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)
	resMeta["resMeta"] = "FromServer"
	time.Sleep(3 * time.Second)
	return nil
}

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
