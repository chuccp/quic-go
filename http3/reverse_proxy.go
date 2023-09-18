package http3

import (
	"crypto/tls"
	"github.com/quic-go/quic-go"

	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	remoteUrl    *url.URL
	reverseProxy *httputil.ReverseProxy
}

func NewReverseProxy(url *url.URL, conn net.PacketConn, config *tls.Config) *ReverseProxy {
	reverseProxy := httputil.NewSingleHostReverseProxy(url)
	transport := &quic.Transport{Conn: conn}
	roundTripper := &RoundTripper{
		transport:       transport,
		TLSClientConfig: config,
	}
	reverseProxy.Transport = roundTripper
	return &ReverseProxy{remoteUrl: url, reverseProxy: reverseProxy}
}

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.reverseProxy.ServeHTTP(rw, req)
}
