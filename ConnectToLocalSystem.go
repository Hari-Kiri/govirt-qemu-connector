package goVirtQemuConnector

import "libvirt.org/go/libvirt"

// Connect to local Qemu socket. This function should be called first to get a connection to the Hypervisor.
// ConnectToLocalSystem().close() should be used to release the resources after the connection is no longer
// needed.
func ConnectToLocalSystem() (*libvirt.Connect, error) {
	result, errorResult := libvirt.NewConnect("qemu:///system")
	if errorResult != nil {
		return nil, errorResult
	}
	defer result.Close()

	return result, nil
}
