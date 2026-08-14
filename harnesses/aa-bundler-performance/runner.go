package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func RunQuoteOnlyBenchmark(cfg *Config) error {
	client := NewBundlerClient(cfg.BundlerRPCURL, 10*time.Second)
	ctx := context.Background()

	log.Println("Running quote-only benchmark (no state-changing UserOperations)")

	eps, latency, err := client.SupportedEntryPoints(ctx)
	if err != nil {
		return fmt.Errorf("supported entry points: %w", err)
	}
	log.Printf("Supported entry points (%v): %v\n", latency, eps)

	// TODO: Implement controlled quote-only loops:
	// - Construct minimal test UserOperation fixtures per test vector
	// - Call eth_estimateUserOperationGas and record latency and gas fields
	// - Emit metrics to aa_bundler_rpc_latency_seconds and aa_user_operation_gas_estimate_error_bps
	// - Do not call eth_sendUserOperation unless AA_ALLOW_STATE_CHANGING=true and spend caps are configured

	return nil
}

func RunStateChangingBenchmark(cfg *Config) error {
	if !cfg.AllowStateChanging {
		return fmt.Errorf("state-changing benchmark disabled; set AA_ALLOW_STATE_CHANGING=true and configure AA_MAX_DAILY_SPEND_WEI")
	}
	if cfg.MaxDailySpendWei == 0 {
		return fmt.Errorf("AA_MAX_DAILY_SPEND_WEI must be set for state-changing runs")
	}

	client := NewBundlerClient(cfg.BundlerRPCURL, 10*time.Second)
	ctx := context.Background()

	log.Println("Running state-changing benchmark (controlled UserOperation submissions)")

	eps, latency, err := client.SupportedEntryPoints(ctx)
	if err != nil {
		return fmt.Errorf("supported entry points: %w", err)
	}
	log.Printf("Supported entry points (%v): %v\n", latency, eps)

	// TODO: Implement safe state-changing loops:
	// - Load deterministic smart-account fixture (no secrets in repo)
	// - For each test vector:
	//   * Call eth_estimateUserOperationGas
	//   * Construct and sign UserOperation using fixture account
	//   * Call eth_sendUserOperation, record acceptance latency and userOpHash
	//   * Poll eth_getUserOperationReceipt until inclusion or timeout
	//   * Record inclusion and confirmation latency, success/failure, and gas used
	//   * Update rolling success ratio and sponsorship ratio metrics
	//   * Enforce AA_MAX_DAILY_SPEND_WEI and halt if exceeded
	// - Classify failures using rejection_class labels per methodology

	return nil
}
