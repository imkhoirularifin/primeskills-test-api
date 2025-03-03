package gcloud_storage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"primeskills-test-api/internal/domain/interfaces"
	"syscall"
	"time"
)

type service struct {
	storageClient *storage.Client
}

func (s *service) UploadFile(ctx *gin.Context, bucket string, object []byte, objectName string) error {
	writer := s.storageClient.Bucket(bucket).Object(objectName).NewWriter(ctx)
	_, err := io.Copy(writer, bytes.NewReader(object))
	if err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	return nil
}

func (s *service) DownloadFile(ctx *gin.Context, bucket, objectName string) ([]byte, error) {
	reader, err := s.storageClient.Bucket(bucket).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *service) DeleteFile(ctx *gin.Context, bucket, objectName string) error {
	err := s.storageClient.Bucket(bucket).Object(objectName).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GenSignedUrl(bucket, objectName string, expire time.Duration) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  http.MethodGet,
		Expires: time.Now().Add(expire),
	}

	signedUrl, err := s.storageClient.Bucket(bucket).SignedURL(objectName, opts)
	if err != nil {
		return "", err
	}

	return signedUrl, nil
}

func (s *service) GetObjectMetadata(ctx *gin.Context, bucket, objectName string) (*storage.ObjectAttrs, error) {
	meta, err := s.storageClient.Bucket(bucket).Object(objectName).Attrs(ctx)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func NewService() interfaces.GcloudStorageService {
	// Init
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	// On Destroy
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		client.Close()
		os.Exit(0)
	}()

	return &service{
		storageClient: client,
	}
}
