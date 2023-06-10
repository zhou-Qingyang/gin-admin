package request

import (
	"github.com/zhou-Qingzhang/gin-admin/model/common/request"
	"github.com/zhou-Qingzhang/gin-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
