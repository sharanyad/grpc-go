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
	"fmt"
	"log"
	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/sendint/sendint"
)

const (
	//address     = "128.105.37.223:50051"
	address       = "localhost:50051"
	defaultNumber = 1
	numberOfRuns  = 100
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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
	var floatF float32 = 1.12
	var doubleD float64 = 3.142
	var longL int64 = 17234234231

	fmt.Printf("\nStarting int")
	for i := 0; i < numberOfRuns; i++ {
		_, err = c.EchoInt(context.Background(), &pb.Wrapper{Number: intI})
	}

	//time.Sleep(2 * time.Second)
	fmt.Printf("\nStarting float")
	for i := 0; i < numberOfRuns; i++ {
		_, err = c.EchoFloat(context.Background(), &pb.WrapperF{Number: floatF})
	}

	fmt.Printf("\nStarting double")
	for i := 0; i < numberOfRuns; i++ {
		_, err = c.EchoDouble(context.Background(), &pb.WrapperD{Number: doubleD})
	}

	fmt.Printf("\nStarting long")
	for i := 0; i < numberOfRuns; i++ {
		_, err = c.EchoLong(context.Background(), &pb.WrapperL{Number: longL})
	}

	for stringLength := 1024; stringLength < 524289; stringLength *= 2 {
		fmt.Printf("\nStarting str len %d", stringLength)
		str := RandStringRunes(stringLength)
		for i := 0; i < numberOfRuns; i++ {
			_, err = c.EchoString(context.Background(), &pb.WrapperS{Number: str})
		}
	}
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
