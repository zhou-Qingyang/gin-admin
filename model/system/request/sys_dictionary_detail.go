package request

import (
	"github.com/zhou-Qingzhang/gin-admin/model/common/request"
	"github.com/zhou-Qingzhang/gin-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
