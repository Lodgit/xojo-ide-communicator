# Xojo IDE Communicator

> JSON HTTP Server to communicate transparently with [Xojo IDE](https://www.xojo.com/).

This is a **Work in process** JSON HTTP Server that works like a proxy server to communicate with Xojo IDE via a more transparent API.

The server also takes care about the underlying implementation of [Unix IPC Socket](https://en.wikipedia.org/wiki/Unix_domain_socket) communication between the two parts involved in order to instrument the Xojo IDE for run, compile and test Xojo applications.

## Build constraints

Only Unix-like systems such as Linux and Darwin are supported.

## Development

#### Install dependencies

```
make install
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
