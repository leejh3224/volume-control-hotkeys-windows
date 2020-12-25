package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/itchyny/volume-go"
)

// Modifier keys, e.g. Alt, Control, etc.
// Complete list of modifiers can be found below
// https://docs.microsoft.com/en-us/dotnet/api/system.windows.input.modifierkeys?view=net-5.0
const (
	ModNone = 0
)

// Secondary keys
// Complete list of keycodes can be found below
// https://docs.microsoft.com/en-us/dotnet/api/system.windows.forms.keys?view=net-5.0
const (
	F10 = 121
	F11 = 122
	F12 = 123
)

const (
	MuteKey       = 1
	VolumeDownKey = 2
	VolumeUpKey   = 3
)

type HotKey struct {
	ID          int
	Modifiers   int
	KeyCode     int
	Description string
}

// Retuan value of `GetMessageW` syscall
type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

func setVolumeTo(newVolume int) {
	if newVolume < 0 || newVolume > 100 {
		return
	}

	err := volume.SetVolume(newVolume)
	if err != nil {
		log.Fatalf("failed to set volume to %d", newVolume)
	}
}

func main() {
	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()

	reghotkey := user32.MustFindProc("RegisterHotKey")
	getmsg := user32.MustFindProc("GetMessageW")

	keys := map[int16]*HotKey{
		MuteKey:       {MuteKey, ModNone, F10, "f10: mute"},
		VolumeDownKey: {VolumeDownKey, ModNone, F11, "f11: volume down"},
		VolumeUpKey:   {VolumeUpKey, ModNone, F12, "f12: volume up"},
	}

	for _, key := range keys {
		result, _, err := reghotkey.Call(0, uintptr(key.ID), uintptr(key.Modifiers), uintptr(key.KeyCode))

		if result == 1 {
			fmt.Printf("Registered: %s\n", key.Description)
		} else {
			fmt.Printf("failed to register: %v\n", err)
		}
	}

	for {
		msg := &MSG{}
		getmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0)

		if id := msg.WPARAM; id != 0 {
			delta := 5
			currentVolume, err := volume.GetVolume()
			if err != nil {
				log.Fatalf("getting volume failed: %+v", err)
			}

			switch id {
			case MuteKey:
				setVolumeTo(0)
				break
			case VolumeDownKey:
				setVolumeTo(currentVolume - delta)
				break
			case VolumeUpKey:
				setVolumeTo(currentVolume + delta)
				break
			}
		}
	}
}
