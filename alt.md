package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	configuration "github.com/tera-insights/go-akka-configuration"
// )

// func main() {
// 	hfile := "/workspaces/confcheck/files/config.conf"
// 	hconf := configuration.LoadConfig(hfile)

// 	file, err := os.Open(hfile)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	line := 1
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		if i := strings.Index(text, "="); i != -1 {
// 			key := strings.TrimRight(strings.TrimLeft(text[:i], " \t\n\r#"), " ")
// 			//fmt.Println("key:", key)
// 			value := hconf.GetString(key)
// 			if err == nil {
// 				fmt.Printf("Key '%s' has value '%s' at line %d, column %d\n", key, value, line, i+1)
// 			} else {
// 				fmt.Println("Error retrieving value for key:", err)
// 			}
// 		}
// 		line++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// }

// func addKeyPosition(keyPositions map[string][2]int, key string, line int, col int) {
// 	keyPositions[key] = [2]int{line, col}
// }

// func printKeyPositions(keyPositions map[string][2]int) {
// 	for key, pos := range keyPositions {
// 		fmt.Printf("Key '%s' found at line %d, column %d\n", key, pos[0], pos[1])
// 	}
// }
