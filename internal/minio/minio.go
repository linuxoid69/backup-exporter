package minio

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	Address   string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

func (c *Client) Client(ctx *context.Context) (minioClient *minio.Client, err error) {
	minioClient, err = minio.New(c.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: c.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return minioClient, nil
}

// GetListTopDir get all top directories from bucket.
func (c *Client) GetListTopDir(ctx *context.Context, minioClient *minio.Client) (dirs []string) {
	for obj := range minioClient.ListObjects(*ctx, c.Bucket, minio.ListObjectsOptions{}) {
		if strings.HasSuffix(obj.Key, "/") {
			dirs = append(dirs, obj.Key)
		}
	}

	return dirs
}

// FilterDirs - remove exclude directories from directory list.
func FilterDirs(sourceDirs []string, excludeDirs []string) (destDirs []string) {
	for _, d := range sourceDirs {
		if slices.Contains(excludeDirs, strings.Trim(d, "/")) {
			destDirs = append(destDirs, d)
		}
	}

	return destDirs
}
