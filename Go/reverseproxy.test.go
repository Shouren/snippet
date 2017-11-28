package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type BuffPool struct {
	pool sync.Pool
}

func (b BuffPool) Get() []byte {
	return b.pool.Get().([]byte)
}

func (b BuffPool) Put(buf []byte) {
	b.pool.Put(buf)
}

func testDirector(req *http.Request) {
	req.URL.Scheme = "http"
	req.URL.Host = "10.110.61.53" + ":" + "80"
}

func newTestTransport() *http.Transport {
	return &http.Transport{
		Proxy: func(_ *http.Request) (*url.URL, error) { return nil, nil },
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second,
			KeepAlive: 300 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 1000,
		TLSHandshakeTimeout: 5 * time.Second,
	}
}

func newReverseProxy() http.Handler {
	bufPool := BuffPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 32*1024)
			},
		},
	}

	return &httputil.ReverseProxy{
		Director:      testDirector,
		Transport:     newTestTransport(),
		FlushInterval: 20 * time.Millisecond,
		BufferPool:    bufPool,
	}
}

func main() {
	urlDispatcher := http.NewServeMux()
	urlDispatcher.Handle("/", newReverseProxy())

	server := &http.Server{
		Addr:         ":10001",
		Handler:      urlDispatcher,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	server.ListenAndServe()
}
