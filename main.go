package main

import (
	"fmt"
	"os"

	"github.com/mbaraa/dotsync/actions"
)

const (
	usageStr = `Usage of Dotsync:
-login string:email
	login into your account, or create an account using a valid email.
-list
	lists the files that are currently synced.
-add string:file_path
	adds a file to the sync list.
-remove string:file_path
	removes a file from the sync list.
-download
	syncs local files with the server's version.
-upload
	syncs server's files with the local version.`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usageStr)
		os.Exit(1)
	}

	action := actions.GetAction(actions.ActionType(os.Args[1]))
	if action == nil {
		fmt.Println("Ivalid argument")
		fmt.Println()
		fmt.Println(usageStr)
		os.Exit(1)
	}

	if action.NeedsRoot() {
		fmt.Println("This action requires superuser to run")
		os.Exit(1)
	}

	var arg string
	if len(os.Args) > 2 {
		arg = os.Args[2]
	}

	if action.HasArgs() && len(arg) == 0 {
		fmt.Println("this action needs an argument")
		fmt.Println(usageStr)
		os.Exit(1)
	}

	err := action.Exec(os.Stdout, arg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
