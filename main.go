package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	inPlace = flag.Bool("i", false, "format in place")
)

func main() {
	flag.Parse()
	for _, fn := range flag.Args() {
		fmt.Fprintf(os.Stderr, "formatting %s\n", fn)
		f, err := os.Open(fn)
		dieIf(err)
		var v interface{}
		json.NewDecoder(f).Decode(&v)
		f.Close()
		j, _ := json.MarshalIndent(v, "", "\t")
		if *inPlace {
			f, err = os.Create(fn)
			dieIf(err)
			fmt.Fprintf(f, "%s\n", j)
			f.Close()
		} else {
			fmt.Printf("%s\n", j)

		}
	}
}

func dieIf(err interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
