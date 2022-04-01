package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func server(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("server", addr, "shutdown")
			s.Shutdown(ctx)
		}
	}()

	return s.ListenAndServe()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// 启动第一个服务
	g.Go(func() error {
		return server(ctx, ":9999", http.DefaultServeMux)
	})

	// 启动第二个服务
	g.Go(func() error {
		return server(ctx, ":9998", http.DefaultServeMux)
	})

	// 监听系统信号
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		done := make(chan struct{}, 1)

		// 注册信号
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			select {
			case sig := <-sigs:
				fmt.Println("接收到系统信号", sig)
				// 发现号退出所有服务
				done <- struct{}{}
			case <-ctx.Done():
				// http服务报错也不再继续监听信号
				done <- struct{}{}
			}
		}()

		<-done
		return errors.New("退出服务")
	})

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
