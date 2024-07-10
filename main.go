package main

import (
	"bufio"
	"educationallsp/analysis"
	"educationallsp/lsp"
	"educationallsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/vinibispo/projects/golang/educational-lsp/educationallsp.txt")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("error decoding message: %v", err)
			continue
		}
		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest

		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("error unmarshalling initialize request: %v", err)
		}

		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		response := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(response)
		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent the reply")
	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification

		if err := json.Unmarshal(contents, &notification); err != nil {
			logger.Printf("error unmarshalling didOpenTextDocument notification: %v", err)
			return
		}

		logger.Printf("Opened document: %s", notification.Params.TextDocument.URI)
		state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)
	case "textDocument/didChange":
		var notification lsp.TextDocumentDidChangeNotification

		if err := json.Unmarshal(contents, &notification); err != nil {
			logger.Printf("error unmarshalling didChangeTextDocument notification: %v", err)
			return
		}

		logger.Printf("Changed document: %s", notification.Params.TextDocument.URI)

		for _, change := range notification.Params.ContentChanges {
			state.UpdateDocument(notification.Params.TextDocument.URI, change.Text)
		}
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("hey, you didn't provide a valid filename")
	}

	return log.New(logfile, "[educationallsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
