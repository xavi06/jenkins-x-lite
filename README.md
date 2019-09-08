# jxl
Jenkins X Lite, Inspired by Jenkins X.


## Install
local install
```bash
make build
```
macos
```bash
make macos
```
linux
```bash
make linux
```

## Pre Deploy
1, We need a file called app.toml. (see example directory)

2, We need some enviroment variables. (see example/env directory)

## Deploy
1, Init all config files we need.
```bash
jxl init
```

We will see Dockerfile,Jenkinsfile,charts directory in our working directory.

2, Install app
```bash
jxl install
```
This step will do
- Upload local helm chart to chartmuseum
- Upload Dockerfile to fileserver(written by go, a light http file server)
- Create jenkins pipeline(Use Jenkinsfile)
- Install helm chart


Show status
```bash
jxl status
```

## Uninstall(Destroy)
Be careful, It will delete all configfile and delete helm release and jenkins pipeline
```
jxl destroy
```

## Other Usage
```
jxl --help
```
Usage:

```
NAME:
   jxl - jenkins x lite, make deploy easy

USAGE:
   jxl [global options] command [command options] [arguments...]

VERSION:
   1.2.0

COMMANDS:
     create       create app.toml
     init         init config
     install      install app
     update       update app
     destroy      destroy app
     status       get app status
     helm         helm operation
     jenkins      jenkins operation
     dockerfile   dockerfile operation
     chartmuseum  chartmuseum operation
     tekton       tekton operation
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
