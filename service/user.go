package service

import (
	"github.com/graphql-stack/backend-go/db"
	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/types"
	"github.com/graphql-stack/backend-go/validator"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/utils"
)

func Register(input types.RegisterInput) (*model.User, error) {
	err := validator.Validate.Struct(&input)
	if err != nil {
		return nil, types.ErrValidateError
	}
	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Avatar:   input.Avatar,
	}

	err = db.ORM.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(input types.LoginInput) (*model.Token, error) {
	err := validator.Validate.Struct(&input)
	if err != nil {
		return nil, types.ErrValidateError
	}

	var u model.User
	err = db.ORM.Where("email = ?", input.Email).First(&u).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, types.ErrEmailOrPasswordError
		}
		return nil, types.ErrInternalError
	}

	if !utils.ComparePassword(input.Password, u.Password) {
		return nil, types.ErrEmailOrPasswordError
	}

	tk, err := GetOrCreateToken(u)
	if err != nil {
		return nil, types.ErrInternalError
	}

	return tk, nil
}
