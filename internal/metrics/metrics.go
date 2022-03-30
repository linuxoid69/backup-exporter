package metrics

import (
	"context"
	"log"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunExporter(ctx context.Context, bucket string, workDirs []string, minioClient *minio.Client) {
	prometheus.MustRegister(timeCreateBackup)

	route := gin.Default()

	route.Use(func(context *gin.Context) {
		collectMetrics(ctx, bucket, workDirs, minioClient)
	})

	route.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	})

	if err := route.Run(":9925"); err != nil {
		log.Println("%w", err)
	}
}

func collectMetrics(ctx context.Context, bucket string, workDirs []string, minioClient *minio.Client) {
	getTimeCreateBackup(ctx, bucket, workDirs, minioClient)
}

func getTimeCreateBackup(ctx context.Context, bucket string, workDirs []string, minioClient *minio.Client) {
	timeCreateBackup.Reset()
	for _, dir := range workDirs {
		for i := range minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true, Prefix: dir}) {
			timestamp := i.LastModified.Unix()
			backupPath := path.Dir(i.Key)
			// TODO here can exclude old backups.
			timeCreateBackup.WithLabelValues(backupPath).Set(float64(timestamp))
		}
	}
}
