package mgr

import "github.com/Azure/azure-sdk-for-go/sdk/azidentity"

type AzureManager struct {
	cred   *azidentity.DefaultAzureCredential
	subsId string
}
