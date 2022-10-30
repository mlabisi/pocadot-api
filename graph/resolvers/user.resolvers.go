package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pocadot-api/graph/generated"
	"pocadot-api/graph/model"
)

// PaymentMethods is the resolver for the paymentMethods field.
func (r *userAccountResolver) PaymentMethods(ctx context.Context, obj *model.UserAccount) ([]model.PaymentMethod, error) {
	panic(fmt.Errorf("not implemented: PaymentMethods - paymentMethods"))
}

// Biases is the resolver for the biases field.
func (r *userAccountResolver) Biases(ctx context.Context, obj *model.UserAccount) ([]model.Talent, error) {
	panic(fmt.Errorf("not implemented: Biases - biases"))
}

// SavedListings is the resolver for the savedListings field.
func (r *userAccountResolver) SavedListings(ctx context.Context, obj *model.UserAccount) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: SavedListings - savedListings"))
}

// SavedProfiles is the resolver for the savedProfiles field.
func (r *userAccountResolver) SavedProfiles(ctx context.Context, obj *model.UserAccount) ([]*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: SavedProfiles - savedProfiles"))
}

// ConnectedAccounts is the resolver for the connectedAccounts field.
func (r *userAccountResolver) ConnectedAccounts(ctx context.Context, obj *model.UserAccount) ([]*model.IdentityProvider, error) {
	panic(fmt.Errorf("not implemented: ConnectedAccounts - connectedAccounts"))
}

// Profile is the resolver for the profile field.
func (r *userAccountResolver) Profile(ctx context.Context, obj *model.UserAccount) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: Profile - profile"))
}

// CreatedListings is the resolver for the createdListings field.
func (r *userAccountResolver) CreatedListings(ctx context.Context, obj *model.UserAccount) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: CreatedListings - createdListings"))
}

// SentOffers is the resolver for the sentOffers field.
func (r *userAccountResolver) SentOffers(ctx context.Context, obj *model.UserAccount) ([]*model.Offer, error) {
	panic(fmt.Errorf("not implemented: SentOffers - sentOffers"))
}

// Socials is the resolver for the socials field.
func (r *userProfileResolver) Socials(ctx context.Context, obj *model.UserProfile) ([]*model.SocialProfile, error) {
	panic(fmt.Errorf("not implemented: Socials - socials"))
}

// UserAccount returns generated.UserAccountResolver implementation.
func (r *Resolver) UserAccount() generated.UserAccountResolver { return &userAccountResolver{r} }

// UserProfile returns generated.UserProfileResolver implementation.
func (r *Resolver) UserProfile() generated.UserProfileResolver { return &userProfileResolver{r} }

type userAccountResolver struct{ *Resolver }
type userProfileResolver struct{ *Resolver }
