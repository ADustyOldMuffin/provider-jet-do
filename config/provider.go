/*
Copyright 2021 The Crossplane Authors.

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

package config

import (
	"github.com/crossplane-contrib/provider-jet-digitalocean/config/database"
	"github.com/crossplane-contrib/provider-jet-digitalocean/config/droplet"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourcePrefix = "digitalocean"
	modulePath     = "github.com/crossplane-contrib/provider-jet-digitalocean"
)

// GetProvider returns provider configuration
func GetProvider(resourceMap map[string]*schema.Resource) *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProvider(resourceMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			"digitalocean_droplet$",
			"digitalocean_database_cluster$",
			"digitalocean_database_connection_pool$",
			"digitalocean_database_db$",
			"digitalocean_database_firewall$",
			"digitalocean_database_replica$",
			"digitalocean_database_user$",
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		droplet.Configure,
		database.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
