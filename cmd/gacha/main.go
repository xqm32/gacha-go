package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xqm32/gacha-go/v2"
)

var rootCmd = &cobra.Command{
	Use:   "gacha",
	Short: "Genshin Impact Gacha Simulator",
	Run: func(cmd *cobra.Command, args []string) {
		times, errTimes := cmd.Flags().GetInt("times")
		charsUp, errCharsUp := cmd.Flags().GetInt("chars-up")
		charsPity, errCharsPity := cmd.Flags().GetInt("chars-pity")
		weapsUp, errWeapsUp := cmd.Flags().GetInt("weaps-up")
		weapsPity, errWeapsPity := cmd.Flags().GetInt("weaps-pity")
		verbose, _ := cmd.Flags().GetBool("verbose")
		if errors.Join(errTimes, errCharsUp, errCharsPity, errWeapsUp, errWeapsPity) != nil {
			os.Exit(1)
		}

		g := &gacha.Gacha{}
		g.SetPity(charsPity, weapsPity)

		if verbose {
			g.OnCharUp = func(g *gacha.Gacha) { fmt.Printf("  UP CHAR %4d %4d\n", g.Pulls(), g.CharAt) }
			g.OnChar = func(g *gacha.Gacha) { fmt.Printf("DOWN CHAR %4d %4d\n", g.Pulls(), g.CharAt) }
			g.OnCharSpec = func(g *gacha.Gacha) { fmt.Printf("LIGH CHAR %4d %4d\n", g.Pulls(), g.CharAt) }
			g.OnWeapUp = func(g *gacha.Gacha) { fmt.Printf("  UP WEAP %4d %4d\n", g.Pulls(), g.WeapAt) }
			g.OnWeapSpec = func(g *gacha.Gacha) { fmt.Printf("ANOT WEAP %4d %4d\n", g.Pulls(), g.WeapAt) }
			g.OnWeap = func(g *gacha.Gacha) { fmt.Printf("DOWN WEAP %4d %4d\n", g.Pulls(), g.WeapAt) }
		}

		pulls := make(chan int)
		for range times {
			t := *g
			go func() {
				pulls <- t.PullUp(charsUp, weapsUp).Pulls()
				if verbose {
					fmt.Println()
				}
			}()
		}

		sum := 0
		for range times {
			sum += <-pulls
		}

		fmt.Printf("%f\n", float64(sum)/float64(times))
	},
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().IntP("times", "t", 1, "Gacha times")
	rootCmd.Flags().IntP("chars-up", "c", 1, "5 star characters up")
	rootCmd.Flags().IntP("chars-pity", "C", 0, "5 star character pity")
	rootCmd.Flags().IntP("weaps-up", "w", 0, "5 star weapons up")
	rootCmd.Flags().IntP("weaps-pity", "W", 0, "5 star weapon pity")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
