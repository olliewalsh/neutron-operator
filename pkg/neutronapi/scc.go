package neutronapi

import corev1 "k8s.io/api/core/v1"

func getNeutronSecurityContext() *corev1.PodSecurityContext {
	trueVal := true
	runAsUser := int64(NeutronUid)
	runAsGroup := int64(NeutronGid)

	return &corev1.PodSecurityContext{
		RunAsUser:    &runAsUser,
		RunAsGroup:   &runAsGroup,
		RunAsNonRoot: &trueVal,
		FSGroup:      &runAsGroup,
	}
}
