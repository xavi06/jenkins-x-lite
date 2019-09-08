package cmd

import (
	"fmt"
	"log"

	"github.com/xavi06/jenkins-x-lite/pkg/helm"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
	"github.com/urfave/cli"
)

func printHelmStatus(release string) {
	hcli := helm.NewHCli()
	hcli.Conn()
	resp, err := hcli.ReleaseStatus(release)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("LAST DEPLOYED: %s\n", util.ConvertUUnixtimeToString(resp.Info.LastDeployed.Seconds))
	fmt.Printf("NAMESPACE: %s\n", resp.Namespace)
	fmt.Printf("STATUS: %s\n\n", resp.Info.Status.Code)
	fmt.Println("RESOURCES:")
	fmt.Println(resp.Info.Status.Resources)
}

// HelmList func
func HelmList(c *cli.Context) error {
	var hcli *helm.HCli
	hcli = helm.NewHCli()
	hcli.Conn()

	res, err := hcli.ListRelease()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("%-32s%-16s%-32s%-16s%-40s%-16s%-32s\n",
		"NAME",
		"REVISION",
		"UPDATED",
		"STATUS",
		"CHART",
		"APP VERSION",
		"NAMESPACE",
	)

	for _, rls := range res.Releases {
		fmt.Printf("%-32s%-16d%-32s%-16s%-40s%-16s%-32s\n",
			rls.Name,
			rls.Version,
			util.ConvertUUnixtimeToString(rls.Info.LastDeployed.Seconds),
			rls.Info.Status.Code,
			rls.Chart.Metadata.Name+"-"+rls.Chart.Metadata.Version,
			rls.Chart.Metadata.AppVersion,
			rls.Namespace,
		)
	}
	return nil
}

// HelmDelete func
func HelmDelete(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Delete helm chart")
	hcli := helm.NewHCli()
	hcli.Conn()
	_, err = hcli.Delete(config.Chart.Name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Delete success")
	return nil
}

// HelmStatus func
func HelmStatus(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Get helm release status")
	printHelmStatus(config.Chart.Name)
	return nil
}

// HelmInstall func
func HelmInstall(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Install helm chart")
	hcli := helm.NewHCli()
	hcli.Conn()
	resp, err := hcli.Install("charts/"+config.Chart.Name, config.Namespace, config.Chart.Name)
	if err != nil {
		log.Fatalln((err))
	}
	fmt.Println((resp))
	return nil
}

// HelmUpgrade func
func HelmUpgrade(c *cli.Context) error {
	config, err := util.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Upgrade helm chart")
	hcli := helm.NewHCli()
	hcli.Conn()
	resp, err := hcli.Upgrade("charts/"+config.Chart.Name, config.Chart.Name)
	if err != nil {
		log.Fatalln((err))
	}
	fmt.Println((resp))
	return nil
}
