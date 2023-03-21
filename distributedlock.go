package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientv3.New err:", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()
	se1, _ := concurrency.NewSession(cli)
	se2, _ := concurrency.NewSession(cli)
	defer se2.Close()
	defer se1.Close()
	mu1 := concurrency.NewMutex(se1, "/service/1")
	if err := mu1.Lock(context.TODO()); err != nil {
		fmt.Println("mu1 lock err:", err)
	} else {
		fmt.Println("mu1 lock")
	}

	mu2 := concurrency.NewMutex(se2, "/service/1")
	m2lock := make(chan struct{})
	go func() {
		defer close(m2lock)
		if err := mu2.Lock(context.TODO()); err != nil {
			fmt.Println("mu2 lock err:", err)
		}
		// session2 获得锁
		fmt.Println("mu2 lock")
	}()
	if err := mu1.Unlock(context.TODO()); err != nil {
		fmt.Println("mu1 unlock err:", err)
	} else {
		fmt.Println("mu1 unlock")
	}
	<-m2lock
}
