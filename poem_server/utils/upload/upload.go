package upload

import "mime/multipart"

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
}

func NewOss() OSS {
	return &AliyunOSS{}
}
