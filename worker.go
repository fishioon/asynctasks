package main

import (
	"context"

	"github.com/fishioon/asynctasks/api"
	"github.com/go-redis/redis"
)

type Worker struct {
	s     *TasksServer
	Queue string
}

func (w *Worker) start(ctx context.Context) {
}

func (w *Worker) fetchPendingTasks() ([]*api.Task, error) {
	return nil, nil
}

func (w *Worker) fetchTasks() (*api.Task, error) {
	res, err := w.s.rds.XReadGroup(&redis.XReadGroupArgs{}).Result()
	if err != nil || len(res) == 0 {
		return nil, nil
	}
	return nil, nil
}

func (w *Worker) execTask(task *api.Task) error {
	w.s.rds.XAck(w.Queue, w.Queue, task.GetName())
	return nil
}
