package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(ServerCmd())
}

func ServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start server",
		Long:  "Start mirror balancer listening on specified port, sending copy requests to each backend",
		RunE: func(cmd *cobra.Command, args []string) error {
			sleepTime, err := cmd.Flags().GetInt("sleep")
			if err != nil {
				return err
			}
			for sleepTime > 0 {
				fmt.Printf("Will sleep for another %d seconds, then exit!\n", sleepTime)
				time.Sleep(time.Second)
				sleepTime -= 1
			}
			return nil
		},
	}

	cmd.Flags().IntP("sleep", "s", 5, "How long to sleep (seconds)?")
	return cmd
}
