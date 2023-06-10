package response

import "github.com/zhou-Qingzhang/gin-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
