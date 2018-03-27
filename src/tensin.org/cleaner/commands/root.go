package commands

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"time"
	"runtime"
    "path/filepath"
	"strings"	
	
	"github.com/spf13/cobra"
	"github.com/karrick/godirwalk"	
    "github.com/shirou/gopsutil/disk"
)

func diskUsage() string {
	path, err := os.Getwd()
	if err != nil {
		return "<undefined>"
	}

	if runtime.GOOS == "windows" {
		p := strings.Index(path, "\\")
		if p != -1 {
			path = path[:p+1]
		}
	} else {
		p := strings.Index(path, "/")
		if p != -1 {
			path = path[:p+1]
		}
	}
	u, err := disk.Usage(path)
	if err != nil {
		return "<undefined>"
	}
	return strconv.FormatUint(u.Free /1024/1024, 10) + " MiB"
}

// RootCmd is the main command = the program itself
var RootCmd = &cobra.Command{
	Use:   "cleaner",
	Short: "cleaner",
	Long:  `cleaner`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		freeSpaceBefore := diskUsage()

		patterns := make([]string, 0)
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		configurationFileName := dir + string(os.PathSeparator) + strings.TrimSuffix(filepath.Base(os.Args[0]),".exe") + ".conf"

		// if configurationFileInfo.Mode().IsRegular() {
                if _, err := os.Stat(configurationFileName); err == nil {
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
		} else if Patterns != "" {
			fmt.Println("Configuration [" + configurationFileName + "] not found, applying patterns provided through command line")
			for _, pattern := range strings.Split(Patterns, ",") {
				patterns = append(patterns, strings.Trim(pattern, " "))
			}
		} else {
			fmt.Println("Configuration [" + configurationFileName + "] not found, applying default java patterns")
			patterns = append(patterns, "bin")
			patterns = append(patterns, "target")
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
					fmt.Println("- [DELETED] path [" + path + "] (matching pattern is [" + matchedPattern + "])")
				} else if Debug && ! matched {
					fmt.Println("- [KEPT]    path [" + path + "]")
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
		freeSpaceAfter := diskUsage()
		fmt.Printf("Execution took %s, went from %s to %s free space, results are :", elapsed, freeSpaceBefore, freeSpaceAfter)
		fmt.Println("\n - " + strconv.Itoa(nbRemovedDirectories) + " removed directories\n - " + strconv.Itoa(nbRemovedFiles) + " removed files")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	  },
}

var Path string
var Debug bool
var Patterns string

func init() {
	cobra.MousetrapHelpText = ""
	RootCmd.PersistentFlags().StringVarP(&Path, "path", "", "", "The path to analyze. Default is current folder (same value than `$(pwd)`)")
	RootCmd.PersistentFlags().StringVarP(&Patterns, "patterns", "", "", "Override default patterns (only used if no configuration file found). Separate values with commas.")
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "Is debug activated (false by default)")
}
