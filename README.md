# Xojo IDE Communicator Client

> CLI client to communicate transparently with [Xojo IDE](https://www.xojo.com/) via [The Xojo IDE Communication Protocol v2](https://docs.xojo.com/UserGuide:IDE_Communicator).

This is a **Work in progress** CLI client written in [Go](https://golang.org/) that communicates with Xojo IDE via a transparent API.
It takes care about the underlying implementation of [Unix IPC Socket](https://en.wikipedia.org/wiki/Unix_domain_socket) communication between the two parties involved in order to instrument Xojo's IDE to run, compile or test Xojo's applications.

## Build constraints

Only Unix-like systems such as Linux and Darwin are supported for now.

## Development

#### Install dependencies

```
make install
```

#### Run application

```
make run
```

#### Testing and code coverage

```
make test
```

#### Build application

```
make build
```

## Resources

- [Xojo IDE Communicator](https://docs.xojo.com/UserGuide:IDE_Communicator)
- [Xojo IDE Scripting Building Commands](https://docs.xojo.com/UserGuide:IDE_Scripting_Building_Commands#BuildApp.28buildType_As_Integer.5B.2C_reveal_As_Boolean.5D.29_As_String)
- [IDE Scripting DoCommand](https://docs.xojo.com/UserGuide:IDE_Scripting_DoCommand)
