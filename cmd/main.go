package main

import (
	"flag"
	"log"

	"github.com/runz0rd/replacer"
)

func main() {
	var flagConfig, flagInput, flagOutput, flagFrom, flagTo string
	flag.StringVar(&flagConfig, "config", "config.yaml", "rules config path")
	flag.StringVar(&flagInput, "input", "", "input file")
	flag.StringVar(&flagOutput, "outut", "", "output file")
	flag.StringVar(&flagFrom, "from", "", "env to find value")
	flag.StringVar(&flagTo, "to", "", "env to replace value")
	flag.Parse()

	c, err := replacer.LoadConfig(flagConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err := replacer.Replace(flagInput, flagOutput, flagFrom, flagTo, c.Rules); err != nil {
		log.Fatal(err)
	}
}
