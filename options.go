package rdsauth

import (
	"net/url"
)

type Options struct {
	URL        *url.URL `kong:"arg='',required,help='Database URL'"`
	Profile    string   `kong:"help='AWS credentials profile name.'"`
	SSORole    string   `kong:"help='Override sso_role_name for every profile, e.g. ReadOnlyAccess.'"`
	DeviceAuth bool     `kong:"help='Sign in to IAM Identity Center with the OAuth 2.0 device authorization grant (no browser needed).'"`
}
