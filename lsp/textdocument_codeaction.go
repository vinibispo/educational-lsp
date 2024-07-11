package lsp

// CodeActionRequest is sent from the client to the server to compute commands for a given text document and range.
type CodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

// TextDocumentCodeActionParams is the parameters for the CodeActionRequest.
type TextDocumentCodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"context"`
}

type TextDocumentCodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

// CodeActionContext contains additional diagnostic information about the context in which a code action is run.
type CodeActionContext struct {
}

type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
}

// Command represents a reference to a command. It is used by code actions and code lens.
type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
