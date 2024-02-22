package main

import (
	"github.com/sirupsen/logrus"
	"zhenxin.me/kook/internal"
)

func main() {
	internal.InitLogger()
	logrus.Info("Hello, World!")
}
