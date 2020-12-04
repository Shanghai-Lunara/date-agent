package date_agent

import (
	"context"
	"fmt"
	pb "github.com/Shanghai-Lunara/date-agent/proto"
	"github.com/golang/protobuf/ptypes"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"k8s.io/klog/v2"
	"time"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

type Client struct {
	hostname     string
	registerAddr string

	currentTaskId int32

	client pb.DateAgentClient
}

// NewClient returns the pointer of the Client structure
func NewClient(addr string) *Client {
	hostname, err := GetHostName()
	if err != nil {
		klog.Fatal(err)
	}

	hostname = fmt.Sprintf("%s%d", hostname, time.Now().Unix())

	c := &Client{
		hostname:      hostname,
		registerAddr:  addr,
		currentTaskId: 0,
	}
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)),
		retry.WithCodes(codes.NotFound, codes.Aborted),
	}
	cc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		klog.Fatal(err)
	}
	c.client = pb.NewDateAgentClient(cc)
	c.register()
	go c.timer()
	return c
}

func (c *Client) Close() {

}

func (c *Client) timer() {
	tick := time.NewTicker(time.Second * 1)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			klog.Info("task")
			_ = c.task()
		}
	}
}

func (c *Client) register() {
	timestamp, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		klog.Fatal(err)
	}
	if _, err := c.client.Register(
		context.Background(),
		&pb.RegisterRequest{Hostname: c.hostname, Time: timestamp},
		retry.WithMax(3),
		retry.WithPerRetryTimeout(1*time.Second),
	); err != nil {
		klog.V(2).Info(err)
	}
}

func (c *Client) task() (err error) {
	var reply *pb.PullTaskReply
	fmt.Println(c.hostname)

	if reply, err = c.client.PullTask(
		context.Background(),
		&pb.PullTaskRequest{Hostname: c.hostname},
		retry.WithMax(3),
		retry.WithPerRetryTimeout(1*time.Second),
	); err != nil {
		klog.V(2).Info(err)
		return err
	}
	klog.V(3).Info("reply:", reply)
	if reply.Task.TaskId == 0 || c.currentTaskId >= reply.Task.TaskId {
		return nil
	}
	c.currentTaskId = reply.Task.TaskId
	if len(reply.Task.Command) == 0 {
		return nil
	}
	var out string
	if out, err = Exec(reply.Task.Command); err != nil {
		klog.V(2).Info(err)
		return err
	}
	if _, err = c.client.CompleteTask(
		context.Background(),
		&pb.CompleteTaskRequest{Hostname: c.hostname, TaskId: reply.Task.TaskId, OutPut: out},
		retry.WithMax(3),
		retry.WithPerRetryTimeout(1*time.Second),
	); err != nil {
		klog.V(2).Info(err)
		return err
	}
	return nil
}
