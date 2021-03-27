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
)

var args = os.Args

var name = makeName()

func log(message string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, message)
}

func main() {
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

func parseTrackOptions() (internal.TrackOptions, error) {
	set := flag.NewFlagSet("", flag.ContinueOnError)

	var interval time.Duration
	set.DurationVar(&interval, "interval", 10*time.Second, "")

	var minStars int
	set.IntVar(&minStars, "min_stars", 0, "Filter out repositories with a star count below the given value")

	var tokenPath string
	set.StringVar(&tokenPath, "token_path", "", "GitHub personal access token will be read from the given file path")

	set.SetOutput(ioutil.Discard)
	args := args[2:]

	if err := set.Parse(args); err != nil {
		return internal.TrackOptions{}, errors.New("got invalid flags")
	}

	if interval <= 0 {
		return internal.TrackOptions{}, errors.New("got invalid interval")
	}

	if minStars < 0 {
		return internal.TrackOptions{}, errors.New("got invalid minimal stars")
	}

	client, err := internal.GetNewClient(tokenPath)
	if err != nil {
		return internal.TrackOptions{}, err
	}

	return internal.TrackOptions{Interval: interval, MinStars: minStars, Client: client}, nil
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
  %[1]s track [-interval=<interval>] [-min_stars=<minStars>] [-token_path=<tokenPath>]

Commands:
  help  Show usage information
  track Track public GitHub repositories

Options:
  -interval=<interval> Repository update interval, greater than zero [default: 10s]
  -min_stars=<minStars> Filter out repositories with a star count below the given value [default: 0]
  -token_path=<tokenPath> File path to GitHub personal access token for authenticated requests [default: "", meaning gothub will use unauthenticated requests]
`
		fmt.Fprintf(os.Stdout, usage, name)
		return nil

	case "track":
		trackOptions, err := parseTrackOptions()
		if err != nil {
			message := fmt.Sprintf("failed parsing track options: %v", err)
			return usageError{message: message}
		}
		if err := internal.Track(trackOptions); err != nil {
			return fmt.Errorf("failed tracking: %v", err)
		}
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
