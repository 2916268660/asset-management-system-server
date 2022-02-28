package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"strings"
	"time"
)

type ManagementLogic struct {
}

// Login 登录
func (m *ManagementLogic) Login(ctx *gin.Context, loginInfo *request.LoginUserInfo) (token string, err error) {
	var user *model.SysUser
	switch loginInfo.Way {
	case global.WayByUserName: // 通过用户名登录
		if user, err = loginByUserName(ctx, loginInfo); err != nil {
			return "", err
		}
	default:
		return "", errors.New("登录方式错误")
	}
	// 获取token
	token, err = utils.GetToken(user)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return token, nil
}

// 通过用户名登录
func loginByUserName(ctx *gin.Context, loginInfo *request.LoginUserInfo) (user *model.SysUser, err error) {
	user, err = userModel.GetUserByUserId(ctx, loginInfo.UserId)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != utils.Encrypt(loginInfo.Password) {
		return nil, errors.New("密码错误,请重新输入")
	}
	return
}

// RegisterUser 注册用户
func (m *ManagementLogic) RegisterUser(ctx *gin.Context, info *request.RegisterUserInfo) error {
	_, err := userModel.FindUserByUserId(ctx, info.UserId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}
	now := time.Now()
	user := &model.SysUser{
		UserName:   info.UserName,
		UserId:     info.UserId,
		Password:   utils.Encrypt(info.Password),
		Department: info.Department,
		Email:      info.Email,
		Phone:      info.Phone,
		Role:       global.User, // 默认创建用户为基本用户
		CreateTime: now,
		UpdateTime: now,
	}
	err = userModel.SaveUser(ctx, user)
	if err != nil {
		global.GLOBAL_LOG.Error("创建用户失败", zap.Any("user", user), zap.Error(err))
		return errors.New("注册失败")
	}
	return nil
}

// Register2 批量注册
func (m *ManagementLogic) Register2(ctx *gin.Context, path string) error {
	rows, err := utils.ReadRowsForExcel(path)
	if err != nil {
		global.GLOBAL_LOG.Error("解析文件失败", zap.String("file", path), zap.Error(err))
		return errors.New("录入失败,解析文件错误")
	}
	if len(rows) <= 0 {
		return errors.New("录入失败,文件内容为空")
	}
	var users []*model.SysUser
	now := time.Now()
	for _, row := range rows {
		// 保证不会数组越界产生panic
		if len(row) < 5 {
			return errors.New("录入失败,数据缺失")
		}
		err = checkRow(row)
		if err != nil {
			return err
		}
		user := &model.SysUser{
			UserId:     row[0],
			UserName:   row[1],
			Password:   utils.Encrypt(row[0][5:]),
			Email:      row[3],
			Phone:      row[2],
			Department: row[4],
			CreateTime: now,
			UpdateTime: now,
			Role:       global.User,
		}
		users = append(users, user)
	}
	if len(users) <= 0 {
		return errors.New("录入失败，信息为空")
	}
	if err = global.GLOBAL_DB.Create(&users).Error; err != nil {
		global.GLOBAL_LOG.Error("批量创建失败", zap.Error(err))
		if strings.Contains(err.Error(), "Duplicate") {
			return errors.New(fmt.Sprintf("录入失败,%s账号已存在", strings.Split(err.Error(), " ")[4]))
		}
		return errors.New("录入失败")
	}
	return nil
}

// checkRow 校验excel每行数据是否合规
func checkRow(row []string) error {
	if row[0] == "" || len(row[0]) > 20 {
		return errors.New("数据格式不合格")
	}
	if row[1] == "" {
		return errors.New("数据格式不合格")
	}
	if row[2] == "" {
		return errors.New("数据格式不合格")
	}
	if row[4] == "" {
		return errors.New("数据格式不合格")
	}
	return nil
}

// GetUserInfo 获取用户详情
func (m *ManagementLogic) GetUserInfo(ctx *gin.Context, userId string) (*response.UserInfo, error) {
	user, err := userModel.GetUserByUserId(ctx, userId)
	if err != nil {
		global.GLOBAL_LOG.Error("获取用户信息失败", zap.String("userId", userId), zap.Error(err))
		return nil, errors.New("获取用户信息失败")
	}
	userInfo := &response.UserInfo{
		UserName: user.UserName,
		UserId:   user.UserId,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	return userInfo, nil
}

// SetRole 给用户设置角色
func (m *ManagementLogic) SetRole(ctx *gin.Context, info *request.UserRole) error {
	if info.Role != global.User && info.Role != global.Charger && info.Role != global.Provider {
		return errors.New("角色错误")
	}

	// 开启事务
	err := global.GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
		var user *model.SysUser
		err := tx.Where("user_id", info.UserId).First(&user).Error
		if err != nil {
			global.GLOBAL_LOG.Error("查询当前角色信息失败", zap.String("userId", info.UserId), zap.Error(err))
			return err
		}
		if user.Department != info.Department {
			global.GLOBAL_LOG.Error("非法请求, 用户部门与请求部门不相同")
			return errors.New("非法错误")
		}
		// 先查出来目前该部门的负责人，然后将该负责人转换为普通角色
		var charger *model.SysUser
		err = tx.Where("department = ? and role = ?", info.Department, global.Charger).First(&charger).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			global.GLOBAL_LOG.Error("获取该部门当前负责人信息失败", zap.String("department", info.Department), zap.Error(err))
			return err
		}
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			charger = nil
		}
		// 如果有负责人, 将负责人的角色改为user
		if charger != nil {
			if err = tx.Model(&charger).Where("user_id = ?", charger.UserId).Update("role", global.User).Error; err != nil {
				global.GLOBAL_LOG.Error("更新用户角色失败", zap.String("userId", info.UserId), zap.String("role", info.Role), zap.Error(err))
				return err
			}
		}
		// 更新用户角色为charger
		if err = tx.Model(&user).Where("user_id = ?", info.UserId).Update("role", info.Role).Error; err != nil {
			global.GLOBAL_LOG.Error("更新用户角色失败", zap.String("userId", info.UserId), zap.String("role", info.Role), zap.Error(err))
			return err
		}
		// 添加负责人部门信息, 没有则创建, 有则更新
		var department *model.Department
		if err = tx.Where("department = ?", info.Department).First(&department).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Create(&model.Department{
					Department: info.Department,
					UserId:     info.UserId,
				})
				return nil
			}
			return err
		}
		if err = tx.Model(&department).Where("department = ?", info.Department).Update("user_id", info.UserId).Error; err != nil {
			global.GLOBAL_LOG.Error("更新部门负责人失败", zap.String("department", info.Department), zap.String("userId", info.UserId), zap.Error(err))
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("更新用户角色失败")
	}
	return nil
}
