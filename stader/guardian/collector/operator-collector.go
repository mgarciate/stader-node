package collector

import (
	"fmt"

	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"

	"github.com/prometheus/client_golang/prometheus"
)

// Represents the collector for the stader network metrics
type OperatorCollector struct {
	ActiveValidators                     *prometheus.Desc
	QueuedValidators                     *prometheus.Desc
	SlashedValidators                    *prometheus.Desc
	ExitingValidators                    *prometheus.Desc
	WithdrawnValidators                  *prometheus.Desc
	TotalSDBonded                        *prometheus.Desc
	UnclaimedClRewards                   *prometheus.Desc
	UnclaimedNonSocializingPoolElRewards *prometheus.Desc
	ETHAPR                               *prometheus.Desc
	SDAPR                                *prometheus.Desc
	CumulativePenalty                    *prometheus.Desc
	ClaimedSDrewards                     *prometheus.Desc
	UnclaimedSocializingPoolELRewards    *prometheus.Desc
	UnclaimedSocializingPoolSdRewards    *prometheus.Desc
	ClaimedSocializingPoolElRewards      *prometheus.Desc
	ClaimedSocializingPoolSdRewards      *prometheus.Desc
	NextRewardCycleTime                  *prometheus.Desc

	// The beacon client
	bc beacon.Client

	// The eth1 client
	ec stader.ExecutionClient

	// The node's address
	nodeAddress common.Address

	// The thread-safe locker for the network state
	stateLocker *StateCache

	// Prefix for logging
	logPrefix string
}

// Create a new NetworkCollector instance
func NewRewardCollector(
	bc beacon.Client,
	ec stader.ExecutionClient,
	nodeAddress common.Address,
	stateLocker *StateCache,
) *OperatorCollector {
	return &OperatorCollector{
		ActiveValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ActiveValidators), "", nil, nil,
		),
		QueuedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, QueuedValidators), "", nil, nil,
		),
		SlashedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, SlashedValidators), "", nil, nil,
		),
		TotalSDBonded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, TotalSDBonded), "", nil, nil,
		),
		ETHAPR: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ETHAPR), "", nil, nil),
		SDAPR: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, SDAPR), "", nil, nil),
		CumulativePenalty: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, CumulativePenalty), "", nil, nil),
		ClaimedSDrewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ClaimedSDrewards), "", nil, nil),

		bc:          bc,
		ec:          ec,
		nodeAddress: nodeAddress,
		stateLocker: stateLocker,
		logPrefix:   "Reward Collector",
	}
}

// Write metric descriptions to the Prometheus channel
func (collector *OperatorCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.ActiveValidators
	channel <- collector.QueuedValidators
	channel <- collector.SlashedValidators
	channel <- collector.TotalSDBonded

	channel <- collector.ETHAPR
	channel <- collector.SDAPR
	channel <- collector.CumulativePenalty
	channel <- collector.ClaimedSDrewards
}

// Collect the latest metric values and pass them to Prometheus
func (collector *OperatorCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetState()

	fmt.Printf("%+v", state)
	channel <- prometheus.MustNewConstMetric(collector.ActiveValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ActiveValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.QueuedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.QueuedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SlashedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SlashedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.CumulativePenalty, prometheus.GaugeValue, float64(state.StaderNetworkDetails.CumulativePenalty.Int64()))
	//channel <- prometheus.MustNewConstMetric(collector.UnclaimedSDRewards, prometheus.GaugeValue, float64(state.StaderNetworkDetails.UnclaimedSDRewards.Int64()))
}

// Log error messages
func (collector *OperatorCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
