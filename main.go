package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "0.1.3"

func main() {
	fmt.Println("hello")

	if err := confirmAndSelfUpdate(version); err != nil {
		log.Fatal(err)
	}
}

func confirmAndSelfUpdate(current string) error {
	latest, found, err := selfupdate.DetectLatest("micnncim/selfupdate-playground")
	if err != nil {
		return err
	}

	v := semver.MustParse(current)
	if !found || latest.Version.LTE(v) {
		return nil
	}

	fmt.Printf("current version: %s, latest version: %s\n", current, latest.Version)
	fmt.Print("Do you want to update to", latest.Version, "? (y/n): ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil || (input != "y\n" && input != "n\n") {
		return errors.New("invalid input")
	}
	if input == "n\n" {
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		return err
	}

	log.Println("successfully updated to version", latest.Version)
	return nil
}
