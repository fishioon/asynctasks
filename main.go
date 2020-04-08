package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	pb "github.com/fishioon/asynctasks/api"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

var (
	Version, Build string
	conf           *Config
)

type Config struct {
	Host  string        `json:"host"`
	Redis redis.Options `json:"redis"`
}

func loadConfig(fname string) {
	c := &Config{}
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("open config file fail %s", err.Error())
	}
	if err = json.NewDecoder(f).Decode(c); err != nil {
		log.Fatalf("parse config file fail %s", err.Error())
	}
	conf = c
}

func main() {
	version := flag.Bool("version", false, "build version")
	confile := flag.String("config", "config.json", "config file")
	flag.Parse()
	if *version {
		fmt.Printf("Version: %s Build: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n", Version, Build, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		return
	}
	loadConfig(*confile)
	client := redis.NewClient(&conf.Redis)
	if err := client.Ping().Err(); err != nil {
		log.Fatalf("redis ping %s", err.Error())
	}
	lis, err := net.Listen("tcp", conf.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAsyncTasksServer(s, newServer(client))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
