package arukas

import (
	"fmt"
	"os"
	"runtime"
	"time"

	API "github.com/arukasio/cli"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/terraform"
)

const (
	JSONTokenParamName   = "ARUKAS_JSON_API_TOKEN"
	JSONSecretParamName  = "ARUKAS_JSON_API_SECRET"
	JSONUrlParamName     = "ARUKAS_JSON_API_URL"
	JSONDebugParamName   = "ARUKAS_DEBUG"
	JSONTimeoutParamName = "ARUKAS_TIMEOUT"
)

type Config struct {
	Token   string
	Secret  string
	URL     string
	Trace   string
	Timeout int
}

func (c *Config) NewClient() (*ArukasClient, error) {

	os.Setenv(JSONTokenParamName, c.Token)
	os.Setenv(JSONSecretParamName, c.Secret)
	os.Setenv(JSONUrlParamName, c.URL)
	os.Setenv(JSONDebugParamName, c.Trace)

	client, err := API.NewClient()
	if err != nil {
		return nil, err
	}
	client.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s (%s/%s)",
		terraform.VersionString(), runtime.GOOS, runtime.GOARCH)

	if logging.IsDebugOrHigher() {
		client.Debug = true
	}

	timeout := time.Duration(0)
	if c.Timeout > 0 {
		timeout = time.Duration(c.Timeout) * time.Second
	}

	return &ArukasClient{
		Client:  client,
		Timeout: timeout,
	}, nil
}

type ArukasClient struct {
	*API.Client
	Timeout time.Duration
}
