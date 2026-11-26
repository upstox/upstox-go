//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox/api"
)

func TestGetOptionContracts(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOptionsApi(c).GetOptionContracts(ctx(t), niftyKey, nil)
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("no option contracts returned")
	}
}

func TestGetOptionContractsByExpiry(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOptionsApi(c).GetOptionContracts(ctx(t), niftyKey, &api.OptionsApiGetOptionContractsOpts{
		ExpiryDate: optional.NewString(niftyExpiry(t, c)),
	})
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("no option contracts for expiry")
	}
}

func TestGetPutCallOptionChain(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOptionsApi(c).GetPutCallOptionChain(ctx(t), niftyKey, niftyExpiry(t, c))
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("empty option chain")
	}
}
