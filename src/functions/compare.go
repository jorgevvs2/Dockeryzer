package functions

import (
	"fmt"
	"github.com/jorgevvs2/dockeryzer/src/utils"
)

func Compare(image1, image2 string) {
	image1Inspect := utils.GetDockerImageInspectByIdOrName(image1)
	image2Inspect := utils.GetDockerImageInspectByIdOrName(image2)

	utils.PrintImageCompareResults(image1, image1Inspect)
	fmt.Println()
	utils.PrintImageCompareResults(image2, image2Inspect)
	fmt.Println()

	fmt.Println("Differences:")
	utils.PrintImageCompareLayersResults(image1, image1Inspect, image2, image2Inspect)
	utils.PrintImageCompareSizeResults(image1, image1Inspect, image2, image2Inspect)
	utils.PrintImageCompareNodeJsResults(image1, image1Inspect, image2, image2Inspect)
}
