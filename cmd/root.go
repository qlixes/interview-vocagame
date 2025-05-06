package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"

	"github.com/qlixes/vocagame/provider"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "vocagame",
}

var poolCmd = &cobra.Command{
	Use:  "create_parking_lot [total_lot]",
	Args: cobra.ExactArgs(1),
	Run:  poolHandle,
}

var parkingCmd = &cobra.Command{
	Use:  "park [car_id]",
	Args: cobra.ExactArgs(1),
	Run:  parkingHandle,
}

var leavingCmd = &cobra.Command{
	Use:  "leave [car_id] [duration]",
	Args: cobra.ExactArgs(2),
	Run:  leavingHandle,
}

var statusCmd = &cobra.Command{
	Use: "status",
	Run: statusHandle,
}

var wg sync.WaitGroup
var mu sync.Mutex

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	initConfig()
	rootCmd.AddCommand(poolCmd)
	rootCmd.AddCommand(parkingCmd)
	rootCmd.AddCommand(leavingCmd)
	rootCmd.AddCommand(statusCmd)
}

func initConfig() {
	fmt.Println(provider.Cfg)
}

func poolHandle(cmd *cobra.Command, args []string) {
	re, err := regexp.Compile(`^[0-9]+$`)

	if err != nil {
		log.Fatalf("Invalid regex : %v \n", err)
	}

	validate := re.MatchString(args[0])

	if !validate {
		log.Fatalf("Args incorrect")
	}

	limit, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalf("Failed parseint : %v \n", err.Error())
	}

}

func parkingHandle(cmd *cobra.Command, args []string) {

}

func leavingHandle(cmd *cobra.Command, args []string) {

}

func statusHandle(cmd *cobra.Command, args []string) {

}
