//go:build !homebrew

package main

func GetKubeCTLInstallCommand() (string, []string) {
	cmd := "curl"
	args := []string{"-LO", "https://dl.k8s.io/release/v1.26.1/bin/linux/amd64/kubectl"}
	return cmd, args
}
