package service

import (
	"github.com/pkg/errors"
	pulumikubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
)

func GetIpAddress(addedKubeService *pulumikubernetescorev1.Service) string {
	loadBalancerIpAddress := addedKubeService.Status.ApplyT(
		func(status *pulumikubernetescorev1.ServiceStatus) (string, error) {
			if status.LoadBalancer.Ingress == nil || len(status.LoadBalancer.Ingress) == 0 {
				return "", errors.New("ingress LoadBalancer not found after service initialization is complete")
			}
			ingressIP := status.LoadBalancer.Ingress[0].Ip
			if ingressIP == nil {
				return "", errors.New("ingress LoadBalancer does not have an ip after service initialization is complete")
			}
			return *ingressIP, nil
		})
	return loadBalancerIpAddress.ElementType().String()
}
