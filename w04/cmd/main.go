package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/wuzehv/geekgo/w04/api"
	"github.com/wuzehv/geekgo/w04/internal/service"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// App app对象依赖db对象
type App struct {
	db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{db: db}
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// wire依赖管理
	app, err := InitializeApp()
	fmt.Println(app, err)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g, ctx := errgroup.WithContext(context.Background())

	s := grpc.NewServer()

	// 启动grpc服务
	g.Go(func() error {
		pb.RegisterUserServer(s, &service.Server{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			return err
		}
		return nil
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
				done <- struct{}{}
			case <-ctx.Done():
				done <- struct{}{}
			}
		}()

		<-done
		s.GracefulStop()
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
