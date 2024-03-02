package utils

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetImageSizeInMBs(imageInspect types.ImageInspect) float32 {
	sizeInMbs := float32(imageInspect.Size) / float32(math.Pow(10.0, 6))

	return sizeInMbs
}

func GetImageSizeString(imageInspect types.ImageInspect) string {
	sizeUnit := "MB"
	sizeInMbs := float32(imageInspect.Size) / float32(math.Pow(10.0, 6))
	sizeInGbs := float32(0.0)

	finalSize := sizeInMbs
	isMoreThanOneGb := sizeInMbs > 1000
	if isMoreThanOneGb {
		sizeUnit = "GB"
		sizeInGbs = sizeInMbs / float32(math.Pow(10.0, 3))
		finalSize = sizeInGbs
	}

	return fmt.Sprintf("%.2f %s", finalSize, sizeUnit)
}

func GetImageNumberOfLayers(imageInspect types.ImageInspect) int {
	return len(imageInspect.RootFS.Layers)
}

func GetImageFormattedCreationDate(imageInspect types.ImageInspect) string {
	parsedTime, err := time.Parse(time.RFC3339Nano, imageInspect.Created)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	return parsedTime.Format("02 Jan 2006")
}

func GetImageAuthor(imageInspect types.ImageInspect) string {
	if imageInspect.Author == "" {
		return "<none>"
	}
	return imageInspect.Author
}

const nodeVersionNoIdentified = "<no-identified>"

func GetImageNodeJsVersionString(imageInspect types.ImageInspect) string {
	envVars := imageInspect.Config.Env
	nodeVersion := nodeVersionNoIdentified

	for _, envVar := range envVars {
		if strings.Contains(envVar, "NODE_VERSION") {
			nodeVersion = strings.Split(envVar, "=")[1]
			break
		}
	}
	return nodeVersion
}

func GetImageNodeJsMajorVersionNumber(imageInspect types.ImageInspect) int {
	nodeVersion := GetImageNodeJsVersionString(imageInspect)
	if nodeVersion == nodeVersionNoIdentified {
		return 0
	}
	num, err := strconv.Atoi(strings.Split(nodeVersion, ".")[0])
	if err != nil {
		fmt.Println("Error on get Node.js major version number", err)
		os.Exit(1)
	}
	return num
}

func GetImageNodeJsVersionWithColor(imageInspect types.ImageInspect) string {
	nodeJsMajorVersion := GetImageNodeJsMajorVersionNumber(imageInspect)
	nodeJsVersionString := GetImageNodeJsVersionString(imageInspect)

	fmt.Printf("  - Node.js version: ")
	if nodeJsMajorVersion < 14 {
		return GetErrorOutput().Sprintf(nodeJsVersionString)

	}
	if nodeJsMajorVersion >= 14 && nodeJsMajorVersion <= 16 {
		return GetWarningOutput().Sprintf(nodeJsVersionString)

	}

	return GetSuccessOutput().Sprintf(nodeJsVersionString)
}

func GetImageSizeWithColor(imageInspect types.ImageInspect) string {
	sizeInMBs := GetImageSizeInMBs(imageInspect)

	fmt.Printf("  - Size: ")
	if sizeInMBs < 250 {
		return GetSuccessOutput().Sprintf("%s", GetImageSizeString(imageInspect))

	}

	if sizeInMBs >= 250 && sizeInMBs <= 500 {
		return GetWarningOutput().Sprintf("%s", GetImageSizeString(imageInspect))

	}

	return GetErrorOutput().Sprintf("%s", GetImageSizeString(imageInspect))

}

func GetImageLayersWithColor(imageInspect types.ImageInspect) string {
	numberOfLayers := GetImageNumberOfLayers(imageInspect)

	fmt.Printf("  - N. of Layers: ")
	if numberOfLayers < 10 {
		return GetSuccessOutput().Sprintf("%d", numberOfLayers)

	}

	if numberOfLayers >= 10 && numberOfLayers <= 20 {
		return GetWarningOutput().Sprintf("%d", numberOfLayers)

	}

	return GetErrorOutput().Sprintf("%d", numberOfLayers)

}

func PrintImageAnalyzeResults(name string, imageInspect types.ImageInspect, minimal bool) {
	fmt.Printf("Details of image ")
	GetBoldOutput().Printf("%s:\n", name)
	fmt.Printf("  - Tags: %s\n", imageInspect.RepoTags)
	fmt.Println(GetImageSizeWithColor(imageInspect))
	fmt.Println(GetImageLayersWithColor(imageInspect))
	fmt.Println(GetImageNodeJsVersionWithColor(imageInspect))
	if !minimal {
		fmt.Printf("  - Author: %s\n", GetImageAuthor(imageInspect))
		fmt.Printf("  - Creation date: %s\n", GetImageFormattedCreationDate(imageInspect))
		fmt.Printf("  - OS: %s\n", imageInspect.Os)
	}

	sizeInMBs := GetImageSizeInMBs(imageInspect)
	numberOfLayers := GetImageNumberOfLayers(imageInspect)
	nodeJsMajorVersion := GetImageNodeJsMajorVersionNumber(imageInspect)

	isBigImage := sizeInMBs > 250
	hasManyLayers := numberOfLayers > 10
	isOutdatedNodeVersion := nodeJsMajorVersion < 16

	shouldShowSuggestions := isBigImage || hasManyLayers || isOutdatedNodeVersion

	if shouldShowSuggestions {
		fmt.Println("\n Improvement suggestions:")
	}

	if isBigImage {
		fmt.Println("  - Consider reducing the size of your image. Try using smaller base images and ensure that no unnecessary files are included.")
	}

	if hasManyLayers {
		fmt.Println("  - Your image has multiple layers. Consider applying a multi-build stage strategy or combining commands to reduce the number of layers.")
	}

	if isOutdatedNodeVersion {
		if nodeJsMajorVersion == 0 {
			fmt.Println("  - No Node.js version detected. Ensure that your image is created correctly and includes a valid Node.js installation.")
		} else {
			fmt.Println("  - An outdated Node.js version is detected. It's recommended to use the latest version to ensure the security of your image.")
		}
	}
}

