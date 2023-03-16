package client

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/Goose97/tiny-http-server/pkg/common"
)

var supportedSchemes = [2]string{"http", "https"}

var help = flag.Bool("help", false, "Show help")
var headerFlags headerFlagsArray
var methodFlag = "GET"

// Custom multi-values flag
type headerFlagsArray []common.ClientRequestHeader

func (i *headerFlagsArray) String() string {
	return ""
}

func (i *headerFlagsArray) Set(value string) error {
	parts := strings.Split(value, ":")

	if len(parts) == 2 {
		*i = append(*i, common.ClientRequestHeader{
			Name:  strings.Trim(parts[0], " "),
			Value: strings.Trim(parts[1], " "),
		})
	}

	return nil
}

func Parse() (common.ClientRequest, error) {
	urlArgument := os.Args[len(os.Args)-1]

	parsedUrl, err := parseUrl(urlArgument)

	if err != nil {
		return common.ClientRequest{}, err
	}

	// Bind the flag
	flag.Var(&headerFlags, "H", "Request headers. You can specify multiple headers, each header must conform to \"<header_name>: <header_value>\" format")
	flag.StringVar(&methodFlag, "X", "GET", "Request method. Supported methods: GET, HEAD, POST, PUT, PATCH, DELETE")

	// Parse the flag
	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	return common.ClientRequest{
		Url:     *parsedUrl,
		Headers: headerFlags,
		Method:  methodFlag,
	}, nil
}

func parseUrl(rawUrl string) (*url.URL, error) {
	var hasValidScheme bool
	for _, scheme := range supportedSchemes {
		if strings.HasPrefix(rawUrl, scheme) {
			hasValidScheme = true
		}
	}

	if !hasValidScheme {
		return nil, fmt.Errorf("URL must start with one of these schemes %v", supportedSchemes)
	}

	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		return nil, err
	}

	// Normalize path
	if parsedUrl.Path == "" {
		parsedUrl.Path = "/"
	}

	return parsedUrl, nil
}
