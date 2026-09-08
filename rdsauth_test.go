package rdsauth

import (
	"testing"
)

func TestRegionFromHost(t *testing.T) {
	tests := []struct {
		host     string
		expected string
	}{
		{"database-1.abcdef012345.ap-northeast-1.rds.amazonaws.com", "ap-northeast-1"},
		{"database-1.cluster-abcdef012345.us-east-1.rds.amazonaws.com", "us-east-1"},
		{"my-proxy.proxy-abcdef012345.eu-west-1.rds.amazonaws.com", "eu-west-1"},
		{"my-db.example.com", ""},
	}

	for _, tt := range tests {
		if actual := regionFromHost(tt.host); actual != tt.expected {
			t.Errorf("regionFromHost(%q) = %q; want %q", tt.host, actual, tt.expected)
		}
	}
}