func PrintImageCompareLayersResults(image1 string, image1Inspect types.ImageInspect, image2 string, image2Inspect types.ImageInspect) {
	numberOfLayers1 := len(image1Inspect.RootFS.Layers)
	numberOfLayers2 := len(image2Inspect.RootFS.Layers)

	minorImage := image1
	minorLayers := numberOfLayers1
	if numberOfLayers2 < numberOfLayers1 {
		minorImage = image2
		minorLayers = numberOfLayers2
	}

	biggerImage := image1
	mostLayers := numberOfLayers1
	if numberOfLayers2 > numberOfLayers1 {
		biggerImage = image2
		mostLayers = numberOfLayers2
	}

	layersDiff := numberOfLayers1 - numberOfLayers2
	if numberOfLayers2 > numberOfLayers1 {
		layersDiff = numberOfLayers2 - numberOfLayers1
	}

	if layersDiff == 0 {
		fmt.Printf("  - Images have the same number of layers: %d\n.", numberOfLayers2)
		return
	}
	fmt.Printf("  - Image ")
	GetSuccessOutput().Printf("%s", minorImage)
	fmt.Printf(" has ")
	GetSuccessOutput().Printf("%d less layers", layersDiff)
	fmt.Printf(" than image ")
	GetErrorOutput().Printf("%s", biggerImage)
	fmt.Printf(" (")
	GetSuccessOutput().Printf("%d", minorLayers)
	fmt.Printf(" < ")
	GetErrorOutput().Printf("%d", mostLayers)
	fmt.Println(").")
}

func PrintImageCompareSizeResults(image1 string, image1Inspect types.ImageInspect, image2 string, image2Inspect types.ImageInspect) {
	size1 := image1Inspect.Size
	size2 := image2Inspect.Size

	size1String := GetImageSizeString(image1Inspect)
	size2String := GetImageSizeString(image2Inspect)

	minorImage := image1
	minorImageString := size1String
	minorSize := size1
	if size2 < size1 {
		minorImage = image2
		minorImageString = size2String
		minorSize = size2
	}

	biggerImage := image1
	biggerImageString := size1String
	biggerSize := size1
	if size2 > size1 {
		biggerImage = image2
		biggerImageString = size2String
		biggerSize = size2
	}

	sizeDiff := size1 - size2
	if size2 > size1 {
		sizeDiff = size2 - size1
	}

	if sizeDiff == 0 {
		fmt.Printf("  - Images have the same size: %s\n", GetImageSizeString(image1Inspect))
		return
	}

	percent := 100 - (float32(minorSize)/float32(biggerSize))*100

	fmt.Printf("  - Image ")
	GetSuccessOutput().Printf("%s", minorImage)
	fmt.Printf(" is ")
	GetSuccessOutput().Printf("%.2f%% smaller", percent)
	fmt.Printf(" than image ")
	GetErrorOutput().Printf("%s", biggerImage)
	fmt.Printf(" (")
	GetSuccessOutput().Printf(minorImageString)
	fmt.Printf(" < ")
	GetErrorOutput().Printf(biggerImageString)
	fmt.Println(").")
}

func PrintImageCompareNodeJsResults(image1 string, image1Inspect types.ImageInspect, image2 string, image2Inspect types.ImageInspect) {
	nodeJsMajorVersion1 := GetImageNodeJsMajorVersionNumber(image1Inspect)
	nodeJsMajorVersion2 := GetImageNodeJsMajorVersionNumber(image2Inspect)

	nodeJsVersionString1 := GetImageNodeJsVersionString(image1Inspect)
	nodeJsVersionString2 := GetImageNodeJsVersionString(image2Inspect)

	minorVersionImage := image1
	minorVersionStringImage := nodeJsVersionString1
	if nodeJsMajorVersion2 < nodeJsMajorVersion1 {
		minorVersionImage = image2
		minorVersionStringImage = nodeJsVersionString2
	}

	latestVersionImage := image1
	biggerVersionStringImage := nodeJsVersionString1
	if nodeJsMajorVersion2 > nodeJsMajorVersion1 {
		latestVersionImage = image2
		biggerVersionStringImage = nodeJsVersionString2
	}

	if nodeJsMajorVersion1 == nodeJsMajorVersion2 {
		fmt.Printf("  - Both images are using the same Node.js version: %d\n.", nodeJsMajorVersion2)
		return
	}
	fmt.Printf("  - Image ")
	GetSuccessOutput().Printf("%s", latestVersionImage)
	fmt.Printf(" is utilizing a ")
	GetSuccessOutput().Printf("more recent")
	fmt.Printf(" version of Node.js compared to image ")
	GetErrorOutput().Printf("%s", minorVersionImage)
	fmt.Printf(" (")
	GetSuccessOutput().Printf(biggerVersionStringImage)
	fmt.Printf(" > ")
	GetErrorOutput().Printf(minorVersionStringImage)
	fmt.Println(").")
}
