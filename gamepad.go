package uinput

import (
	"fmt"
	"io"
	"os"
)

const (
	absRx    = 0x03
	absRy    = 0x04
	absHat0x = 0x10
	absHat0y = 0x11

	evBtnSouth         = 0x130
	evBtnEast          = 0x131
	evBtnNorth         = 0x133
	evBtnWest          = 0x134
	evBtnTrigger1Left  = 0x136
	evBtnTrigger1Right = 0x137
	evBtnTrigger2Left  = 0x138
	evBtnTrigger2Right = 0x139
	evBtnSelect        = 0x13a
	evBtnStart         = 0x13b
	evBtnMode          = 0x13c
	evBtnThumbLeft     = 0x13d
	evBtnThumbRight    = 0x13e
	evBtnDpadUp        = 0x220
	evBtnDpadDown      = 0x221
	evBtnDpadLeft      = 0x222
	evBtnDpadRight     = 0x223
)

type GamePad interface {
	SouthPress() error
	SouthRelease() error

	EastPress() error
	EastRelease() error

	NorthPress() error
	NorthRelease() error

	WestPress() error
	WestRelease() error

	Trigger1LeftPress() error
	Trigger1LeftRelease() error

	Trigger1RightPress() error
	Trigger1RightRelease() error

	Trigger2LeftPress() error
	Trigger2LeftRelease() error

	Trigger2RightPress() error
	Trigger2RightRelease() error

	SelectPress() error
	SelectRelease() error

	StartPress() error
	StartRelease() error

	ModePress() error
	ModeRelease() error

	ThumbLeftPress() error
	ThumbLeftRelease() error

	ThumbRightPress() error
	ThumbRightRelease() error

	Move(x, y int32) error
	MoveLeftStick(x, y int32) error
	MoveRightStick(x, y int32) error

	io.Closer
}

type vGamePad struct {
	name       []byte
	deviceFile *os.File
}

func (vgp vGamePad) SouthPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnSouth}, btnStatePressed)
}

func (vgp vGamePad) SouthRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnSouth}, btnStateReleased)
}

func (vgp vGamePad) EastPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnEast}, btnStatePressed)
}

func (vgp vGamePad) EastRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnEast}, btnStateReleased)
}

func (vgp vGamePad) NorthPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnNorth}, btnStatePressed)
}

func (vgp vGamePad) NorthRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnNorth}, btnStateReleased)
}

func (vgp vGamePad) WestPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnWest}, btnStatePressed)
}

func (vgp vGamePad) WestRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnWest}, btnStateReleased)
}

func (vgp vGamePad) Trigger1LeftPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger1Left}, btnStatePressed)
}

func (vgp vGamePad) Trigger1LeftRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger1Left}, btnStateReleased)
}

func (vgp vGamePad) Trigger1RightPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger1Right}, btnStatePressed)
}

func (vgp vGamePad) Trigger1RightRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger1Right}, btnStateReleased)
}

func (vgp vGamePad) Trigger2LeftPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger2Left}, btnStatePressed)
}

func (vgp vGamePad) Trigger2LeftRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger2Left}, btnStateReleased)
}

func (vgp vGamePad) Trigger2RightPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger2Right}, btnStatePressed)
}

func (vgp vGamePad) Trigger2RightRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnTrigger2Right}, btnStateReleased)
}

func (vgp vGamePad) SelectPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnSelect}, btnStatePressed)
}

func (vgp vGamePad) SelectRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnSelect}, btnStateReleased)
}

func (vgp vGamePad) StartPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnStart}, btnStatePressed)
}

func (vgp vGamePad) StartRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnStart}, btnStateReleased)
}

func (vgp vGamePad) ModePress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnMode}, btnStatePressed)
}

func (vgp vGamePad) ModeRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnMode}, btnStateReleased)
}

func (vgp vGamePad) ThumbLeftPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnThumbLeft}, btnStatePressed)
}

func (vgp vGamePad) ThumbLeftRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnThumbLeft}, btnStateReleased)
}

func (vgp vGamePad) ThumbRightPress() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnThumbRight}, btnStatePressed)
}

func (vgp vGamePad) ThumbRightRelease() error {
	return sendBtnEvent(vgp.deviceFile, []int{evBtnThumbRight}, btnStateReleased)
}

func (vgp vGamePad) Move(x, y int32) error {
	return sendArbitraryAbsEvent(vgp.deviceFile, absHat0x, absHat0y, x, y)
}

func (vgp vGamePad) MoveLeftStick(x, y int32) error {
	return sendArbitraryAbsEvent(vgp.deviceFile, absX, absY, x, y)
}

func (vgp vGamePad) MoveRightStick(x, y int32) error {
	return sendArbitraryAbsEvent(vgp.deviceFile, absRx, absRy, x, y)
}

func (vgp vGamePad) Close() error {
	return closeDevice(vgp.deviceFile)
}

