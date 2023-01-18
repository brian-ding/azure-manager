package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/refresh", func(c *gin.Context) {
		c.JSON(200, refresh())
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func refresh() gin.H {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println(err.Error())

		return gin.H{"message": err.Error()}
	}

	client, err := armnetwork.NewPublicIPAddressesClient("d2997c3a-85a3-4d1f-ba7c-e0aa9fea5d1e", cred, nil)
	if err != nil {
		fmt.Println(err.Error())

		return gin.H{"message": err.Error()}
	}

	resp, err := client.Get(context.Background(), "vpn", "vpn-us566", nil)
	if err != nil {
		fmt.Println(err.Error())

		return gin.H{"message": err.Error()}
	}

	fmt.Println(resp)
	// client.BeginCreateOrUpdate(context.Background(), "vpn", "vpn-us-ip-backup")

	return gin.H{"message": "pong"}
}
