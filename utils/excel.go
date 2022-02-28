package utils

import (
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"server/global"
)

var sheet = "Sheet1"

// ReadRowsForExcel 读取sheet1中每行数据
func ReadRowsForExcel(path string) (res [][]string, err error) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		global.GLOBAL_LOG.Error("打开文件失败", zap.String("file", path), zap.Error(err))
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			global.GLOBAL_LOG.Error("关闭文件失败", zap.String("file", path), zap.Error(err))
		}
	}()
	res, err = file.GetRows(sheet)
	if err != nil {
		global.GLOBAL_LOG.Error("获取行数据失败", zap.String("file", path), zap.Error(err))
		return nil, err
	}
	return
}
