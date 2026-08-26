package rdsauth

import (
	"context"
	"net"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/ssocreds"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
)

func GetToken(options *Options) (string, error) {
	awsopts := []func(*config.LoadOptions) error{}
	if options.Profile != "" {
		awsopts = append(awsopts, config.WithSharedConfigProfile(options.Profile))
	}
	if options.SSORole != "" {
		awsopts = append(awsopts, config.WithSSOProviderOptions(func(o *ssocreds.Options) {
			o.RoleName = options.SSORole
		}))
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), awsopts...)

	if err != nil {
		return "", err
	}

	url := options.URL
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
