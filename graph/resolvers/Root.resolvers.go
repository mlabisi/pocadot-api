package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pocadot-api/graph/generated"
	"pocadot-api/graph/model"
)

// AddListing is the resolver for the addListing field.
func (r *mutationResolver) AddListing(ctx context.Context, input model.AddListingInput) (*model.Listing, error) {
	panic(fmt.Errorf("not implemented: AddListing - addListing"))
}

// SaveListing is the resolver for the saveListing field.
func (r *mutationResolver) SaveListing(ctx context.Context, input string) (*model.Listing, error) {
	panic(fmt.Errorf("not implemented: SaveListing - saveListing"))
}

// UnsaveListing is the resolver for the unsaveListing field.
func (r *mutationResolver) UnsaveListing(ctx context.Context, input string) (*model.Listing, error) {
	panic(fmt.Errorf("not implemented: UnsaveListing - unsaveListing"))
}

// DeleteListings is the resolver for the deleteListings field.
func (r *mutationResolver) DeleteListings(ctx context.Context, input []string) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: DeleteListings - deleteListings"))
}

// SkipSuggestedListing is the resolver for the skipSuggestedListing field.
func (r *mutationResolver) SkipSuggestedListing(ctx context.Context, input string) (*model.Listing, error) {
	panic(fmt.Errorf("not implemented: SkipSuggestedListing - skipSuggestedListing"))
}

// CreateProfile is the resolver for the createProfile field.
func (r *mutationResolver) CreateProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: CreateProfile - createProfile"))
}

// SaveProfile is the resolver for the saveProfile field.
func (r *mutationResolver) SaveProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: SaveProfile - saveProfile"))
}

// UnsaveProfile is the resolver for the unsaveProfile field.
func (r *mutationResolver) UnsaveProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: UnsaveProfile - unsaveProfile"))
}

// BlockProfile is the resolver for the blockProfile field.
func (r *mutationResolver) BlockProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: BlockProfile - blockProfile"))
}

// ReportProfile is the resolver for the reportProfile field.
func (r *mutationResolver) ReportProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: ReportProfile - reportProfile"))
}

// MakeOffer is the resolver for the makeOffer field.
func (r *mutationResolver) MakeOffer(ctx context.Context) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented: MakeOffer - makeOffer"))
}

// EditOffer is the resolver for the editOffer field.
func (r *mutationResolver) EditOffer(ctx context.Context) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented: EditOffer - editOffer"))
}

// RescindOffer is the resolver for the rescindOffer field.
func (r *mutationResolver) RescindOffer(ctx context.Context) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented: RescindOffer - rescindOffer"))
}

// AcceptOffer is the resolver for the acceptOffer field.
func (r *mutationResolver) AcceptOffer(ctx context.Context) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented: AcceptOffer - acceptOffer"))
}

// NegotiateOffer is the resolver for the negotiateOffer field.
func (r *mutationResolver) NegotiateOffer(ctx context.Context) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented: NegotiateOffer - negotiateOffer"))
}

// RejectOffer is the resolver for the rejectOffer field.
func (r *mutationResolver) RejectOffer(ctx context.Context) (*model.Offer, error) {
	panic(fmt.Errorf("not implemented: RejectOffer - rejectOffer"))
}

// SendMessage is the resolver for the sendMessage field.
func (r *mutationResolver) SendMessage(ctx context.Context, input model.SendMessageInput) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: SendMessage - sendMessage"))
}

// MakePayment is the resolver for the makePayment field.
func (r *mutationResolver) MakePayment(ctx context.Context) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: MakePayment - makePayment"))
}

// DisputeCharge is the resolver for the disputeCharge field.
func (r *mutationResolver) DisputeCharge(ctx context.Context) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: DisputeCharge - disputeCharge"))
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: ChangePassword - changePassword"))
}

// ForgetPassword is the resolver for the forgetPassword field.
func (r *mutationResolver) ForgetPassword(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: ForgetPassword - forgetPassword"))
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: ResetPassword - resetPassword"))
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: UpdateAccount - updateAccount"))
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: UpdateProfile - updateProfile"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: Logout - logout"))
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: CreateAccount - createAccount"))
}

// DeleteAccount is the resolver for the deleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: DeleteAccount - deleteAccount"))
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, input string) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: Account - account"))
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context, input string) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: Profile - profile"))
}

// MyAccount is the resolver for the myAccount field.
func (r *queryResolver) MyAccount(ctx context.Context) (*model.UserAccount, error) {
	panic(fmt.Errorf("not implemented: MyAccount - myAccount"))
}

// MyProfile is the resolver for the myProfile field.
func (r *queryResolver) MyProfile(ctx context.Context) (*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: MyProfile - myProfile"))
}

// Listings is the resolver for the listings field.
func (r *queryResolver) Listings(ctx context.Context, input model.ListingFilters) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: Listings - listings"))
}

// ListingsFeed is the resolver for the listingsFeed field.
func (r *queryResolver) ListingsFeed(ctx context.Context) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: ListingsFeed - listingsFeed"))
}

// UserSuggestions is the resolver for the userSuggestions field.
func (r *queryResolver) UserSuggestions(ctx context.Context, input string) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: UserSuggestions - userSuggestions"))
}

// FeaturedListings is the resolver for the featuredListings field.
func (r *queryResolver) FeaturedListings(ctx context.Context) ([]*model.Listing, error) {
	panic(fmt.Errorf("not implemented: FeaturedListings - featuredListings"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, input model.UserFilters) ([]*model.UserProfile, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// UsersFeed is the resolver for the usersFeed field.
func (r *queryResolver) UsersFeed(ctx context.Context, page int) (*model.ProfileFeed, error) {
	panic(fmt.Errorf("not implemented: UsersFeed - usersFeed"))
}

// Idols is the resolver for the idols field.
func (r *queryResolver) Idols(ctx context.Context, input model.IdolFilters) ([]*model.Idol, error) {
	panic(fmt.Errorf("not implemented: Idols - idols"))
}

// IdolsFeed is the resolver for the idolsFeed field.
func (r *queryResolver) IdolsFeed(ctx context.Context, page int) (*model.IdolFeed, error) {
	panic(fmt.Errorf("not implemented: IdolsFeed - idolsFeed"))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, input model.GroupFilters) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// GroupsFeed is the resolver for the groupsFeed field.
func (r *queryResolver) GroupsFeed(ctx context.Context, page int) (*model.GroupFeed, error) {
	panic(fmt.Errorf("not implemented: GroupsFeed - groupsFeed"))
}

// Talent is the resolver for the talent field.
func (r *queryResolver) Talent(ctx context.Context, input model.TalentFilters) ([]model.Talent, error) {
	panic(fmt.Errorf("not implemented: Talent - talent"))
}

// TalentFeed is the resolver for the talentFeed field.
func (r *queryResolver) TalentFeed(ctx context.Context) ([]model.Talent, error) {
	panic(fmt.Errorf("not implemented: TalentFeed - talentFeed"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
