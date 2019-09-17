# Deployment Patterns

## controller in seed controlplane accesses ressources via shoot kube-apiserver

* Deploy Controller in seed controlplane with generated kubeconfig for the shoot-apiserver (charts/internal/control-plane and add charts to valuesprovider.go#controlPlaneChart and provide prefixed values)
* Generate Certificate and Kubeconfig as Secret

```
&secrets.ControlPlaneSecretConfig{
    CertificateSecretConfig: &secrets.CertificateSecretConfig{
        Name:       accountingExporterName,
        CommonName: "system:accounting-exporter",
        // Groupname of user
        Organization: []string{accountingExporterName},
        CertType:     secrets.ClientCert,
        SigningCA:    cas[gardencorev1alpha1.SecretNameCACluster],
    },
    KubeConfigRequest: &secrets.KubeConfigRequest{
        ClusterName:  clusterName,
        APIServerURL: gardencorev1alpha1.DeploymentNameKubeAPIServer,
    },
},
```

* Configure Controller with Kubeconfig via commandline and volume/mounts

```
    spec:
      containers:
      - image: {{ index .Values.images "accounting-exporter" }}
        name: accounting-exporter
        env:
        - name: KUBE_COUNTER_KUBECONFIG
          value: /var/lib/accounting-exporter/kubeconfig
        volumeMounts:
        - name: accounting-exporter
          mountPath: /var/lib/accounting-exporter
      restartPolicy: Always
      volumes:
      - name: accounting-exporter
        secret:
          secretName: accounting-exporter
```

* Deploy (Cluster-)Role / (Cluster-)RoleBindings in shoot (!) using /charts/internal/shoot-control-plane
* Bind to the User with the CommonName of your Certificate (see above)

```
subjects:
- kind: User
  name: system:accounting-exporter
  apiGroup: ""
```