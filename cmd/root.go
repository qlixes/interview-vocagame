package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/qlixes/vocagame/manager"
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(poolCmd)
	rootCmd.AddCommand(parkingCmd)
	rootCmd.AddCommand(leavingCmd)
	rootCmd.AddCommand(statusCmd)
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

	for i := 1; i <= limit; i++ {
		key := fmt.Sprintf("pool_%d", i)
		manager.Mgr.Update(key, "")
		fmt.Printf("Allocated slot number: %d \n", i)
	}
}

func parkingHandle(cmd *cobra.Command, args []string) {

}

func leavingHandle(cmd *cobra.Command, args []string) {

}

func statusHandle(cmd *cobra.Command, args []string) {

}
