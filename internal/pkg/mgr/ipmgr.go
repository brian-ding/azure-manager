package mgr

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

type IPAddressManager struct {
	AzureManager
	client *armnetwork.PublicIPAddressesClient
}

func NewIPAddressManager(subsId string) (*IPAddressManager, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	return &IPAddressManager{AzureManager{cred: cred, subsId: subsId}, nil}, nil
}

func (mgr *IPAddressManager) ensureClient() (bool, error) {
	if mgr.client == nil {
		client, err := armnetwork.NewPublicIPAddressesClient(mgr.subsId, mgr.cred, nil)
		if err != nil {
			return false, err
		}

		mgr.client = client
	}

	return true, nil
}

func (mgr *IPAddressManager) GetIPAddress(ctx context.Context, resGrNm, resNm string) (*armnetwork.PublicIPAddress, error) {
	_, err := mgr.ensureClient()
	if err != nil {
		return nil, err
	}

	resp, err := mgr.client.Get(ctx, resGrNm, resNm, nil)
	if err != nil {
		return nil, err
	}

	return &resp.PublicIPAddress, nil
}
