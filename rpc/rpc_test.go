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
	expectedContentLength := 15
	expectedMethod := "hi"
	method, actual, err := rpc.DecodeMessage([]byte("Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"))
	contentLength := len(actual)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if contentLength != expectedContentLength {
		t.Fatalf("Expected %d, got %d", expectedContentLength, contentLength)
	}

	if method != expectedMethod {
		t.Fatalf("Expected %s, got %s", expectedMethod, method)
	}
}
