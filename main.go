package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "uu-gfast/internal/app/boot"

	"github.com/gogf/gf/v2/os/gctx"
	_ "uu-gfast/internal/app/system/packed"
	"uu-gfast/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
