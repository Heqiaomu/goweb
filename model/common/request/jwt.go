package request

import (
	uuid "github.com/satori/go.uuid"
)

// CustomClaims Custom claims structure

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
