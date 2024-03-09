package functions

import (
	"github.com/jorgevvs2/dockeryzer/src/utils"
)

func Analyze(name string) {
	imageInspect := utils.GetDockerImageInspectByIdOrName(name)
	utils.PrintImageAnalyzeResults(name, imageInspect)
}
