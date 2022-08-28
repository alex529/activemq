module github.com/alex529/activemq/notifier

go 1.18

require (
	github.com/alex529/activemq v0.0.0-20220828000127-3e4214b7b70d
	github.com/alex529/activemq/schema v0.0.0
	github.com/aymerick/raymond v2.0.2+incompatible
	golang.org/x/exp v0.0.0-20220827204233-334a2380cb91
)

require gopkg.in/yaml.v2 v2.4.0 // indirect

replace github.com/alex529/activemq/schema v0.0.0 => ../schema
