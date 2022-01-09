package asset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"server/global"
	"server/models/common"
	"server/models/common/request"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type ManagementLogic struct {
}

// ApplyReceive 申请领用资产
func (m ManagementLogic) ApplyReceive(ctx *gin.Context, applyInfo *request.ApplyReceiveForm) (taskId int64, err error) {
	// 前置校验
	if err := preCheckOne(ctx, applyInfo); err != nil {
		return 0, err
	}
	// 获取当前用户信息
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("get userInfo failed, err=%v", err))
		return 0, err
	}
	now := time.Now()
	task := &common.Task{
		UserId:       claims.UserId,
		UserName:     claims.UserName,
		UserPhone:    claims.Phone,
		Department:   claims.Department,
		Category:     applyInfo.Category,
		Nums:         applyInfo.Nums,
		Days:         applyInfo.Days,
		Remake:       applyInfo.Remake,
		Status:       global.WaitApprove,
		Property:     global.Receive,
		ExpireTime:   now.AddDate(0, 0, applyInfo.Days),
		CreateTime:   now,
		UpdateTime:   now,
		ProvideTime:  time.Unix(0, 0),
		AgreeTime:    time.Unix(0, 0),
		RollbackTime: time.Unix(0, 0),
	}

	// 发送申请
	taskId, err = assetModel.CreateTask(ctx, task)
	if err != nil {
		return 0, errors.New("发起申请失败")
	}
	return
}

// 前置校验
func preCheckOne(ctx *gin.Context, info *request.ApplyReceiveForm) error {
	if info == nil {
		return errors.New("参数错误")
	}
	if info.Category <= 0 {
		return errors.New("资产品类不存在")
	}
	if info.Nums <= 0 {
		return errors.New("申请领用资产数量必须大于0")
	}
	if len(info.Remake) > 600 {
		return errors.New("备注信息超出数字限制")
	}
	if info.Days <= 1 {
		return errors.New("申请天数不能低于1天")
	}
	return nil
}

// ApplyRevert 申请归还资产
func (m *ManagementLogic) ApplyRevert(ctx *gin.Context, applyInfo *request.ApplyRevertForm) (taskId int64, err error) {
	if err = preCheckTwo(ctx, applyInfo); err != nil {
		return 0, err
	}

	// 获取当前用户信息
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("get userInfo failed, err=%v", err))
		return 0, err
	}
	now := time.Now()
	assets, err := json.Marshal(applyInfo.Assets)
	if err != nil {
		log.Println(fmt.Sprintf("json Unmarshal failed err=%v", err))
		return 0, err
	}
	task := &common.Task{
		UserId:       claims.UserId,
		UserName:     claims.UserName,
		UserPhone:    claims.Phone,
		Department:   claims.Department,
		Nums:         len(applyInfo.Assets),
		Assets:       string(assets),
		Remake:       applyInfo.Remake,
		Property:     global.Revert,
		CreateTime:   now,
		UpdateTime:   now,
		ProvideTime:  time.Unix(0, 0),
		AgreeTime:    time.Unix(0, 0),
		RollbackTime: time.Unix(0, 0),
		ExpireTime:   time.Unix(0, 0),
	}
	taskId, err = assetModel.CreateTask(ctx, task)
	if err != nil {
		return 0, errors.New("发起申请失败")
	}
	return
}

func preCheckTwo(ctx *gin.Context, info *request.ApplyRevertForm) error {
	if info == nil {
		return errors.New("参数有误")
	}
	if len(info.Assets) <= 0 {
		return errors.New("申请归还资产数量必须大于0")
	}
	if len(info.Remake) > 600 {
		return errors.New("备注信息超出数字限制")
	}
	return nil
}

func (m *ManagementLogic) ApplyRepair(ctx *gin.Context, applyInfo *request.ApplyRepairForm) (repairId int64, err error) {
	if err = preCheckThree(ctx, applyInfo); err != nil {
		return 0, err
	}
	// todo err 处理
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("get userInfo failed, err=%v", err))
		return 0, err
	}
	now := time.Now()
	repair := &common.Repairs{
		UserId:       claims.UserId,
		UserName:     claims.UserName,
		UserPhone:    claims.Phone,
		Address:      applyInfo.Address,
		Remake:       applyInfo.Remake,
		Status:       global.WaitRepair,
		CreateTime:   now,
		UpdateTime:   time.Unix(0, 0),
		ReceiveTime:  time.Unix(0, 0),
		RepairedTime: time.Unix(0, 0),
		RollbackTime: time.Unix(0, 0),
	}
	repairId, err = assetModel.CreateRepair(ctx, repair)
	if err != nil {
		log.Println(fmt.Sprintf("create repair failed. userId=%d||err=%v", claims.UserId, err))
		return 0, errors.New("发起申请失败")
	}
	return
}

func preCheckThree(ctx *gin.Context, info *request.ApplyRepairForm) error {
	if info == nil {
		return errors.New("参数有误")
	}
	if len(info.Assets) <= 0 {
		return errors.New("申请报修资产数量必须大于0")
	}
	if info.Address == "" && len(info.Address) > 300 {
		return errors.New("地址字数不能超过100字")
	}
	if len(info.Remake) > 600 {
		return errors.New("备注信息超出数字限制")
	}
	return nil
}
