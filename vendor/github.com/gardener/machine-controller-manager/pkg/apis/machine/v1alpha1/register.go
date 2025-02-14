/*
Copyright (c) 2017 SAP SE or an SAP affiliate company. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeBuilder used to register the Machine resource.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	localSchemeBuilder = &SchemeBuilder

	// AddToScheme is a pointer to SchemeBuilder.AddToScheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// GroupName is the group name use in this package
const GroupName = "machine.sapcloud.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// func Init() {
// 	// We only register manually written functions here. The registration of the
// 	// generated functions takes place in the generated files. The separation
// 	// makes the code compile even when the generated files are missing.
// 	SchemeBuilder.Register(addKnownTypes)
// }

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&OpenStackMachineClass{},
		&OpenStackMachineClassList{},

		&AWSMachineClass{},
		&AWSMachineClassList{},

		&AzureMachineClass{},
		&AzureMachineClassList{},

		&GCPMachineClass{},
		&GCPMachineClassList{},

		&AlicloudMachineClass{},
		&AlicloudMachineClassList{},

		&PacketMachineClass{},
		&PacketMachineClassList{},

		&MetalMachineClass{},
		&MetalMachineClassList{},

		&Machine{},
		&MachineList{},

		&MachineSet{},
		&MachineSetList{},

		&MachineDeployment{},
		&MachineDeploymentList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
