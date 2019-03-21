/*
 * Copyright 2018-Present the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/projectriff/command-function-invoker/pkg/function"
	"github.com/projectriff/command-function-invoker/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 10382, "The server port")

	flag.Parse()

	fnUri := os.Getenv("FUNCTION_URI")
	if fnUri == "" {
		log.Fatal("Environment variable $FUNCTION_URI not defined")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRpcServer := grpc.NewServer()
	function.RegisterMessageFunctionServer(gRpcServer, server.New(fnUri))

	// Handle shutdown gracefully
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals
		log.Println("Shutting Down...")
		gRpcServer.GracefulStop()
	}()

	gRpcServer.Serve(listener)
}
