//go:build homebrew

package main

func GetKubeCTLInstallCommand() (string, []string) {
	cmd := "brew"
	args := []string{"install", "kubectl"}
	return cmd, args
}
