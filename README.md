# OpenChainBench AA Bundler Benchmark

Production-grade ERC-4337 bundler and paymaster performance benchmark harness.

## Architecture

- `benchmarks/aa-bundler-performance.yml` – OpenChainBench spec defining metrics, providers, and methodology
- `harnesses/aa-bundler-performance/` – Go harness implementing the spec
- Prometheus `/metrics` endpoint exposing bounded-cardinality metrics
- Health endpoint `/healthz` for liveness/readiness checks

## Safety defaults

- **No private keys, API keys, provider credentials, or funded-account information committed**
- **State-changing UserOperation submission disabled by default**
- `AA_ALLOW_STATE_CHANGING=true` required to enable submissions
- `AA_MAX_DAILY_SPEND_WEI` must be configured for state-changing execution
- Quote-only measurements are safe to run in continuous mode

## Configuration

Required environment variables:

- `AA_BUNDLER_RPC_URL` – ERC-4337 bundler JSON-RPC endpoint
- `AA_CHAIN_ID` – Target chain ID (e.g., 1, 8453, 42161)
- `AA_REGION` – Vantage region label (e.g., `us-east-1`)
- `AA_ENTRY_POINT_VERSION` – EntryPoint version (`0.6` or `0.7`)

Optional for state-changing runs:

- `AA_ALLOW_STATE_CHANGING` – Set to `true` to enable submissions (default `false`)
- `AA_MAX_DAILY_SPEND_WEI` – Maximum daily spend cap in wei
- `AA_FUNDING_ACCOUNT` – Public address of the funded benchmark account (no secrets)

## Running

```bash
go mod tidy
go run ./harnesses/aa-bundler-performance
```

Access points:

- Metrics: http://localhost:8080/metrics
- Health: http://localhost:8080/healthz

## Deployment

- Deploy behind HTTPS with managed secrets for provider credentials
- Configure Prometheus scrape for `/metrics`
- Set conservative `AA_MAX_DAILY_SPEND_WEI` and monitor daily spend
- Use separate deployments per chain and per EntryPoint version

## License

MIT
