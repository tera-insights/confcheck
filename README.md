# Confcheck
Configuration File Checker

## State of ConfCheck Repository 

[Branch : Main]
- The code reads a TOML configuration file, validates it according to a predefined specification, and generates logs and reports using Go programming language.
- It defines two custom types "node" and "Position" and a function "dfs" to traverse a tree structure representing the configuration specification and validate each node in the configuration file.
- The main function sets the path to the TOML specification file, reads the default configuration file, and passes the visited nodes map to several functions to print logs, generate JSON logs, and export the map to a JSON file.

[Branch : Dev]
- Similar Implementation but loads configuration files in TOML and HOCON formats into tree data structures.
- It creates a map of configuration keys and their locations (line and column) performs value validation on the configuration data. [attempts - tbd]
- The code includes functions for building and printing the configuration tree for testing purposes.

[Branch : Test]
- Testing ways to build config tree for hocon using multiple libraries

## Hocon Parsing Library (Go-Akka-Configuration)

Link to HOCON Parsing Library [https://github.com/tera-insights/go-akka-configuration] [Branch : Test]

Functions Added :
- ValueAt()
- TraverseTree()
- traverseHoconTreeValue()
- Updated GetKey() SetKey()
- testing functions

Current State : Parser unable to identify HoconType from the parse String for int and boolean values, and extracting their literals as 'nil' values.

## TOML Parsing Library (Go-Toml)

Link to TOML Parsing Library [https://github.com/tera-insights/go-toml] [Branch : Main]
