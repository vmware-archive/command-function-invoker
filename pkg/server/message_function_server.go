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

package server

import (
	"io"
	"log"

	"github.com/projectriff/shell-function-invoker/pkg/function"
)

type messageFunctionServer struct {
	fnUri string
}

func New(fnUri string) function.MessageFunctionServer {
	return &messageFunctionServer{
		fnUri: fnUri,
	}
}

func (mfs *messageFunctionServer) Call(callServer function.MessageFunction_CallServer) error {
	for {
		in, err := callServer.Recv()
		if err == io.EOF {
			log.Println("EOF returned from callServer.Recv")
			return nil
		}
		if err != nil {
			log.Printf("Error returned from callServer.Recv: %v\n", err)
			return err
		}

		out, err := invoke(mfs.fnUri, in)
		if err != nil {
			continue
		}

		err = callServer.Send(out)
		if err == io.EOF {
			log.Println("EOF returned from callServer.Send")
			return nil
		}
		if err != nil {
			log.Printf("Error returned from callServer.Send: %v\n", err)
			return err
		}
	}
}
