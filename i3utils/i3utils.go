package i3utils

import (
	"fmt"

	"go.i3wm.org/i3/v4"
)

const NEOVIDE_WM_CLASS = "i3-kraken-neovide"

func GenerateNeoVideWmClassForDir(directory string) string {
	return NEOVIDE_WM_CLASS + "-" + directory
}

func MoveToWorkspace(selector string, workspace i3.Workspace) ([]i3.CommandResult, error) {
	command := fmt.Sprintf("[%s] move to workspace %v", selector, workspace.Name)
	fmt.Println(command)
	return i3.RunCommand(command)
}
func MoveToWorkspaceByWmClass(wmClass string, workspace i3.Workspace) ([]i3.CommandResult, error) {
	selector := fmt.Sprintf("class=\"%s\"", wmClass)
	return MoveToWorkspace(selector, workspace)
}

func GetActiveWorkspace() (i3.Workspace, error) {
	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		return i3.Workspace{}, err
	}
	for _, workspace := range workspaces {
		if workspace.Focused {
			return workspace, nil
		}
	}
	return i3.Workspace{}, fmt.Errorf("No active workspace found")
}
