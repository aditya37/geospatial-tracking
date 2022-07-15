package service

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	getenv "github.com/aditya37/get-env"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func serve(grpc *grpc.Server, httpHandler http.Handler) error {
	lis, err := net.Listen("tcp",
		fmt.Sprintf(":%s", getenv.GetString("PORT", "")),
	)
	if err != nil {
		return err
	}
	m := cmux.New(lis)
	grpcl := m.Match(cmux.HTTP2(), cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
	httpl := m.Match(cmux.HTTP1Fast())

	http := &http.Server{
		Handler: routeHandler(grpc, httpHandler),
	}
	log.Println("Geotracking service run on port", getenv.GetString("PORT", "1111"))
	// serve grpc
	go grpc.Serve(grpcl)
	// serve http
	go http.Serve(httpl)
	return m.Serve()
}

func routeHandler(grpc http.Handler, other http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if r.ProtoMajor == 2 && strings.HasPrefix(
			r.Header.Get("Content-Type"), "application/grpc") {
			grpc.ServeHTTP(rw, r)
		} else {
			other.ServeHTTP(rw, r)
		}
	})
}

func healthcheck() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, `{"status": "UP"}`)
	})
}
