package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type redisConfig struct {
	RedisUrl      []string `json:"redis_url"`
	RedisUername  string   `json:"redis_username"`
	redisPassword string   `json:"redis_password"`
}

type etcdConfig struct {
	EtcdEndpoints []string `json:"etcd_endpoints"`
	EtcdUsername  string   `json:"etcd_username"`
	EtcdPassword  string   `json:"etcd_password"`
}

type Config struct {
	EtcdOptions  etcdConfig  `json:"etcd_options"`
	RedisOptions redisConfig `json:"redis_options"`
}

func main() {
	var jsonBlob = []byte(`{"etcd_options": {"etcd_endpoints": ["http://localhost:2379", "http://103.37.149.30:2379"],"etcd_username": "Shouren","etcd_password": "test"}}`)
	data := Config{}
	err := json.Unmarshal(jsonBlob, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
	fmt.Printf("%v\n", len(data.RedisOptions.RedisUrl))
	if data.RedisOptions.RedisUername == "" {
		data.RedisOptions.RedisUername = "test"
		fmt.Printf("RedisUername is %v\n", data.RedisOptions.RedisUername)
	}
	var test string
	fmt.Printf("%t\n", test)
}
