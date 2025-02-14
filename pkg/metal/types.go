// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metal

import "path/filepath"

const (
	// Name is the name of the Metal provider.
	Name = "provider-metal"
	// StorageProviderName is the name of the Metal storage provider.
	StorageProviderName = "S3"

	// MachineControllerManagerImageName is the name of the MachineControllerManager image.
	MachineControllerManagerImageName = "machine-controller-manager"
	// CCMImageName is the name of the cloud controller manager image.
	CCMImageName = "metalccm"
	// ETCDBackupRestoreImageName is the name of the etcd backup and restore image.
	ETCDBackupRestoreImageName = "etcd-backup-restore"
	// GroupRolebindingControllerImageName is the name of the GroupRolebindingController image
	GroupRolebindingControllerImageName = "group-rolebinding-controller"
	// AccountingExporterImageName is the name of the accounting exporter image
	AccountingExporterImageName = "accounting-exporter"
	// AuthNWebhookImageName is the name of the AuthN Webhook configured with the shoot kube-apiserver
	AuthNWebhookImageName = "authn-webhook"
	// DroptailerImageName is the name of the Droptailer to deploy to the shoot.
	DroptailerImageName = "droptailer"
	// LimitValidatingWebhookImageName is the name of the limit validating webhook to deploy to the seed's shoot namespace.
	LimitValidatingWebhookImageName = "limit-validating-webhook"

	// APIURL is a constant for the url of metal-api.
	APIURL = "metalAPIURL"
	// APIKey is a constant for the key in a cloud provider secret.
	APIKey = "metalAPIKey"
	// APIHMac is a constant for the key in a cloud provider secret.
	APIHMac = "metalAPIHMac"
	// Region is a constant for the key in a backup secret that holds the AWS region.
	Region = "region"
	// BucketName is a constant for the key in a backup secret that holds the bucket name.
	// The bucket name is written to the backup secret by Gardener as a temporary solution.
	// TODO In the future, the bucket name should come from a BackupBucket resource (see https://github.com/gardener/gardener/blob/master/docs/proposals/02-backupinfra.md)
	BucketName = "bucketName"
	// SSHKeyName key for accessing SSH key name from outputs in terraform
	SSHKeyName = "keyName"

	// CloudProviderConfigName is the name of the configmap containing the cloud provider config.
	CloudProviderConfigName = "cloud-provider-config"
	// MachineControllerManagerName is a constant for the name of the machine-controller-manager.
	MachineControllerManagerName = "machine-controller-manager"
	// BackupSecretName is the name of the secret containing the credentials for storing the backups of Shoot clusters.
	BackupSecretName = "etcd-backup"

	// AuthN Webhook
	AuthNWebHookConfigName        = "authn-webhook-config"
	AuthNWebHookCertName          = "authn-webhook-cert"
	ShootExtensionTypeTokenIssuer = "tokenissuer"

	// Shoot Annotations
	ShootAnnotationCreatedBy   = "garden.sapcloud.io/createdBy"
	ShootAnnotationPurpose     = "garden.sapcloud.io/purpose"
	ShootAnnotationProject     = "cluster.metal-pod.io/project"
	ShootAnnotationDescription = "cluster.metal-pod.io/description"
	ShootAnnotationClusterName = "cluster.metal-pod.io/name"
	ShootAnnotationTenant      = "cluster.metal-pod.io/tenant"
)

var (
	// ChartsPath is the path to the charts
	ChartsPath = filepath.Join("controllers", Name, "charts")
	// InternalChartsPath is the path to the internal charts
	InternalChartsPath = filepath.Join(ChartsPath, "internal")
)

// Credentials stores Metal credentials.
type Credentials struct {
	MetalAPIURL  string
	MetalAPIKey  string
	MetalAPIHMac string
}
