package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/antoinevanluchem/elimity-backend-exercise/internal"
	// "github.com/elimity-com/backend-intern-exercise/internal"
)

var args = os.Args

var tokenFile string = ""

var name = makeName()

func log(message string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, message)
}

func main() {
	fmt.Println("Dit is main")
	if err := run(); err != nil {
		message := err.Error()
		log(message)
		if _, ok := err.(usageError); ok {
			message := fmt.Sprintf("run '%s help' for usage information", name)
			log(message)
		}
	}
}

func makeName() string {
	path := args[0]
	return filepath.Base(path)
}

func parseInterval() (time.Duration, int, error) {
	set := flag.NewFlagSet("", flag.ExitOnError)
	var interval time.Duration
	set.DurationVar(&interval, "interval", 10*time.Second, "")

	var minStars int
	set.IntVar(&minStars, "min_stars", 0, "Filter out repositories with a star count below the given value")

	fmt.Println("minStars is:", minStars)

	var tF string
	set.StringVar(&tF, "token_file", "", "GitHub personal access token will be read from the given file path")
	if tF != "" {
		tokenFile = tF
	}

	set.SetOutput(ioutil.Discard)
	args := args[2:]

	if err := set.Parse(args); err != nil {
		return 0, 0, errors.New("got invalid flags")
	}

	if interval <= 0 {
		return 0, 0, errors.New("got invalid interval")
	}

	if minStars < 0 {
		return 0, 0, errors.New("got invalid minimal stars")
	}

	return interval, minStars, nil
}

func run() error {
	if nbArgs := len(args); nbArgs < 2 {
		return usageError{message: "missing command"}
	}
	switch args[1] {
	case "help":
		const usage = `
Simple CLI for tracking public GitHub repositories.

Usage:
  %[1]s help
  %[1]s track [-interval=<interval>] [-min_stars=<minStars>] [-token_file=<tokenFile>]

Commands:
  help  Show usage information
  track Track public GitHub repositories

Options:
  -interval=<interval> Repository update interval, greater than zero [default: 10s]
  -min_stars=<minStars> Filter out repositories with a star count below the given value
  -token_file=<tokenFile> File path to GitHub personal access token
`
		fmt.Fprintf(os.Stdout, usage, name)
		return nil

	case "track":
		interval, minStars, err := parseInterval()
		if err != nil {
			message := fmt.Sprintf("failed parsing interval: %v", err)
			return usageError{message: message}
		}
		if err := internal.Track(interval, minStars, tokenFile); err != nil {
			return fmt.Errorf("failed tracking: %v", err)
		}
		return nil

	case "token_file":
		tF, err := ioutil.ReadFile(args[2])
		if err != nil {
			message := fmt.Sprintf("failed reading from the path: %v", err)
			return usageError{message: message}
		}
		tokenFile = string(tF)
		return nil

	default:
		return usageError{message: "got invalid command"}
	}
}

type usageError struct {
	message string
}

func (e usageError) Error() string {
	return e.message
}
