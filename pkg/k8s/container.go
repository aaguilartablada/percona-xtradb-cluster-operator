package k8s

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
)

const (
	// errCreateContainerConfig - failed to create container config
	errCreateContainerConfig string = "CreateContainerConfigError"
	// errPreCreateHook - failed to execute PreCreateHook
	errPreCreateHook string = "PreCreateHookError"
	// errCreateContainer - failed to create container
	errCreateContainer string = "CreateContainerError"
	// errPreStartHook - failed to execute PreStartHook
	errPreStartHook string = "PreStartHookError"
	// errPostStartHook - failed to execute PostStartHook
	errPostStartHook string = "PostStartHookError"
)

// PodContainerError return an error if some pod containers face an error
func PodContainerError(pod corev1.Pod) error {
	for _, status := range pod.Status.ContainerStatuses {
		if state := status.State.Waiting; state != nil {
			switch state.Reason {
			case errCreateContainer, errCreateContainerConfig, errPreCreateHook, errPreStartHook, errPostStartHook:
				return errors.New(state.Message)
			}
		}
	}
	return nil
}
