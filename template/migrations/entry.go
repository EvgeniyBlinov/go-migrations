package main

import (
	gmEngine "github.com/EvgeniyBlinov/go-migrations/engine"
	gmStore "github.com/EvgeniyBlinov/go-migrations/engine/store"
	"github.com/EvgeniyBlinov/go-migrations/template/migrations/list"
)

func main() {
	e := gmEngine.NewEngine()
	e.Run(getMigrationsList())
	e.GetConnector().Close()
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateExampleTable{},
		// Register you migrations here
	}
}
