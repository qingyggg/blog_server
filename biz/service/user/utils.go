package service

import (
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/pkg/utils"
)

// UserAssign: convert orm_gen.User to common.UserBase
func (s *UserService) UserAssign(payload *orm_gen.User) *common.UserBase {
	base := &common.UserBase{
		Id:     payload.ID,
		Name:   payload.UserName,
		HashId: utils.ConvertByteHashToString(payload.HashID),
		Profile: &common.UserProfile{
			Avatar:          utils.URLconvert(s.ctx, s.c, payload.Avatar),
			BackgroundImage: utils.URLconvert(s.ctx, s.c, payload.BackgroundImage),
			Signature:       payload.Signature,
		},
	}
	return base
}
