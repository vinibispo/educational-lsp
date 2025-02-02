package analysis

import (
	"educationallsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "Common Sense",
				Message:  "Please make sure we use good language in this video",
			})
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]
	return lsp.HoverResponse{
		Response: lsp.Response{
			ID:  &id,
			RPC: "2.0",
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Character: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			ID:  &id,
			RPC: "2.0",
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) CodeAction(id int, uri string, range_ lsp.Range, context lsp.CodeActionContext) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	return lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			ID:  &id,
			RPC: "2.0",
		},
		Result: actions,
	}
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "Fun to watch videos. Don't forget to like & subscribe to streamers using it:)",
		},
	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}

func LineRange(row, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      row,
			Character: start,
		},
		End: lsp.Position{
			Line:      row,
			Character: end,
		},
	}
}
