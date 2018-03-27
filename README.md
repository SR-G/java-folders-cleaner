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

# Development

## TODO

- [x] New `--patterns` flag to override default flags
- [ ] Better logging system instead of standard output
- [x] Dump disk space retrieved
- [x] Handle by default an external file with patterns ("<binary_name>.conf", with one pattern per line)