/*
Copyright 2017 Caicloud authors. All rights reserved.

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

package provider

import (
	"github.com/caicloud/loadbalancer-controller/pkg/plugin"
	"github.com/caicloud/loadbalancer-controller/pkg/provider/azure"
	"github.com/caicloud/loadbalancer-controller/pkg/provider/external"
	"github.com/caicloud/loadbalancer-controller/pkg/provider/f5"
	"github.com/caicloud/loadbalancer-controller/pkg/provider/ipvsdr"
)

var localRegistryBuilder = plugin.RegistryBuilder{
	azure.AddToRegistry,
	external.AddToRegistry,
	ipvsdr.AddToRegistry,
	f5.AddToRegistry,
}

var AddToRegistry = localRegistryBuilder.AddToRegistry
