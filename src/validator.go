package main

import (
	"strconv"

	"github.com/go-akka/configuration"
	"github.com/go-akka/configuration/hocon"
)

func validate(config *configuration.Config, node *node) {
	if !config.HasPath(node.NodeName) {
		// node.Error = append(node.Error, "Config value not found!")
		return
	}
	val := config.GetNode(node.NodeName)

	switch node.Type {
	case "string":
		stringCheck(val, node)
	case "boolean":
		booleanCheck(val, node)
	case "hostname":
		hostnameCheck(val, node)
	case "port":
		portCheck(val, node)
	case "int":
		intCheck(val, node)
	default:
		//node.Error = append(node.Error, "Provided Type:'"+node.Type+"' is not valid!")
	}
}

func stringCheck(hoconValue *hocon.HoconValue, node *node) {
	val := hoconValue.GetString()
	if !requiredAndNullCheck(val, node, "Expected a string but got empty!") {
		return
	}
}

func booleanCheck(hoconValue *hocon.HoconValue, node *node) {
	val := hoconValue.GetString()
	if !requiredAndNullCheck(val, node, "Expected a boolean value but got empty!") {
		return
	}
	if !isBoolean(val) {
		node.Error = append(node.Error, "Value is not a boolean. Expected 'true' or 'false' but got '"+val+"'")
		return
	}

}

func hostnameCheck(hoconValue *hocon.HoconValue, node *node) {
	val := hoconValue.GetString()
	if !requiredAndNullCheck(val, node, "Expected an hostname but got empty!") {
		return
	}
	// TODO: "regex check"
}

func portCheck(hoconValue *hocon.HoconValue, node *node) {
	val := hoconValue.GetString()
	if !requiredAndNullCheck(val, node, "Expected port number but got empty!") {
		return
	}
	if !isPort(val) {
		node.Error = append(node.Error, "Expected valid port number in the range [1, 65535] but got '"+val+"'")
	}
}

func intCheck(hoconValue *hocon.HoconValue, node *node) {
	val := hoconValue.GetString()
	if !requiredAndNullCheck(val, node, "Expected an integer value but got empty!") {
		return
	}
	if !isInt(val) {
		node.Error = append(node.Error, "Expected an integer value but got '"+val+"'")
		return
	}
}

func requiredAndNullCheck(val string, node *node, errorMessage string) bool {
	if len(val) == 0 {
		if node.Required {
			node.Error = append(node.Error, errorMessage)
			return false
		} else {
			node.Warning = append(node.Warning, errorMessage)
		}
	}
	return true
}

func isBoolean(val string) bool {
	if val != "true" && val != "false" {
		return false
	}
	return true
}

func isSwitch(val string) bool {
	if val != "on" && val != "off" {
		return false
	}
	return true
}

func isInt(value string) bool {
	_, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return false
	}
	return true
}

func isPort(value string) bool {
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil || val < 1 || val > 65535 {
		return false
	}
	return true
}
