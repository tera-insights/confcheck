package main

// import (
// 	"fmt"

// 	"github.com/tera-insights/go-akka-configuration"
// )

// type node struct {
// 	NodeName  string
// 	Value     interface{}
// 	Position  interface{}
// 	Validated bool
// }

// func extractEntries(config *configuration.Config) map[string]interface{} {
// 	entries := make(map[string]interface{})
// 	for _, key := range config.GetKeys() {
// 		value := config.GetString(key)
// 		entries[key] = value
// 	}
// 	return entries
// }

// func createPaths(config *configuration.Config, nodeName string, visited map[string]node) {
// 	entries := extractEntries(config)
// 	for key, value := range entries {
// 		if subConfig, ok := value.(*configuration.Config); ok {
// 			newNodeName := nodeName + "." + key
// 			createPaths(subConfig, newNodeName, visited)
// 		} else {
// 			currNode, ok := visited[nodeName]
// 			if !ok {
// 				currNode = node{}
// 				currNode.NodeName = nodeName
// 				visited[nodeName] = currNode
// 			}
// 			fmt.Println(nodeName + "." + key)
// 		}
// 	}
// }

// func main() {
// 	hoconFile := "/workspaces/confcheck/files/config.conf"
// 	hoconConf := configuration.LoadConfig(hoconFile)

// 	visited := make(map[string]node)
// 	createPaths(hoconConf, "", visited)
// }
