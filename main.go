package main

import (
	"fmt"

	"github.com/mogottsch/i3kraken/i3utils"
	"github.com/mogottsch/i3kraken/osutils"
)

const NEOVIDE_WM_CLASS = "i3-kraken-neovide"
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
	fmt.Println(sessionizerRes)

	neoVideWmClass := i3utils.GenerateNeoVideWmClassForDir(sessionizerRes)
	_, err = i3utils.LaunchNeoVide(
		activeWorkspace,
		sessionizerRes,
		neoVideWmClass)
	check(err)
}
