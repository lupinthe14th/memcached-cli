package cmd

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/cobra"
)

var (
	hostip   string
	hostport int
	repert   uint
	interval time.Duration
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping for memcached",
	RunE: func(cmd *cobra.Command, args []string) error {
		for i := repert; i > 0; i-- {
			if err := ping(); err != nil {
				return err
			}
			time.Sleep(interval)
		}
		return nil
	},
}

func init() {
	pingCmd.PersistentFlags().StringVarP(&hostip, "host", "H", "127.0.0.1", "Server hostname")
	pingCmd.PersistentFlags().IntVarP(&hostport, "port", "p", 11211, "Server port")
	pingCmd.PersistentFlags().UintVarP(&repert, "repert", "r", 1, "Execute specified command N times.")
	pingCmd.PersistentFlags().DurationVarP(&interval, "interval", "i", 1*time.Second,
		"When -r is used, waits <interval> duration per command.")
}

func ping() error {
	mc := memcache.New(fmt.Sprintf("%s:%d", hostip, hostport))
	if err := mc.Ping(); err != nil {
		return err
	}
	fmt.Println("PONG")
	return nil
}
