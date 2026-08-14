package main

import (
	"os"
	"strconv"
)

type Config struct {
	BundlerRPCURL       string
	ChainID             uint64
	Region              string
	EntryPointVersion   string
	AllowStateChanging  bool
	MaxDailySpendWei    uint64
	FundingAccount      string
}

func LoadConfig() (*Config, error) {
	bundlerRPCURL := os.Getenv("AA_BUNDLER_RPC_URL")
	if bundlerRPCURL == "" {
		return nil, ErrMissingEnv{Key: "AA_BUNDLER_RPC_URL"}
	}

	chainIDStr := os.Getenv("AA_CHAIN_ID")
	if chainIDStr == "" {
		return nil, ErrMissingEnv{Key: "AA_CHAIN_ID"}
	}
	chainID, err := strconv.ParseUint(chainIDStr, 10, 64)
	if err != nil {
		return nil, ErrInvalidEnv{Key: "AA_CHAIN_ID", Reason: err.Error()}
	}

	region := os.Getenv("AA_REGION")
	if region == "" {
		return nil, ErrMissingEnv{Key: "AA_REGION"}
	}

	entryPointVersion := os.Getenv("AA_ENTRY_POINT_VERSION")
	if entryPointVersion == "" {
		return nil, ErrMissingEnv{Key: "AA_ENTRY_POINT_VERSION"}
	}

	allowStateChanging := os.Getenv("AA_ALLOW_STATE_CHANGING") == "true"

	maxDailySpendWei := uint64(0)
	if allowStateChanging {
		maxSpendStr := os.Getenv("AA_MAX_DAILY_SPEND_WEI")
		if maxSpendStr == "" {
			return nil, ErrMissingEnv{Key: "AA_MAX_DAILY_SPEND_WEI"}
		}
		maxDailySpendWei, err = strconv.ParseUint(maxSpendStr, 10, 64)
		if err != nil {
			return nil, ErrInvalidEnv{Key: "AA_MAX_DAILY_SPEND_WEI", Reason: err.Error()}
		}
	}

	fundingAccount := os.Getenv("AA_FUNDING_ACCOUNT")

	return &Config{
		BundlerRPCURL:      bundlerRPCURL,
		ChainID:            chainID,
		Region:             region,
		EntryPointVersion:  entryPointVersion,
		AllowStateChanging: allowStateChanging,
		MaxDailySpendWei:   maxDailySpendWei,
		FundingAccount:     fundingAccount,
	}, nil
}

type ErrMissingEnv struct {
	Key string
}

func (e ErrMissingEnv) Error() string {
	return "missing required environment variable: " + e.Key
}

type ErrInvalidEnv struct {
	Key    string
	Reason string
}

func (e ErrInvalidEnv) Error() string {
	return "invalid environment variable " + e.Key + ": " + e.Reason
}
