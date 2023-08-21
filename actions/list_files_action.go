package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type ListFilesAction struct {
	output io.Writer
}

func NewListFilesAction() IAction {
	return &ListFilesAction{}
}

func (l *ListFilesAction) Exec(output io.Writer, args ...any) error {
	l.output = output

	err := l.listFiles()
	if err != nil {
		return err
	}

	return nil
}

func (l *ListFilesAction) NeedsRoot() bool {
	return false
}

func (l *ListFilesAction) HasArgs() bool {
	return false
}

func (l *ListFilesAction) listFiles() error {
	req, err := http.NewRequest("GET", config.ServerAddress+"/file", nil)
	token, err := configfile.GetValue("token")
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", token)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// TODO: handle different occuring errors
	if resp.StatusCode != 200 {
		return errors.New("something went wrong...")
	}

	files, err := json.ParseFromReader[[]string](resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	fmt.Fprintln(l.output, "Your current synced files:")
	for _, file := range files {
		fmt.Fprintln(l.output, file)
	}

	return nil
}
