package ysdb

import (
	"hia/core/types"
)

type YsDb interface {
	UserAdd(user *types.User) error
	UserUpdate(user *types.User) error
	UserQuery(user *types.User) (*[]types.User, error)
	UserDelete(user *types.User) error

	VideoAdd(video *types.Video) error
	VideoUpdate(video *types.Video) error
	VideoQuery(video *types.Video) (*[]types.Video, error)
	VideoDelete(video *types.Video) error

	VideoTransactionAdd(vt *types.VideoTransaction) error
	VideoTransactionQuery(vt *types.VideoTransaction) (*[]types.VideoTransaction, error)
}
