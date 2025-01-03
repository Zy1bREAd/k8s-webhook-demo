package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

// 定义TLS证书相关变量
var (
	tlsCert, tlsKey, tlsPort string
)

func StartServer() {
	// 使用命令参数化传递TLS文件
	flag.StringVar(&tlsCert, "tlscert", "ocean.crt", "Path to the TLS certificate")
	flag.StringVar(&tlsKey, "tlskey", "ocean.key", "Path to the TLS key")
	flag.StringVar(&tlsPort, "port", "17443", "The port to listen")
	flag.Parse()
	fmt.Println(tlsCert, tlsKey, tlsPort)

	// goroutine 启动HTTP Server
	srv := &http.Server{Addr: "localhost:" + tlsPort}
	go func() {
		err := srv.ListenAndServeTLS(tlsCert, tlsKey)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("启动TLS HTTP服务器失败,", err)
		}
	}()

	// 设置优雅退出服务器(使用context方式)
	// signalChannel := make(chan os.Signal, 1)
	// signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	defer cancel()
	// 由于ctx.Done()返回一个只读channel，从该通道读取消息（阻塞）即是等待信号量的出现。
	<-ctx.Done()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println("关闭Server出现错误,", err)
		return
	}
	log.Println("优雅顺利关闭服务器！！！")
}