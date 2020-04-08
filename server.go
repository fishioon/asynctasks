package main

import (
	"github.com/go-redis/redis"
)

type AsyncTasksServer struct {
	rds *redis.Client
}

func newServer(rds *redis.Client) *AsyncTasksServer {
	return &AsyncTasksServer{}
}

