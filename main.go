package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	_ "github.com/zcyc/gf-grpc-template/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/zcyc/gf-grpc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
