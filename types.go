package date_agent

import "time"

type NodeStatus int

const (
	NodeOnline  NodeStatus = iota
	NodeOffline NodeStatus = 1
	NodeRemoved NodeStatus = 2
)

type Node struct {
	Hostname string     `json:"hostname"`
	Status   NodeStatus `json:"status"`
	Time     time.Time  `json:"time"`
}
