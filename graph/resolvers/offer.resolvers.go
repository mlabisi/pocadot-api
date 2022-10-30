package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pocadot-api/graph/generated"
	"pocadot-api/graph/model"
)

// Currency is the resolver for the currency field.
func (r *amountResolver) Currency(ctx context.Context, obj *model.Amount) (string, error) {
	panic(fmt.Errorf("not implemented: Currency - currency"))
}

// Amount is the resolver for the amount field.
func (r *amountResolver) Amount(ctx context.Context, obj *model.Amount) (int, error) {
	panic(fmt.Errorf("not implemented: Amount - amount"))
}

// Listing is the resolver for the listing field.
func (r *offerResolver) Listing(ctx context.Context, obj *model.Offer) (*model.Listing, error) {
	panic(fmt.Errorf("not implemented: Listing - listing"))
}

// MadeBy is the resolver for the madeBy field.
func (r *offerResolver) MadeBy(ctx context.Context, obj *model.Offer) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: MadeBy - madeBy"))
}

// Transaction is the resolver for the transaction field.
func (r *offerResolver) Transaction(ctx context.Context, obj *model.Offer) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: Transaction - transaction"))
}

// Amount returns generated.AmountResolver implementation.
func (r *Resolver) Amount() generated.AmountResolver { return &amountResolver{r} }

// Offer returns generated.OfferResolver implementation.
func (r *Resolver) Offer() generated.OfferResolver { return &offerResolver{r} }

type amountResolver struct{ *Resolver }
type offerResolver struct{ *Resolver }
