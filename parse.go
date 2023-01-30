package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type treeStruct struct {
	treeMap interface{}
}

func (treeStruct *treeStruct) UnmarshalTOML(data interface{}) error {
	d, _ := data.(map[string]interface{})

	//d has all the data as nested maps
	keys := make([]interface{}, len(d))
	i := 0
	for k := range d {
		keys[i] = k
		fmt.Printf("%s\n %s", keys[i], d["type"])
		i++
	}

	return nil
}

//changed for running parser.go file (toml implementation)
func main1() {
	f := "../tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	if _, err := os.Stat(f); err != nil {
		f = "../tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	}

	os.Stdout.WriteString("Starting Decoding config file\n")
	var config treeStruct

	meta, err := toml.DecodeFile(f, &config.treeMap)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.WriteString("Decoding config file completed\n")

	keys := meta.Keys()
	mapping := meta.Undecoded()
	for i, k := range keys {
		fmt.Printf("%d %-10s %s\n", i, meta.Type(k...), k)
	}

	fmt.Println(mapping[10])
}
