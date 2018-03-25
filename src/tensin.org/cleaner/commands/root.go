package commands

import (
	"fmt"
	"os"
    "path/filepath"
	"strings"	
	
	"github.com/spf13/cobra"
)

/*
func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}
*/

// RootCmd is the main command = the program itself
var RootCmd = &cobra.Command{
	Use:   "java-folders-cleaner",
	Short: "java-folders-cleaner",
	Long:  `java-folders-cleaner`,
	Run: func(cmd *cobra.Command, args []string) {
		searchDir := "."
		patterns := make([]string, 0)
		/*
		*/
		patterns = append(patterns, "bin")
		patterns = append(patterns, "target")
		patterns = append(patterns, "*.log")
		patterns = append(patterns, "*.class")

		fmt.Println("Now cleaning java useless items from [" + searchDir + "] with patterns [" + strings.Join(patterns, ", ") + "]")

		fileList := make([]string, 0)
		e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return err
		})

		if e != nil {
			panic(e)
		}

		for _, path := range fileList {
			// fmt.Println("path [" + path + "]")
			for _, pattern := range patterns {
				matched, err := filepath.Match(pattern, path)
				if err != nil {
					fmt.Println(err)
				} 
				if matched {
					// fmt.Println("- pattern [" + pattern + "] / [" + path + "]")
				} else {
					// fmt.Println(". pattern [" + pattern + "] / [" + path + "]")
				}
				if matched {
					fullPath, _ := filepath.Abs(path)
					fi, _ := os.Stat(fullPath)
					switch mode := fi.Mode(); {
					case mode.IsDir():
						fmt.Println("Remove directory [" + path + "]")
						os.RemoveAll(fullPath)
					case mode.IsRegular():
						fmt.Println("Remove file [" + path + "]")
						os.Remove(fullPath)
					}					
				}
			}
		}

		// os.RemoveAll()
	  },
}

func init() {
	// RootCmd.PersistentFlags().StringVarP(&ConfigurationFileName, "configuration", "", "watchthatpage.json", "Configuration file name. Default is binary name + .json (e.g. 'watchthatpage.json'), in the same folder than the binary itself")
	// daemon.Server.Start()
	// fmt.Println("start")
}
