package main

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/liyinhgqw/typesafe-config/parse"
)

type Config struct {
	MongoConnection struct {
		URI      string
		Hostname string
		Port     int64
		//	User     string
		Password string
	}
	underlyingData *parse.Tree
}

func (c *Config) ParseFile(path string) error {
	os.Stdout.WriteString("Parsing config file\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("failed to read config file")
	}
	return c.Parse(bytes)
}

func (c *Config) Parse(configFileData []byte) (err error) {
	os.Stdout.WriteString("Parsing config string\n")
	c.underlyingData, err = parse.Parse("ticrypt-auth", string(configFileData))
	if err == nil {
		populateConfigVars(c)
	}
	return
}

func populateConfigVars(c *Config) error {
	typesafeConf := c.underlyingData.GetConfig()
	os.Stdout.WriteString("Printing the extracted structure:\n" + typesafeConf.String())
	val, err := typesafeConf.GetArray("ticrypt.auth")
	//os.Stdout.WriteString(val[0].String())
	if err == nil {
		os.Stdout.WriteString("Val_0:\n" + val[0].String())
	} else {
		os.Stdout.WriteString(err.Error())
	}
	if val, err := typesafeConf.GetString("ticrypt.auth.mongodb.uri"); err == nil {
		c.MongoConnection.URI = unquoteString(val)
		if c.MongoConnection.URI == "" {
			return errors.New("mongodb.uri: can not be empty")
		}
	}

	if val, err := typesafeConf.GetString("ticrypt.auth.mongodb.hostname"); err == nil {
		c.MongoConnection.Hostname = unquoteString(val)
		if c.MongoConnection.Hostname == "" {
			return errors.New("ticrypt.auth.mongodb.hostname: can not be empty")
		}
	}

	if val, err := typesafeConf.GetInt("ticrypt.auth.mongodb.port"); err == nil {
		c.MongoConnection.Port = val
		if c.MongoConnection.Port == 0 {
			return errors.New("ticrypt.auth.mongodb.port: can not be zero")
		}
	}

	/*if val, err := typesafeConf.GetString("ticrypt.auth.mongodb.user"); err == nil {
		c.MongoConnection.User = unquoteString(val)
		if c.MongoConnection.User == "" {
			return errors.New("ticrypt.auth.mongodb.user: can not be empty")
		}
	}*/

	if val, err := typesafeConf.GetString("ticrypt.auth.mongodb.password"); err == nil {
		c.MongoConnection.Password = unquoteString(val)
		if c.MongoConnection.Password == "" {
			return errors.New("ticrypt.auth.mongodb.password: can not be empty")
		}
	}

	return nil
}

func unquoteString(value string) string {
	re := regexp.MustCompile("^\"(.*)\"$")
	if strippedVal := re.FindStringSubmatch(value); strippedVal != nil {
		return strippedVal[1]
	} else {
		return value
	}
}
