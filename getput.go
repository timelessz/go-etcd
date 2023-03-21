package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	// etcd 中写入数据
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientV3.New err:", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()
	keyPrefix := "/service/"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, keyPrefix+"2", "127.0.0.2")
	cancel()
	if err != nil {
		fmt.Println("put to etcd failed, err:", err)
		return
	}
	// etcd 中读取数据
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, keyPrefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s", ev.Key, ev.Value)
	}
}
