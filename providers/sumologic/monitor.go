package sumologic

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	sumologic "github.com/saurabh-agarwals/sumologic-api-client-go/openapi"
)

// MonitorGenerator ...
type MonitorGenerator struct {
	SumologicService
}

func (g *MonitorGenerator) createResources(monitors []sumologic.MonitorsLibraryBaseResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, monitor := range monitors {
		resources = append(resources, g.createResource(monitor.GetId()))
	}
	return resources
}

func (g *MonitorGenerator) createResource(monitorID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		monitorID,
		fmt.Sprintf("monitor_%s", monitorID),
		"sumologic_monitor",
		"sumologic",
		[]string{
			fmt.Sprintf("first_name = %s", "gahana"),
		},
	)
}

func (g *MonitorGenerator) InitResources() error {
	var monitors []sumologic.MonitorsLibraryBaseResponse
	cf := sumologic.NewConfiguration()
	sumologicClient := sumologic.NewAPIClient(cf)
	auth := context.WithValue(context.Background(), sumologic.ContextBasicAuth, sumologic.BasicAuth{
		UserName: os.Getenv("SUMOLOGIC_ACCESSID"),
		Password: os.Getenv("SUMOLOGIC_ACCESSKEY"),
	})

	resp, _, err := sumologicClient.MonitorsLibraryManagementApi.GetMonitorsLibraryRoot(auth).Execute()
	// fmt.Println(resp)
	if err != nil {
		return err
	}
	rootFolderChildren := resp.GetChildren()
	for _, child := range rootFolderChildren {
		if child.GetType() == "MonitorsLibraryMonitor" {
			monitors = append(monitors, child)
		}
	}

	g.Resources = g.createResources(monitors)
	return nil
}
