package main

// import (
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strconv"
// 	"strings"

// 	configuration "github.com/tera-insights/go-akka-configuration"
// 	"github.com/tera-insights/go-akka-configuration/hocon"
// )

// func validate(config *configuration.Config, node *Node) {
// 	if !config.HasPath(node.Key) {
// 		// node.Error = append(node.Error, "Config value not found!")
// 		return
// 	}
// 	val := config.GetNode(node.Key)

// 	switch node.Type {
// 	case "string":
// 		stringCheck(val.GetString(), node)
// 	case "[]string":
// 		stringArrayCheck(val, node)
// 	case "boolean":
// 		booleanCheck(val, node)
// 	case "switch":
// 		switchCheck(val, node)
// 	case "hostname":
// 		hostnameCheck(val, node)
// 	case "port":
// 		portCheck(val, node)
// 	case "int":
// 		intCheck(val, node)
// 	case "file":
// 		fileCheck(val, node)
// 	default:
// 		complexTypeCheck(val, node)
// 	}
// }

// func stringCheck(val string, node *Node) {
// 	if !requiredAndNullCheck(val, node, "Expected a string but got empty!") {
// 		return
// 	}

// 	if node.Type == "string" {
// 		return
// 	}
// 	key := strings.ReplaceAll(node.Type, " ", "")
// 	min, _ := strconv.Atoi(key[strings.Index(key, "[")+1 : strings.Index(key, ":")])
// 	max, _ := strconv.Atoi(key[strings.Index(key, ":")+1 : strings.Index(key, "]")])

// 	if len(val) < min || len(val) > max {
// 		errorMessage := fmt.Sprintf("'%v' has length %v but expected length should be in the range [%v:%v]", val, len(val), min, max)
// 		node.Error = append(node.Error, errorMessage)
// 		return
// 	}
// }

// func stringArrayCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	if !hoconValue.IsArray() && len(hoconValue.GetString()) > 0 {
// 		node.Error = append(node.Error, "Expected []string but got an unexpected value.")
// 	}
// 	// stringArrayRegex := "\\[(\\s*\"[a-zA-Z0-9!@#$%^&*() <>?:{}|]*\"\\s*,)*\\s*\"[a-zA-Z0-9!@#$%^&*() <>?:{}|]*\"\\s*]"
// 	// match, _ := regexp.MatchString(stringArrayRegex, val)
// 	// if !match {
// 	// 	node.Error = append(node.Error, "'"+val+"' is not of type []string")
// 	// 	return
// 	// }
// }

// func booleanCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	val := hoconValue.GetString()
// 	if !requiredAndNullCheck(val, node, "Expected a boolean value but got empty!") {
// 		return
// 	}
// 	if !isBoolean(val) {
// 		node.Error = append(node.Error, "Value is not a boolean. Expected 'true' or 'false' but got '"+val+"'")
// 		return
// 	}

// }
// func switchCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	val := hoconValue.GetString()
// 	if !requiredAndNullCheck(val, node, "Expected a switch value but got empty!") {
// 		return
// 	}
// 	if !isSwitch(val) {
// 		node.Error = append(node.Error, "Value is not a switch. Expected 'on' or 'off' but got '"+val+"'")
// 		return
// 	}

// }

// func fileCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	path := hoconValue.GetString()
// 	pathRegex := "^(\\/[0-9a-zA-Z_-]+)+\\.\\w+$"
// 	match, _ := regexp.MatchString(pathRegex, path)
// 	if !match {
// 		node.Error = append(node.Error, "File path is not valid: "+path)
// 		return
// 	}
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return
// 	}
// 	if os.IsNotExist(err) {
// 		node.Error = append(node.Error, "File does not exist: "+path)
// 		return
// 	}

// 	node.Error = append(node.Error, "There was an error accessing file: "+path+"Error: "+err.Error())
// }

// func hostnameCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	val := hoconValue.GetString()
// 	if !requiredAndNullCheck(val, node, "Expected an hostname but got empty!") {
// 		return
// 	}
// 	if !isHostname(val) {
// 		node.Error = append(node.Error, "Hostname "+val+" does not have a valid format. Couldn't pass regex check!")
// 		return
// 	}
// }

// func portCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	val := hoconValue.GetString()
// 	if !requiredAndNullCheck(val, node, "Expected port number but got empty!") {
// 		return
// 	}
// 	if !isPort(val) {
// 		node.Error = append(node.Error, "Expected valid port number in the range [1, 65535] but got '"+val+"'")
// 	}
// }

