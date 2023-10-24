/*
Copyright © 2023 Bernard Cooke

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bernardcooke53/textwave/wave"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "textwave [text]",
	Short: "Convert text to waves for meme value",
	Long: `Convert text on the command line to waves
  for added emphasis. Especially useful in professional settings where
	remote work mandates that your text-formatting skills evolve rapidly
	to ensure your continued effectiveness at communication.
  `,
	Run:  textwaveMain,
	Args: cobra.MaximumNArgs(1),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type RootCmdOptions struct {
	SpongebobMocking, AllCaps   bool
	ColumnSize, NumberOfColumns int
}

var rootCmdOptions RootCmdOptions = RootCmdOptions{}

func init() {
	// cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().
	// 	StringVar(&cfgFile, "config", "", "config file (default is $HOME/.textwave.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().
		BoolVarP(&rootCmdOptions.SpongebobMocking, "mock", "m", false, "hOw dO i UsE ThIs FlAg?")

	rootCmd.PersistentFlags().
		BoolVarP(&rootCmdOptions.AllCaps, "all-caps", "U", false, "Output in all caps for EXTRA EMPHASIS")

	rootCmd.PersistentFlags().
		IntVarP(&rootCmdOptions.ColumnSize, "column-size", "w", 1, "Set a max width if you want a skinny wave")

	rootCmd.PersistentFlags().
		IntVarP(&rootCmdOptions.NumberOfColumns, "columns", "n", 0, "Set a max width if you want a skinny wave")
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// if cfgFile != "" {
// 	// Use config file from the flag.
// 	viper.SetConfigFile(cfgFile)
// } else {
// 	// Find home directory.
// 	home, err := os.UserHomeDir()
// 	cobra.CheckErr(err)
//
// 	// Search config in home directory with name ".textwave" (without extension).
// 	viper.AddConfigPath(home)
// 	viper.SetConfigType("yaml")
// 	viper.SetConfigName(".textwave")
// }

// viper.AutomaticEnv() // read in environment variables that match
//
// // If a config file is found, read it in.
// if err := viper.ReadInConfig(); err == nil {
// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// }
// }

func textwaveMain(cmd *cobra.Command, args []string) {
	var inputText string
	if len(args) >= 1 {
		inputText = args[0]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		stdin := scanner.Text()
		if len(stdin) == 0 {
			_ = cmd.Usage()
			os.Exit(1)
		}
		inputText = stdin
	}

	waveMaker := wave.WaveMaker{
		SpongebobMocking: rootCmdOptions.SpongebobMocking && !rootCmdOptions.AllCaps,
		AllCaps:          rootCmdOptions.AllCaps,
		ColumnSize:       rootCmdOptions.ColumnSize,
		NumberOfColumns:  rootCmdOptions.NumberOfColumns,
	}

	lines := waveMaker.MakeWave(inputText)

	for _, line := range lines {
		fmt.Println(line)
	}
}
