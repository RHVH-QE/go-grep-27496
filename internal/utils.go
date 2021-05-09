package internal

import (
	"flag"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

var UsageText = `NAME
    go-grep

SYNOPSIS
    go-grep [-e] [pattern] [file]

DESCRIPTION
    only implement basic features of original "grep"

EXAMPLES
    1.
      $ go-grep 'pattern' myfile
    2.
      $ go-grep -e '^\.Pp' myfile
    3.
      $ cat myfile | go-grep 'pattern'

The following options are available:

`

func ReadData() []byte {
	data, _ := os.Stdin.Stat()
	if (data.Mode() & os.ModeCharDevice) == 0 {
		log.Debug("data is from pipe")
		if flag.NArg() != 1 {
			log.Fatal("need 1 position arg")
		}
		bytes, _ := ioutil.ReadAll(os.Stdin)
		return bytes
	}
	log.Debug("data is not from pipe, continue")

	if flag.NArg() != 2 {
		log.Fatal("need 2 postition args")
	}
	data2, err := ioutil.ReadFile(flag.Arg(1))
	if err != nil {
		log.Error(err)
		return []byte{}
	}
	return data2
}
