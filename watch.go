package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{""},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("clientv3.New err:", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()
	key := "/service"
	rch := cli.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
