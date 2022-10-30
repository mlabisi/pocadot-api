package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pocadot-api/graph/generated"
	"pocadot-api/graph/model"
)

// Idols is the resolver for the idols field.
func (r *groupResolver) Idols(ctx context.Context, obj *model.Group) ([]*model.Idol, error) {
	panic(fmt.Errorf("not implemented: Idols - idols"))
}

// Groups is the resolver for the groups field.
func (r *idolResolver) Groups(ctx context.Context, obj *model.Idol) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

// Idol returns generated.IdolResolver implementation.
func (r *Resolver) Idol() generated.IdolResolver { return &idolResolver{r} }

type groupResolver struct{ *Resolver }
type idolResolver struct{ *Resolver }
