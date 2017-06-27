package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/awgh/markov"
)

var inputDir = flag.String("inputdir", "text", "Input Directory, reads all txt files")
var outputFile = flag.String("outfile", "markov.chain", "Output File")

func main() {
	flag.Parse()
	files, err := ioutil.ReadDir(*inputDir) // open all txt files in this directory, generate chain, save to outfile
	if err != nil {
		log.Fatal(err)
	}

	chain := markov.NewChain(2) //prefixLen should likely always be 2

	for _, file := range files {
		fmt.Println(file.Name())

		if strings.HasSuffix(file.Name(), ".txt") {
			dat, err := ioutil.ReadFile(*inputDir + string(os.PathSeparator) + file.Name())
			if err != nil {
				log.Fatal(err.Error())
			}

			sentences := strings.Split(string(dat), ".")

			for _, sentence := range sentences {
				chain.Write(sentence)
			}
		}
	}
	if err := chain.Save(*outputFile); err != nil {
		log.Fatal(err.Error())
	}
}
