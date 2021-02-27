package repository

import (
	"fmt"
	"go-web-mini/common"
	"go-web-mini/model"
	"go-web-mini/vo"
	"strings"
)

type IOperationLogRepository interface {
	GetOperationLogs(req *vo.OperationLogListRequest) ([]model.OperationLog, int64, error)
	BatchDeleteOperationLogByIds(ids []uint) error
	SaveOperationLogChannel(olc <-chan *model.OperationLog) //处理OperationLogChan将日志记录到数据库
}

type OperationLogRepository struct {
}

func NewOperationLogRepository() IOperationLogRepository {
	return OperationLogRepository{}
}

func (o OperationLogRepository) GetOperationLogs(req *vo.OperationLogListRequest) ([]model.OperationLog, int64, error) {
	var list []model.OperationLog
	db := common.DB.Model(&model.OperationLog{}).Order("start_time DESC")

	username := strings.TrimSpace(req.Username)
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	ip := strings.TrimSpace(req.Ip)
	if ip != "" {
		db = db.Where("ip LIKE ?", fmt.Sprintf("%%%s%%", ip))
	}
	path := strings.TrimSpace(req.Path)
	if path != "" {
		db = db.Where("path LIKE ?", fmt.Sprintf("%%%s%%", path))
	}
	status := req.Status
	if status != 0 {
		db = db.Where("status = ?", status)
	}

	// 分页
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := req.PageNum
	pageSize := req.PageSize
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}

	return list, total, err

}

func (o OperationLogRepository) BatchDeleteOperationLogByIds(ids []uint) error {
	err := common.DB.Where("id IN (?)", ids).Unscoped().Delete(&model.OperationLog{}).Error
	return err
}

//var Logs []model.OperationLog //全局变量多个线程需要加锁，所以每个线程自己维护一个
//处理OperationLogChan将日志记录到数据库
func (o OperationLogRepository) SaveOperationLogChannel(olc <-chan *model.OperationLog) {
	// 只会在线程开启的时候执行一次
	Logs := make([]model.OperationLog, 0)

	// 一直执行--收到olc就会执行
	for log := range olc {
		Logs = append(Logs, *log)
		// 每10条记录到数据库
		if len(Logs) > 5 {
			common.DB.Create(&Logs)
			Logs = make([]model.OperationLog, 0)
		}
	}
}
