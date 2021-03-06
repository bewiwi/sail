package application

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/runabove/sail/internal"
)

var cmdApplicationDomain = &cobra.Command{
	Use:     "domain",
	Short:   "Application Domain commands: sail application domain --help",
	Long:    `Application Domain commands: sail application domain <command>`,
	Aliases: []string{"domains"},
}

var cmdApplicationDomainList = &cobra.Command{
	Use:     "list",
	Short:   "List domains and routes on the HTTP load balancer: sail application domain list <applicationName>",
	Aliases: []string{"ls", "ps"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see sail application domain list --help")
		} else {
			// cmdApplicationDomainList TODO ? Tab view with headers ['DOMAIN', 'SERVICE', 'METHOD', 'PATTERN']
			internal.FormatOutputDef(internal.GetWantJSON(fmt.Sprintf("/applications/%s/attached-domains", args[0])))
		}
	},
}

var cmdApplicationDomainAttach = &cobra.Command{
	Use:     "attach",
	Short:   "Attach a domain on the HTTP load balancer: sail application domain attach <applicationName> <domainName>",
	Aliases: []string{"add"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see sail application domain attach --help")
		} else {
			internal.FormatOutputDef(internal.PostWantJSON(fmt.Sprintf("/applications/%s/attached-domains/%s", args[0], args[1])))
		}
	},
}

var cmdApplicationDomainDetach = &cobra.Command{
	Use:     "detach",
	Short:   "Detach a domain from the HTTP load balancer: sail application domain detach <applicationName> <domainName>",
	Aliases: []string{"add"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "Invalid usage. Please see sail application domain attach --help")
		} else {
			internal.FormatOutputDef(internal.DeleteWantJSON(fmt.Sprintf("/applications/%s/attached-domains/%s", args[0], args[1])))
		}
	},
}
