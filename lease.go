package main

import (
	"fmt"
	"time"
)

// etcd lease

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()
	cli.Put(context.TODO(), "/nazha1/2", "dsb")
	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
	}

	// 5秒钟之后, /nazha/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/nazha/1", "dsb", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}
