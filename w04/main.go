package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	x, err := time.ParseDuration("+1s")
	fmt.Println(int(x)/int(time.Second))
	fmt.Println(x, err)
	fmt.Println(time.Now())
	a := time.FixedZone("UTC", 1)
	fmt.Println(time.Now().In(a))

	b, _ := time.LoadLocation("Asia/Bangkok")
	fmt.Println(time.Now().In(b))
	return
	g, _ := errgroup.WithContext(context.Background())

	// 启动第一个服务
	g.Go(func() error {
		time.Sleep(3* time.Second)
		fmt.Println("1 goroutine", time.Now())
		return nil
	})

	// 启动第二个服务
	g.Go(func() error {
		time.Sleep(1* time.Second)
		fmt.Println("2 goroutine", time.Now(), "timeout")
		return errors.New("timeout")
	})

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
