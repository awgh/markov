package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/awgh/markov"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
}

var inputFile = flag.String("outfile", "markov.chain", "Output File")

func main() {
	flag.Parse()

	chain := markov.NewChain(2) //prefix length should likely always be 2
	if err := chain.Load(*inputFile); err != nil {
		log.Fatal(err)
	}
	for {
		log.Println(chain.Generate(23))
	}
}
