package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	dieIf(err)
	var v interface{}
	dieIf(json.NewDecoder(f).Decode(&v))
	f.Close()
	j, _ := json.MarshalIndent(v, "", "\t")
	os.Stdout.Write(j)
}

func dieIf(err interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
