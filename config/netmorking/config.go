package compute

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("openstack_networking_network_v2", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("openstack_networking_addressscope_v2", func(r *config.Resource) {
	})
}
