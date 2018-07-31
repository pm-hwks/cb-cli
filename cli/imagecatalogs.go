package cli

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/cli/cloud"
	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/hortonworks/cb-cli/client_cloudbreak/v1imagecatalogs"
	"github.com/hortonworks/cb-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var imagecatalogHeader = []string{"Name", "Default", "URL"}

type imagecatalogOut struct {
	Name    string `json:"Name" yaml:"Name"`
	Default bool   `json:"Default" yaml:"Default"`
	URL     string `json:"URL" yaml:"URL"`
}

func (r *imagecatalogOut) DataAsStringArray() []string {
	return []string{r.Name, strconv.FormatBool(r.Default), r.URL}
}

var imageHeader = []string{"Date", "Description", "Version", "ImageID"}

type imageOut struct {
	Date        string `json:"Date" yaml:"Date"`
	Description string `json:"Description" yaml:"Description"`
	Version     string `json:"Version" yaml:"Version"`
	ImageID     string `json:"ImageID" yaml:"ImageID"`
}

func (r *imageOut) DataAsStringArray() []string {
	return []string{r.Date, r.Description, r.Version, r.ImageID}
}

var imageDetailsHeader = []string{"Date", "Description", "Version", "ImageID", "OS", "OS Type", "Images", "Package Versions"}

type imageDetailsOut struct {
	Date            string                       `json:"Date" yaml:"Date"`
	Description     string                       `json:"Description" yaml:"Description"`
	Version         string                       `json:"Version" yaml:"Version"`
	ImageID         string                       `json:"ImageID" yaml:"ImageID"`
	Os              string                       `json:"OS" yaml:"OS"`
	OsType          string                       `json:"OSType" yaml:"OSType"`
	Images          map[string]map[string]string `json:"Images" yaml:"Images"`
	PackageVersions map[string]string            `json:"PackageVersions" yaml:"PackageVersions"`
}

func (r *imageDetailsOut) DataAsStringArray() []string {
	var images string
	for prov, imgs := range r.Images {
		images += fmt.Sprintf("%s:\n", prov)
		for region, image := range imgs {
			images += fmt.Sprintf("  %s: %s\n", region, image)
		}
	}

	var packageVersions string
	for pkg, ver := range r.PackageVersions {
		packageVersions += fmt.Sprintf("%s: %s\n", pkg, ver)
	}

	return []string{r.Date, r.Description, r.Version, r.ImageID, r.Os, r.OsType, images, packageVersions}
}

func CreateImagecatalogFromUrl(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)

	log.Infof("[CreateImagecatalogFromUrl] creating imagecatalog from a URL")
	cbClient := NewCloudbreakHTTPClientFromContext(c)
	createImagecatalogImpl(
		cbClient.Cloudbreak.V1imagecatalogs,
		c.String(FlName.Name),
		c.Bool(FlPublicOptional.Name),
		c.String(FlURL.Name))
}

type imagecatalogClient interface {
	PostPrivateImageCatalog(params *v1imagecatalogs.PostPrivateImageCatalogParams) (*v1imagecatalogs.PostPrivateImageCatalogOK, error)
	PostPublicImageCatalog(params *v1imagecatalogs.PostPublicImageCatalogParams) (*v1imagecatalogs.PostPublicImageCatalogOK, error)
}

func createImagecatalogImpl(client imagecatalogClient, name string, public bool, imagecatalogURL string) {
	defer utils.TimeTrack(time.Now(), "create imagecatalog")
	imagecatalogRequest := &models_cloudbreak.ImageCatalogRequest{
		Name: &name,
		URL:  &imagecatalogURL,
	}
	var ic *models_cloudbreak.ImageCatalogResponse
	if public {
		log.Infof("[createImagecatalogImpl] sending create public imagecatalog request")
		resp, err := client.PostPublicImageCatalog(v1imagecatalogs.NewPostPublicImageCatalogParams().WithBody(imagecatalogRequest))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		ic = resp.Payload
	} else {
		log.Infof("[createImagecatalogImpl] sending create private imagecatalog request")
		resp, err := client.PostPrivateImageCatalog(v1imagecatalogs.NewPostPrivateImageCatalogParams().WithBody(imagecatalogRequest))
		if err != nil {
			utils.LogErrorAndExit(err)
		}
		ic = resp.Payload
	}
	log.Infof("[createImagecatalogImpl] imagecatalog created: %s (id: %d)", ic.Name, ic.ID)
}

func ListImagecatalogs(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list imagecatalogs")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	listImagecatalogsImpl(cbClient.Cloudbreak.V1imagecatalogs, output.WriteList)
}

type getPublicsImagecatalogsClient interface {
	GetPublicsImageCatalogs(*v1imagecatalogs.GetPublicsImageCatalogsParams) (*v1imagecatalogs.GetPublicsImageCatalogsOK, error)
}

func listImagecatalogsImpl(client getPublicsImagecatalogsClient, writer func([]string, []utils.Row)) {
	log.Infof("[listImagecatalogsImpl] sending imagecatalog list request")
	imagecatalogResp, err := client.GetPublicsImageCatalogs(v1imagecatalogs.NewGetPublicsImageCatalogsParams())
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, ic := range imagecatalogResp.Payload {
		tableRows = append(tableRows, &imagecatalogOut{*ic.Name, ic.UsedAsDefault, *ic.URL})
	}

	writer(imagecatalogHeader, tableRows)
}

