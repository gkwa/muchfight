package muchfight

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	LogFormat string `long:"log-format" choice:"text" choice:"json" default:"text" description:"Log format"`
	Verbose   []bool `short:"v" long:"verbose" description:"Show verbose debug information, each -v bumps log level"`
	logLevel  slog.Level
}

func Execute() int {
	if err := parseFlags(); err != nil {
		return 1
	}

	if err := setLogLevel(); err != nil {
		return 1
	}

	if err := setupLogger(); err != nil {
		return 1
	}

	if err := run(); err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}

	return 0
}

func parseFlags() error {
	_, err := flags.Parse(&opts)
	return err
}

func run() error {
	mdfindCmdSlice := []string{
		"mdfind",
		`kMDItemFSContentChangeDate >= $time.now(-7200) && kMDItemFSName == '*.go'`,
		"-onlyin", "/Users/mtm/pdev",
	}
	xargsCmdSlice := []string{
		"xargs",
		"-d", "\n",
		"-a", "-",
		"rg", "-l", "Command",
	}

	mdfindCmd := exec.Command(mdfindCmdSlice[0], mdfindCmdSlice[1:]...)
	xargsCmd := exec.Command(xargsCmdSlice[0], xargsCmdSlice[1:]...)

	pipe, err := mdfindCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating pipe: %w", err)
	}
	xargsCmd.Stdin = pipe

	xargsCmd.Stdout = os.Stdout
	xargsCmd.Stderr = os.Stderr

	err = mdfindCmd.Start()
	if err != nil {
		return fmt.Errorf("error starting mdfind: %w", err)
	}

	err = xargsCmd.Start()
	if err != nil {
		return fmt.Errorf("error starting xargs: %w", err)
	}

	err = mdfindCmd.Wait()
	if err != nil {
		return fmt.Errorf("error waiting for mdfind: %w", err)
	}

	err = xargsCmd.Wait()
	if err != nil {
		return fmt.Errorf("error waiting for xargs: %w", err)
	}

	return nil
}
