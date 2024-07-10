package rpc_test

import (
	"educationallsp/rpc"
	"testing"
)

type EncodingExample struct {
  Testing bool `json:"testing"`
}

func TestEncode(t *testing.T) {
  expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
  actual := rpc.EncodeMessage(EncodingExample{Testing: true})

  if actual != expected {
    t.Fatalf("Expected %s, got %s", expected, actual)
  }
}
