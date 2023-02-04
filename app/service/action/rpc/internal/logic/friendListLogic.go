package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/action/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

	"SimpleDouYin/app/service/action/rpc/internal/svc"
	"SimpleDouYin/app/service/action/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *pb.FriendListReq) (*pb.FriendListResp, error) {
	resp := new(pb.FriendListResp)

	//查询关注列表
	var (
		follows []model.Favorite
		foCt    int64
	)

	err := l.svcCtx.GormDB.Where("target_uid = ?", in.UserId).
		Find(&follows).
		Count(&foCt)
	if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
		log.Println("关注列表查询出错:", err.Error, " count:", foCt)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	if foCt == 0 {
		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "成功"
		return resp, nil
	}
	log.Println("follows查询成功", foCt)

	//查询用户
	var (
		users []model.User
		UIDS  []int64
	)
	for idx := 0; int64(idx) < foCt; idx++ {
		UIDS = append(UIDS, follows[idx].UserID)
	}

	err = l.svcCtx.GormDB.Where("user_id in ?", UIDS).
		Order(tool.FiledStringBuild("user_id", UIDS)).
		Find(&users)
	if err.Error != nil {
		log.Println("用户查询出错:", err.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, err.Error
	}
	log.Println("用户查询成功:", users)

	//当前用户与UIDS关注关系
	var Tfollows []model.Follow
	followRela := make([]bool, foCt)
	for i := 0; int64(i) < foCt; i++ {
		followRela[i] = false
	}

	errr1 := l.svcCtx.GormDB.
		Where("uid = ? and target_uid in ?", in.CurUser, UIDS).
		Find(&Tfollows)
	if errr1.Error != nil && (!errors.Is(errr1.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr1.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, err.Error
	}
	for _, v := range Tfollows {
		for ku, vu := range UIDS {
			if vu == v.TargetUID {
				followRela[ku] = true
			}
		}
	}
	log.Println("关注关系查询成功:", followRela)

	//整合
	for i := 0; int64(i) < foCt; i++ {
		var Author pb.Author
		//author
		Author.Id = users[i].UserID
		Author.FollowerCount = users[i].FollowCount
		Author.FollowerCount = users[i].FollowerCount
		Author.Name = users[i].Name
		Author.IsFollow = followRela[i]

		resp.UserList = append(resp.UserList, &Author)
	}
	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "获取列表成功"
	return resp, nil
}
