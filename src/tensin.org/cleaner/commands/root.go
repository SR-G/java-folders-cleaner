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
	"io/ioutil"
	
	"github.com/spf13/cobra"
	"github.com/shirou/gopsutil/disk"
	"github.com/bmatcuk/doublestar"
	// "tensin.org/cleaner/core/path"
)

func diskUsage(path string) string {
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

func buildFoldersToProcess(folders string) []string {
	results := make([]string, 0)
	if (folders == "") {
		results = append(results, "./")
	} else {
		for _, folder := range strings.Split(folders, ",") {
			results = append(results, strings.TrimSpace(folder))
		}
	}
	return results
}

func retrieveCurrentType(currentFolder string) string {
	result := ""

	// No types defined, let's see if we can auto-discover where we are : eclipse installation or workspace
	files, _ := ioutil.ReadDir(currentFolder)
	for _, f := range files {
		s := f.Name()
		if (s == "eclipse.exe") {
			result = "eclipse"
			break
		} else if (s == ".metadata") {
			result = "workspaces"
			break
		} else if (s == ".classpath" || s == ".project" ) {
			result = "projects"
			break
		}
	}			

	return result
}

func contains(array []string, searched string) bool {
    for _, item := range array {
        if item == searched {
            return true
        }
    }
    return false
}

func sanitizePathPerOS(s string) string {
	return strings.Replace(s, OTHER_PATH_SEPARATOR, PATH_SEPARATOR, -1)
}

func purgeDirectoryContent(currentFolder string, allPatterns map[string][]string, patterns []string) {
	currentPatterns := make([]string, 0)
	for _, pattern := range patterns {
		if (!contains(currentPatterns, pattern)) {
			currentPatterns = append(currentPatterns, pattern)
		}
	}

	file, _ := os.Open(currentFolder)
	defer file.Close()
	if fi, err := file.Stat(); err != nil || fi.IsDir() {
		currentType := retrieveCurrentType(currentFolder)
		if (currentType != "") {
			for _, pattern := range retrievePatternsForType(currentType, allPatterns) {
				if (!contains(currentPatterns, pattern)) {
					currentPatterns = append(currentPatterns, pattern)
				}
			}
			fmt.Println("- [BROWSE]  path [" + currentFolder + "] (detected as type [" + currentType + "]) purge patterns will be [" + strings.Join(currentPatterns, ", ") + "]")
		}
	}

	files, _ := ioutil.ReadDir(currentFolder)
	for _, path := range files {
		fullPath, _ := filepath.Abs(filepath.Join(currentFolder, path.Name()))

		matched := false
		matchedPattern := ""
		for _, pattern := range currentPatterns {
			name := sanitizePathPerOS(fullPath)
			// fmt.Println(" >> " + pattern + " | " + name)
			// m, err := filepath.Match(pattern, name)
			m, err := doublestar.PathMatch(pattern, name)
			if err != nil {
				fmt.Println(err)
			} 
			if m {
				matched = true
				matchedPattern = pattern
				break
			}
		}

		file, _ := os.Open(fullPath)
		fi, _ := file.Stat();
		file.Close()			
		if matched {
			if (fi.IsDir()) {
				fmt.Println("- [DELETED] path [" + fullPath + "] (matched pattern is [" + matchedPattern + "])")
				err := os.RemoveAll(fullPath)
				if err != nil {
					fmt.Println("  [ERROR]   path [" + fullPath + "] : " + err.Error())
				} else {
					nbRemovedDirectories++
				}
			} else {
				fmt.Println("- [DELETED] file [" + fullPath + "] (matched pattern is [" + matchedPattern + "])")
				err := os.RemoveAll(fullPath)
				if err != nil {
					fmt.Println("  [ERROR]   path [" + fullPath + "] : " + err.Error())
				} else {
					nbRemovedFiles++
				}
			}
		} else {
			if (fi.IsDir()) {
				if Debug {
					fmt.Println("- [RECURSE] path [" + fullPath + "]")
				}
				purgeDirectoryContent(fullPath, allPatterns, currentPatterns)
			} else {
				if Debug {
					fmt.Println("- [KEPT]    file [" + fullPath + "]")
				}				
			}
		}
	}
}

func loadAllPatternsFromConfiguration(configurationFileName string) map[string][]string {
	patterns := make(map[string][]string)
	if _, err := os.Stat(configurationFileName); err == nil {
		fmt.Println("Loading configuration from [" + configurationFileName + "]")
		file, err := os.Open(configurationFileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	
		scanner := bufio.NewScanner(file)
		currentSection := "default"
		for scanner.Scan() {
			currentLine := strings.TrimSpace(scanner.Text())
			if (currentLine != "") {
				if (strings.HasPrefix(currentLine, "[")) {
					// Header / section found
					currentSection = strings.Replace(currentLine, "[", "", 1)
					currentSection = strings.Replace(currentSection, "]", "", 1)
					currentSection = strings.TrimSpace(currentSection)

					// Remove current line as it is a header [...]
					currentLine = ""
				}
				
				if (currentLine != "") {
					patterns[currentSection] = append(patterns[currentSection], sanitizePathPerOS(currentLine))
				}
			}
		}
	
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}			
	} else if Patterns != "" {
		fmt.Println("Configuration [" + configurationFileName + "] not found, applying patterns provided through command line")
		for _, pattern := range strings.Split(Patterns, ",") {
			patterns["default"] = append(patterns["default"], strings.Trim(pattern, " "))
		}
	} else {
		fmt.Println("Configuration [" + configurationFileName + "] not found, applying default java patterns")
		patterns["default"] = append(patterns["default"], "bin")
		patterns["default"] = append(patterns["default"], "target")
	}

	fmt.Println("Patterns taken in account are : ")
	for key, value := range patterns {
		fmt.Println(" - " + key)
		for _, p := range value {
			fmt.Println("   - " + p)
		}
	}

	return patterns
}

func retrievePatternsForType(currentType string, allPatterns map[string][]string) []string {
	results := make([]string, 0)
	for _, pattern := range allPatterns["default"] {
		results = append(results, pattern)
	}
	if (currentType != "") {
		for _, pattern := range allPatterns[currentType] {
			results = append(results, pattern)
		}
	}
	return results
}

// RootCmd is the main command = the program itself
var RootCmd = &cobra.Command{
	Use:   "cleaner",
	Short: "cleaner",
	Long:  `cleaner`,
	Run: func(cmd *cobra.Command, args []string) {

		folders := buildFoldersToProcess(Folders)
		fmt.Println("Detected folders to purge are [" + strings.Join(folders, ", ") + "]")
		
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		configurationFileName := dir + string(os.PathSeparator) + strings.TrimSuffix(filepath.Base(os.Args[0]),".exe") + ".conf"

		// allPatterns["default"] = default / generic patterns
		// allPattersn["<section>"] = patterns loaded from [<section>] from the configuration file
		allPatterns := loadAllPatternsFromConfiguration(configurationFileName)

		for _, currentFolder := range folders {
			fmt.Println("Now cleaning java useless items from [" + currentFolder + "]")

			start := time.Now()
			freeSpaceBefore := diskUsage(currentFolder)
			purgeDirectoryContent(currentFolder, allPatterns, allPatterns["default"]);
			freeSpaceAfter := diskUsage(currentFolder)
			elapsed := time.Since(start)

			fmt.Printf("Execution took %s, went from %s to %s free space, results are :", elapsed, freeSpaceBefore, freeSpaceAfter)
			fmt.Println("\n - " + strconv.Itoa(nbRemovedDirectories) + " removed directories\n - " + strconv.Itoa(nbRemovedFiles) + " removed files")
		}
		os.Exit(1)
	},
}

// var Path string
var Debug bool
var Patterns string
var Folders string

var nbRemovedFiles int 
var nbRemovedDirectories int

func init() {
	cobra.MousetrapHelpText = ""
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "Is debug activated (false by default)")
	RootCmd.PersistentFlags().StringVarP(&Patterns, "patterns", "", "", "Override default patterns (only used if no configuration file found). Separate values with commas.")
	RootCmd.PersistentFlags().StringVarP(&Folders, "folders", "", "", "Folders to clean, separated by commas. Current folder will be used as a default, if --folders is not defined.")
}
