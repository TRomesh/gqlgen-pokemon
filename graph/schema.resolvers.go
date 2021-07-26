package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/TRomesh/gqlgen-pokemon/graph/generated"
	"github.com/TRomesh/gqlgen-pokemon/graph/model"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	pokemon := &model.Pokemon{
		Name:        input.Name,
		ID:          input.ID,
		Power:       input.Power,
		Description: input.Description,
	}
	r.pokemon = append(r.pokemon, pokemon)
	return pokemon, nil
}

func (r *mutationResolver) CreateBattle(ctx context.Context, input model.NewBattle) (*model.Battles, error) {
	battle := &model.Battles{
		ID:       input.ID,
		Pokemon:  input.Pokemon,
		Opponent: input.Opponent,
		Location: input.Location,
		Date:     nil,
	}
	r.battle = append(r.battle, battle)
	return battle, nil
}

func (r *queryResolver) Pokemons(ctx context.Context) ([]*model.Pokemon, error) {
	return r.pokemon, nil
}

func (r *queryResolver) Battles(ctx context.Context) ([]*model.Battles, error) {
	return r.battle, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
