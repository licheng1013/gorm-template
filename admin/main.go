package main

import (
	"admin/api"
	"common/app"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	data, _ := os.ReadFile("app.yml")
	app.ParseAppConfig(data)
	app.MysqlInit()
	api.Init()
	log.Println("http://localhost:" + fmt.Sprint(app.Config.Port))
	//下面都是优雅停机方式
	srv := &http.Server{
		Addr:    "0.0.0.0:" + fmt.Sprint(app.Config.Port),
		Handler: app.R,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	// 使用 kill -2 pid 或 kill pid 命令进行优雅停机
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
