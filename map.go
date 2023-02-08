package main

import (
	"fmt"
	"strings"

	configuration "github.com/tera-insights/go-akka-configuration"
	"github.com/tera-insights/go-akka-configuration/hocon"
	"github.com/mitchellh/mapstructure"
)
type Position struct {
	Line int
	Col int
	Len int
}

func NewPosition (pos interface{}) *Position {
	res := Position{}
	switch t := pos.(type) {
	case hocon.Position:
		x := pos.(hocon.Position)
		res.Line = x.Line
		res.Col = x.Col
		res.Len = x.Len
	default:
		fmt.Printf("type %v is not supported \n", t)
	}
	return &res
}

func addToMap(mapp map[string][2]int, key string, value [2]int) {
    mapp[key] = value
}

func printMapValues(mapp map[string][2]int) {
    for _, value := range mapp {
        fmt.Println(value)
    }
}

func dfs(tree ConfigTree, config *configuration.Config, key string, parentNode ConfigTree, visited map[string]Node , tempMap map[string][2]int) {


	for i, j := range tree {
		_ = j
		nestedtree, ok := tree[i].(map[string]interface{})
		if ok {
			c := key
			if len(key) > 0 {
				c = key + "."
			}
			dfs(nestedtree, config, c+i, tree, visited, tempMap)
		} else {
			currNode, ok := visited[key]
			if !ok {
				currNode = Node{}
				currMap := parentNode[key[1+strings.LastIndex(key, "."):]]
				mapstructure.Decode(currMap, &currNode)
				currNode.Key = key
				currNode.Position = NewPosition(config.GetPosition(key))
				
				//add in map
				addToMap(tempMap, key, [2]int{config.Root().GetPosition().Line, config.Root().GetPosition().Col})

				//validate
				//validate(config, &currNode)

				visited[key] = currNode
				fmt.Printf("%v %v \n", config.GetPosition(key), key)
			}
		}
	}
}