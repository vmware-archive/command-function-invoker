/*
 * Copyright 2018-Present the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/projectriff/command-function-invoker/pkg/function"
	"github.com/projectriff/command-function-invoker/pkg/server"
	"google.golang.org/grpc"
	"strconv"
	"net/http"
	"context"
)

func main() {
	var err error
	grpcPort := 10382
	httpPort := 8080

	sGrpcPort := os.Getenv("GRPC_PORT")
	if sGrpcPort != "" {
		grpcPort, err = strconv.Atoi(sGrpcPort)
		if err != nil {
			log.Fatal("Unable to parse GRPC_PORT: ", err)
		}
	}

	sHttpPort := os.Getenv("HTTP_PORT")
	if sHttpPort != "" {
		httpPort, err = strconv.Atoi(sHttpPort)
		if err != nil {
			log.Fatal("Unable to parse HTTP_PORT: ", err)
		}
	}

	fnUri := os.Getenv("FUNCTION_URI")
	if fnUri == "" {
		log.Fatal("Environment variable $FUNCTION_URI not defined")
	}

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRpcServer := grpc.NewServer()
	function.RegisterMessageFunctionServer(gRpcServer, server.New(fnUri))
	go func() {
		gRpcServer.Serve(grpcListener)
		log.Printf("GRPC Server shut down properly")
	}()

	mux := http.NewServeMux()
	httpServer := &http.Server{Addr: fmt.Sprintf(":%d", httpPort),
		Handler: mux,
	}
	mux.HandleFunc("/", server.NewHttpAdapter(fnUri))
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		log.Printf("HTTP Server shut down properly")
	}()

	// Handle shutdown gracefully
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	log.Println("Shutting Down...")
	gRpcServer.GracefulStop()
	httpServer.Shutdown(context.Background())

}
