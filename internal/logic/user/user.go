package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zcyc/gf-grpc-template/internal/model/entity"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zcyc/gf-grpc-template/api/pbentity"
	"github.com/zcyc/gf-grpc-template/internal/dao"
	"github.com/zcyc/gf-grpc-template/internal/model/do"
	"github.com/zcyc/gf-grpc-template/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) GetById(ctx context.Context, uid uint64) (*pbentity.User, error) {
	var (
		user   *entity.User
		userPB *pbentity.User
	)

	// 从数据库中获取用户信息
	err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&user)
	if user == nil {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, g.I18n().T(ctx, "{#user-not-found}"))
	}

	// 转换 pb 对象
	if err := gconv.Struct(user, &userPB); err != nil {
		return nil, gerror.NewCodef(gcode.CodeInternalError, g.I18n().T(ctx, "{#system-busy}"))
	}

	// 处理 pb 时间
	userPB.CreateAt = timestamppb.New(user.CreateAt.Time)
	userPB.UpdateAt = timestamppb.New(user.UpdateAt.Time)

	return userPB, err
}

func (s *sUser) DeleteById(ctx context.Context, uid uint64) error {
	_, err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Delete()
	return err
}
