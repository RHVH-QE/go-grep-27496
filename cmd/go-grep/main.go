package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/dracher/go-grep/internal"
)

var regexpFlag = flag.Bool("e", false, "Interpret pattern as an regular expression")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", internal.UsageText)
		flag.PrintDefaults()
	}
	log.SetLevel(log.InfoLevel)
	flag.Parse()
}

func main() {
	internal.Grep(
		[]byte(flag.Arg(0)),
		internal.ReadData(),
		*regexpFlag)
}
