package response

import "github.com/zhou-Qingzhang/gin-admin/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