func CreateGamePad(
	path string,
	name []byte,
	minLeftStickX, maxLeftStickX, minLeftStickY, maxLeftStickY,
	minRightStickX, maxRightStickX, minRightStickY, maxRightStickY,
	minDpadX, maxDpadX, minDpadY, maxDpadY int32, // FIXME: should i use a struct for this?
) (GamePad, error) {
	err := validateDevicePath(path)
	if err != nil {
		return nil, err
	}
	err = validateUinputName(name)
	if err != nil {
		return nil, err
	}

	fd, err := createGamePad(
		path,
		name,
		minLeftStickX, maxLeftStickX, minLeftStickY, maxLeftStickY,
		minRightStickX, maxRightStickX, minRightStickY, maxRightStickY,
		minDpadX, maxDpadX, minDpadY, maxDpadY,
	)
	if err != nil {
		return nil, err
	}

	return vGamePad{name: name, deviceFile: fd}, nil
}

func createGamePad(
	path string,
	name []byte,
	minLeftStickX, maxLeftStickX, minLeftStickY, maxLeftStickY,
	minRightStickX, maxRightStickX, minRightStickY, maxRightStickY,
	minDpadX, maxDpadX, minDpadY, maxDpadY int32,
) (fd *os.File, err error) {
	deviceFile, err := createDeviceFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not create gamepad input device: %v", err)
	}

	err = registerDevice(deviceFile, uintptr(evKey))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register key device: %v", err)
	}

	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnSouth))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register south gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnEast))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register east gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnNorth))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register north gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnWest))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register west gamepad button event: %v", err)
	}

	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnTrigger1Left))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register trigger 1 left gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnTrigger1Right))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register trigger 1 right gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnTrigger2Left))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register trigger 2 left gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnTrigger2Right))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register trigger 2 right gamepad button event: %v", err)
	}

	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnSelect))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register select gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnStart))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register start gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnMode))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register mode gamepad button event: %v", err)
	}

	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnThumbLeft))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register thumb left gamepad button event: %v", err)
	}
	err = ioctl(deviceFile, uiSetKeyBit, uintptr(evBtnThumbRight))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register thumb right gamepad button event: %v", err)
	}

	err = registerDevice(deviceFile, uintptr(evAbs))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register absolute axis input device: %v", err)
	}

	err = ioctl(deviceFile, uiSetAbsBit, uintptr(absX))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register left stick x gamepad event: %v", err)
	}
	err = ioctl(deviceFile, uiSetAbsBit, uintptr(absY))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register left stick y gamepad event: %v", err)
	}
	err = ioctl(deviceFile, uiSetAbsBit, uintptr(absRx))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register right stick x gamepad event: %v", err)
	}
	err = ioctl(deviceFile, uiSetAbsBit, uintptr(absRy))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register right stick y gamepad event: %v", err)
	}
	err = ioctl(deviceFile, uiSetAbsBit, uintptr(absHat0x))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register dpad x gamepad event: %v", err)
	}
	err = ioctl(deviceFile, uiSetAbsBit, uintptr(absHat0y))
	if err != nil {
		deviceFile.Close()
		return nil, fmt.Errorf("failed to register dpad y gamepad event: %v", err)
	}

	var absMin [absSize]int32
	absMin[absX] = minLeftStickX
	absMin[absY] = minLeftStickY
	absMin[absRx] = minRightStickX
	absMin[absRy] = minRightStickY
	absMin[absHat0x] = minDpadX
	absMin[absHat0y] = minDpadY

	var absMax [absSize]int32
	absMax[absX] = maxLeftStickX
	absMax[absY] = maxLeftStickY
	absMax[absRx] = maxRightStickX
	absMax[absRy] = maxRightStickY
	absMax[absHat0x] = maxDpadX
	absMax[absHat0y] = maxDpadY

	return createUsbDevice(deviceFile, uinputUserDev{
		Name: toUinputName(name),
		ID: inputID{
			Bustype: busUsb,
			Vendor:  0x4711,
			Product: 0x0819,
			Version: 1,
		},
		Absmin: absMin,
		Absmax: absMax,
	})
}

func sendArbitraryAbsEvent(deviceFile *os.File, codeX, codeY uint16, xPos, yPos int32) error {
	var ev [2]inputEvent
	ev[0].Type = evAbs
	ev[0].Code = codeX
	ev[0].Value = xPos

	ev[1].Type = evAbs
	ev[1].Code = codeY
	ev[1].Value = yPos

	for _, iev := range ev {
		buf, err := inputEventToBuffer(iev)
		if err != nil {
			return fmt.Errorf("writing abs event failed: %v", err)
		}

		_, err = deviceFile.Write(buf)
		if err != nil {
			return fmt.Errorf("failed to write abs event to device file: %v", err)
		}
	}

	return syncEvents(deviceFile)
}
