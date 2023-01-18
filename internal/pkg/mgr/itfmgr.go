package mgr

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

type InterfaceManager struct {
	AzureManager
	client *armnetwork.InterfacesClient
}

func NewInterfaceManager(subsId string) (*InterfaceManager, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	return &InterfaceManager{AzureManager{cred: cred, subsId: subsId}, nil}, nil
}

func (mgr *InterfaceManager) ensureClient() (bool, error) {
	if mgr.client == nil {
		client, err := armnetwork.NewInterfacesClient(mgr.subsId, mgr.cred, nil)
		if err != nil {
			return false, err
		}

		mgr.client = client
	}

	return true, nil
}

func (mgr *InterfaceManager) Get(ctx context.Context, resGrNm, resNm string) (*armnetwork.Interface, error) {
	_, err := mgr.ensureClient()
	if err != nil {
		return nil, err
	}

	resp, err := mgr.client.Get(ctx, resGrNm, resNm, nil)
	if err != nil {
		return nil, err
	}

	return &resp.Interface, nil
}

func (mgr *InterfaceManager) Update(ctx context.Context, resGrNm, resNm string, itf *armnetwork.Interface) (bool, error) {
	_, err := mgr.ensureClient()
	if err != nil {
		return false, err
	}

	poller, err := mgr.client.BeginCreateOrUpdate(ctx, resGrNm, resNm, *itf, nil)
	if err != nil {
		return false, err
	}

	resp, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		return false, err
	}

	itf = &resp.Interface

	return true, nil
}
