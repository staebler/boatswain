/*
Copyright 2017 The Kubernetes Authors.

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

package validation

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/staebler/boatswain/pkg/apis/boatswain"
)

func TestValidateNode(t *testing.T) {
	cases := []struct {
		name  string
		node  *boatswain.Node
		valid bool
	}{
		{
			name: "invalid node - node without namespace",
			node: &boatswain.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-node",
				},
			},
			valid: false,
		},
	}

	for _, tc := range cases {
		errs := ValidateNode(tc.node)
		if len(errs) != 0 && tc.valid {
			t.Errorf("%v: unexpected error: %v", tc.name, errs)
			continue
		} else if len(errs) == 0 && !tc.valid {
			t.Errorf("%v: unexpected success", tc.name)
		}
	}
}
