package main

import (
	"testing"
)

func TestParsedStructure(t *testing.T) {
	conf := Config{}
	//name := "Gladys"
	//want := regexp.MustCompile(`\b` + name + `\b`)

	if err := conf.ParseFile("../examples/tiCrypt-backend/default-config/ticrypt-auth.conf"); err != nil {
		t.Error("Failed to read the conf file:" + err.Error())
	}

	if conf.MongoConnection.URI != "mongodb://user:password@localhost:27017/${database}" {
		t.Error("Incorrect URI value, Value received: " + conf.MongoConnection.URI)
	}

	if conf.MongoConnection.Port != int64(27017) {
		t.Error("Incorrect Port value, Value received:", conf.MongoConnection.Port)
	}

	t.Logf("Conf object after: %v", conf)

}
