package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuekcc/fns"
	"github.com/yuekcc/fns/fnspec"
	"sigs.k8s.io/yaml"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "fns",
		Short: "a function service",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install",
			Short: "install a function",
			Run:   install,
		},

		&cobra.Command{
			Use:   "uninstall",
			Short: "kill and uninstall a function",
			Run:   uninstall,
		},

		&cobra.Command{
			Use:   "ps",
			Short: "list all functions",
			Run:   listFunctions,
		},
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func listFunctions(cmd *cobra.Command, args []string) {
	machine := fns.GetMachine()
	defer machine.Shutdown()

	ctx := context.Background()

	containers, err := machine.List(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for i := 0; i < len(containers); i++ {
		c := containers[i]
		fmt.Println(c.ID, c.Labels["fns-route"], c.State)
	}
}

func install(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("require a functionspec file")
	}

	specFile, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	functionSpec := fnspec.FunctionSpec{}
	err = yaml.Unmarshal(specFile, &functionSpec)
	if err != nil {
		log.Fatalln("unable to read yaml file, ", err)
	}

	machine := fns.GetMachine()
	defer machine.Shutdown()

	ctx := context.Background()

	log.Printf("start install function: %s %d\n", functionSpec.Metadata.Name, functionSpec.Metadata.Version)
	id, err := machine.Install(ctx, &functionSpec)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("install success\n")
	log.Printf("launch app\n")
	err = machine.Spawn(ctx, id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s run on %s\n", functionSpec.Metadata.Name, id)
}

func uninstall(cmd *cobra.Command, args []string) {
	machine := fns.GetMachine()
	defer machine.Shutdown()

	ctx := context.Background()

	for _, id := range args {
		err := machine.Uninstall(ctx, id)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("%s uninstalled\n", id)
	}
}
