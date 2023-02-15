package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/BurntSushi/toml"
// 	//configuration "github.com/tera-insights/go-akka-configuration"
// )

// // Node Struct for Hocon
// type Node struct {
// 	Key         string
// 	Value       interface{}
// 	Parent      *Node
// 	Child       []*Node
// 	Required    bool
// 	//Position    *Position
// 	Description string
// 	Type        string
// 	Error       []string
// 	Warning     []string
// 	Verbose     []string
// }

// // Map Struct for TOML
// type ConfigTree map[string]interface{}

// func main1() {

// 	// Parsing TOML Config File into a Tree DS
// 	var config ConfigTree
// 	specFile := "/workspaces/confcheck/files/spec.toml"
// 	_, err := toml.DecodeFile(specFile, &config)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Error decoding toml file:", err)
// 		os.Exit(1)
// 	}
// 	//PrintConfigTree(config, "")

// 	// Parsing HOCON Config File into a Tree DS
// 	//hoconFile := "/workspaces/confcheck/files/config.conf"
// 	// hoconConf := configuration.LoadConfig(hoconFile)

// 	// tempMap := map[string][2]int{}
// 	//BuildConfTree("hoconFile", hoconConf, nil , tempMap)
// 	//fmt.Println(root)



// 	// Creating a Map of Keys and Location (line / col)
// 	//(testing for toml tree)
// 	// keyMap := map[string]Node{}
// 	// tempMap := map[string][2]int{}
// 	// Dfs(config, hoconConf, "", nil, keyMap, tempMap)
// 	// PrintMapValues(tempMap)

// }

// // Building Hocon Config Tree
// func (n *Node) addChild(node *Node) {
// 	n.Child = append(n.Child, node)
// 	node.Parent = n
// }

// // func buildTree(key string, value interface{}, parent *Node, config *akka.Config) *Node {
// // 	node := &Node{
// // 		Key:    key,
// // 		Value:  value,
// // 		Parent: parent,
// // 		Child:  []*Node{},
// // 	}

// // 	if parent != nil {
// // 		parent.addChild(node)
// // 	}

// // 	loc, _ := config.GetPosition(key)
// // 	if loc != hocon.Position{} {
// // 		node.Data = [2]int{loc.Line, loc.Column}
// // 	}

// // 	if v, ok := value.(map[string]interface{}); ok {
// // 		for k, val := range v {
// // 			buildTree(k, val, node, config)
// // 		}
// // 	}

// // 	return node
// // }



// // Print Config Tree - Test
// func PrintConfigTree(tree ConfigTree, indent string) {
// 	for key, value := range tree {
// 		switch v := value.(type) {
// 		case map[string]interface{}:
// 			fmt.Printf("%s%s:\n", indent, key)
// 			PrintConfigTree(v, indent+"  ")
// 		default:
// 			fmt.Printf("%s%s = %v\n", indent, key, v)
// 		}
// 	}
// }

// //Add values to Map
// func AddToMap(mapp map[string][2]int, key string, value [2]int) {
//     mapp[key] = value
// }

// //Print values to Map
// func PrintMapValues(mapp map[string][2]int) {
//     for _, value := range mapp {
//         fmt.Println(value)
//     }
// 	fmt.Println("Printed")
// }
