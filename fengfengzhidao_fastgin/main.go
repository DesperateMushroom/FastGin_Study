package main

import (
	"fengfengzhidao_fastgin/core"
	"fmt"
)

func main() {
	cfg := core.ReadConfig()
	fmt.Println(cfg.DB)
}
