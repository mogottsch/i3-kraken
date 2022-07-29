package main

import (
	"fmt"
	"os"

	"github.com/mogottsch/i3kraken/i3utils"
	"github.com/mogottsch/i3kraken/osutils"
)

const SESSIONIZER_COMMAND = "find " +
	"~/dev " +
	"~/dev/privat " +
	"~/dev/uni " +
	"~/dev/zendri " +
	"~ " +
	"~/OneDrive/Uni/Master/1.\\ Semester " +
	"~/go/src/moritz " +
	"~/Sync/ " +
	"-mindepth 1 -maxdepth 1 -type d -not -name .ansible " +
	"| fzf" +
	" > " + osutils.SESSIONIZER_RES_FILE

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	activeWorkspace, err := i3utils.GetActiveWorkspace()
	check(err)
	_, err = i3utils.LaunchTerminalWithCommand(
		activeWorkspace,
		SESSIONIZER_COMMAND)
	check(err)

	i3utils.WaitForClose()

	sessionizerRes, err := osutils.ReadSessionizerRes()
	check(err)

	if sessionizerRes == "" {
		fmt.Println("Please select a directory for the session")
		os.Exit(1)
	}
	neoVideWmClass := i3utils.GenerateNeoVideWmClassForDir(sessionizerRes)

	_, err = i3utils.MoveToWorkspaceByWmClass(
		neoVideWmClass,
		activeWorkspace)
	// if the workspace was moved it already exists and therefore we are
	// finished
	if err == nil {
		os.Exit(0)
	}

	_, err = i3utils.LaunchNeoVide(
		activeWorkspace,
		sessionizerRes,
		neoVideWmClass)
	check(err)
}
