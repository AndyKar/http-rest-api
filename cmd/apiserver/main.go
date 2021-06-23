package main

import (
	"github.com/AndyKar/http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	if err :=s.Start(); err != nill{
		log.Fatal(err)
	}
}