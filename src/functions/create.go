package functions

import (
	"github.com/jorgevvs2/dockeryzer/src/utils"
)

func Create(imageName string, ignoreComments bool) {
	utils.CreateDockerfileContent(ignoreComments)
	utils.CreateDockerignoreContent()

	utils.ShowCreateSuccessfulOutput(imageName)

	if imageName != "" {
		cmd := utils.ExecDockerBuildCommand(imageName)
		utils.HandleCommandOutput(cmd)
	}
}
