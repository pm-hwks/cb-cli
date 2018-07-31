package cli

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1users"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var userListHeader = []string{"Name"}

type userListOut models_cloudbreak.User

func (u *userListOut) DataAsStringArray() []string {
	return []string{u.Username}
}

func ListUsers(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list users")
	log.Infof("[ListUsers] List all users in a tenant")
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	resp, err := cbClient.Cloudbreak.V1users.GetAllUsers(v1users.NewGetAllUsersParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, user := range resp.Payload {
		tableRows = append(tableRows, &userListOut{user.Username})
	}
	output.WriteList(userListHeader, tableRows)
}
