package main

import (
	"context"

	pb "github.com/fishioon/asynctasks/api"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *AsyncTasksServer) ListQueues(ctx context.Context, in *pb.ListQueuesRequest) (*pb.ListQueuesResponse, error) {
	return nil, nil
}
func (s *AsyncTasksServer) GetQueue(ctx context.Context, in *pb.GetQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *AsyncTasksServer) CreateQueue(ctx context.Context, in *pb.CreateQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *AsyncTasksServer) UpdateQueue(ctx context.Context, in *pb.UpdateQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *AsyncTasksServer) DeleteQueue(ctx context.Context, in *pb.DeleteQueueRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *AsyncTasksServer) PurgeQueue(ctx context.Context, in *pb.PurgeQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *AsyncTasksServer) PauseQueue(ctx context.Context, in *pb.PauseQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *AsyncTasksServer) ResumeQueue(ctx context.Context, in *pb.ResumeQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *AsyncTasksServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	return nil, nil
}
func (s *AsyncTasksServer) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.Task, error) {
	return nil, nil
}
func (s *AsyncTasksServer) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.Task, error) {
	return nil, nil
}
func (s *AsyncTasksServer) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *AsyncTasksServer) RunTask(ctx context.Context, in *pb.RunTaskRequest) (*pb.Task, error) {
	return nil, nil
}
