package main

import (
	"context"

	pb "github.com/fishioon/asynctasks/api"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *TasksServer) ListQueues(ctx context.Context, in *pb.ListQueuesRequest) (*pb.ListQueuesResponse, error) {
	return nil, nil
}
func (s *TasksServer) GetQueue(ctx context.Context, in *pb.GetQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *TasksServer) CreateQueue(ctx context.Context, in *pb.CreateQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *TasksServer) UpdateQueue(ctx context.Context, in *pb.UpdateQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *TasksServer) DeleteQueue(ctx context.Context, in *pb.DeleteQueueRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *TasksServer) PurgeQueue(ctx context.Context, in *pb.PurgeQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *TasksServer) PauseQueue(ctx context.Context, in *pb.PauseQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *TasksServer) ResumeQueue(ctx context.Context, in *pb.ResumeQueueRequest) (*pb.Queue, error) {
	return nil, nil
}
func (s *TasksServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	return nil, nil
}
func (s *TasksServer) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.Task, error) {
	return nil, nil
}
func (s *TasksServer) CreateTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.Task, error) {
	return nil, nil
}
func (s *TasksServer) DeleteTask(ctx context.Context, in *pb.DeleteTaskRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *TasksServer) RunTask(ctx context.Context, in *pb.RunTaskRequest) (*pb.Task, error) {
	return nil, nil
}
