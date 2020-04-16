package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// todo Flags
	// imgur bool
	// streamable
	// sendvid
	// ffsend
	// gofile bool

	// Root cmd
	rootCmd = &cobra.Command{
		Use:   "shareit <file_path>",
		Short: "Share media anonymously with ease!",
		Long:  `Share images and videos anonymously with a single command! No need to worry or wonder where the media is being stored, the URL is added directly to your clipboard so just paste and go!`,
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Print the usage if no args are passed in :)
			if len(args) < 1 {
				if err := cmd.Usage(); err != nil {
					log.Fatal(err)
				}

				return
			}

			// shareit!
			shareit(args)
		},
	}
)

// Execute the shareit command
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
