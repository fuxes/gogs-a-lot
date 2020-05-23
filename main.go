// Copyright 2020 Martín Becerra. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.
//
// Gogs-a-lot is tool to print random log messages to stdout
// Useful for debugging logging systems
//
// Inspired from: https://github.com/themoosman/logs-a-lot
//
package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	chars       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	infoMessage = "{\"message\": %q, \"severity\": \"INFO\"}"
)

// randSeq returns a random sequence of alphanumeric characters of length n
func randSeq(n int) string {
	str := make([]byte, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func main() {
	log.SetFlags(0)

	for {
		log.Printf(infoMessage, randSeq(64))
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}
