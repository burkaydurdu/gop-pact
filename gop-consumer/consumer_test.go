//go:build consumer
// +build consumer

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gofiber/fiber/v2"
	"github.com/pact-foundation/pact-go/dsl"
)

type ClientSuite struct {
	suite.Suite
}

var (
	pact *dsl.Pact
)

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (s *ClientSuite) SetupSuite() {
	pact = &dsl.Pact{
		Host:                     "localhost",
		Consumer:                 "GopClient",
		Provider:                 "GopServer",
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
		LogDir:                   "./pacts/logs",
	}
}

func (s *ClientSuite) TearDownSuite() {
	pact.Teardown()
}

func (s *ClientSuite) Test_IGetNewProductPriceWithSpecifiedDiscountRate() {
	const productID = 1
	const discountRate = 30

	pact.
		AddInteraction().
		Given("i get new product price with specified discount rate").
		UponReceiving("A request for campaign").
		WithRequest(
			dsl.Request{
				Method: http.MethodGet,
				Path:   dsl.String(fmt.Sprintf("/products/%d/discount", productID)),
				Query: map[string]dsl.Matcher{
					"rate": dsl.Like(strconv.Itoa(discountRate)),
				},
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusOK,
				Headers: dsl.MapMatcher{
					fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
				},
				Body: dsl.StructMatcher{
					"id":    dsl.Integer(),
					"price": dsl.Decimal(),
					"name":  dsl.Like(""),
				},
			},
		)

	var test = func() error {
		return makeRequest(pact.Server.Port, productID, discountRate)
	}

	err := pact.Verify(test)

	if err != nil {
		s.Error(err)
	}
}
