package rdsauth

import (
	"context"
	"net"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
)

func GetToken(url *url.URL) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		return "", err
	}

	host := url.Hostname()

	if !strings.HasSuffix(host, ".rds.amazonaws.com") {
		host, err = net.LookupCNAME(host)

		if err != nil {
			return "", err
		}

		host = strings.TrimSuffix(host, ".")
	}

	port := url.Port()

	if port == "" {
		switch url.Scheme {
		case "mysql":
			port = "3306"
		case "postgres", "postgresql":
			port = "5432"
		}
	}

	token, err := auth.BuildAuthToken(context.Background(), host+":"+port, cfg.Region, url.User.Username(), cfg.Credentials)

	if err != nil {
		return "", err
	}

	return token, nil
}
