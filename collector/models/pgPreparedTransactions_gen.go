/*generated by binding gen*/
package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PgPreparedTransactionsSlice []PgPreparedTransactions

func (r PgPreparedTransactionsSlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgPreparedTransactions(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgPreparedTransactions) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["database"] = r.Database

	// optional labels

	// metrics
	// oldest (CounterValue)
	var oldest float64
	if r.Oldest.IsZero() {
		oldest = float64(0)
	} else {
		oldest = float64(r.Oldest.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `oldest`), `Time at which the oldest transaction was prepared for commit`, nil, labels,
		), prometheus.CounterValue, oldest,
	)

	// count (GaugeValue)
	count := float64(r.Count)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `count`), `Number of prepared transactions`, nil, labels,
		), prometheus.GaugeValue, count,
	)

	// optional metrics

	return nil
}
