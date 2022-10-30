package graph

import (
	"pocadot-api/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) Amount() generated.AmountResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Group() generated.GroupResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Idol() generated.IdolResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Listing() generated.ListingResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Mutation() generated.MutationResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Offer() generated.OfferResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Query() generated.QueryResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) UserAccount() generated.UserAccountResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) UserProfile() generated.UserProfileResolver {
	//TODO implement me
	panic("implement me")
}
