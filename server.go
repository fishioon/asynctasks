package main

import (
	"github.com/go-redis/redis"
	pb "github.com/fishioon/asynctasks/api"
)

type TasksServer struct {
	rds     *redis.Client
	workers map[string]*Worker
}

func newServer(rds *redis.Client) *TasksServer {
	return &TasksServer{
		rds: rds,
	}
}

func (ts *TasksServer) Start() error {
	return nil
}

func (ts *TasksServer) watchQueue() error {
	return nil
}
