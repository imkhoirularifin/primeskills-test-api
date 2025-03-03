package interfaces

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"time"
)

type GcloudStorageService interface {
	UploadFile(ctx *gin.Context, bucket string, object []byte, objectName string) error
	DownloadFile(ctx *gin.Context, bucket, objectName string) ([]byte, error)
	DeleteFile(ctx *gin.Context, bucket, objectName string) error
	GenSignedUrl(bucket, objectName string, expire time.Duration) (string, error)
	GetObjectMetadata(ctx *gin.Context, bucket, objectName string) (*storage.ObjectAttrs, error)
}
