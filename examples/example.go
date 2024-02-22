package main

import (
	"fmt"
	"os"
	"zhenxin.me/kook"
)

func main() {
	token := os.Getenv("KOOK_BOT_TOKEN")
	if token == "" {
		fmt.Println("请检查环境变量 KOOK_BOT_TOKEN 是否设置")
		return
	}
	client := kook.NewClient(token)

	client.Start()
}
