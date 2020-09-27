package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	bitly "github.com/yuzuy/bitly-go"
)

var (
	tFlag                = flag.String("t", "", `input your access token. if it not specified, read the token set by "bitly config token"`)
	vFlag                = flag.Bool("v", false, "print the result in detail")
	shortenDomainFlag    = flag.String("sdomain", "", "[shorten] input your branded short domain")
	shortenGroupGUIDFlag = flag.String("sgguid", "", "[shorten] input your group guid")

	configFilePath = os.Getenv("HOME") + "/.bitly"
)

func main() {
	if err := run(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	switch flag.Arg(0) {
	case "shorten":
		return shorten(flag.Arg(1))
	case "config":
		if flag.Arg(1) == "token" {
			return setDefaultToken(flag.Arg(2))
		}
	}

	return fmt.Errorf("commnd %q is not found", strings.Join(flag.Args(), " "))
}

func shorten(url string) error {
	if url == "" {
		return errors.New("input an url to shorten")
	}

	var client bitly.Client
	if *tFlag != "" {
		client = bitly.New(*tFlag)
	} else {
		f, err := os.Open(configFilePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return errors.New(`set your access token with "bitly config token"`)
			}
			return err
		}
		defer f.Close()
		token, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		client = bitly.New(string(token))
	}

	config := bitly.ShortenConfig{
		Domain:    *shortenDomainFlag,
		GroupGUID: *shortenGroupGUIDFlag,
	}
	resp, err := client.Shorten(url, config)
	if err != nil {
		return err
	}

	if *vFlag {
		v, _ := json.MarshalIndent(resp, "", "\t")
		fmt.Println(string(v))
		return nil
	}
	fmt.Println(resp.Link)

	return nil
}

func setDefaultToken(token string) error {
	f, err := os.Open(configFilePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(configFilePath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	if err := f.Truncate(0); err != nil {
		return err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	_, err = f.WriteString(token)

	return err
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("bitly: ")
}
