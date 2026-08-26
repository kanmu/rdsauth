package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/kanmu/rdsauth"
)

var version string

func init() {
	log.SetFlags(0)
}

func parseArgs() *rdsauth.Options {
	var cli struct {
		rdsauth.Options
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
	token, err := rdsauth.GetToken(options)

	if err != nil {
		log.Fatal(err)
	}

	if options.Export {
		switch options.URL.Scheme {
		case "mysql":
			fmt.Printf("export MYSQL_PWD=%s\n", token)
		case "postgres", "postgresql":
			fmt.Printf("export PGPASSWORD=%s\n", token)
		default:
			log.Fatalf("unimplemented database: %s", options.URL.Scheme)
		}
	} else {
		fmt.Println(token)
	}
}
