package main

import (
	"context"
	"github.com/coreos/etcd/client"
	"log"
	"reflect"
	"strings"
	"time"
)

var ETCD_EP = "http://localhost:2379/"

func main() {
	// Setup Endpoint
	var ep []string
	ep = append(ep, ETCD_EP)

	// Setup Client
	cfg := client.Config{
		Endpoints:               ep,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
		Username:                "",
		Password:                "",
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Setup APIs
	kapi := client.NewKeysAPI(c)

	// Setup API's Options
	GetOpt := client.GetOptions{
		Recursive: true,
		Sort:      false,
		Quorum:    false,
	}

	// Get Data
	key := "namespace"
	resp, err := kapi.Get(context.Background(), key, &GetOpt)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Get Response from etcd: \nAction: %v\nKey: %v\nValue: %v\n",
			resp.Action,
			resp.Node.Key,
			resp.Node.Value,
		)
		for _, service := range resp.Node.Nodes {
			keys := strings.Split(service.Key, "/")
			log.Printf("Service: %v\n", keys[len(keys)-1])
			for _, upstream := range service.Nodes {
				keys := strings.Split(upstream.Key, "/")
				log.Printf("  upstreams: %v\n", keys[len(keys)-1])
			}
		}
	}
	GetOpt.Recursive = false
	resp, err = kapi.Get(context.Background(), "namespace/serv_2", &GetOpt)
	if err != nil {
		log.Printf("%+v\n", err)
		// errCode := strings.Split(err.Error(), ":")[0]
		log.Printf("%+v", reflect.ValueOf(err).Interface().(client.Error).Code)
	}

}
