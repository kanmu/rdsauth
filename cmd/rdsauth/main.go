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

type Options struct {
	URL    *url.URL `kong:"arg='',required,help='Database URL'"`
	Export bool     `kong:"short='e',help='Output as environment variable.'"`
}

func parseArgs() *Options {
	var cli struct {
		Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&cli, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &cli.Options
}

func main() {
	options := parseArgs()
	token, err := rdsauth.GetToken(options.URL)

	if err != nil {
		log.Fatal(err)
	}

	if options.Export {
		switch options.URL.Scheme {
		case "mysql":
			fmt.Printf("export MYSQL_PWD='%s'\n", token)
		case "postgres", "postgresql":
			fmt.Printf("export PGPASSWORD='%s'\n", token)
		default:
			log.Fatalf("unimplemented database: %s", options.URL.Scheme)
		}
	} else {
		fmt.Println(token)
	}
}
