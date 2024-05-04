**Description:**

mcinstaller is a command-line tool for installing minecraft servers. It allows you to list supported servers, list available server versions, and install minecraft servers with ease.

**Usage:**

```
Usage:
  mcinstaller [command]

Available Commands:
  help        Help about any command
  install     Install a minecraft server
  servers     List supported servers
  versions    List versions for the specified server

Flags:
  -h, --help   help for mcinstaller
```

**Examples:**

List supported servers:
```sh
mcinstaller servers
```

List available versions for a server:
```sh
mcinstaller versions vanilla
```

Install a minecraft server:
```sh
mcinstaller install vanilla 1.20 /path/to/server-directory
```
