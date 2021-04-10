// Copyright 2021 Antrea Authors
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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	controlplanev1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta1"
	controlplanev1beta2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/controlplane/v1beta2"
	crdv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha1"
	crdv1alpha2 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/crd/v1alpha2"
	crdv1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/crd/v1beta1"
	statsv1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/stats/v1alpha1"
	systemv1beta1 "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/typed/system/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ControlplaneV1beta1() controlplanev1beta1.ControlplaneV1beta1Interface
	ControlplaneV1beta2() controlplanev1beta2.ControlplaneV1beta2Interface
	CrdV1alpha1() crdv1alpha1.CrdV1alpha1Interface
	CrdV1alpha2() crdv1alpha2.CrdV1alpha2Interface
	CrdV1beta1() crdv1beta1.CrdV1beta1Interface
	StatsV1alpha1() statsv1alpha1.StatsV1alpha1Interface
	SystemV1beta1() systemv1beta1.SystemV1beta1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	controlplaneV1beta1 *controlplanev1beta1.ControlplaneV1beta1Client
	controlplaneV1beta2 *controlplanev1beta2.ControlplaneV1beta2Client
	crdV1alpha1         *crdv1alpha1.CrdV1alpha1Client
	crdV1alpha2         *crdv1alpha2.CrdV1alpha2Client
	crdV1beta1          *crdv1beta1.CrdV1beta1Client
	statsV1alpha1       *statsv1alpha1.StatsV1alpha1Client
	systemV1beta1       *systemv1beta1.SystemV1beta1Client
}

// ControlplaneV1beta1 retrieves the ControlplaneV1beta1Client
func (c *Clientset) ControlplaneV1beta1() controlplanev1beta1.ControlplaneV1beta1Interface {
	return c.controlplaneV1beta1
}

// ControlplaneV1beta2 retrieves the ControlplaneV1beta2Client
func (c *Clientset) ControlplaneV1beta2() controlplanev1beta2.ControlplaneV1beta2Interface {
	return c.controlplaneV1beta2
}

// CrdV1alpha1 retrieves the CrdV1alpha1Client
func (c *Clientset) CrdV1alpha1() crdv1alpha1.CrdV1alpha1Interface {
	return c.crdV1alpha1
}

// CrdV1alpha2 retrieves the CrdV1alpha2Client
func (c *Clientset) CrdV1alpha2() crdv1alpha2.CrdV1alpha2Interface {
	return c.crdV1alpha2
}

// CrdV1beta1 retrieves the CrdV1beta1Client
func (c *Clientset) CrdV1beta1() crdv1beta1.CrdV1beta1Interface {
	return c.crdV1beta1
}

// StatsV1alpha1 retrieves the StatsV1alpha1Client
func (c *Clientset) StatsV1alpha1() statsv1alpha1.StatsV1alpha1Interface {
	return c.statsV1alpha1
}

// SystemV1beta1 retrieves the SystemV1beta1Client
func (c *Clientset) SystemV1beta1() systemv1beta1.SystemV1beta1Interface {
	return c.systemV1beta1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.controlplaneV1beta1, err = controlplanev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.controlplaneV1beta2, err = controlplanev1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.crdV1alpha1, err = crdv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.crdV1alpha2, err = crdv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.crdV1beta1, err = crdv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.statsV1alpha1, err = statsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.systemV1beta1, err = systemv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.controlplaneV1beta1 = controlplanev1beta1.NewForConfigOrDie(c)
	cs.controlplaneV1beta2 = controlplanev1beta2.NewForConfigOrDie(c)
	cs.crdV1alpha1 = crdv1alpha1.NewForConfigOrDie(c)
	cs.crdV1alpha2 = crdv1alpha2.NewForConfigOrDie(c)
	cs.crdV1beta1 = crdv1beta1.NewForConfigOrDie(c)
	cs.statsV1alpha1 = statsv1alpha1.NewForConfigOrDie(c)
	cs.systemV1beta1 = systemv1beta1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.controlplaneV1beta1 = controlplanev1beta1.New(c)
	cs.controlplaneV1beta2 = controlplanev1beta2.New(c)
	cs.crdV1alpha1 = crdv1alpha1.New(c)
	cs.crdV1alpha2 = crdv1alpha2.New(c)
	cs.crdV1beta1 = crdv1beta1.New(c)
	cs.statsV1alpha1 = statsv1alpha1.New(c)
	cs.systemV1beta1 = systemv1beta1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
