# Xojo IDE Communicator Client ![CI](https://github.com/Lodgit/xojo-ide-communicator/workflows/CI/badge.svg)

> CLI application to communicate transparently with [Xojo IDE](https://www.xojo.com/) via [The Xojo IDE Communication Protocol v2](https://docs.xojo.com/UserGuide:IDE_Communicator).

CLI application written in [Go](https://golang.org/) that communicates with Xojo IDE using [The Xojo IDE Communication Protocol v2](https://docs.xojo.com/UserGuide:IDE_Communicator) taking care about the underlying implementation of [Unix IPC Socket](https://en.wikipedia.org/wiki/Unix_domain_socket) communication between the two parties involved in order to instrument Xojo's IDE to run, compile or test Xojo applications.

This CLI tool makes possible to instrument the Xojo IDE for different purposes like automation. For example on [Continuous Integration and Deployment](https://www.atlassian.com/continuous-delivery/principles/continuous-integration-vs-delivery-vs-deployment) systems.

## The XojoUnit Testing Server

Additionally it implements a communication mechanism with [The XojoUnit Testing Server](https://github.com/Lodgit/XojoUnit) which makes possible to run unit tests using [XojoUnit framework](https://github.com/xojo/XojoUnit) through a convenient JSON API.

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

__Note:__ The XojoUnit needs to implement the Testing Server API only if `run` command is wanted to run with `--with-tests` flag.

```sh
xojo-ide-com run --with-tests /Users/MyUser/XojoUnit/Desktop\ Project/XojoUnitDesktop.xojo_project
# 2020/12/17 15:36:30 close project command sent: {"tag":"build","script":"CloseProject(False) \nprint \"Default app closed.\""}
# 2020/12/17 15:36:31 data received: {"tag":"build","response":"Default app closed."}
# 2020/12/17 15:36:31 open project command sent: {"tag":"build","script":"OpenFile(\"/Users/MyUser/XojoUnit/Desktop Project/XojoUnitDesktop.xojo_project\") \nprint \"Project is opened.\""}
# 2020/12/17 15:36:32 data received: {"tag":"build","response":"Project is opened."}
# 2020/12/17 15:36:32 run project command sent: {"tag":"build","script":"DoCommand(\"RunApp\") \nprint \"App is running.\""}
# 2020/12/17 15:36:35 data received: {"tag":"build","response":"App is running."}
# 2020/12/17 15:36:35 Waiting 5 second(s) for the application to start...
# 2020/12/17 15:36:40 Running project tests via XojoUnit...
#
#  XojoUnit Testing Server
# =========================
# XojoUnit version: 6.6
# Xojo IDE version: 2020r2.1
#       Start time: 17.12.20 15:36
#
# === RUN   Assertion
# --- PASS: Assertion/AreDifferentObject (0.00ms)
# --- PASS: Assertion/AreDifferentString (0.00ms)
# --- PASS: Assertion/AreDifferentText (0.00ms)
# --- PASS: Assertion/AreEqualColor (0.00ms)
# --- PASS: Assertion/AreEqualCurrency (0.01ms)
# --- PASS: Assertion/AreEqualDate (0.00ms)
# --- PASS: Assertion/AreEqualDoubleDefault (0.00ms)
# --- PASS: Assertion/AreEqualDouble (0.00ms)
# --- PASS: Assertion/AreEqualInt64 (0.00ms)
# --- PASS: Assertion/AreEqualIntegerArray (0.01ms)
# --- PASS: Assertion/AreEqualInteger (0.00ms)
# --- PASS: Assertion/AreEqualMemoryBlock (0.00ms)
# --- PASS: Assertion/AreEqualNewDate (0.00ms)
# --- PASS: Assertion/AreEqualNewMemoryBlock (0.00ms)
# --- PASS: Assertion/AreEqualStringArray (0.01ms)
# --- PASS: Assertion/AreEqualString (0.00ms)
# --- PASS: Assertion/AreEqualTextArray (0.00ms)
# --- PASS: Assertion/AreEqualUInteger (0.00ms)
# --- PASS: Assertion/AreNotEqualColor (0.00ms)
# --- PASS: Assertion/AreNotEqualDate (0.01ms)
# --- PASS: Assertion/AreNotEqualDouble (0.00ms)
# --- PASS: Assertion/AreNotEqualMemoryBlock (0.00ms)
# --- PASS: Assertion/AreNotEqualNewDate (0.00ms)
# --- PASS: Assertion/AreNotEqualNewMemoryBlock (0.00ms)
# --- PASS: Assertion/AreSameObject (0.00ms)
# --- PASS: Assertion/AreSameStringArray (0.01ms)
# --- PASS: Assertion/AreSameString (0.00ms)
# --- PASS: Assertion/AreSameTextArray (0.00ms)
# --- PASS: Assertion/AreSameText (0.00ms)
# --- PASS: Assertion/AssertFailed (0.00ms)
# --- PASS: Assertion/Async (0.50ms)
# --- PASS: Assertion/CleanSlate1 (0.01ms)
# --- PASS: Assertion/CleanSlate2 (0.00ms)
# --- PASS: Assertion/DoesNotMatchString (0.00ms)
# --- PASS: Assertion/DoEvents (0.99ms)
# --- PASS: Assertion/IsFalse (0.00ms)
# --- PASS: Assertion/IsNil (0.00ms)
# --- PASS: Assertion/IsNotNil (0.00ms)
# --- PASS: Assertion/IsTrue (0.00ms)
# --- PASS: Assertion/MatchesString (0.00ms)
# --- SKIP: Assertion/NotImplemented (0.00ms)
# --- PASS: Assertion/OverriddenMethod (0.00ms)
# --- PASS: Assertion/Pass (0.00ms)
# --- PASS: Assertion/Setup1 (0.00ms)
# --- PASS: Assertion/Setup2 (0.00ms)
# --- PASS: Assertion/SuperClassMethod (0.00ms)
# --- SKIP: Assertion/TestTimers (0.25ms)
# --- PASS: Assertion/UnhandledException (0.00ms)
# PASS: Assertion (1.60s)
#           Total: 48
#          Passed: 46
#          Failed: 0
#         Skipped: 0
# Not implemented: 2
#
# ---
#
# === RUN   Always Fail
# --- PASS: Always Fail/AreDifferentObject (0.00ms)
# --- FAIL: Always Fail/AreDifferentString (0.00ms)
#              : String '48 65 6C 6C 6F' is the same
#              : Expected 2 failures, but had 1
# --- PASS: Always Fail/AreDifferentText (0.00ms)
# --- PASS: Always Fail/AreEqualColor (0.00ms)
# --- PASS: Always Fail/AreEqualCurrency (0.00ms)
# --- PASS: Always Fail/AreEqualDate (0.00ms)
# --- PASS: Always Fail/AreEqualDoubleDefault (0.00ms)
# --- PASS: Always Fail/AreEqualDouble (0.00ms)
# --- PASS: Always Fail/AreEqualInt64 (0.00ms)
# --- PASS: Always Fail/AreEqualIntegerArray (0.00ms)
# --- PASS: Always Fail/AreEqualInteger (0.00ms)
# --- PASS: Always Fail/AreEqualMemoryBlock (0.00ms)
# --- PASS: Always Fail/AreEqualNewDate (0.00ms)
# --- PASS: Always Fail/AreEqualNewMemoryBlock (0.00ms)
# --- PASS: Always Fail/AreEqualStringArray (0.00ms)
# --- PASS: Always Fail/AreEqualString (0.00ms)
# --- PASS: Always Fail/AreEqualTextArray (0.00ms)
# --- PASS: Always Fail/AreEqualUInteger (0.00ms)
# --- PASS: Always Fail/AreNotEqualColor (0.00ms)
# --- PASS: Always Fail/AreNotEqualDate (0.00ms)
# --- PASS: Always Fail/AreNotEqualDouble (0.00ms)
# --- PASS: Always Fail/AreNotEqualMemoryBlock (0.00ms)
# --- PASS: Always Fail/AreNotEqualNewDate (0.00ms)
# --- PASS: Always Fail/AreNotEqualNewMemoryBlock (0.00ms)
# --- PASS: Always Fail/AreSameObject (0.00ms)
# --- PASS: Always Fail/AreSameStringArray (0.00ms)
# --- PASS: Always Fail/AreSameString (0.00ms)
# --- PASS: Always Fail/AreSameTextArray (0.00ms)
# --- PASS: Always Fail/AreSameText (0.00ms)
# --- PASS: Always Fail/Async (1.00s)
# --- PASS: Always Fail/DoesNotMatchString (0.00ms)
# --- PASS: Always Fail/Fail (0.00ms)
# --- PASS: Always Fail/IsFalse (0.00ms)
# --- PASS: Always Fail/IsNil (0.00ms)
# --- PASS: Always Fail/IsNotNil (0.00ms)
# --- PASS: Always Fail/IsTrue (0.00ms)
# --- PASS: Always Fail/MatchesString (0.00ms)
# --- FAIL: Always Fail/WillTrulyFail (0.00ms)
#              : A RuntimeException occurred and was caught.
#              Message: This is an on purpose exception!
#    Error Type: RuntimeException
#  Error Number: 0
# Error Message: This is an on purpose exception!
#  Error Reason: This is an on purpose exception!
#   Error Stack:
#             0: RuntimeRaiseException
#             1: XojoUnitFailTests.WillTrulyFailTest%%o<XojoUnitFailTests>
#             2: Introspection.Invoke%%
#             3: Xojo.Introspection.MethodInfo.Invoke%x%o<Xojo.Introspection.MethodInfo>o<Object>A1x
#             4: Xojo.Introspection.MethodInfo.Invoke%%o<Xojo.Introspection.MethodInfo>o<Object>A1x
#             5: TestGroup.RunTestsTimer_Action%%o<TestGroup>o<Xojo.Core.Timer>
#             6: Delegate.IM_Invoke%%o<Xojo.Core.Timer>
#             7: _ZN10TimerImpCF15FireTimerActionEv
#             8: _ZN10TimerImpCF13TimerCallbackEv
#             9: __CFRUNLOOP_IS_CALLING_OUT_TO_A_TIMER_CALLBACK_FUNCTION__
#             10: __CFRunLoopDoTimer
#             11: __CFRunLoopDoTimers
#             12: __CFRunLoopRun
#             13: CFRunLoopRunSpecific
#             14: RunCurrentEventLoopInMode
#             15: ReceiveNextEventCommon
#             16: _BlockUntilNextEventMatchingListInModeWithFilter
#             17: _DPSNextEvent
#             18: -[NSApplication(NSEvent) _nextEventMatchingEventMask:untilDate:inMode:dequeue:]
#             19: XojoFramework$4860
#             20: XojoFramework$4861
#             21: Application._CallFunctionWithExceptionHandling%%o<Application>p
#             22: _Z33CallFunctionWithExceptionHandlingPFvvE
#             23: XojoFramework$4860
#             24: -[NSApplication run]
#             25: RuntimeRun
#             26: REALbasic._RuntimeRun
#             27: _Main
#             28: main
# FAIL: Always Fail (1.04s)
#           Total: 38
#          Passed: 36
#          Failed: 2
#         Skipped: 0
# Not implemented: 0
#
# ---
#
# Final test results:
#           Total: 84 of 86 tests in 2 groups were run
#          Passed: 82 (97,62%)
#          Failed: 2 (2,38%)
#         Skipped: 0
# Not implemented: 2
#
# ✗ Tests has been failed / 2 (2,38%)
#
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
#  -t --targets    Operating systems with their architectures. Coma-separated list with one or more target pairs in lower case. E.g linux-amd64,darwin-arm64,windows-386.
#  -r --reveal    Open the built application directory using the operating system file manager available. [default: false]
#  -h --help      Prints help information
#
# Run 'xojo-ide-com build COMMAND --help' for more information on a command
```

#### Targets supported

The following list describe operating system and architecture combination pairs supported by `--targets` option.

| **Operating system** | **amd64**       | **386**       | **arm64**      | **Universal**  |
| -------------------- | --------------- | ------------- | -------------- | -------------- |
| Linux                | `linux-amd64`   | `linux-386`   | `linux-arm64`  |                |
| macOS                | `darwin-amd64`  | `darwin-386`  | `darwin-arm64` | `darwin-universal` * |
| Windows              | `windows-amd64` | `windows-386` | ?              |                |
| iOS                  | `ios-amd64`     | ?             | `ios-arm64`    |                |

**Notes:**

- \* Xojo [macOS Universal](https://developer.apple.com/documentation/apple-silicon/porting-your-macos-apps-to-apple-silicon) provides support for Apple Silicon (ARM64) and Intel-based (adm64) Mac computers.

More details at https://docs.xojo.com/UserGuide:IDE_Scripting_Building_Commands

#### Usage

The following command opens a local Xojo project and build executables for Linux (amd64), MacOs (arm64) and Windows (386).

```sh
xojo-ide-com build --targets linux-amd64,darwin-arm64,windows-386 \
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
# 2021/01/06 10:52:19 build project options chosen: windows/386
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

### Single build

```sh
make build
```

### Multiple builds and archiving

```sh
make release
```

See `Makefile` for more details

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
