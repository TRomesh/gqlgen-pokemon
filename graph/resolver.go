package graph

import "github.com/TRomesh/gqlgen-pokemon/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	pokemon []*model.Pokemon
	battle []*model.Battles
}
