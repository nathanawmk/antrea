// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import "fmt"

// Conversion functions between GroupMember and GroupMemberPod
func (g *GroupMember) ToGroupMemberPod() *GroupMemberPod {
	return &GroupMemberPod{
		Pod:   g.Pod,
		IP:    g.Endpoints[0].IP,
		Ports: g.Endpoints[0].Ports,
	}
}

func (p *GroupMemberPod) ToGroupMember() *GroupMember {
	return &GroupMember{
		Pod: p.Pod,
		Endpoints: []Endpoint{
			{IP: p.IP, Ports: p.Ports},
		},
	}
}

func (r *NetworkPolicyReference) ToString() string {
	if r.Type == AntreaClusterNetworkPolicy {
		return fmt.Sprintf("%s:%s", r.Type, r.Name)
	}
	return fmt.Sprintf("%s:%s/%s", r.Type, r.Namespace, r.Name)
}
