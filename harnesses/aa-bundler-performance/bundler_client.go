package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BundlerClient struct {
	httpClient *http.Client
	rpcURL     string
}

func NewBundlerClient(rpcURL string, timeout time.Duration) *BundlerClient {
	return &BundlerClient{
		httpClient: &http.Client{Timeout: timeout},
		rpcURL:     rpcURL,
	}
}

func (c *BundlerClient) SupportedEntryPoints(ctx context.Context) ([]string, time.Duration, error) {
	req := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_supportedEntryPoints",
		Params:  []interface{}{},
		ID:      1,
	}
	start := time.Now()
	respBytes, err := c.callRPC(ctx, req)
	latency := time.Since(start)
	if err != nil {
		return nil, latency, err
	}
	var rpcResp RPCResponse
	if err := json.Unmarshal(respBytes, &rpcResp); err != nil {
		return nil, latency, fmt.Errorf("unmarshal supported entry points: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, latency, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	var eps []string
	if err := json.Unmarshal(rpcResp.Result, &eps); err != nil {
		return nil, latency, fmt.Errorf("parse supported entry points: %w", err)
	}
	return eps, latency, nil
}

func (c *BundlerClient) EstimateUserOperationGas(ctx context.Context, userOp *UserOperation, entryPoint string) (map[string]interface{}, time.Duration, error) {
	req := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_estimateUserOperationGas",
		Params:  []interface{}{userOp, entryPoint},
		ID:      2,
	}
	start := time.Now()
	respBytes, err := c.callRPC(ctx, req)
	latency := time.Since(start)
	if err != nil {
		return nil, latency, err
	}
	var rpcResp RPCResponse
	if err := json.Unmarshal(respBytes, &rpcResp); err != nil {
		return nil, latency, fmt.Errorf("unmarshal estimate gas: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, latency, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(rpcResp.Result, &result); err != nil {
		return nil, latency, fmt.Errorf("parse estimate result: %w", err)
	}
	return result, latency, nil
}

func (c *BundlerClient) SendUserOperation(ctx context.Context, userOp *UserOperation, entryPoint string) (string, time.Duration, error) {
	req := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_sendUserOperation",
		Params:  []interface{}{userOp, entryPoint},
		ID:      3,
	}
	start := time.Now()
	respBytes, err := c.callRPC(ctx, req)
	latency := time.Since(start)
	if err != nil {
		return "", latency, err
	}
	var rpcResp RPCResponse
	if err := json.Unmarshal(respBytes, &rpcResp); err != nil {
		return "", latency, fmt.Errorf("unmarshal send user op: %w", err)
	}
	if rpcResp.Error != nil {
		return "", latency, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	var userOpHash string
	if err := json.Unmarshal(rpcResp.Result, &userOpHash); err != nil {
		return "", latency, fmt.Errorf("parse userOpHash: %w", err)
	}
	return userOpHash, latency, nil
}

func (c *BundlerClient) GetUserOperationReceipt(ctx context.Context, userOpHash string) (*UserOperationReceipt, time.Duration, error) {
	req := RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getUserOperationReceipt",
		Params:  []interface{}{userOpHash},
		ID:      4,
	}
	start := time.Now()
	respBytes, err := c.callRPC(ctx, req)
	latency := time.Since(start)
	if err != nil {
		return nil, latency, err
	}
	var rpcResp RPCResponse
	if err := json.Unmarshal(respBytes, &rpcResp); err != nil {
		return nil, latency, fmt.Errorf("unmarshal receipt: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, latency, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	var receipt *UserOperationReceipt
	if err := json.Unmarshal(rpcResp.Result, &receipt); err != nil {
		return nil, latency, fmt.Errorf("parse receipt: %w", err)
	}
	return receipt, latency, nil
}

func (c *BundlerClient) callRPC(ctx context.Context, req RPCRequest) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.rpcURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() { _ = httpResp.Body.Close() }()
	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", httpResp.StatusCode, string(respBytes))
	}
	return respBytes, nil
}
