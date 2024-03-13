package server

import (
	"fmt"
	"testing"

	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/identity/v2/tokens"
)

func PointerOf[T any](t T) *T {
	return &t
}

var (
	projectID = ""
	serverID  = "ins-26029e6f-0a29-4ed3-a72a-6bfc53381b40"
)

func NewSC() *client.ServiceClient {
	var (
		identityURL  = "https://iamapis.vngcloud.vn/accounts-api/v2"
		vServerURL   = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2"
		clientID     = ""
		clientSecret = ""
	)

	provider, _ := vngcloud.NewClient(identityURL)
	vngcloud.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vngcloud.NewServiceClient(
		vServerURL,
		provider,
		"vlb")
	return vlb
}

func TestGetInstance(t *testing.T) {
	vserverClient := NewSC()

	opt := NewGetOpts(projectID, serverID)

	resp, err := Get(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}

func TestListInstances(t *testing.T) {
	vserverClient := NewSC()

	opt := NewListOpts(projectID, "cuongdm3-vayne-master-0", 0, 0)

	resp, err := List(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: , haha %d", len(resp))

	fmt.Printf("resp: %+v\n", resp)
}

func TestCreateServer(t *testing.T) {
	vserverClient := NewSC()

	opt := NewCreateOpts(projectID, &CreateOpts{
		Name:             "cuongdm3-uchiha-master-0",
		EncryptionVolume: false,
		ImageId:          "img-a98db89e-ef87-44fd-9567-f1877cbd4cd2",
		FlavorId:         "flav-3929c073-9da9-486f-a96f-9282dbb8d83f",
		NetworkId:        "net-ef3a97fc-3d82-4356-b1d8-9cdbcc1dd80b",
		RootDiskSize:     30,
		RootDiskTypeId:   "vtype-7a7a8610-34f5-11ee-be56-0242ac120002",
		SubnetId:         "sub-55ec620d-943d-4285-bdfd-a3428b37d306",
		UserData:         "echo haha > /tmp/haha.txt",
	})

	resp, err := Create(vserverClient, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("resp: %+v\n", resp)
}