package http3

import (
	"crypto/tls"
	"github.com/quic-go/quic-go"
	"net"
	"net/http"
)

func NewClient(conn net.PacketConn, config *tls.Config) *http.Client {
	transport := &quic.Transport{Conn: conn}
	roundTripper := &RoundTripper{
		transport:       transport,
		TLSClientConfig: config,
	}
	return &http.Client{Transport: roundTripper}
}
