package actions

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type AddFileAction struct {
	output io.Writer
}

func NewAddFileAction() IAction {
	return &AddFileAction{}
}

func (a *AddFileAction) Exec(output io.Writer, args ...any) error {
	a.output = output

	err := a.addFile(args[0].(string))
	if err != nil {
		return err
	}

	return nil
}

func (a *AddFileAction) NeedsRoot() bool {
	return false
}

func (a *AddFileAction) HasArgs() bool {
	return true
}

func (a *AddFileAction) addFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return a.addDirectory(filePath)
	}

	if stat.Size() > 256*1024 {
		return errors.New("file is larger than 256KiB...")
	}

	fmt.Fprintf(a.output, "Adding %s to your synced files...\n", filePath)

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return err
	}

	encContent := base64.StdEncoding.EncodeToString(content)

	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]string{
		"path":    filePath,
		"content": encContent,
	})

	req, err := http.NewRequest("POST", config.ServerAddress+"/file", reqBody)
	req.Header.Add("Content-Type", "application/json")
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

	if resp.StatusCode != 200 {
		respBody, _ := json.ParseFromReader[json.Json](resp.Body)
		resp.Body.Close()
		return errors.New(respBody["error"].(string))
	}

	fmt.Fprintln(a.output, "Done, you can check the synced files list by using `dotsync -list`")

	return nil
}

func (a *AddFileAction) addDirectory(dirPath string) error {
	uploadList := make([]struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}, 0)

	fmt.Fprintf(a.output, "reading files inside %s...\n", dirPath)
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if info.Size() > 256*1024 {
			return errors.New(fmt.Sprintf("the file %s is larger than 256KiB...", path))
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		uploadList = append(uploadList, struct {
			Path    string `json:"path"`
			Content string `json:"content"`
		}{
			Path:    path,
			Content: string(content),
		})

		return nil
	})
	if err != nil {
		return err
	}
	fmt.Fprintln(a.output, "done 👍")

	fmt.Fprintln(a.output, "uploading your new files...")
	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]any{
		"files": uploadList,
	})

	req, err := http.NewRequest("POST", config.ServerAddress+"/file/add-directory", reqBody)
	req.Header.Add("Content-Type", "application/json")
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

	if resp.StatusCode != 200 {
		respBody, _ := json.ParseFromReader[json.Json](resp.Body)
		resp.Body.Close()
		return errors.New(respBody["error"].(string))
	}
	fmt.Fprintln(a.output, "done 👍")

	return nil
}
