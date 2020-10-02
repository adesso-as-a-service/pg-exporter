/*generated by binding gen*/
package models

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type PgStatReplicationSlice []PgStatReplication

func (r PgStatReplicationSlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgStatReplication(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgStatReplication) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["pid"] = strconv.Itoa(r.PID)
	labels["application_name"] = r.ApplicationName

	// optional labels
	if r.ClientAddr.Valid {
		labels["client_addr"] = r.ClientAddr.String
	}

	// metrics
	// backend_xmin (CounterValue)
	backendXmin := float64(r.BackendXmin)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `backend_xmin`), `This standby's xmin horizon reported by hot_standby_feedback.`, nil, labels,
		), prometheus.CounterValue, backendXmin,
	)

	// sent_lag_bytes (GaugeValue)
	sentLagBytes := float64(r.SentBytesLag)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `sent_lag_bytes`), `Number of bytes not yet sent on this connection`, nil, labels,
		), prometheus.GaugeValue, sentLagBytes,
	)

	// write_lag_bytes (GaugeValue)
	writeLagBytes := float64(r.WriteBytesLag)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `write_lag_bytes`), `Number of bytes not yet written to this on this standby server`, nil, labels,
		), prometheus.GaugeValue, writeLagBytes,
	)

	// flush_lag_bytes (GaugeValue)
	flushLagBytes := float64(r.FlushBytesLag)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `flush_lag_bytes`), `Number of bytes not yet flushed to disk by this standby server`, nil, labels,
		), prometheus.GaugeValue, flushLagBytes,
	)

	// replay_lag_bytes (GaugeValue)
	replayLagBytes := float64(r.ReplayBytesLag)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `replay_lag_bytes`), `Number of bytes not yet replayed on this standby server`, nil, labels,
		), prometheus.GaugeValue, replayLagBytes,
	)

	// optional metrics
	// write_lag_seconds (GaugeValue)
	if r.WriteLag.Valid {
		writeLagSeconds := r.WriteLag.Float64

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, `write_lag_seconds`), `Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written it`, nil, labels,
			), prometheus.GaugeValue, writeLagSeconds,
		)
	}
	// flush_lag_seconds (GaugeValue)
	if r.FlushLag.Valid {
		flushLagSeconds := r.FlushLag.Float64

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, `flush_lag_seconds`), `Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written and flushed it`, nil, labels,
			), prometheus.GaugeValue, flushLagSeconds,
		)
	}
	// replay_lag_seconds (GaugeValue)
	if r.ReplayLag.Valid {
		replayLagSeconds := r.ReplayLag.Float64

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, subsystem, `replay_lag_seconds`), `Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written, flushed and applied it`, nil, labels,
			), prometheus.GaugeValue, replayLagSeconds,
		)
	}

	return nil
}
