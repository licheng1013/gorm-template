package main

import (
	"context"
	_ "gin-web-template/api" //通过包初始化注册
	"gin-web-template/common"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	//设置日志打印
	log.SetPrefix("系统日志: ")
	log.SetFlags(log.LstdFlags + log.Lshortfile)
}

func main() {
	port := "8088"
	log.Println("http://localhost:" + port)
	//下面都是优雅停机方式
	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: common.R,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务器...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器关闭：", err)
	}
	log.Println("服务器退出")
}
