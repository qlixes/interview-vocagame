package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "vocagame",
}

var poolCmd = &cobra.Command{
	Use:  "create_parking_lot [total_lot]",
	Run:  poolHandle,
}

var parkingCmd = &cobra.Command{
	Use:  "park [car_id]",
	Run:  parkingHandle,
}

var leavingCmd = &cobra.Command{
	Use:  "leave [car_id] [duration]",
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
		log.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(poolCmd)
	rootCmd.AddCommand(parkingCmd)
	rootCmd.AddCommand(leavingCmd)
	rootCmd.AddCommand(statusCmd)
}


func poolHandle(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
	}

	re, err := regexp.Compile(`^[0-9]+$`)

	if err != nil {
		log.Fatalf("Invalid regex : %v \n", err)
	}

	validate := re.MatchString(args[0])

	if !validate {
		log.Fatalf("Args incorrect")
	}

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
