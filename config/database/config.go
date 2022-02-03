package database

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("digitalocean_database_cluster", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "digitalocean"
		r.ShortGroup = "database"
		r.Kind = "DatabaseCluster"
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("digitalocean_database_connection_pool", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "digitalocean"
		r.ShortGroup = "database"

		r.Kind = "DatabaseConnectionPool"
		r.ExternalName = config.IdentifierFromProvider
		r.References["database"] = config.Reference{
			Type: "DatabaseCluster",
		}
	})

	p.AddResourceConfigurator("digitalocean_database_db", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "digitalocean"
		r.ShortGroup = "database"
		r.Kind = "Database"
		r.ExternalName = config.IdentifierFromProvider
		r.References["database"] = config.Reference{
			Type: "DatabaseCluster",
		}
	})

	p.AddResourceConfigurator("digitalocean_database_firewall", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "digitalocean"
		r.ShortGroup = "database"
		r.Kind = "DatabaseFirewall"
		r.ExternalName = config.IdentifierFromProvider
		r.References["database"] = config.Reference{
			Type: "DatabaseCluster",
		}
		r.References["droplet"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-digitalocean/apis/droplet/v1apha1.Droplet",
		}
	})

	p.AddResourceConfigurator("digitalocean_database_replica", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "digitalocean"
		r.ShortGroup = "database"
		r.Kind = "DatabaseReplica"
		r.ExternalName = config.IdentifierFromProvider
		r.References["database"] = config.Reference{
			Type: "DatabaseCluster",
		}
	})

	p.AddResourceConfigurator("digitalocean_database_user", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "digitalocean"
		r.ShortGroup = "database"
		r.Kind = "DatabaseUser"
		r.ExternalName = config.IdentifierFromProvider
		r.References["database"] = config.Reference{
			Type: "DatabaseCluster",
		}
	})
}
