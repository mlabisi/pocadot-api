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

// Offers is the resolver for the offers field.
func (r *listingResolver) Offers(ctx context.Context, obj *model.Listing) ([]*model.Offer, error) {
	panic(fmt.Errorf("not implemented: Offers - offers"))
}

// Listing returns generated.ListingResolver implementation.
func (r *Resolver) Listing() generated.ListingResolver { return &listingResolver{r} }

type listingResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *listingResolver) Idols(ctx context.Context, obj *model.Listing) ([]*model.Idol, error) {
	panic(fmt.Errorf("not implemented: Idols - idols"))
}
func (r *listingResolver) Groups(ctx context.Context, obj *model.Listing) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}
