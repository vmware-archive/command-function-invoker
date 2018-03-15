package server_test

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"path"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/projectriff/command-function-invoker/pkg/function"
	"github.com/projectriff/command-function-invoker/pkg/server"
	"google.golang.org/grpc"
)

var _ = Describe("MessageFunctionServer", func() {

	var (
		fnUri  string
		client function.MessageFunctionClient
	)

	BeforeEach(func() {
		// The test currently expects *ix echo behaviour.
		Expect(runtime.GOOS).NotTo(Equal("windows"))

		fnUri = path.Join("fixtures", "echo.sh")

		addr := randomAddress()

		listener, err := net.Listen("tcp", addr)
		Expect(err).NotTo(HaveOccurred())
		gRpcServer := grpc.NewServer()
		function.RegisterMessageFunctionServer(gRpcServer, server.New(fnUri))
		go func() {
			gRpcServer.Serve(listener)
		}()

		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		Expect(err).NotTo(HaveOccurred())

		client = function.NewMessageFunctionClient(conn)
	})

	It("should respond to a gRPC request with no headers", func() {
		callClient, err := client.Call(context.Background())
		Expect(err).NotTo(HaveOccurred())

		err = callClient.Send(&function.Message{[]byte("hello"), map[string]*function.Message_HeaderValue{}})
		Expect(err).NotTo(HaveOccurred())

		response, err := callClient.Recv()
		Expect(err).NotTo(HaveOccurred())
		Expect(string(response.Payload)).To(Equal("echoed hello"))
		Expect(response.Headers).To(BeEmpty())

	})

	It("should respond to a gRPC request with a correlation id header by returning the same correlation id", func() {
		callClient, err := client.Call(context.Background())
		Expect(err).NotTo(HaveOccurred())

		correlationId := "9876543210"

		err = callClient.Send(&function.Message{[]byte("hello"), createHeadersWithCorrelationId(correlationId)})
		Expect(err).NotTo(HaveOccurred())

		response, err := callClient.Recv()
		Expect(err).NotTo(HaveOccurred())
		Expect(string(response.Payload)).To(Equal("echoed hello"))
		Expect(response.Headers[server.CorrelationId].Values).To(ConsistOf(correlationId))

	})

})

func randomAddress() string {
	port := 1024 + rand.Intn(32768-1024)
	return fmt.Sprintf("localhost:%d", port)
}

func createHeadersWithCorrelationId(correlationId string) map[string]*function.Message_HeaderValue {
	outHeaders := make(map[string]*function.Message_HeaderValue)

	headerValue := function.Message_HeaderValue{
		Values: []string{correlationId},
	}
	outHeaders[server.CorrelationId] = &headerValue

	return outHeaders
}
