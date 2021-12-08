package sumologic

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"os"

	sumologic "github.com/saurabh-agarwals/sumologic-api-client-go/openapi"
)

// UserGenerator ...
type UserGenerator struct {
	SumologicService
}

func (g *UserGenerator) createResources(users []sumologic.UserModel) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, user := range users {
		resources = append(resources, g.createResource(user.GetId()))
	}
	return resources
}

func (g *UserGenerator) createResource(userID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		userID,
		fmt.Sprintf("user_%s", userID),
		"sumologic_user",
		"sumologic",
		[]string{},
	)
}

func (g *UserGenerator) InitResources() error {
	var users []sumologic.UserModel
	cf := sumologic.NewConfiguration()
	sumologicClient := sumologic.NewAPIClient(cf)
	auth := context.WithValue(context.Background(), sumologic.ContextBasicAuth, sumologic.BasicAuth{
		UserName: os.Getenv("SUMOLOGIC_ACCESSID"),
		Password: os.Getenv("SUMOLOGIC_ACCESSKEY"),
	})

	resp, _, err := sumologicClient.UserManagementApi.ListUsers(auth).Execute()
	fmt.Println(resp)
	if err != nil {
		return err
	}
	users = append(users, resp.GetData()...)

	g.Resources = g.createResources(users)
	return nil
}
