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

func TestDecode(t *testing.T) {
  expected := 16
  actual, err := rpc.DecodeMessage([]byte("Content-Length: 16\r\n\r\n{\"testing\":true}"))

  if err != nil {
    t.Fatalf("Unexpected error: %s", err)
  }

  if actual != expected {
    t.Fatalf("Expected %d, got %d", expected, actual)
  }
}
