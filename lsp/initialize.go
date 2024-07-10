package lsp

type InitializeRequest struct {
  Request
  Params InitializeParams `json:"params"`
}

type InitializeParams struct {
  ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
  Name    string `json:"name"`
  Version string `json:"version"`
}
