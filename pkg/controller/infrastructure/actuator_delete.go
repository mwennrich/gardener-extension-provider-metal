package infrastructure

import (
	"context"
	"time"

	metalclient "github.com/metal-pod/gardener-extension-provider-metal/pkg/metal/client"

	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	controllererrors "github.com/gardener/gardener-extensions/pkg/controller/error"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

func (a *actuator) delete(ctx context.Context, infrastructure *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	infrastructureConfig, infrastructureStatus, err := a.decodeInfrastructure(infrastructure)
	if err != nil {
		return err
	}

	mclient, err := metalclient.NewClient(ctx, a.client, &infrastructure.Spec.SecretRef)
	if err != nil {
		return err
	}

	firewallStatus := infrastructureStatus.Firewall
	if firewallStatus.MachineID != "" {
		machineID := decodeMachineID(firewallStatus.MachineID)
		if machineID != "" {
			_, err = mclient.MachineDelete(machineID)
			if err != nil {
				a.logger.Error(err, "failed to delete firewall", "infrastructure", infrastructure.Name, "firewallID", machineID)
				return &controllererrors.RequeueAfterError{
					Cause:        err,
					RequeueAfter: 30 * time.Second,
				}
			}

			firewallStatus.MachineID = ""
			err = a.updateProviderStatus(ctx, infrastructure, infrastructureConfig, firewallStatus)
			if err != nil {
				a.logger.Error(err, "unable to update provider status after firewall deletion", "infrastructure", infrastructure.Name)
				return &controllererrors.RequeueAfterError{
					Cause:        err,
					RequeueAfter: 30 * time.Second,
				}
			}
		}
	}

	projectID := cluster.Shoot.Spec.Cloud.Metal.ProjectID
	nodeCIDR := cluster.Shoot.Spec.Cloud.Metal.Networks.Nodes
	clusterID := string(cluster.Shoot.ObjectMeta.UID)

	ipsToFree, ipsToUpdate, err := metalclient.GetEphemeralIPsFromCluster(mclient, projectID, clusterID)
	if err != nil {
		a.logger.Error(err, "failed to query ephemeral cluster ips", "infrastructure", infrastructure.Name, "clusterID", clusterID)
		return &controllererrors.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}

	for _, ip := range ipsToFree {
		_, err := mclient.IPFree(*ip.Ipaddress)
		if err != nil {
			a.logger.Error(err, "failed to release ephemeral cluster ip", "infrastructure", infrastructure.Name, "ip", *ip.Ipaddress)
			return &controllererrors.RequeueAfterError{
				Cause:        err,
				RequeueAfter: 30 * time.Second,
			}
		}
	}

	for _, ip := range ipsToUpdate {
		err := metalclient.UpdateIPInCluster(mclient, ip, clusterID)
		if err != nil {
			a.logger.Error(err, "failed to remove cluster tags from ip which is member of other clusters", "infrastructure", infrastructure.Name, "ip", *ip.Ipaddress)
			return &controllererrors.RequeueAfterError{
				Cause:        err,
				RequeueAfter: 30 * time.Second,
			}
		}
	}

	privateNetworks, err := metalclient.GetPrivateNetworksFromNodeNetwork(mclient, projectID, nodeCIDR)
	if err != nil {
		a.logger.Error(err, "failed to query private network", "infrastructure", infrastructure.Name, "nodeCIDR", nodeCIDR)
		return &controllererrors.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}

	for _, pn := range privateNetworks {
		_, err := mclient.NetworkFree(*pn.ID)
		if err != nil {
			a.logger.Error(err, "failed to release private network", "infrastructure", infrastructure.Name, "networkID", *pn.ID)
			return &controllererrors.RequeueAfterError{
				Cause:        err,
				RequeueAfter: 30 * time.Second,
			}
		}
	}

	return nil
}
