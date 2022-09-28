package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-akka/configuration"
	"github.com/mitchellh/mapstructure"
)

type node struct {
	Default     interface{}
	Description interface{}
	Required    interface{}
	Type        interface{}
}

type nodes struct {
	Node []node
}

type tree map[string]interface{}

type treeNode interface{}

func main() {
	specFilePath := "../examples/tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	if _, err := os.Stat(specFilePath); err != nil {
		specFilePath = "../examples/tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	}
	buff, err := os.ReadFile("../examples/tiCrypt-backend/default-config/ticrypt-auth.conf") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	conf := configuration.ParseString(string(buff))

	var tree tree

	_, err = toml.DecodeFile(specFilePath, &tree)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	Nodes := map[string]node{}
	dfs(tree, conf, "", nil, Nodes)
	fmt.Printf("Nodes count: %v\n", len(Nodes))
	// count := 0
	// for k, v := range Nodes {
	// 	count++
	// 	fmt.Printf("%v %v: %v\n", count, k, v)
	// }
}

func dfs(tree tree, config *configuration.Config, nodeName string, parentNode tree, visited map[string]node) {
	for i, j := range tree {
		_ = j
		nestedtree, ok := tree[i].(map[string]interface{})
		if ok {
			c := nodeName
			if len(nodeName) > 0 {
				c = nodeName + "."
			}
			dfs(nestedtree, config, c+i, tree, visited)
		} else {
			currNode, ok := visited[nodeName]
			if !ok {
				currNode = node{}
				currMap := parentNode[nodeName[1+strings.LastIndex(nodeName, "."):]]
				mapstructure.Decode(currMap, &currNode)
				visited[nodeName] = currNode
				fmt.Printf("%v: %v\n", nodeName, config.GetValue(nodeName))
				// validate(currNode, config.GetString(nodeName)) validation check
			}
		}
	}

}
