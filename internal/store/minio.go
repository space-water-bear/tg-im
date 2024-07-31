package store

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type MinioClient struct {
	*minio.Client
	Bucket string
}

type MinioOption struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewMinio(option MinioOption, bucket string) *MinioClient {
	minioClient, err := minio.New(option.Endpoint, &minio.Options{
		Secure: option.UseSSL,
		Creds:  credentials.NewStaticV4(option.AccessKeyID, option.SecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println(option)
		panic(err)
	}

	return &MinioClient{Client: minioClient, Bucket: bucket}
}

func (m *MinioClient) Upload(ctx context.Context, objectName string, reader io.Reader, size int64) error {

	info, err := m.PutObject(ctx, m.Bucket, objectName, reader, size, minio.PutObjectOptions{})
	fmt.Println("upload: ", info)
	return err
}

func (m *MinioClient) Download(objectName string) ([]byte, error) {

	return nil, nil
}
