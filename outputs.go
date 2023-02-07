package main

import (
	"fmt"
)

func printConfigTree(tree ConfigTree, indent string) {
	for key, value := range tree {
		switch v := value.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s:\n", indent, key)
			printConfigTree(v, indent+"  ")
		default:
			fmt.Printf("%s%s = %v\n", indent, key, v)
		}
	}
}
