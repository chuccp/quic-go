package http3

import (
	"github.com/quic-go/quic-go"
	"net"
	"net/http"
)

func NewClient(conn net.PacketConn) *http.Client {
	transport := &quic.Transport{Conn: conn}
	roundTripper := &RoundTripper{transport: transport}
	return &http.Client{Transport: roundTripper}
}
