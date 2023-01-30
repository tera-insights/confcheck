# confcheck
Configuration file checker

This is a Go program that reads a config file and converts it to a map data structure. The program is using the toml and hocon packages for parsing the config file. The program uses DFS (depth-first search) to traverse the config data, collects the metadata of each node, and stores it in a map structure. The collected data can then be used for error and warning reporting, as well as for exporting desired outputs. The config file can be either in hocon or toml format.