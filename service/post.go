package service

import (
	"github.com/graphql-stack/backend-go/db"
	"github.com/graphql-stack/backend-go/model"
	"github.com/zcong1993/libgo/mysql"
)

func GetPosts(limit, offset int) (*model.PostsList, error) {
	var posts []model.Post
	totalCount, err := mysql.PaginationQuery(db.ORM.Model(new(model.Post)), limit, offset, &posts)
	if err != nil {
		return nil, err
	}

	return &model.PostsList{TotalCount: totalCount, Posts: posts}, nil
}
