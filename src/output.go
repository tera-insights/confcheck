package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
)

type Response struct {
	NodeName string
	Errors   []string
	Position *Position
}

type JsonResponse struct {
	Count  int
	Errors []Response
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
		fmt.Printf("Error at line:%v col:%v len:%v \n", v.Position.Line, v.Position.Col, v.Position.Len)
		errCount := 0
		for err := range v.Error {
			errCount++
			// colored := fmt.Sprintf(Red+"\t%v. %v", errCount, v.Error[error]+Reset)
			colored := fmt.Sprintf(Red+"\t%v", v.Error[err]+Reset)
			fmt.Println(colored)
		}
		fmt.Println()
	}
}

func exportToJson(nodes map[string]node) {
	b, err := getJsonResponse(nodes)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("output.json", b, 0644)
	//fmt.Println(string(b))

}

func getJsonResponse(nodes map[string]node) ([]byte, error) {
	count := 0
	var res []Response
	for _, v := range nodes {
		if len(v.Error) == 0 {
			continue
		}
		count++
		temp := Response{
			v.NodeName,
			v.Error,
			v.Position,
		}

		res = append(res, temp)
	}
	jsonResponse := JsonResponse{
		count,
		res,
	}
	return json.MarshalIndent(jsonResponse, "", "	")
}

func printJsonLog(nodes map[string]node) {
	jsonResponse, err := getJsonResponse(nodes)
	if err != nil {
		panic(err)
	}
	var obj interface{}
	json.Unmarshal(jsonResponse, &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
