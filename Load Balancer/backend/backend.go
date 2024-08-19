package backend

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend interface {
	SetAlive(bool)
	IsAlive() bool
	GetUrl() *url.URL
	Serve(http.ResponseWriter, *http.Request)
	GetLoad() int
	GetCapacity() int
	GetProxy() *httputil.ReverseProxy
}

type backend struct {
	Url          *url.URL
	Alive        bool
	Mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
	Connections  int
	Capacity     int
}

// sets backend alive
func (b *backend) SetAlive(alive bool) {
	b.Mux.Lock()
	b.Alive = alive
	b.Mux.Unlock()
}

// returns current status of backend
func (b *backend) IsAlive() (alive bool) {
	b.Mux.RLock()
	defer b.Mux.RUnlock()
	alive = b.Alive
	return
}

func (b *backend) GetLoad() int {
	b.Mux.RLock()
	defer b.Mux.RUnlock()
	return b.Connections
}

func (b *backend) GetCapacity() int {
	b.Mux.RLock()
	defer b.Mux.RUnlock()
	return b.Capacity
}

func (b *backend) GetUrl() *url.URL {
	return b.Url
}

func (b *backend) GetProxy() *httputil.ReverseProxy {
	return b.ReverseProxy
}

func (b *backend) Serve(rw http.ResponseWriter, req *http.Request) {
	// Increment connections
	b.Mux.Lock()
	b.Connections++
	b.Mux.Unlock()

	b.ReverseProxy.ServeHTTP(rw, req)

	b.Mux.Lock()
	b.Connections--
	b.Mux.Unlock()

}

func NewBackend(u *url.URL, rp *httputil.ReverseProxy, weight int) Backend {
	return &backend{
		Url:          u,
		ReverseProxy: rp,
		Alive:        true,
		Capacity:     weight,
	}
}
