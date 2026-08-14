# AA Bundler Performance Harness

This harness implements the `aa-bundler-performance` OpenChainBench spec for ERC-4337 bundlers and paymasters.

## Configuration

Required environment variables:

- `AA_BUNDLER_RPC_URL` – ERC-4337 bundler JSON-RPC endpoint
- `AA_CHAIN_ID` – Target chain ID (e.g., `1`, `8453`, `42161`)
- `AA_REGION` – Vantage region label (e.g., `us-east-1`)
- `AA_ENTRY_POINT_VERSION` – EntryPoint version (`0.6` or `0.7`)

Optional for state-changing runs:

- `AA_ALLOW_STATE_CHANGING` – Set to `true` to enable submissions (default `false`)
- `AA_MAX_DAILY_SPEND_WEI` – Maximum daily spend cap in wei
- `AA_FUNDING_ACCOUNT` – Public address of the funded benchmark account

## Running

Quote-only (safe default):

```bash
export AA_BUNDLER_RPC_URL="https://..."
export AA_CHAIN_ID="8453"
export AA_REGION="us-east-1"
export AA_ENTRY_POINT_VERSION="0.7"

go run ./harnesses/aa-bundler-performance
```

State-changing (use with caution):

```bash
export AA_ALLOW_STATE_CHANGING="true"
export AA_MAX_DAILY_SPEND_WEI="10000000000000000" # 0.01 ETH
export AA_FUNDING_ACCOUNT="0x..."

go run ./harnesses/aa-bundler-performance
```

## Metrics

- `aa_bundler_rpc_latency_seconds`
- `aa_user_operation_acceptance_latency_seconds`
- `aa_user_operation_inclusion_latency_seconds`
- `aa_user_operation_confirmation_latency_seconds`
- `aa_user_operation_success_ratio`
- `aa_paymaster_sponsorship_success_ratio`
- `aa_user_operation_gas_estimate_error_bps`
- `aa_user_operation_rejection_total`
- `aa_bundler_supported_entry_points`

## Safety

- Never commit private keys, API keys, or funded-account secrets
- Use dedicated benchmark accounts with small balances
- Enforce daily spend caps and monitor continuously
- Start with quote-only runs before enabling state-changing submissions
