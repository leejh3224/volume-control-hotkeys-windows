# volume-control-hotkeys-windows

Register set of hotkeys for media volume control for Windows

- f10: mute
- f11: volume down
- f12: volume up

## Getting Started

```
go get
go run main.go
```

## Build

```
make build # if make is installed

go build # otherwise run this
```

## Caveat

Unless you changed [UserDebuggerHotKey](https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2003/cc786263(v=ws.10)?redirectedfrom=MSDN) into something other than the default one (`f12`), `f12` key must be in use and therefore can't be registered for global hot key.

In order to use this program, you need to update it to `0x2F`. ([reference](http://muzso.hu/2011/12/13/setting-f12-as-a-global-hotkey-in-windows))
