package main

import (
	"flag"
	"log"
	"os"
)

var (
	h          bool
	debug      bool
	fquery     string
	fwfile     string
	fofaDomain string
	fofaSize   int
	chromePath string
	Version    = "v0.0.1"
)

func init() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.StringVar(&fquery, "q", "", "fofa query")
	flag.StringVar(&fofaDomain, "fofa-domain", "soall.org", "fofa domain")
	flag.IntVar(&fofaSize, "size", 10000, "fofa query size")
	flag.StringVar(&chromePath, "p", "", "chrome path")
	flag.Parse()

	showBanner()
	if h || len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}
	if fquery == "" {
		log.Print("fquery not empty")
		os.Exit(1)
	}
}
func main() {
	getResult(fquery)
}
