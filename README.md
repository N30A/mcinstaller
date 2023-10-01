**Description:**

mcinstaller is a command-line tool for installing minecraft servers. It allows you to list supported servers, list available server versions, and install minecraft servers with ease.

**Usage:**

```
mcinstaller <command> [options]

list servers
  List all supported minecraft servers.

list versions <server>
  List all supported versions for the specified server.

install <server> <version> <server-dir>
  Install a minecraft server with the supplied information.

help
  Show this help message.
```

**Examples:**

List all supported servers:
```sh
$ mcinstaller list servers
```

List available versions for a server:
```sh
$ mcinstaller list versions vanilla
```

Install a minecraft server:
```sh
$ mcinstaller install vanilla 1.17.1 /path/to/server-directory
```

**Notes:**
```
The <server> argument should be a valid server (e.g. vanilla or paper).
The <version> argument should be a valid version number for the selected server.
The <server-dir> argument should be the path where you want to install the server.
```