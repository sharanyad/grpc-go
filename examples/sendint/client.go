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
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/sendint/sendint"
)

const (
	//address     = "128.105.37.223:50051"
	address       = "localhost:50051"
	defaultNumber = 1
	numberOfRuns  = 100000
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
	//var stringS string = "erhwkjejkhfjkhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"
	var start time.Time
	var elapsed time.Duration
	var r *pb.Wrapper
	var rf *pb.WrapperF
	var rd *pb.WrapperD
	var rl *pb.WrapperL
	//var rs *pb.WrapperS

	var elapsedTime int64
	elapsedTime = 0
	for i := 0; i < numberOfRuns; i++ {
		start = time.Now()
		r, err = c.EchoInt(context.Background(), &pb.Wrapper{Number: intI})
		elapsed = time.Since(start)
		if i == 0 {
			fmt.Printf("RTT for first packet of int: %d", elapsed.Nanoseconds())
		}
		elapsedTime += elapsed.Nanoseconds()
	}
	fmt.Printf("\n Average RTT for int: %d", elapsedTime/numberOfRuns)
	/*if err != nil {
		log.Fatalf("could not greet: %v", err)
	}*/
	fmt.Printf("\n Number: %d", r.Number)

	elapsedTime = 0
	for i := 0; i < numberOfRuns; i++ {
		start = time.Now()
		rf, err = c.EchoFloat(context.Background(), &pb.WrapperF{Number: floatF})
		elapsed = time.Since(start)
		if i == 0 {
			fmt.Printf("RTT for first packet of float: %d\n", elapsed.Nanoseconds())
		}
		elapsedTime += elapsed.Nanoseconds()
	}
	fmt.Printf("\n Average RTT for float: %d\n", elapsedTime/numberOfRuns)
	/*if err != nil {
		log.Fatalf("could not greet: %v", err)
	}*/
	fmt.Printf("\n Number: %d\n", rf.Number)

	elapsedTime = 0
	//time.Sleep(2 * time.Second)
	for i := 0; i < numberOfRuns; i++ {
		start = time.Now()
		rd, err = c.EchoDouble(context.Background(), &pb.WrapperD{Number: doubleD})
		elapsed = time.Since(start)
		if i == 0 {
			fmt.Printf("RTT for first packet of double: %d\n", elapsed.Nanoseconds())
		}
		elapsedTime += elapsed.Nanoseconds()
	}
	fmt.Printf("\n Average RTT for double: %d", elapsedTime/numberOfRuns)
	fmt.Printf("\n Number: %d\n", rd.Number)

	elapsedTime = 0
	//time.Sleep(2 * time.Second)
	for i := 0; i < numberOfRuns; i++ {
		start = time.Now()
		rl, err = c.EchoLong(context.Background(), &pb.WrapperL{Number: longL})
		elapsed = time.Since(start)
		if i == 0 {
			fmt.Printf("RTT for first packet of long: %d", elapsed.Nanoseconds())
		}
		elapsedTime += elapsed.Nanoseconds()
	}
	fmt.Printf("\n Average RTT for long: %d", elapsedTime/numberOfRuns)
	fmt.Printf("\n Number: %d\n", rl.Number)

	for stringLength := 1024; stringLength < 524289; stringLength *= 2 {
		elapsedTime = 0
		//time.Sleep(2 * time.Second)
		str := RandStringRunes(stringLength)
		for i := 0; i < numberOfRuns; i++ {
			start = time.Now()
			_, err = c.EchoString(context.Background(), &pb.WrapperS{Number: str})
			elapsed = time.Since(start)
			if i == 0 {
				fmt.Printf("\nRTT for first packet of string of length %d: %d", stringLength, elapsed.Nanoseconds())
			}
			elapsedTime += elapsed.Nanoseconds()
		}
		fmt.Printf("\n Average RTT for string of length %d: %d", stringLength, elapsedTime/numberOfRuns)
		//fmt.Printf("\n Number: %d\n", rs.Number)
	}
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