// func intCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	val := hoconValue.GetString()
// 	if !requiredAndNullCheck(val, node, "Expected an integer value but got empty!") {
// 		return
// 	}
// 	if !isInt(val) {
// 		node.Error = append(node.Error, "Expected an integer value but got '"+val+"'")
// 		return
// 	}
// }

// func complexTypeCheck(hoconValue *hocon.HoconValue, node *Node) {
// 	val := hoconValue.GetString()
// 	if !requiredAndNullCheck(val, node, "Expected "+node.Type+" but got empty!") {
// 		return
// 	}
// 	durationRegex := "^duration\\[\\s*[0-9]+\\s*(sec|min|hr|days)\\s*:\\s*[0-9]+\\s*(sec|min|hr|days)\\s*]$"
// 	match, _ := regexp.MatchString(durationRegex, node.Type)
// 	if match {
// 		durationCheck(val, node)
// 		return
// 	}

// 	stringRegex := "^string\\[\\s*[0-9]+\\s*:\\s*[0-9]+\\s*]$"
// 	match, _ = regexp.MatchString(stringRegex, node.Type)
// 	if match {
// 		stringCheck(val, node)
// 		return
// 	}

// 	uriRegex := "^uri<\\s*\\w*\\s*>$"
// 	match, _ = regexp.MatchString(uriRegex, node.Type)
// 	if match {
// 		uriCheck(val, node)
// 		return
// 	}
// 	node.Error = append(node.Error, "Type "+node.Type+" is not valid!!")
// }

// func durationCheck(val string, node *Node) {
// 	//TODO:
// 	//node.Error = append(node.Error, "Type "+node.Type+" is not supported")
// }

// func uriCheck(val string, node *Node) {
// 	uriType := node.Type[1+strings.Index(node.Type, "<") : strings.Index(node.Type, ">")]
// 	switch uriType {
// 	case "mongo":
// 		mongoUriCheck(val, node)
// 	case "https":
// 		httpsUriCheck(val, node)
// 	default:
// 		node.Error = append(node.Error, "URI of type: '"+uriType+"' is not supported!")
// 	}
// }

// func mongoUriCheck(val string, node *Node) {
// 	regex := "^(mongodb:(?:\\/{2})?)((\\w+?):(\\w+?)@|:?@?)(\\w+?):(\\d+)\\/(\\w+?)$"
// 	match, _ := regexp.MatchString(regex, val)
// 	if !match {
// 		node.Error = append(node.Error, "uri '"+val+"' is not valid mongo uri")
// 		return
// 	}
// }

// func httpsUriCheck(val string, node *Node) {
// 	regex := "^https://(www.)?[-a-zA-Z0-9@:%._+~#=]{1,256}.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)$"
// 	match, _ := regexp.MatchString(regex, val)
// 	if !match {
// 		node.Error = append(node.Error, "uri '"+val+"' is not valid https uri")
// 		return
// 	}
// }

// func requiredAndNullCheck(val string, node *Node, errorMessage string) bool {
// 	if len(val) == 0 {
// 		if node.Required {
// 			node.Error = append(node.Error, errorMessage)
// 			return false
// 		} else {
// 			node.Warning = append(node.Warning, errorMessage)
// 		}
// 	}
// 	return true
// }

// func isBoolean(val string) bool {
// 	if val != "true" && val != "false" {
// 		return false
// 	}
// 	return true
// }

// func isSwitch(val string) bool {
// 	if val != "on" && val != "off" {
// 		return false
// 	}
// 	return true
// }

// func isInt(value string) bool {
// 	_, err := strconv.ParseInt(value, 10, 32)
// 	return err == nil
// }

// func isPort(value string) bool {
// 	val, err := strconv.ParseInt(value, 10, 32)
// 	if err != nil || val < 1 || val > 65535 {
// 		return false
// 	}
// 	return true
// }

// // func isIP(value string) bool {
// // 	ValidIpAddressRegex := "^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$"
// // 	match, _ := regexp.MatchString(ValidIpAddressRegex, value)
// // 	return match
// // }

// func isHostname(value string) bool {
// 	ValidHostnameRegex := "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$"
// 	match, _ := regexp.MatchString(ValidHostnameRegex, value)
// 	return match
// }
