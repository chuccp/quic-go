package http3

import (
	"crypto/tls"
	"github.com/quic-go/quic-go/internal/testdata"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
)

func TestName2(t *testing.T) {

	udp, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Println("qqqqqq", err)
		return
	}
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
	}
	cl := NewClient(udp, tlsConf)

	res, err := cl.Get("https://127.0.0.1:4253/")
	if err != nil {
		log.Println("))))))))", err)
		return
	}
	log.Println(io.ReadAll(res.Body))
}
func TestName(t *testing.T) {
	fullpem, privkey := testdata.GetCertificatePaths()

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("1111111111111111"))

	})

	err := ListenAndServeQUIC("0.0.0.0:4253", fullpem, privkey, serveMux)
	if err != nil {
		return
	}

}
