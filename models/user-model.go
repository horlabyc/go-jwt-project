package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"firstName" validate:"required,min=2,max=100"`
	LastName     *string            `json:"lastName" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"email,required"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	Usertype     *string            `json:"userType" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refreshToken"`
	CreatedAt    *string            `json:"createdAt"`
	UpdatedAt    *string            `json:"updatedAt"`
	UserId       *string            `json:"userId"`
}
