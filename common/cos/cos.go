package cos

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"server/common/config"
)

var _ Client = (*cosClient)(nil)

type Client interface {
	UploadFile(ctx context.Context) (err error)
}

func (c cosClient) UploadFile(ctx context.Context) (err error) {
	uploadInfo, err := c.client.FPutObject(ctx, config.RunData.MinioConfig.BucketName, "tes11t/sh.sh", "/Users/xu756/开发/小程序/appServer/idl.sh", minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return err
	}
	log.Print(uploadInfo.Key)
	return nil
}