func DeleteImagecatalog(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "delete imagecatalog")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	name := c.String(FlName.Name)
	log.Infof("[DeleteImagecatalog] sending delete imagecatalog request with name: %s", name)
	if err := cbClient.Cloudbreak.V1imagecatalogs.DeletePublicImageCatalogByName(v1imagecatalogs.NewDeletePublicImageCatalogByNameParams().WithName(name)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteImagecatalog] imagecatalog deleted, name: %s", name)
}

func SetDefaultImagecatalog(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "set default imagecatalog")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	name := c.String(FlName.Name)
	log.Infof("[SetDefautlImagecatalog] sending set default imagecatalog request with name: %s", name)
	if _, err := cbClient.Cloudbreak.V1imagecatalogs.PutSetDefaultImageCatalogByName(v1imagecatalogs.NewPutSetDefaultImageCatalogByNameParams().WithName(name)); err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[SetDefaultImagecatalog] imagecatalog is set as default, name: %s", name)
}

func ListAwsImages(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	listImages(c)
}

func ListAzureImages(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	listImages(c)
}

func ListGcpImages(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	listImages(c)
}

func ListOpenstackImages(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	listImages(c)
}

func ListImagesValidForUpgrade(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list images valid for stack upgrade")

	cbClient := NewCloudbreakHTTPClientFromContext(c)

	clusterName := c.String(FlClusterToUpgrade.Name)
	imageCatalogName := c.String(FlImageCatalog.Name)

	log.Infof("[ListImagesValidForUpgrade] sending list images request")
	imageResp, err := cbClient.Cloudbreak.V1imagecatalogs.GetImagesByClusterNameAndCustomImageCatalog(v1imagecatalogs.NewGetImagesByClusterNameAndCustomImageCatalogParams().WithName(imageCatalogName).WithClusterName(clusterName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}

	for _, i := range imageResp.Payload.BaseImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}
	for _, i := range imageResp.Payload.HdfImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}
	for _, i := range imageResp.Payload.HdpImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}

	output.WriteList(imageHeader, tableRows)
}

func listImages(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "list available images")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	listImagesImpl(cbClient.Cloudbreak.V1imagecatalogs, output.WriteList, c.String(FlImageCatalog.Name))
}

func DescribeAwsImage(c *cli.Context) {
	cloud.SetProviderType(cloud.AWS)
	describeImage(c)
}

func DescribeAzureImage(c *cli.Context) {
	cloud.SetProviderType(cloud.AZURE)
	describeImage(c)
}

func DescribeGcpImage(c *cli.Context) {
	cloud.SetProviderType(cloud.GCP)
	describeImage(c)
}

func DescribeOpenstackImage(c *cli.Context) {
	cloud.SetProviderType(cloud.OPENSTACK)
	describeImage(c)
}

func describeImage(c *cli.Context) {
	checkRequiredFlagsAndArguments(c)
	defer utils.TimeTrack(time.Now(), "describe image")

	cbClient := NewCloudbreakHTTPClientFromContext(c)
	output := utils.Output{Format: c.String(FlOutputOptional.Name)}
	describeImageImpl(cbClient.Cloudbreak.V1imagecatalogs, output.WriteList, c.String(FlImageCatalog.Name), c.String(FlImageId.Name))
}

func describeImageImpl(client getPublicImagesClient, writer func([]string, []utils.Row), imagecatalog string, imageid string) {
	log.Infof("[listImagesImpl] sending list images request")
	provider := cloud.GetProvider().GetName()
	imageResp, err := client.GetPublicImagesByProviderAndCustomImageCatalog(v1imagecatalogs.NewGetPublicImagesByProviderAndCustomImageCatalogParams().WithName(imagecatalog).WithPlatform(*provider))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	image := findBaseImageByUUID(imageResp.Payload.BaseImages, imageid)

	if image == nil {
		utils.LogErrorMessageAndExit(fmt.Sprintf("Image not found by id: %s", imageid))
	}

	listImageInformation(writer, image)
}

func findBaseImageByUUID(baseImages []*models_cloudbreak.BaseImageResponse, imageid string) *models_cloudbreak.BaseImageResponse {
	for i := range baseImages {
		if baseImages[i].UUID == imageid {
			return baseImages[i]
		}
	}
	return nil
}

func listImageInformation(writer func([]string, []utils.Row), image *models_cloudbreak.BaseImageResponse) {
	tableRows := []utils.Row{}
	tableRows = append(tableRows, &imageDetailsOut{image.Date, image.Description, image.Version, image.UUID, image.Os, image.OsType, image.Images, image.PackageVersions})
	writer(imageDetailsHeader, tableRows)
}

type getPublicImagesClient interface {
	GetPublicImagesByProviderAndCustomImageCatalog(*v1imagecatalogs.GetPublicImagesByProviderAndCustomImageCatalogParams) (*v1imagecatalogs.GetPublicImagesByProviderAndCustomImageCatalogOK, error)
}

func listImagesImpl(client getPublicImagesClient, writer func([]string, []utils.Row), imagecatalog string) {
	log.Infof("[listImagesImpl] sending list images request")
	provider := cloud.GetProvider().GetName()
	imageResp, err := client.GetPublicImagesByProviderAndCustomImageCatalog(v1imagecatalogs.NewGetPublicImagesByProviderAndCustomImageCatalogParams().WithName(imagecatalog).WithPlatform(*provider))
	if err != nil {
		utils.LogErrorAndExit(err)
	}

	tableRows := []utils.Row{}
	for _, i := range imageResp.Payload.BaseImages {
		tableRows = append(tableRows, &imageOut{i.Date, i.Description, i.Version, i.UUID})
	}

	writer(imageHeader, tableRows)

}
