/*
Copyright 2019 The Kubernetes Authors.
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

package v1alpha3

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBareMetalMachineDefault(t *testing.T) {
	// No-op because we do not default anything in BMM yet
	c := &BareMetalMachine{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "fooboo",
		},
		Spec: BareMetalMachineSpec{},
	}
	c.Default()
}

func TestBareMetalMachineValidation(t *testing.T) {
	valid := &BareMetalMachine{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "foo",
		},
		Spec: BareMetalMachineSpec{
			Image: Image{
				URL:      "http://abc.com/image",
				Checksum: "http://abc.com/image.md5sum",
			},
		},
	}
	invalidURL := valid.DeepCopy()
	invalidURL.Spec.Image.URL = ""

	invalidChecksum := valid.DeepCopy()
	invalidChecksum.Spec.Image.Checksum = ""

	tests := []struct {
		name      string
		expectErr bool
		c         *BareMetalMachine
	}{
		{
			name:      "should return error when url empty",
			expectErr: true,
			c:         invalidURL,
		},
		{
			name:      "should return error when checksum empty",
			expectErr: true,
			c:         invalidChecksum,
		},
		{
			name:      "should succeed when image correct",
			expectErr: false,
			c:         valid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			if tt.expectErr {
				g.Expect(tt.c.ValidateCreate()).NotTo(Succeed())
				g.Expect(tt.c.ValidateUpdate(nil)).NotTo(Succeed())
			} else {
				g.Expect(tt.c.ValidateCreate()).To(Succeed())
				g.Expect(tt.c.ValidateUpdate(nil)).To(Succeed())
			}
		})
	}
}
