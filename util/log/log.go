package log

import (
	"HydraClient/util/conf"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Info  *log.Logger
	CMD   *log.Logger
	Debug *log.Logger
	Report *log.Logger
)

func init() {
	prefix := strings.Join([]string{conf.GetVal("main", "function"), ":"}, "")
	clog, err := os.OpenFile("cmd.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	ilog, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	dlog, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	rlog,err := os.OpenFile("report.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	Info = log.New(io.MultiWriter(ilog, os.Stdout), prefix, log.Ltime|log.Lshortfile)
	Debug = log.New(io.MultiWriter(dlog, os.Stdout), prefix, log.Ltime|log.Llongfile)
	CMD = log.New(io.MultiWriter(clog, os.Stdout), prefix, log.Ltime|log.Llongfile)
	Report = log.New(io.MultiWriter(rlog, os.Stdout),"",64)
}
