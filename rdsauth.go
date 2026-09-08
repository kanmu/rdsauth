package rdsauth

import (
	"context"
	"net"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/ssocreds"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/winebarrel/awsdag"
)

func GetToken(options *Options) (string, error) {
	ctx := context.Background()
	awsopts := []func(*config.LoadOptions) error{}
	if options.Profile != "" {
		awsopts = append(awsopts, config.WithSharedConfigProfile(options.Profile))
	}
	if options.DeviceAuth {
		creds, err := deviceAuth(ctx, options)

		if err != nil {
			return "", err
		}

		awsopts = append(awsopts, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken),
		))
	} else if options.SSORole != "" {
		awsopts = append(awsopts, config.WithSSOProviderOptions(func(o *ssocreds.Options) {
			o.RoleName = options.SSORole
		}))
	}
	cfg, err := config.LoadDefaultConfig(ctx, awsopts...)

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

	token, err := auth.BuildAuthToken(ctx, host+":"+port, cfg.Region, url.User.Username(), cfg.Credentials)

	if err != nil {
		return "", err
	}

	return token, nil
}

func deviceAuth(ctx context.Context, options *Options) (*awsdag.Credentials, error) {
	name := options.Profile

	if name == "" {
		name = os.Getenv("AWS_PROFILE")
	}

	if name == "" {
		name = "default"
	}

	profile, err := awsdag.LoadProfile(ctx, name)

	if err != nil {
		return nil, err
	}

	session, err := awsdag.Auth(ctx, &awsdag.Options{
		StartURL: profile.StartURL,
		Region:   profile.Region,
	})

	if err != nil {
		return nil, err
	}

	accountID, err := session.ChooseAccount(ctx, os.Stdin, os.Stderr, profile.AccountID)

	if err != nil {
		return nil, err
	}

	role := profile.Role

	if options.SSORole != "" {
		role = options.SSORole
	}

	role, err = session.ChooseRole(ctx, os.Stdin, os.Stderr, accountID, role)

	if err != nil {
		return nil, err
	}

	return session.Credentials(ctx, accountID, role)
}
