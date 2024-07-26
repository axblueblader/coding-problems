package main

import (
	"fmt"
)

// LoadBalancer represents a round-robin load balancer
type LoadBalancer struct {
	servers      []string
	currentIndex uint
}

// NewLoadBalancer creates a new load balancer instance
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers:      servers,
		currentIndex: 0,
	}
}

// Next returns the next server in the round-robin fashion
func (lb *LoadBalancer) Next() string {
	lb.currentIndex = (lb.currentIndex + 1) % uint(len(lb.servers))
	return lb.servers[lb.currentIndex]
}

// Distribute distributes requests across servers
func (lb *LoadBalancer) Distribute(requests []string) map[string][]string {
	result := make(map[string][]string)
	for _, request := range requests {
		server := lb.Next()
		result[server] = append(result[server], request)
	}
	return result
}

func main() {
	servers := []string{"server1", "server2", "server3"}
	requests := []string{"request1", "request2", "request3", "request4", "request5"}

	lb := NewLoadBalancer(servers)
	distributedRequests := lb.Distribute(requests)

	for server, reqs := range distributedRequests {
		fmt.Printf("Server %s: %v\n", server, reqs)
	}
}
