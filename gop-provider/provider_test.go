//go:build provider
// +build provider

package main

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/pact-foundation/pact-go/types"

	"github.com/pact-foundation/pact-go/utils"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/suite"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string // Basic authentication
	BrokerPassword  string // Basic authentication
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ConsumerTag     string // dev, staging, prod
	ProviderVersion string
}

func (s *Settings) getPactURL(useLocal bool) string {
	// Local pact file or remote based urls (Pact Broker)
	var pactURL string

	if useLocal {
		pactURL = "../gop-consumer/pacts/gopclient-gopserver.json"
		return pactURL
	}

	if s.ConsumerVersion == "" {
		pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/latest/master.json", s.BrokerBaseURL, s.ProviderName, s.ConsumerName)
	} else {
		pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json", s.BrokerBaseURL, s.ProviderName, s.ConsumerName, s.ConsumerVersion)
	}

	return pactURL
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "GopServer"
	s.ConsumerName = "GopClient"
	s.BrokerBaseURL = "http://localhost"
	s.ConsumerTag = "master"
	s.ProviderVersion = "2.5.2"
	s.ConsumerVersion = ""
}

type ServerSuite struct {
	suite.Suite
}

var (
	pact *dsl.Pact
)

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) SetupSuite() {
	settings := Settings{}
	settings.create()

	pact = &dsl.Pact{
		Host:                     settings.Host,
		Consumer:                 settings.ConsumerName,
		Provider:                 settings.ProviderName,
		DisableToolValidityCheck: true,
	}
}

func (s *ServerSuite) TearDownSuite() {
	pact.Teardown()
}

func (s *ServerSuite) TestProvider() {
	port, _ := utils.GetFreePort()

	go startServer(port)

	settings := Settings{}
	settings.create()

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://%s:%d", settings.Host, port),
		ProviderVersion: settings.ProviderVersion,
		BrokerUsername:  settings.BrokerUsername,
		BrokerURL:       settings.BrokerBaseURL,
		BrokerPassword:  settings.BrokerPassword,
		//Tags:            []string{settings.ConsumerTag},
		PactURLs: []string{settings.getPactURL(false)},
		StateHandlers: map[string]types.StateHandler{
			"i get new product price with specified discount rate": func() error {
				return nil
			},
		},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	verifyResponses, err := pact.VerifyProvider(s.T(), verifyRequest)

	if err != nil {
		s.Error(err)
	}

	pp.Println(len(verifyResponses), "pact tests run")
}
