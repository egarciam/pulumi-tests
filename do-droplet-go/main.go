package main

import (
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a tiny DigitalOcean Droplet
		droplet, err := digitalocean.NewDroplet(ctx, "tiny-droplet", &digitalocean.DropletArgs{
			Name:   pulumi.String("tiny-droplet"),
			Region: pulumi.String("fra1"),               // Or another region slug, e.g., "fra1"
			Size:   pulumi.String("s-1vcpu-512mb-10gb"), // The smallest Droplet size
			Image:  pulumi.String("ubuntu-22-04-x64"),   // A common Ubuntu image
		})
		if err != nil {
			return err
		}

		// Export the Droplet's IP address
		ctx.Export("dropletIp", droplet.Ipv4Address)

		return nil
	})
}
