// +build !linux

package lib

import "github.com/kubernetes-incubator/cri-o/lib/sandbox"

func (c *ContainerServer) addSandboxPlatform(sb *sandbox.Sandbox) {
	// nothin' doin'
}

func (c *ContainerServer) removeSandboxPlatform(sb *sandbox.Sandbox) {
	// nothin' doin'
}