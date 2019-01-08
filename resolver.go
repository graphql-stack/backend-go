package backend_go

import (
	"context"
	"github.com/graphql-stack/backend-go/resolver"
	"github.com/graphql-stack/backend-go/service"

	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/types"
)

type Resolver struct{}

func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) PostDetail() PostDetailResolver {
	return &postDetailResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return resolver.CommentAuthor(ctx, obj)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Register(ctx context.Context, registerInput types.RegisterInput) (model.User, error) {
	return resolver.Register(ctx, registerInput)
}
func (r *mutationResolver) Login(ctx context.Context, loginInput types.LoginInput) (model.Token, error) {
	return resolver.Login(ctx, loginInput)
}
func (r *mutationResolver) CreatePost(ctx context.Context, postInput types.PostInput) (model.Post, error) {
	return resolver.CreatePost(ctx, postInput)
}
func (r *mutationResolver) CreateComment(ctx context.Context, commentInput types.CommentInput) (model.Comment, error) {
	return resolver.CreateComment(ctx, commentInput)
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	return resolver.PostAuthor(ctx, obj)
}

type postDetailResolver struct{ *Resolver }

func (r *postDetailResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	return resolver.PostAuthor(ctx, obj)
}
func (r *postDetailResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	return service.GetComments(obj)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (model.User, error) {
	return resolver.Me(ctx)
}
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) (model.PostsList, error) {
	return resolver.Posts(ctx, limit, offset)
}
func (r *queryResolver) Post(ctx context.Context, id string) (model.Post, error) {
	return resolver.Post(ctx, id)
}
