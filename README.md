# java-folders-cleaner

`java-folders-cleaner` is a standalone program allowing to easily and automatically remove useless files from :

- eclipse workspaces
- eclipse projects
- eclipse installations

Files to remove are customizables.

Benefits are for example :

- remove OSGI temporary files in order to have an eclipse installation faster to start
- remove useless files in order to retrieve some disk space usage

Program is written in `golang` and is working fine under Windows or Linux.

## Easy launch 

Just put the .exe and its associated configuration file `launcher.conf` in your workspace folder and execute it (through double-click, or from a `cmd.exe` command line execution).

## Advanced launch

Launch manually the tool with (`.bat` or from `cmd.exe`) with :

```
java-folders-cleaner --folders <folder1>,<folder2>,...
```

## Configuration

Configuration file has just to be located in the same folder than the executable, with the same name ending by `.conf`.

![Directory content](java-folders-cleaner-setup-directory-content.png)

Example of java configuration :

```
**/*.tmp

[projects]
**/bin
**/build
**/classes
**/web/WEB-INF/classes
**/dist
**/target
**/*.class
**/*.log
# **/.cache
# **/work

[workspaces]
**/.metadata/.plugins/org.eclipse.core.resources/.history
**/.metadata/.plugins/org.eclipse.jdt.core/*

[eclipse]
**/*.log
**/configuration/org.eclipse.osgi/*

# This will define additional paths to purge
[paths]
c:\Tools\eclipse\
c:\Workspaces\
```

First lines (the ones without sections) are applyed everywhere (whatever the kind of detected folder, and in addition to patterns declared in a section).

Comments are possible through lines starting with `#`.

Possible sections are :

- `projects` : folder containing an eclipse project, that means containing a `.project` or `.classpath` file
- `workspaces` : folder containing an eclipse workspace, that means containing a `.metadata` sub-folder
- `eclipse` : folder containing an eclipse installation, that means containing an `eclipse.exe` file

Additionnaly a `[paths]` section may also be defined to specify which full paths have to be purged (will be merged with `--folders` values, if defined).

## Usage

```bash
cleaner

Usage:
  cleaner [flags]
  cleaner [command]

Available Commands:
  help        Help about any command
  version     Prints the version of the chainer command.

Flags:
      --debug             Is debug activated (false by default)
      --folders string    Folders to clean, separated by commas. Current folder will be used as a default, if --folders is not defined.
  -h, --help              help for cleaner
      --patterns string   Override default patterns (only used if no configuration file found). Separate values with commas.

Use "cleaner [command] --help" for more information about a command.

```

## Log output

Example :

```
$ ./cleaner.exe --folders "/c/Tools/eclipse/"
Detected folders to purge are [C:/Tools/eclipse/]
Loading configuration from [C:\Tools\java-folders-cleaner\cleaner.conf]
Patterns taken in account are :
 - workspaces
   - **\.metadata\.plugins\org.eclipse.core.resources\.history
   - **\.metadata\.plugins\org.eclipse.jdt.core\*
 - eclipse
   - **\*.log
   - **\configuration\org.eclipse.osgi\*
 - default
   - **\*.tmp
 - projects
   - **\bin
   - **\build
   - **\classes
   - **\web\WEB-INF\classes
   - **\dist
   - **\target
   - **\work
   - **\*.class
   - **\*.log
Now cleaning java useless items from [C:/Tools/eclipse/]
- [BROWSE]  path [C:\Tools\eclipse\2018-12-R\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [BROWSE]  path [C:\Tools\eclipse\indigo\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [BROWSE]  path [C:\Tools\eclipse\luna\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [BROWSE]  path [C:\Tools\eclipse\mars\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [BROWSE]  path [C:\Tools\eclipse\neon\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [BROWSE]  path [C:\Tools\eclipse\oxygen\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [BROWSE]  path [C:\Tools\eclipse\photon\java] (detected as type [eclipse]) purge patterns will be [**\*.tmp, **\*.log, **\configuration\org.eclipse.osgi\*]
- [DELETED] path [C:\Tools\eclipse\photon\java\configuration\org.eclipse.osgi\138] (matched pattern is [**\configuration\org.eclipse.osgi\*])
- [DELETED] path [C:\Tools\eclipse\photon\java\configuration\org.eclipse.osgi\150] (matched pattern is [**\configuration\org.eclipse.osgi\*])
- [DELETED] path [C:\Tools\eclipse\photon\java\configuration\org.eclipse.osgi\151] (matched pattern is [**\configuration\org.eclipse.osgi\*])
(...)
Execution took 4m17.2027311s, went from 68085 MiB to 68106 MiB free space, results are :
 - 1006 removed directories
 - 12 removed files
```

# Development

## How to develop

On the host :

- Start a docker container : `make docker`

Inside the container :

- Compile and create linux binary : `make quick`
- Build whole distribution : `make clean distribution`

## TODO

- [x] New `--patterns` flag to override default flags
- [x] Dump disk space retrieved
- [x] Handle by default an external file with patterns ("<binary_name>.conf", with one pattern per line)
- [x] Have patterns by kind of folders (eclipse installation, workspace, ...)
- [x] Allow several paths to be configured at the same time, through `--folders <folder1>,<folder2>,...`
- [ ] Better logging system instead of standard output
- [ ] Improve performances : @see https://github.com/golang/go/issues/16399 & https://github.com/s12chung/fastwalk
