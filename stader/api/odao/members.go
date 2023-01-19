package odao

import (
	"github.com/rocket-pool/rocketpool-go/dao/trustednode"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getMembers(c *cli.Context) (*api.TNDAOMembersResponse, error) {

	// Get services
	if err := services.RequireRocketStorage(c); err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.TNDAOMembersResponse{}

	// Get members
	members, err := trustednode.GetMembers(rp, nil)
	if err != nil {
		return nil, err
	}
	response.Members = members

	// Return response
	return &response, nil

}
