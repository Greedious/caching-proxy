package cmd

import (
	"caching-proxy/internal/proxy"
	"caching-proxy/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	Port       string
	Origin     string
	ClearCache bool
)

var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Short: "caching proxy server, it handles forwarding http requests, and caching received content from origin",
	Long: `This is a caching proxy server which act as a reverse proxy, it receives client, requests, forward it
	and cache data (if it is not cached) and returning result to the client.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateFlags(cmd); err != nil {
			return err
		}
		return runCachingProxy(cmd)
	},
}

func runCachingProxy(cmd *cobra.Command) error {
	proxy := proxy.GetNewProxy(Origin)
	http.HandleFunc("/", proxy.RequestHandler)
	port, _ := cmd.Flags().GetString("port")
	connAddress := ":" + port
	fmt.Printf("caching proxy is running on port %s...\n", port)
	return http.ListenAndServe(connAddress, nil)
}

func validateFlags(cmd *cobra.Command) error {
	// Validate port: must be a valid integer between 1 and 65535
	portNum, err := strconv.Atoi(Port)
	if err != nil || portNum < 1 || portNum > 65535 {
		return fmt.Errorf("invalid port: %s. Must be a number between 1 and 65535", Port)
	}
	if isPortTaken := utils.IsPortTaken(Port); isPortTaken {
		return fmt.Errorf("error: port %s is already in use", Port)
	}
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&Port, "port", "p", "", "port which the caching proxy will run on.")
	rootCmd.Flags().StringVarP(&Origin, "origin", "o", "", "the origin which the traffic will be forwarded to.")
	rootCmd.MarkFlagRequired("port")
	rootCmd.MarkFlagRequired("origin")
}
