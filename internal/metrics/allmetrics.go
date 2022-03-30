package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var timeCreateBackup = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "veeam_backup_timestamp",
		Help: "Time creating backup",
	}, []string{"backup_path"})
