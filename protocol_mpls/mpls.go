package protocol

import (
	"github.com/mhersson/glsp"
	protocol316 "github.com/mhersson/glsp/protocol_3_16"
)

const (
	MethodMplsEditorDidChangeFocus = protocol316.Method("mpls/editorDidChangeFocus")
)

type (
	MplsEditorDidChangeFocusFunc func(context *glsp.Context, params *EditorDidChangeFocusParams) error
)

type EditorDidChangeFocusParams struct {
	/**
	 * The text document's URI.
	 */
	URI protocol316.DocumentUri `json:"uri"`
}
