# Xojo IDE Communicator Client [![Build Status](https://travis-ci.com/Lodgit/xojo-ide-communicator.svg?branch=master)](https://travis-ci.com/Lodgit/xojo-ide-communicator)

> CLI application to communicate transparently with [Xojo IDE](https://www.xojo.com/) via [The Xojo IDE Communication Protocol v2](https://docs.xojo.com/UserGuide:IDE_Communicator).

CLI application written in [Go](https://golang.org/) that communicates with Xojo IDE using [The Xojo IDE Communication Protocol v2](https://docs.xojo.com/UserGuide:IDE_Communicator) via a more transparent API. It takes care about the underlying implementation of [Unix IPC Socket](https://en.wikipedia.org/wiki/Unix_domain_socket) communication between the two parties involved in order to instrument Xojo's IDE to run, compile or test Xojo's applications.

This CLI tool makes possible to instrument the Xojo IDE for different purposes like automation. For example on [Continuous Integration and Deployment](https://www.atlassian.com/continuous-delivery/principles/continuous-integration-vs-delivery-vs-deployment) systems.

## The XojoUnit Testing Server

Additionally it implements a communication mechanism with [The XojoUnit Testing Server](https://github.com/Lodgit/XojoUnit) which makes possible to run unit tests via [XojoUnit framework](https://github.com/xojo/XojoUnit) via a convenient JSON API.

## Requirements

- Only Unix-like systems such as Linux and Darwin are supported for now.
- It requires Go [v1.15](https://blog.golang.org/go1.15) or later.

## API

Command-line commands and arguments available:

```sh
xojo-ide-com --help
# xojo-ide-com -h
# NAME: xojo-ide-com [OPTIONS] COMMAND
#
# CLI client to communicate transparently with Xojo IDE using The Xojo IDE Communication Protocol v2.
#
# OPTIONS:
#   -w --use-current-workdir    Use the current working directory as the base path for "PROJECT_FILE_PATH" argument on `run` and `build` commands. [default: false]
#   -h --help                   Prints help information
#   -v --version                Prints version information
#
# COMMANDS:
#   run      Runs a Xojo project in debug mode. Example: xojo-ide-com run [OPTIONS] PROJECT_FILE_PATH
#   build    Builds a Xojo project. Example: xojo-ide-com build [OPTIONS] PROJECT_FILE_PATH
#   test     Runs Xojo project tests after a project has been started via `run` command.
#
# Run 'xojo-ide-com COMMAND --help' for more information on a comman
```

### Run

It runs a Xojo project in debug mode.

```sh
xojo-ide-com run --help
# NAME: xojo-ide-com run [OPTIONS] COMMAND
# 
# Runs a Xojo project in debug mode. Example: xojo-ide-com run [OPTIONS] PROJECT_FILE_PATH
# 
# OPTIONS:
#   -d --delay         Workaround delay in seconds to wait until the current application is finally displayed on screen. [default: 5]
#   -t --with-tests    Run available unit tests through XojoUnit testing server. The testing TCP server should run on 127.0.0.1:8123 [default: false]
#   -h --help          Prints help information
# 
# Run 'xojo-ide-com run COMMAND --help' for more information on a command
```

#### Usage

The following command opens a local Xojo project, starts it and runs the unit tests via XojoUnit.

```sh
xojo-ide-com run --with-tests /Users/MyUser/XojoUnit/Desktop\ Project/XojoUnitDesktop.xojo_project
# 2020/12/17 15:36:30 close project command sent: {"tag":"build","script":"CloseProject(False) \nprint \"Default app closed.\""}
# 2020/12/17 15:36:31 data received: {"tag":"build","response":"Default app closed."}
# 2020/12/17 15:36:31 open project command sent: {"tag":"build","script":"OpenFile(\"/Users/MyUser/XojoUnit/Desktop Project/XojoUnitDesktop.xojo_project\") \nprint \"Project is opened.\""}
# 2020/12/17 15:36:32 data received: {"tag":"build","response":"Project is opened."}
# 2020/12/17 15:36:32 run project command sent: {"tag":"build","script":"DoCommand(\"RunApp\") \nprint \"App is running.\""}
# 2020/12/17 15:36:35 data received: {"tag":"build","response":"App is running."}
# 2020/12/17 15:36:35 waiting 5 second(s) for the application...
# 2020/12/17 15:36:40 running project tests via XojoUnit...
#
#      XojoUnit version: 6.6
#      Xojo IDE version: 2019.021
#            Start time: 16.12.20 15:36
#           Total tests: 84 of 86 tests in 2 groups were run
#          Failed tests: 1 (1,19%)
#          Passed tests: 83 (98,81%)
#         Skipped tests: 0
#         Skipped tests: 2
# Not implemented tests: 2
#
# === RUN   Assertion
# --- PASS: Assertion/AreDifferentObject (0.000444 sec)
# --- PASS: Assertion/AreDifferentString (0.000974 sec)
# --- PASS: Assertion/AreDifferentText (0.001056 sec)
# --- PASS: Assertion/AreEqualColor (0.001063 sec)
# --- PASS: Assertion/AreEqualCurrency (0.001169 sec)
# --- PASS: Assertion/AreEqualDate (0.001170 sec)
# --- PASS: Assertion/AreEqualDoubleDefault (0.001042 sec)
# --- PASS: Assertion/AreEqualDouble (0.004553 sec)
# --- PASS: Assertion/AreEqualInt64 (0.001123 sec)
# --- PASS: Assertion/AreEqualIntegerArray (0.001169 sec)
# --- PASS: Assertion/AreEqualInteger (0.000996 sec)
# --- PASS: Assertion/AreEqualMemoryBlock (0.001059 sec)
# --- PASS: Assertion/AreEqualNewDate (0.001132 sec)
# --- PASS: Assertion/AreEqualNewMemoryBlock (0.001125 sec)
# --- PASS: Assertion/AreEqualStringArray (0.001128 sec)
# --- PASS: Assertion/AreEqualString (0.001077 sec)
# --- PASS: Assertion/AreEqualTextArray (0.011069 sec)
# --- PASS: Assertion/AreEqualUInteger (0.001047 sec)
# --- PASS: Assertion/AreNotEqualColor (0.001177 sec)
# --- PASS: Assertion/AreNotEqualDate (0.001138 sec)
# --- PASS: Assertion/AreNotEqualDouble (0.001022 sec)
# --- PASS: Assertion/AreNotEqualMemoryBlock (0.009912 sec)
# --- PASS: Assertion/AreNotEqualNewDate (0.001047 sec)
# --- PASS: Assertion/AreNotEqualNewMemoryBlock (0.001146 sec)
# --- PASS: Assertion/AreSameObject (0.001142 sec)
# --- PASS: Assertion/AreSameStringArray (0.001167 sec)
# --- PASS: Assertion/AreSameString (0.001103 sec)
# --- PASS: Assertion/AreSameTextArray (0.003447 sec)
# --- PASS: Assertion/AreSameText (0.001162 sec)
# --- PASS: Assertion/AssertFailed (0.001133 sec)
# --- PASS: Assertion/Async (0.502278 sec)
# --- PASS: Assertion/CleanSlate1 (0.001392 sec)
# --- PASS: Assertion/CleanSlate2 (0.001175 sec)
# --- PASS: Assertion/DoesNotMatchString (0.001137 sec)
# --- PASS: Assertion/DoEvents (0.992153 sec)
# --- PASS: Assertion/IsFalse (0.001147 sec)
# --- PASS: Assertion/IsNil (0.001233 sec)
# --- PASS: Assertion/IsNotNil (0.001079 sec)
# --- PASS: Assertion/IsTrue (0.001200 sec)
# --- PASS: Assertion/MatchesString (0.001191 sec)
# --- PASS: Assertion/NotImplemented (0.001136 sec)
# --- PASS: Assertion/OverriddenMethod (0.001051 sec)
# --- PASS: Assertion/Pass (0.001135 sec)
# --- PASS: Assertion/Setup1 (0.001207 sec)
# --- PASS: Assertion/Setup2 (0.000886 sec)
# --- PASS: Assertion/SuperClassMethod (0.001134 sec)
# --- PASS: Assertion/TestTimers (0.250893 sec)
# --- PASS: Assertion/UnhandledException (0.004731 sec)
# PASS: Assertion (1.571123 sec)
# Total tests: 48
# Not implemented: 2
# Passed tests: 46
# Failed tests: 0
# Skipped tests: 0
#
# === RUN   Always Fail
# --- PASS: Always Fail/AreDifferentObject (0.001139 sec)
# --- PASS: Always Fail/AreDifferentString (0.001172 sec)
# --- PASS: Always Fail/AreDifferentText (0.001173 sec)
# --- PASS: Always Fail/AreEqualColor (0.001136 sec)
# --- PASS: Always Fail/AreEqualCurrency (0.001144 sec)
# --- PASS: Always Fail/AreEqualDate (0.001058 sec)
# --- PASS: Always Fail/AreEqualDoubleDefault (0.001173 sec)
# --- PASS: Always Fail/AreEqualDouble (0.001181 sec)
# --- PASS: Always Fail/AreEqualInt64 (0.001072 sec)
# --- PASS: Always Fail/AreEqualIntegerArray (0.001019 sec)
# --- PASS: Always Fail/AreEqualInteger (0.001054 sec)
# --- PASS: Always Fail/AreEqualMemoryBlock (0.000921 sec)
# --- PASS: Always Fail/AreEqualNewDate (0.001165 sec)
# --- PASS: Always Fail/AreEqualNewMemoryBlock (0.001169 sec)
# --- PASS: Always Fail/AreEqualStringArray (0.001136 sec)
# --- PASS: Always Fail/AreEqualString (0.001173 sec)
# --- PASS: Always Fail/AreEqualTextArray (0.001202 sec)
# --- PASS: Always Fail/AreEqualUInteger (0.000939 sec)
# --- PASS: Always Fail/AreNotEqualColor (0.001148 sec)
# --- PASS: Always Fail/AreNotEqualDate (0.001188 sec)
# --- PASS: Always Fail/AreNotEqualDouble (0.001043 sec)
# --- PASS: Always Fail/AreNotEqualMemoryBlock (0.001106 sec)
# --- PASS: Always Fail/AreNotEqualNewDate (0.000968 sec)
# --- PASS: Always Fail/AreNotEqualNewMemoryBlock (0.000957 sec)
# --- PASS: Always Fail/AreSameObject (0.001108 sec)
# --- PASS: Always Fail/AreSameStringArray (0.001173 sec)
# --- PASS: Always Fail/AreSameString (0.001171 sec)
# --- PASS: Always Fail/AreSameTextArray (0.001107 sec)
# --- PASS: Always Fail/AreSameText (0.001099 sec)
# --- PASS: Always Fail/Async (1.001193 sec)
# --- PASS: Always Fail/DoesNotMatchString (0.001135 sec)
# --- PASS: Always Fail/Fail (0.001136 sec)
# --- PASS: Always Fail/IsFalse (0.001176 sec)
# --- PASS: Always Fail/IsNil (0.000978 sec)
# --- PASS: Always Fail/IsNotNil (0.001044 sec)
# --- PASS: Always Fail/IsTrue (0.001013 sec)
# --- PASS: Always Fail/MatchesString (0.001119 sec)
# --- FAIL: Always Fail/WillTrulyFail (0.001129 sec)
# FAIL: Always Fail (1.042017 sec)
# Total tests: 38
# Not implemented: 0
# Passed tests: 37
# Failed tests: 1
# Skipped tests: 0
#
# ✗ Tests has been failed.
# exit status 1
```

### Build

It builds a Xojo project to specific target(s).

```sh
xojo-ide-com build --helps
# NAME: xojo-ide-com build [OPTIONS] COMMAND
#
# Builds a Xojo project to specific target(s). Example: xojo-ide-com build [OPTIONS] PROJECT_FILE_PATH
#
# OPTIONS:
#  -t --targets    Operating systems with their architectures. Coma-separated list with one or more target pairs in lower case. E.g linux-amd64,darwin-arm64,windows-i386.
#  -r --reveal    Open the built application directory using the operating system file manager available. [default: false]
#  -h --help      Prints help information
#
# Run 'xojo-ide-com build COMMAND --help' for more information on a command
```

#### Targets supported

The following list describe operating system and architecture combination pairs supported by `--targets` option.

| **Operating system** | **amd64**       | **arm64**      | **i386**       |
| -------------------- | --------------- | -------------- | -------------- |
| Linux                | `linux-amd64`   | `linux-arm64`  | `linux-i386`   |
| macOS                | `darwin-amd64`  | `darwin-arm64` | `darwin-i386`  |
| Windows              | `windows-amd64` | ?              | `windows-i386` |
| iOS                  | `ios-amd64`     | `ios-arm64`    | ?              |

More details at https://docs.xojo.com/UserGuide:IDE_Scripting_Building_Commands

#### Usage

The following command opens a local Xojo project and build executables for Linux (amd64), MacOs (arm64) and Windows (i386).

```sh
xojo-ide-com build --targets linux-amd64,darwin-arm64,windows-i386 \
    /Users/MyUser/XojoUnit/Desktop\ Project/XojoUnitDesktop.xojo_project
# 2021/01/06 10:52:07 close project command sent: {"tag":"build","script":"CloseProject(False)
# print \"Default app closed.\""}
# 2021/01/06 10:52:07 data received: {"tag":"build","response":"Default app closed."}
# 2021/01/06 10:52:07 open project command sent: {"tag":"build","script":"OpenFile(\"/Users/MyUser/XojoUnit/Desktop Project/XojoUnitDesktop.xojo_project\")
# print \"Project is opened.\""}
# 2021/01/06 10:52:09 data received: {"tag":"build","response":"Project is opened."}
# 2021/01/06 10:52:09 build project options chosen: linux/amd64
# 2021/01/06 10:52:09 build project command sent: {"script":"Print BuildApp(17,False)", "tag":"build"}
# 2021/01/06 10:52:13 data received: {"tag":"build","response":"\/Users\/MyUser\/XojoUnit\/Desktop\\ Project\/Builds\\ \\-\\ XojoUnitDesktop\/Linux\\ 64\\ bit\/# XojoUnit\/XojoUnit"}
# 2021/01/06 10:52:13 build project options chosen: darwin/arm64
# 2021/01/06 10:52:13 build project command sent: {"script":"Print BuildApp(24,False)", "tag":"build"}
# 2021/01/06 10:52:19 data received: {"tag":"build","response":"\/Users\/MyUser\/XojoUnit\/Desktop\\ Project\/Builds\\ \\-\\ XojoUnitDesktop\/macOS\\ ARM\\ # 64\\ bit\/XojoUnit.app"}
# 2021/01/06 10:52:19 build project options chosen: windows/i386
# 2021/01/06 10:52:19 build project command sent: {"script":"Print BuildApp(3,False)", "tag":"build"}
# 2021/01/06 10:52:21 data received: {"tag":"build","response":"\/Users\/MyUser\/XojoUnit\/Desktop\\ Project\/Builds\\ \\-\\ XojoUnitDesktop\/Windows\/# XojoUnit\/XojoUnit.exe"}
# 2021/01/06 10:52:21 close project command sent: {"tag":"build","script":"CloseProject(False)
# print \"Default app closed.\""}
# 2021/01/06 10:52:21 data received: {"tag":"build","response":"Default app closed."}
```

## Development

### Install dependencies

```sh
make install
```

### Run application

```sh
go run main.go [CLI APP OPTIONS]
```

### Testing and code coverage

```sh
make test
```

## Production

### Build

```sh
make build
```

## Resources

- [Xojo IDE Communicator](https://docs.xojo.com/UserGuide:IDE_Communicator)
- [Xojo IDE Scripting Building Commands](https://docs.xojo.com/UserGuide:IDE_Scripting_Building_Commands#BuildApp.28buildType_As_Integer.5B.2C_reveal_As_Boolean.5D.29_As_String)
- [IDE Scripting DoCommand](https://docs.xojo.com/UserGuide:IDE_Scripting_DoCommand)
- [IDE Scripting Input Output Commands](https://docs.xojo.com/UserGuide:IDE_Scripting_Input_Output_Commands)

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/Lodgit/xojo-ide-communicator/pulls) or [issue](https://github.com/Lodgit/xojo-ide-communicator/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Lodgit Hotelsoftware GmbH](https://www.lodgit-hotel-software.com/)
