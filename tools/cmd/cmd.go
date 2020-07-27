package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/Grimkey/cmdban/todolist"
	"github.com/spf13/cobra"
)

func Execute(backend todolist.Backend) error {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(addListCmd(backend), boardCmd(backend))
	rootCmd.Execute()

	return nil
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
