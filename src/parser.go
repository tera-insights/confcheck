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
	NodeName    string
	Default     interface{}
	Description string
	Required    bool
	Type        string
	Error       []string
	Warning     []string
	Verbose     []string
	Errors      map[string][]string
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
	printLog(Nodes)
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
				currNode.NodeName = nodeName
				validate(config, &currNode) //validation check
				visited[nodeName] = currNode
			}
		}
	}

}

func printLog(nodes map[string]node) {
	Red := "\033[31m"
	Reset := "\033[0m"
	count := 0
	for _, v := range nodes {
		if len(v.Error) == 0 {
			continue
		}
		count++
		fmt.Printf("%v> %v \n", count, v.NodeName)
		errCount := 0
		for error := range v.Error {
			errCount++
			// colored := fmt.Sprintf(Red+"\t%v. %v", errCount, v.Error[error]+Reset)
			colored := fmt.Sprintf(Red+"\t%v", v.Error[error]+Reset)
			fmt.Println(colored)
		}
		fmt.Println()
	}
}
