package main

import (
	"aidb_go/internal/aidb"
	"aidb_go/internal/config"
	"aidb_go/internal/handler"
	"aidb_go/internal/svc"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/aidb.yaml", "the config file")

func main() {

	// export LD_LIBRARY_PATH=$PWD/internal/aidb/lib:$LD_LIBRARY_PATH

	aidb.AiDBCreate()
	stopC := make(chan os.Signal)
	signal.Notify(stopC, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stopC
		aidb.AiDBFree()
		os.Exit(1)
	}()

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
