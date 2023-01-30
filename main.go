package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-akka/configuration"
	"github.com/go-akka/configuration/hocon"
	"github.com/mitchellh/mapstructure"
	// 	"github.com/en-vee/aconf"
	// 	validator "gopkg.in/go-playground/validator.v9"
)

type node struct {
	NodeName    string
	Default     interface{}
	Description string
	Required    bool
	Type        string
	Position    *Position
	Error       []string
	Warning     []string
	Verbose     []string
}

type nodes struct {
	Node []node
}

type Position struct {
	Line int
	Col  int
	Len  int
}

type tree map[string]interface{}

func NewPosition(pos interface{}) *Position {
	res := Position{}
	switch t := pos.(type) {
	case hocon.Position: //find this
		x := pos.(hocon.Position)
		res.Line = x.Line
		res.Col = x.Col
		res.Len = x.Len
	default:
		fmt.Printf("type %v is not supported \n", t)
	}
	return &res
}

func main() {
	specFilePath := "../tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	if _, err := os.Stat(specFilePath); err != nil {
		specFilePath = "../tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	}
	buff, err := os.Readlink("../tiCrypt-backend/default-config/ticrypt-auth.conf")
	if err != nil {
		fmt.Print(err)
	}

	//for hocon file
	conf := configuration.ParseString(string(buff))

	//for the toml
	var tree tree
	meta, err := toml.DecodeFile(specFilePath, &tree)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	Nodes := map[string]node{}
	dfs(tree, conf, "", nil, Nodes, meta)
	fmt.Printf("Nodes count: %v\n", len(Nodes))
	printLog(Nodes)
	printJsonLog(Nodes)
	exportToJson(Nodes)

	// 	var parser ConfigParser
	// 	var filepath string
	// 	var extension string
	// 	fmt.Println("Enter config file path:")
	// 	fmt.Scan(&filepath)
	// 	fmt.Println("Enter config file extension:")
	// 	fmt.Scan(&extension)

	// 	switch extension {
	// 	case "hocon":
	// 		parser = HoconParser{}
	// 	case "toml":
	// 		parser = TomlParser{}
	// 	default:
	// 		fmt.Println("Invalid config file extension.")
	// 		return
	//	}

	// 	cfg, err := parser.Parse(filepath)
	// 	if err != nil {
	// 		log.Fatalf("Failed to parse config file: %s", err)
	// 	}
	// 	fmt.Println("Config:", cfg)

}

func dfs(tree tree, config *configuration.Config, nodeName string, parentNode tree, visited map[string]node, meta toml.MetaData) {
	for i, j := range tree {
		_ = j
		nestedtree, ok := tree[i].(map[string]interface{})
		if ok {
			c := nodeName
			if len(nodeName) > 0 {
				c = nodeName + "."
			}
			dfs(nestedtree, config, c+i, tree, visited, meta)
		} else {
			currNode, ok := visited[nodeName]
			if !ok {
				currNode = node{}
				currMap := parentNode[nodeName[1+strings.LastIndex(nodeName, "."):]]
				mapstructure.Decode(currMap, &currNode)
				currNode.NodeName = nodeName
				currNode.Position = NewPosition(config.GetPosition(nodeName))
				validate(config, &currNode)
				//validation check

				visited[nodeName] = currNode
				fmt.Printf("%v %v \n", config.GetPosition(nodeName), nodeName)
			}
		}
	}
}

// // Temp struct - change later
// type Config struct {
// 	Username string `config:"username" validate:"required"`
// 	Password string `config:"password" validate:"required"`
// }

// type ConfigParser interface {
// 	Parse(filepath string) (*Config, error)
// }

// type HoconParser struct{}

// func (p HoconParser) Parse(filepath string) (*Config, error) {
// 	// var appConfig = &ConfigFile{}
// 	// if err := parser.Parse(reader, appConfig); err != nil {
// 	// 			fmt.Printf("Error %v", err)
// 	// }
// 	var cfg Config
// 	err := aconf.Load(&cfg, "/home/siddhant/workspace/test/hoconfiles/test.conf")
// 	if err != nil {
// 		return nil, err
// 	}
// 	validate := validator.New()
// 	err = validate.Struct(cfg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &cfg, nil
// }

// type TomlParser struct{}

// func (p TomlParser) Parse(filepath string) (*Config, error) {
// 	var cfg Config
// 	_, err := toml.DecodeFile(filepath, &cfg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &cfg, nil
// }
