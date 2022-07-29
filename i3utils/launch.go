package i3utils

import (
	"fmt"

	"go.i3wm.org/i3/v4"
)

func LaunchNeoVide(workspace i3.Workspace, directory string, wmClass string) ([]i3.CommandResult, error) {
	command := fmt.Sprintf(
		"workspace %v; exec neovide --x11-wm-class %v %s",
		workspace.Name,
		wmClass,
		directory)
	return i3.RunCommand(command)
}

func LaunchTerminalWithCommand(workspace i3.Workspace, terminalCommand string) ([]i3.CommandResult, error) {
	command := fmt.Sprintf(
		"workspace %v; exec urxvt -e bash -c '%s'",
		workspace.Name,
		terminalCommand)
	return i3.RunCommand(command)
}
