package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/rdsauth"
)

var version string

func init() {
	log.SetFlags(0)
}

var cli struct {
	URL     *url.URL `kong:"arg='',required,help='Database URL'"`
	Version kong.VersionFlag
}

func parseArgs() *url.URL {
	parser := kong.Must(&cli, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)
	return cli.URL
}

func main() {
	url := parseArgs()
	token, err := rdsauth.GetToken(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
}
