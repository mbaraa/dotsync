package actions

import "io"

type IAction interface {
	// Exec executes the selected action, and prints action's output on the given io.Writer
	// or exits with an error
	Exec(output io.Writer, args ...any) error

	// NeedsRoot return true if the action needs root to be executed
	NeedsRoot() bool

	// HasArgs returns true if the action requires arguments to run
	HasArgs() bool
}

type ActionType string

const (
	LoginActionType         ActionType = "-login"
	DeleteUserActionType    ActionType = "-delete-user"
	AddFileActionType       ActionType = "-add"
	RemoveFileActionType    ActionType = "-remove"
	ListFilesActionType     ActionType = "-list"
	DownloadFilesActionType ActionType = "-download"
	UploadFilesActionType   ActionType = "-upload"
)

func GetAction(at ActionType) IAction {
	switch at {
	case LoginActionType:
		return NewLoginAction()
	case DeleteUserActionType:
		return NewDeleteUserAction()
	case AddFileActionType:
		return NewAddFileAction()
	case RemoveFileActionType:
		return NewRemoveFileAction()
	case ListFilesActionType:
		return NewListFilesAction()
	case DownloadFilesActionType:
		return NewDownloadFilesAction()
	case UploadFilesActionType:
		return NewUploadFilesAction()
	default:
		return nil
	}
}
