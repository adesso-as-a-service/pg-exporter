/*generated by binding gen*/
package models

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type PgStatDatabaseSlice []PgStatDatabase

func (r PgStatDatabaseSlice) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	for _, row := range []PgStatDatabase(r) {
		if err := row.ToMetrics(namespace, subsystem, ch, labelsKV...); err != nil {
			return err
		}
	}
	return nil
}

func (r *PgStatDatabase) ToMetrics(namespace string, subsystem string, ch chan<- prometheus.Metric, labelsKV ...string) error {
	labels := newLabels(labelsKV...)
	// labels
	labels["database_id"] = strconv.FormatInt(r.DatID, 10)
	labels["database"] = r.DatName

	// optional labels

	// metrics
	// backends (GaugeValue)
	backends := float64(r.NumBackends)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `backends`), `Number of backends currently connected to this database`, nil, labels,
		), prometheus.GaugeValue, backends,
	)

	// xact_commited_total (CounterValue)
	xactCommitedTotal := float64(r.XactCommit)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `xact_commited_total`), `Number of transactions in this database that have been committed`, nil, labels,
		), prometheus.CounterValue, xactCommitedTotal,
	)

	// xact_rolledback_total (CounterValue)
	xactRolledbackTotal := float64(r.XactRollback)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `xact_rolledback_total`), `Number of transactions in this database that have been rolled back`, nil, labels,
		), prometheus.CounterValue, xactRolledbackTotal,
	)

	// blocks_read_total (CounterValue)
	blocksReadTotal := float64(r.BlksRead)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `blocks_read_total`), `Number of disk blocks read in this database`, nil, labels,
		), prometheus.CounterValue, blocksReadTotal,
	)

	// blocks_hit_total (CounterValue)
	blocksHitTotal := float64(r.BlksHit)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `blocks_hit_total`), `Number of times disk blocks were found already in the buffer cache, so that a read was not necessary`, nil, labels,
		), prometheus.CounterValue, blocksHitTotal,
	)

	// rows_returned_total (CounterValue)
	rowsReturnedTotal := float64(r.TupReturned)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_returned_total`), `Number of rows returned by queries in this database`, nil, labels,
		), prometheus.CounterValue, rowsReturnedTotal,
	)

	// rows_fetched_total (CounterValue)
	rowsFetchedTotal := float64(r.TupFetched)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_fetched_total`), `Number of rows fetched by queries in this database`, nil, labels,
		), prometheus.CounterValue, rowsFetchedTotal,
	)

	// rows_inserted_total (CounterValue)
	rowsInsertedTotal := float64(r.TupInserted)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_inserted_total`), `Number of rows inserted by queries in this database`, nil, labels,
		), prometheus.CounterValue, rowsInsertedTotal,
	)

	// rows_updated_total (CounterValue)
	rowsUpdatedTotal := float64(r.TupUpdated)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_updated_total`), `Number of rows updated by queries in this database`, nil, labels,
		), prometheus.CounterValue, rowsUpdatedTotal,
	)

	// rows_deleted_total (CounterValue)
	rowsDeletedTotal := float64(r.TupDeleted)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `rows_deleted_total`), `Number of rows deleted by queries in this database`, nil, labels,
		), prometheus.CounterValue, rowsDeletedTotal,
	)

	// conflicts_total (CounterValue)
	conflictsTotal := float64(r.Conflicts)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `conflicts_total`), `Number of queries canceled due to conflicts with recovery in this database`, nil, labels,
		), prometheus.CounterValue, conflictsTotal,
	)

	// temp_files_total (CounterValue)
	tempFilesTotal := float64(r.TempFiles)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `temp_files_total`), `Number of temporary files created by queries in this database`, nil, labels,
		), prometheus.CounterValue, tempFilesTotal,
	)

	// temp_bytes_total (CounterValue)
	tempBytesTotal := float64(r.TempBytes)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `temp_bytes_total`), `Total amount of data written to temporary files by queries in this database`, nil, labels,
		), prometheus.CounterValue, tempBytesTotal,
	)

	// deadlocks_total (CounterValue)
	deadlocksTotal := float64(r.Deadlocks)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `deadlocks_total`), `Number of deadlocks detected in this database`, nil, labels,
		), prometheus.CounterValue, deadlocksTotal,
	)

	// blk_read_seconds_total (CounterValue)
	blkReadSecondsTotal := r.BlkReadTime.Seconds()
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `blk_read_seconds_total`), `Time spent reading data file blocks by backends in this database`, nil, labels,
		), prometheus.CounterValue, blkReadSecondsTotal,
	)

	// blk_write_seconds_total (CounterValue)
	blkWriteSecondsTotal := r.BlkWriteTime.Seconds()
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `blk_write_seconds_total`), `Time spent writing data file blocks by backends in this database`, nil, labels,
		), prometheus.CounterValue, blkWriteSecondsTotal,
	)

	// stats_reset (CounterValue)
	var statsReset float64
	if r.StatsReset.IsZero() {
		statsReset = float64(0)
	} else {
		statsReset = float64(r.StatsReset.Unix())
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, `stats_reset`), `Time at which these statistics were last reset`, nil, labels,
		), prometheus.CounterValue, statsReset,
	)

	// optional metrics

	return nil
}
