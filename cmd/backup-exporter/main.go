package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/linuxoid69/backup-exporter/internal/metrics"
	"github.com/linuxoid69/backup-exporter/internal/minio"
)

func main() {
	client := minio.Client{
		Address:   os.Getenv("S3_ADDRESS"),
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
		Bucket:    os.Getenv("S3_BUCKET"),
		UseSSL:    true,
	}

	ctx := context.Background()

	minioClient, err := client.Client(&ctx)
	if err != nil {
		log.Printf("%v", err)
	}

	allTopDirFromBucket := client.GetListTopDir(&ctx, minioClient)

	excludeDirs := strings.Split(strings.ReplaceAll(os.Getenv("S3_EXCLUDE_DIRS"), " ", ""), ",")

	workDirs := minio.FilterDirs(allTopDirFromBucket, excludeDirs)

	metrics.RunExporter(ctx, client.Bucket, workDirs, minioClient)
}
