package cmd

import (
	"github.com/urfave/cli"
)

// Commands var
var Commands = []cli.Command{
	{
		Name:   "create",
		Usage:  "create app.toml",
		Action: Create,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name", Usage: "chart(app) name"},
			cli.StringFlag{Name: "apptype", Usage: "app type, default maven"},
			cli.StringFlag{Name: "gitrepo", Usage: "git repo url"},
			cli.StringFlag{Name: "gitbranch", Usage: "git branch", Value: "develop_docker"},
			cli.IntFlag{Name: "pid", Usage: "project id", Value: 100},
			cli.IntFlag{Name: "replica", Usage: "repicaCount, default 1", Value: 1},
		},
	},
	{
		Name:   "init",
		Usage:  "init config",
		Action: Init,
	},
	{
		Name:   "install",
		Usage:  "install app",
		Action: Install,
	},
	{
		Name:   "update",
		Usage:  "update app",
		Action: Update,
	},
	{
		Name:   "destroy",
		Usage:  "destroy app",
		Action: Destroy,
	},
	{
		Name:   "status",
		Usage:  "get app status",
		Action: Status,
	},
	{
		Name:  "helm",
		Usage: "helm operation",
		Subcommands: []cli.Command{
			{
				Name:   "status",
				Usage:  "get helm release status",
				Action: HelmStatus,
			},
			{
				Name:   "delete",
				Usage:  "delete a release",
				Action: HelmDelete,
			},
			{
				Name:   "install",
				Usage:  "install a release [helm client needed]",
				Action: HelmInstall,
			},
			{
				Name:   "upgrade",
				Usage:  "upgrade a release [helm client needed]",
				Action: HelmUpgrade,
			},
		},
	},
	{
		Name:  "jenkins",
		Usage: "jenkins operation",
		Subcommands: []cli.Command{
			{
				Name:   "init",
				Usage:  "init Jenkinsfile",
				Action: JenkinsInit,
			},
			{
				Name:   "add",
				Usage:  "add a job",
				Action: JenkinsAdd,
			},
			{
				Name:   "delete",
				Usage:  "delete a job",
				Action: JenkinsDelete,
			},
			{
				Name:   "update",
				Usage:  "update a job",
				Action: JenkinsUpdate,
			},
			{
				Name:   "status",
				Usage:  "get job build status",
				Action: JenkinsStatus,
			},
		},
	},
	{
		Name:  "dockerfile",
		Usage: "dockerfile operation",
		Subcommands: []cli.Command{
			{
				Name:   "init",
				Usage:  "init Dockerfile",
				Action: DockerfileInit,
			},
			{
				Name:   "upload",
				Usage:  "upload Dockerfile to fileserver",
				Action: DockerfileUpload,
			},
		},
	},
	{
		Name:  "chartmuseum",
		Usage: "chartmuseum operation",
		Subcommands: []cli.Command{
			{
				Name:   "upload",
				Usage:  "upload to chartmuseum",
				Action: UploadChart,
			},
			{
				Name:   "delete",
				Usage:  "delete remote chart",
				Action: DeleteChart,
			},
			{
				Name:   "update",
				Usage:  "update remote chart",
				Action: UpdateChart,
			},
		},
	},
	{
		Name:  "tekton",
		Usage: "tekton operation",
		Subcommands: []cli.Command{
			{
				Name:   "create",
				Usage:  "create tekton pipeline",
				Action: CreateTekton,
			},
			{
				Name:   "update",
				Usage:  "update tekton pipeline",
				Action: UpdateTekton,
			},
			{
				Name: "logs",
				Usage: "get tekton taskRun logs",
				Action: GetTektonTaskRunLogs,
			},
		},
	},
}
