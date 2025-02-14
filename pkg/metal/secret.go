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

import (
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// ReadCredentialsSecret reads a secret containing credentials.
func ReadCredentialsSecret(secret *corev1.Secret) (*Credentials, error) {
	if secret.Data == nil {
		return nil, fmt.Errorf("secret does not contain any data")
	}

	url, ok := secret.Data[APIURL]
	if !ok {
		return nil, fmt.Errorf("missing %q field in secret", APIURL)
	}

	hmac, hmacOK := secret.Data[APIHMac]
	key, keyOK := secret.Data[APIKey]

	if hmacOK && keyOK {
		return nil, errors.New("metalAPIHMac and metalAPIKey given, only one allowed")
	}
	if !hmacOK && !keyOK {
		return nil, errors.New("neither metalAPIHMac nor metalAPIKey given")
	}

	return &Credentials{
		MetalAPIURL:  string(url),
		MetalAPIHMac: string(hmac),
		MetalAPIKey:  string(key),
	}, nil
}
