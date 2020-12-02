# Xojo IDE Communicator Client

> CLI client to communicate transparently with [Xojo IDE](https://www.xojo.com/) via [The Xojo IDE Communication Protocol v2](https://docs.xojo.com/UserGuide:IDE_Communicator).

This is a **Work in progress** CLI client written in [Go](https://golang.org/) that communicates with Xojo IDE via a transparent API.
It takes care about the underlying implementation of [Unix IPC Socket](https://en.wikipedia.org/wiki/Unix_domain_socket) communication between the two parties involved in order to instrument Xojo's IDE to run, compile or test Xojo's applications.

### Build constraints

Only Unix-like systems such as Linux and Darwin are supported for now.

## API

Command-line commands and arguments available:

### run

It runs a Xojo project in debug mode.

```sh
xojo-ide-com run -h
# NAME: xojo-ide-com run [OPTIONS] ARGS
# 
# Runs a Xojo project in debug mode. Example: xojo-ide-com run [OPTIONS] PROJECT_FILE_PATH
# 
# OPTIONS:
#   -h --help    Prints help information
# 
# Run 'xojo-ide-com run COMMAND --help' for more information on a command
```

### build

It builds a Xojo project.

```sh
xojo-ide-com build -h
# NAME: xojo-ide-com build [OPTIONS] ARGS
# 
# Builds a Xojo project. Example: xojo-ide-com build [OPTIONS] PROJECT_FILE_PATH
# 
# OPTIONS:
#      --os        Operating system target(s) such as `linux`, `darwin`, `windows` and `ios`. For multiple targets use a coma-separated list.
#      --arch      Target architecture such as `i386`, `amd64` and `arm64`.
#      --reveal    Open the built application directory using the operating system file manager available. [default: false]
#   -h --help      Prints help information
# 
# Run 'xojo-ide-com build COMMAND --help' for more information on a command
```

## Development

#### Install dependencies

```sh
make install
```

#### Run application

```sh
go run main.go [CLI APP OPTIONS]
```

#### Testing and code coverage

```sh
make test
```

#### Build application

```sh
make build
```

## Resources

- [Xojo IDE Communicator](https://docs.xojo.com/UserGuide:IDE_Communicator)
- [Xojo IDE Scripting Building Commands](https://docs.xojo.com/UserGuide:IDE_Scripting_Building_Commands#BuildApp.28buildType_As_Integer.5B.2C_reveal_As_Boolean.5D.29_As_String)
- [IDE Scripting DoCommand](https://docs.xojo.com/UserGuide:IDE_Scripting_DoCommand)
- [IDE Scripting Input Output Commands](https://docs.xojo.com/UserGuide:IDE_Scripting_Input_Output_Commands)

## Contributions

TODO

## License

TODO

© 2020-present [Lodgit Hotelsoftware GmbH](https://www.lodgit-hotel-software.com/)
