package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	aaBundlerRPCLatencySeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "aa_bundler_rpc_latency_seconds",
			Help:    "JSON-RPC response latency for ERC-4337 bundler methods",
			Buckets: prometheus.ExponentialBuckets(0.01, 2, 12),
		},
		[]string{"provider", "chain", "region", "entry_point_version", "rpc_method", "outcome"},
	)

	aaUserOperationAcceptanceLatencySeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "aa_user_operation_acceptance_latency_seconds",
			Help:    "Time from eth_sendUserOperation request until UserOperation hash returned",
			Buckets: prometheus.ExponentialBuckets(0.05, 2, 12),
		},
		[]string{"provider", "chain", "region", "entry_point_version", "account_type", "sponsorship_mode", "outcome", "failure_reason"},
	)

	aaUserOperationInclusionLatencySeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "aa_user_operation_inclusion_latency_seconds",
			Help:    "Time from UserOperation submission to inclusion in handleOps transaction",
			Buckets: prometheus.ExponentialBuckets(1, 2, 14),
		},
		[]string{"provider", "chain", "region", "entry_point_version", "account_type", "sponsorship_mode", "outcome", "failure_reason"},
	)

	aaUserOperationConfirmationLatencySeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "aa_user_operation_confirmation_latency_seconds",
			Help:    "Time from UserOperation submission to configured confirmation threshold",
			Buckets: prometheus.ExponentialBuckets(5, 2, 12),
		},
		[]string{"provider", "chain", "region", "entry_point_version", "confirmations", "outcome"},
	)

	aaUserOperationSuccessRatio = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "aa_user_operation_success_ratio",
			Help: "Rolling fraction of submitted UserOperations that execute successfully",
		},
		[]string{"provider", "chain", "region", "entry_point_version", "account_type", "sponsorship_mode", "window_seconds"},
	)

	aaPaymasterSponsorshipSuccessRatio = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "aa_paymaster_sponsorship_success_ratio",
			Help: "Rolling fraction of valid sponsored UserOperations that execute successfully",
		},
		[]string{"provider", "paymaster", "chain", "region", "entry_point_version", "window_seconds"},
	)

	aaUserOperationGasEstimateErrorBps = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "aa_user_operation_gas_estimate_error_bps",
			Help:    "Difference between estimated and actual gas cost in basis points",
			Buckets: prometheus.ExponentialBuckets(10, 2, 10),
		},
		[]string{"provider", "chain", "region", "entry_point_version", "gas_component", "account_type"},
	)

	aaUserOperationRejectionTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "aa_user_operation_rejection_total",
			Help: "Total rejected UserOperations, grouped by normalized reason",
		},
		[]string{"provider", "chain", "region", "entry_point_version", "rejection_class"},
	)

	aaBundlerSupportedEntryPoints = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "aa_bundler_supported_entry_points",
			Help: "Number of EntryPoint contracts reported as supported",
		},
		[]string{"provider", "chain", "region"},
	)
)
