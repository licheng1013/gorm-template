package tool

import (
	"log"
	"os"
)

var MyLog *log.Logger

func init() {
	MyLog = log.New(os.Stdout, "系统日志: ", log.Lshortfile+log.LstdFlags)
}
