package main

import "encoding/json"

type UserOperation struct {
	Sender               string          `json:"sender"`
	Nonce                string          `json:"nonce"`
	InitCode             string          `json:"initCode"`
	CallData             string          `json:"callData"`
	CallGasLimit         string          `json:"callGasLimit"`
	VerificationGasLimit string          `json:"verificationGasLimit"`
	PreVerificationGas   string          `json:"preVerificationGas"`
	MaxFeePerGas         string          `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string          `json:"maxPriorityFeePerGas"`
	PaymasterAndData     string          `json:"paymasterAndData"`
	Signature            string          `json:"signature"`
	EntryPoint           json.RawMessage `json:"-"`
}

type UserOperationReceipt struct {
	UserOpHash    string `json:"userOpHash"`
	Sender        string `json:"sender"`
	Nonce         string `json:"nonce"`
	ActualGasUsed string `json:"actualGasUsed"`
	Success       bool   `json:"success""
	Logs          []Log  `json:"logs"`
}

type Log struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      interface{}   `json:"id"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
	ID      interface{}     `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}
