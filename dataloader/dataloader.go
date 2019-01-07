package dataloader

import (
	"errors"
	"github.com/graphql-stack/backend-go/db"
	"github.com/graphql-stack/backend-go/model"
	"time"
)

func userToDataLoaderResp(keys []string, data []model.User) ([]*model.User, []error) {
	notFound := errors.New("NOT_FOUND")
	l := len(keys)

	tmpMap := make(map[string]model.User, l)
	resp := make([]*model.User, l)
	errs := make([]error, l)

	for _, v := range data {
		tmpMap[v.ID] = v
	}

	for i, key := range keys {
		d, ok := tmpMap[key]
		if ok {
			resp[i] = &d
		} else {
			resp[i] = nil
			errs[i] = notFound
		}
	}

	return resp, errs
}

func NewUserLoader(maxBatch int, wait time.Duration) *UserLoader {
	return &UserLoader{
		maxBatch: maxBatch,
		wait:     wait,
		fetch: func(keys []string) (users []*model.User, errors []error) {
			var data []model.User
			err := db.ORM.Where("id in (?)", keys).Find(&data).Error
			if err != nil {
				panic(err)
			}

			return userToDataLoaderResp(keys, data)
		},
	}
}
