module main

go 1.17

//change location here
replace github.com/go-akka/configuration => /workspaces/confcheck/configuration

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/TylerBrock/colorjson v0.0.0-20200706003622-8a50f05110d2
	github.com/fatih/color v1.14.1 // indirect
	github.com/go-akka/configuration v0.0.0-20200606091224-a002c0330665
	github.com/hokaccha/go-prettyjson v0.0.0-20211117102719-0474bc63780f // indirect
	github.com/mitchellh/mapstructure v1.5.0
)
