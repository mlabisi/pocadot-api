package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pocadot-api/graph/generated"
	"pocadot-api/graph/model"
)

// ListedBy is the resolver for the listedBy field.
func (r *listingResolver) ListedBy(ctx context.Context, obj *model.Listing) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: ListedBy - listedBy"))
}

// Idols is the resolver for the idols field.
func (r *listingResolver) Idols(ctx context.Context, obj *model.Listing) ([]*model.Idol, error) {
	panic(fmt.Errorf("not implemented: Idols - idols"))
}

// Groups is the resolver for the groups field.
func (r *listingResolver) Groups(ctx context.Context, obj *model.Listing) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// Offers is the resolver for the offers field.
func (r *listingResolver) Offers(ctx context.Context, obj *model.Listing) ([]*model.Offer, error) {
	panic(fmt.Errorf("not implemented: Offers - offers"))
}

// Listing returns generated.ListingResolver implementation.
func (r *Resolver) Listing() generated.ListingResolver { return &listingResolver{r} }

type listingResolver struct{ *Resolver }
