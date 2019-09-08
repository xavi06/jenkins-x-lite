package helm

import (
	gohelm "k8s.io/helm/pkg/helm"
	rls "k8s.io/helm/pkg/proto/hapi/services"

	"github.com/xavi06/jenkins-x-lite/pkg/kubernetes"
	"github.com/xavi06/jenkins-x-lite/pkg/util"
)

// HCli struct
type HCli struct {
	client  *gohelm.Client
	helmbin string
	//	host   string
}

// NewHCli func
func NewHCli() *HCli {
	helmbin, err := util.GetExcPath("helm")
	if err != nil {
		helmbin = "/usr/local/bin/helm"
	}
	return &HCli{helmbin: helmbin}
}

// SetHost func
/*
func (h *HCli) SetHost(host string) {
	os.Setenv("TILLER_ADDR", host)
	h.host = host
}
*/

// Conn func
func (h *HCli) Conn() {
	host := getTillerHost()
	h.client = gohelm.NewClient(gohelm.Host(host))
}

// ListRelease func
func (h *HCli) ListRelease() (*rls.ListReleasesResponse, error) {
	return h.client.ListReleases()
}

// Install func
func (h *HCli) Install(ch, ns, release string) (string, error) {
	var args []string
	kubeconfig := kubernetes.GetKubeconfig()
	args = append(args, "--kubeconfig", kubeconfig, "install", "--name", release, "--namespace", ns, ch)
	outStr, _, err := util.RunCommand(h.helmbin, args...)
	if err != nil {
		return "", err
	}
	return outStr, nil
}

// Upgrade func
func (h *HCli) Upgrade(ch, release string) (string, error) {
	var args []string
	kubeconfig := kubernetes.GetKubeconfig()
	args = append(args, "--kubeconfig", kubeconfig, "upgrade", release, ch)
	outStr, _, err := util.RunCommand(h.helmbin, args...)
	if err != nil {
		return "", err
	}
	return outStr, nil
}

// Delete func
func (h *HCli) Delete(release string) (*rls.UninstallReleaseResponse, error) {
	return h.client.DeleteRelease(release, gohelm.DeletePurge(true))
}

// ReleaseStatus func
func (h *HCli) ReleaseStatus(release string) (*rls.GetReleaseStatusResponse, error) {
	return h.client.ReleaseStatus(release)
}
