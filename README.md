# java-folders-cleaner

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
      --debug         Is debug activated (false by default)
  -h, --help          help for cleaner
      --path $(pwd)   The path to analyze. Default is current folder (same value than $(pwd))

Use "cleaner [command] --help" for more information about a command.
```

## Log output

Example :

```
$ ../saig/trunk/cleaner.exe
Loading configuration from [C:\Users\sersimon\workspaces\projects-psa\saig\trunk\cleaner.conf]
Now cleaning java useless items from [.] with patterns [.cache, bin, build, classes, dist, target, work, *.class]
- [DELETE] path [trunk\pmm-archetype-occurrences\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pmm-core\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pmm-core\work] (matching pattern is [work])
- [DELETE] path [trunk\pmm-logplayer\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pmm-site\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pmm-war-backoffice\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pmm-war-webapp\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pmm-web-services-forcagescheminement\WebContent\WEB-INF\classes] (matching pattern is [classes])
- [DELETE] path [trunk\pmm-webapp\web\WEB-INF\classes] (matching pattern is [classes])
- [DELETE] path [trunk\pyr01\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr02\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr03\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr04\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr05\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr06\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr07\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr08\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr29\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr66\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyr67\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrm1\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrm2\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrmu\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrpy\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrrj\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrrj\work] (matching pattern is [work])
- [DELETE] path [trunk\pyrsx\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrta\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrv1\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrv2\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrvf\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrvh\bin] (matching pattern is [bin])
- [DELETE] path [trunk\pyrvl\bin] (matching pattern is [bin])
Execution took 34.9460339s, went from 9176 MiB to 9265 MiB free space, results are :
- 33 removed directories
- 0 removed files
```

# Development

## TODO

- [x] New `--patterns` flag to override default flags
- [ ] Better logging system instead of standard output
- [x] Dump disk space retrieved
- [x] Handle by default an external file with patterns ("<binary_name>.conf", with one pattern per line)

