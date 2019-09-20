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
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/projectriff/command-function-invoker/pkg/server"
	"strconv"
	"net/http"
	"context"
)

func main() {
	var err error
	httpPort := 8080

	sHttpPort := os.Getenv("PORT")
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
	httpServer.Shutdown(context.Background())

}
