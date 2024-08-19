package lb

import (
	"context"
	"net/http"

	"github.com/rimo10/load_balancer/backend"
	"github.com/rimo10/load_balancer/serverpool"
)

const (
	Attempts int = iota
	Retry
)

type LoadBalancer interface {
	Serve(http.ResponseWriter, *http.Request)
	GetRetryFromContext(*http.Request) int
	ChooseBackend(context.Context) (*backend.Backend, error)
}

type loadbalancer struct {
	sp serverpool.Serverpool
}

func NewLoadBalancer(spl serverpool.Serverpool) LoadBalancer {
	return &loadbalancer{
		sp: spl,
	}

}

func (lb *loadbalancer) Serve(w http.ResponseWriter, r *http.Request) {
	// peer := lb.sp.GetNextPeerRoundRobin()
	peer := *lb.sp.GetNextPeerByLoad()
	if peer != nil {
		peer.Serve(w, r)
		return
	}
	http.Error(w, "service not available", http.StatusServiceUnavailable)
}

func (lb *loadbalancer) ChooseBackend(ctx context.Context) (*backend.Backend, error) {
	return lb.sp.GetNextPeerByLoad(), nil
}

func (lb *loadbalancer) GetRetryFromContext(r *http.Request) int {
	if retry, ok := r.Context().Value(Retry).(int); ok {
		return retry
	}
	return 0
}
