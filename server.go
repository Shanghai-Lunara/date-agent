package date_agent

import (
	"context"
	"net/http"

	pb "github.com/Shanghai-Lunara/date-agent/proto"
	"k8s.io/klog"
)

type Server struct {
	hub  *Hub
	http *http.Server
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if err := s.hub.Register(req.Hostname); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return &pb.RegisterReply{}, nil
}

func (s *Server) PullTask(ctx context.Context, req *pb.PullTaskRequest) (*pb.PullTaskReply, error) {
	task := s.hub.PullTask(req.Hostname)
	return &pb.PullTaskReply{Task: &pb.Task{TaskId: task.Id, Command: task.Command}}, nil
}

func (s *Server) CompleteTask(ctx context.Context, req *pb.CompleteTaskRequest) (*pb.CompleteTaskReply, error) {
	if err := s.hub.CompleteTask(req.Hostname, req.TaskId, req.OutPut); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return &pb.CompleteTaskReply{}, nil
}

func (s *Server) Close() {
}

func NewServer(grpcAddr string, httpAddr string) *Server {
	hub := NewHub(10)
	s := &Server{
		hub:  hub,
		http: InitHttp(httpAddr, hub),
	}
	return s
}