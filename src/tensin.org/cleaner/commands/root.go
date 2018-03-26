package commands

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"time"
    "path/filepath"
	"strings"	
	
	"github.com/spf13/cobra"
	"github.com/karrick/godirwalk"	
)

// RootCmd is the main command = the program itself
var RootCmd = &cobra.Command{
	Use:   "cleaner",
	Short: "cleaner",
	Long:  `cleaner`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		patterns := make([]string, 0)
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		configurationFileName := dir + string(os.PathSeparator) + filepath.Base(os.Args[0]) + ".conf"
		configurationFileInfo, _ := os.Stat(configurationFileName)
		if configurationFileInfo.Mode().IsRegular() {
			fmt.Println("Loading configuration from [" + configurationFileName + "]")
			file, err := os.Open(configurationFileName)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
		
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				patterns = append(patterns, scanner.Text())
			}
		
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}			
		} else {
			patterns = append(patterns, "bin")
			patterns = append(patterns, "target")
			patterns = append(patterns, "*.log")
			patterns = append(patterns, "*.class")
		}

		searchDir := "."
		if Path != "" {
			searchDir = Path
		}

		if _, err := os.Stat(searchDir); os.IsNotExist(err) {
			fmt.Println("Can't analyze [" + searchDir + "], directory not found")
			os.Exit(1)
		}

		fmt.Println("Now cleaning java useless items from [" + searchDir + "] with patterns [" + strings.Join(patterns, ", ") + "]")

		nbRemovedFiles := 0
		nbRemovedDirectories := 0
		err := godirwalk.Walk(searchDir, &godirwalk.Options{
			// Unsorted: true, // set true for faster yet non-deterministic enumeration (see godoc)
			Callback: func(path string, de *godirwalk.Dirent) error {
				// fmt.Printf("%s %s\n", de.ModeType(), osPathname)
				matched := false
				matchedPattern := ""
				for _, pattern := range patterns {
					name := filepath.Base(path)
					m, err := filepath.Match(pattern, name)
					if err != nil {
						fmt.Println(err)
					} 
					if m {
						matched = true
						matchedPattern = pattern
						break
					}
				}

				if matched {
					fmt.Println("- [DELETE] path [" + path + "] (matching pattern is [" + matchedPattern + "])")
				} else if Debug && ! matched {
					fmt.Println("- [KEEP]   path [" + path + "]")
				}

				if matched {
					fullPath, _ := filepath.Abs(path)
					switch mode := de.ModeType(); {
					case mode.IsDir():
						// fmt.Println("Remove directory [" + path + "]")
						os.RemoveAll(fullPath)
						nbRemovedDirectories++
						return filepath.SkipDir					
					case mode.IsRegular():
						// fmt.Println("Remove file [" + path + "]")
						os.Remove(fullPath)
						nbRemovedFiles++
						return nil
					}
				}

				return nil
			},
			ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
				// Your program may want to log the error somehow.
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	
				// For the purposes of this example, a simple SkipNode will suffice,
				// although in reality perhaps additional logic might be called for.
				return godirwalk.SkipNode
			},
		})
		elapsed := time.Since(start)
		fmt.Printf("Execution took %s, results are :", elapsed)
		fmt.Println("\n - " + strconv.Itoa(nbRemovedDirectories) + " removed directories\n - " + strconv.Itoa(nbRemovedFiles) + " removed files")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	  },
}

var Path string
var Debug bool

func init() {
	cobra.MousetrapHelpText = ""
	RootCmd.PersistentFlags().StringVarP(&Path, "path", "", "", "The path to analyze. Default is current folder (same value than `$(pwd)`)")
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "Is debug activated (false by default)")
}
