package app

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func NewCompletionCommand(f Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "completion",
		Short: "Create completion script for kubectl (or alias)",
		Run: func(cmd *cobra.Command, args []string) {
			RunAlias(f, cmd, args)
		},
	}

	AddCommonFlags(cmd)
	cmd.Flags().String("kubectl-alias", "kubectl", "Alias of your kubectl command")
	cmd.Flags().String("kubemrr-path", "kubemrr", "Path to the kubemrr command, if it is outside $PATH variable")

	return cmd
}

func RunAlias(f Factory, cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		fmt.Fprintf(f.StdErr(), "Arguments are not expected. Use flags")
		return
	}

	var err error
	c := replacement{
		kubectlAlias:   "kubectl",
		kubemrrPort:    33033,
		kubemrrAddress: "0.0.0.0",
		kubemrrPath:    "kubemrr",
	}

	if c.kubemrrPort, err = cmd.Flags().GetInt("port"); err != nil {
		log.Fatal(err)
	}
	if c.kubemrrAddress, err = cmd.Flags().GetString("address"); err != nil {
		log.Fatal(err)
	}
	if c.kubectlAlias, err = cmd.Flags().GetString("kubectl-alias"); err != nil {
		log.Fatal(err)
	}
	if c.kubemrrPath, err = cmd.Flags().GetString("kubemrr-path"); err != nil {
		log.Fatal(err)
	}

	in := bash_template
	in = strings.Replace(in, "[[kubectl_alias]]", c.kubectlAlias, -1)
	in = strings.Replace(in, "[[kubemrr_path]]", c.kubemrrPath, -1)
	in = strings.Replace(in, "[[kubemrr_address]]", c.kubemrrAddress, -1)
	in = strings.Replace(in, "[[kubemrr_port]]", strconv.Itoa(c.kubemrrPort), -1)

	fmt.Fprint(f.StdOut(), in)
}

type replacement struct {
	kubectlAlias   string
	kubemrrPort    int
	kubemrrAddress string
	kubemrrPath    string
}