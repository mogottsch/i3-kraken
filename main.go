package main

import (
	"errors"
	"os"

	"github.com/mogottsch/i3kraken/i3utils"
	"github.com/mogottsch/i3kraken/osutils"
	"go.i3wm.org/i3/v4"
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

func readI3State() (i3.Node, i3.Workspace) {
	focusedNode, err := i3utils.GetFocusedNode()
	check(err)

	activeWorkspace, err := i3utils.GetActiveWorkspace()
	check(err)
	return focusedNode, activeWorkspace
}

func selectNeoVideSession(activeWorkspace i3.Workspace) error {
	_, err := i3utils.LaunchTerminalWithCommand(
		activeWorkspace,
		SESSIONIZER_COMMAND)
	if err != nil {
		return err
	}

	i3utils.WaitForClose()

	sessionizerRes, err := osutils.ReadSessionizerRes()
	if err != nil {
		return err
	}

	if sessionizerRes == "" {
		if err != nil {
			return errors.New("Please select a directory for the session")
		}
	}
	neoVideWmClass := i3utils.GenerateNeoVideWmClassForDir(sessionizerRes)

	_, err = i3utils.MoveToWorkspaceByWmClass(
		neoVideWmClass,
		activeWorkspace)

	// if the workspace was moved it already exists and therefore we are
	// finished
	if err == nil {
		i3utils.DisableFloatingByWmClass(neoVideWmClass)
		return nil
	}

	_, err = i3utils.LaunchNeoVide(
		activeWorkspace,
		sessionizerRes,
		neoVideWmClass)

	return err
}

func main() {
	focusedNode, activeWorkspace := readI3State()

	_, err := i3utils.MoveNodeToScratchpad(focusedNode)
	check(err)

	err = selectNeoVideSession(activeWorkspace)
	if err != nil {
		i3utils.MoveNodeToWorkspace(focusedNode, activeWorkspace)
		panic(err)
	}
	os.Exit(0)
}
