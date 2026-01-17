package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type Role string

const (
	RoleOwner  Role = "Owner"
	RoleClient Role = "Client"
	RoleAdmin  Role = "Admin"
)

type Claims struct {
	UserID uint `json:"user_id"`
	Role   Role `json:"role"`
	jwt.RegisteredClaims
}
