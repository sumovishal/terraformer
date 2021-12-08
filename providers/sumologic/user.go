package sumologic

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
	"os"
	"strings"

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
	cf.Servers = sumologic.ServerConfigurations{
		sumologic.ServerConfiguration{
			URL:         strings.TrimSuffix(os.Getenv("SUMOLOGIC_BASE_URL"), "/"),
			Description: fmt.Sprintf("%s deployment API server", os.Getenv("SUMOLOGIC_BASE_URL")),
		},
	}
	sumologicClient := sumologic.NewAPIClient(cf)
	auth := context.WithValue(context.Background(), sumologic.ContextBasicAuth, sumologic.BasicAuth{
		UserName: os.Getenv("SUMOLOGIC_ACCESSID"),
		Password: os.Getenv("SUMOLOGIC_ACCESSKEY"),
	})
	log.Printf("%s deployment API server", os.Getenv("SUMOLOGIC_BASE_URL"))
	resp, _, err := sumologicClient.UserManagementApi.ListUsers(auth).Execute()
	if err != nil {
		return err
	}
	users = append(users, resp.GetData()...)

	g.Resources = g.createResources(users)
	return nil
}
