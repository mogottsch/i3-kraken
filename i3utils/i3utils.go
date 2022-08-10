package i3utils

import (
	"fmt"
	"hash/crc32"

	"go.i3wm.org/i3/v4"
)

const NEOVIDE_WM_CLASS = "i3_kraken_neovide"

func GenerateNeoVideWmClassForDir(directory string) string {
	hashDecimal := crc32.ChecksumIEEE([]byte(directory))
	hashHex := fmt.Sprintf("%x", hashDecimal)

	return NEOVIDE_WM_CLASS + "_" + hashHex
}

func MoveToWorkspace(selector string, workspace i3.Workspace) ([]i3.CommandResult, error) {
	command := fmt.Sprintf("[%s] move to workspace %v", selector, workspace.Name)
	return i3.RunCommand(command)
}
func MoveToWorkspaceByWmClass(wmClass string, workspace i3.Workspace) ([]i3.CommandResult, error) {
	selector := fmt.Sprintf("class=\"%s\"", wmClass)
	return MoveToWorkspace(selector, workspace)
}

func MoveToScratchpad(selector string) ([]i3.CommandResult, error) {
	command := fmt.Sprintf("[%s] move scratchpad", selector)
	return i3.RunCommand(command)
}

func MoveNodeToScratchpad(node i3.Node) ([]i3.CommandResult, error) {
	selector := fmt.Sprintf("con_id=%d", node.ID)
	return MoveToScratchpad(selector)
}

func MoveNodeToWorkspace(node i3.Node, workspace i3.Workspace) ([]i3.CommandResult, error) {
	selector := fmt.Sprintf("con_id=%d", node.ID)
	return MoveToWorkspace(selector, workspace)
}

func DisableFloating(selector string) ([]i3.CommandResult, error) {
	command := fmt.Sprintf("[%s] floating disable", selector)
	return i3.RunCommand(command)
}

func DisableFloatingByWmClass(wmClass string) ([]i3.CommandResult, error) {
	selector := fmt.Sprintf("class=\"%s\"", wmClass)
	return DisableFloating(selector)
}
func DisableFloatingByNode(node i3.Node) ([]i3.CommandResult, error) {
	selector := fmt.Sprintf("con_id=%d", node.ID)
	return DisableFloating(selector)
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

func GetFocusedNode() (i3.Node, error) {
	tree, err := i3.GetTree()
	if err != nil {
		return i3.Node{}, err
	}
	root := tree.Root
	focusedNode := root.FindChild(
		func(node *i3.Node) bool {
			return node.Focused
		},
	)
	if focusedNode == nil {
		return i3.Node{}, fmt.Errorf("No focused node found")
	}
	return *focusedNode, nil

}
