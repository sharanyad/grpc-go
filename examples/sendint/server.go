/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/sendint/sendint"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer

func (s *server) EchoFloat(ctx context.Context, in *pb.WrapperF) (*pb.WrapperF, error) {
	return &pb.WrapperF{Number: in.Number}, nil
}

func (s *server) EchoInt(ctx context.Context, in *pb.Wrapper) (*pb.Wrapper, error) {
      return &pb.Wrapper{Number: in.Number}, nil
  }

func (s *server) EchoString(ctx context.Context, in *pb.WrapperS) (*pb.WrapperS, error) {
      return &pb.WrapperS{Number: in.Number}, nil
  }

func (s *server) EchoComplex(ctx context.Context, in *pb.WrapperComplex) (*pb.WrapperComplex, error) {
        return &pb.WrapperComplex{Inti: in.Inti, Floatf: in.Floatf, Strings: in.Strings}, nil
    }


/*func (s *server) EchoComplex(in *pb.WrapperComplex, stream pb.SendInt_EchoComplexServer) error {
        err:= stream.Send(&pb.WrapperComplex{Inti: in.Inti, Floatf: in.Floatf, Strings: in.Strings})
		if err != nil {
			return err
		}
		return nil
}*/


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSendIntServer(s, &server{})
	s.Serve(lis)
}
