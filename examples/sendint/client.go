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
		"time"
		"fmt"
		"golang.org/x/net/context"
		"google.golang.org/grpc"
		pb "google.golang.org/grpc/examples/sendint/sendint"
	   )

const (
		//address     = "128.105.37.223:50051"
		address     = "localhost:50051"
		defaultNumber = 1
		numberOfRuns = 1000000
	  )

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
	defer conn.Close()
		c := pb.NewSendIntClient(conn)

		// Contact the server and print out its response.
		var intI int32 = 1
		//var floatF float32 = 1.12
		//var stringS string = "erhwkjejkhfjkhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"
		/*if len(os.Args) > 1 {
		  number = strconv.Atoi(os.Args[1])
		  }*/
		var start time.Time
		var elapsed time.Duration
		//var elapsed int64
		//var r pb.SendInt_EchoComplexClient
		//var r2 *pb.WrapperComplex
		var r *pb.Wrapper

		var elapsedTime int64
		elapsedTime = 0

		for i:=0; i < numberOfRuns; i++ {
			start = time.Now()
			//r, err = c.EchoComplex(context.Background(), &pb.WrapperComplex{Inti: intI, Floatf: floatF, Strings: stringS})
			//r2, err = r.Recv()
			r, err = c.EchoInt(context.Background(), &pb.Wrapper{Number: intI})
			elapsed = time.Since(start)
			//fmt.Printf("\n RTT : %v", elapsed)
			elapsedTime += elapsed.Nanoseconds()
		}
	fmt.Printf("\n RTT : %d", elapsedTime/numberOfRuns)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("\n Number: %d", r.Number)
	//fmt.Printf("\n Number:%d, float: %d, string: %s", r2.Inti, r2.Floatf, r2.Strings)
}
