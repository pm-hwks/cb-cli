package cli

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1organizations"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var orgListHeader = []string{"Name", "Description"}

type orgListOut models_cloudbreak.OrganizationResponse

func (o *orgListOut) DataAsStringArray() []string {
	return []string{o.Name, utils.SafeStringConvert(o.Description)}
}

func ListOrgs(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list organizations")
	log.Infof("[ListOrgs] List all organizations in a tenant")
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V1organizations.GetOrganizations(v1organizations.NewGetOrganizationsParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, org := range resp.Payload {
		tableRows = append(tableRows, &orgListOut{org.Description, org.ID, org.Name, org.Users})
	}
	output.WriteList(orgListHeader, tableRows)
}
