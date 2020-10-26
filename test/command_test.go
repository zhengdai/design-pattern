package test

import (
	"design-pattern/command"
	"testing"
)

func TestCommand(t *testing.T) {
	tv := &command.Tv{}
	onCommand := &command.OnCommand{
		Device: tv,
	}
	offCommand := &command.OffCommand{
		Device: tv,
	}
	onButton := &command.Button{
		Command: onCommand,
	}
	offButton := &command.Button{
		Command: offCommand,
	}
	onButton.Press()
	offButton.Press()
}
