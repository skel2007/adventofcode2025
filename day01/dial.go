package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type Dial struct {
	Position         uint8
	ZeroCounter      uint32
	TotalZeroCounter uint32
}

func NewDial() *Dial {
	return &Dial{
		Position:         50,
		ZeroCounter:      0,
		TotalZeroCounter: 0,
	}
}

func (d *Dial) Rotate(str string) error {
	rotation, err := parseRotation(str)
	if err != nil {
		return err
	}

	switch rotation.direction {
	case directionLeft:
		newPosition := (d.Position + 100 - rotation.distance) % 100
		if newPosition == 0 || (newPosition > d.Position && d.Position > 0) {
			d.TotalZeroCounter++
		}

		d.Position = newPosition
	case directionRight:
		newPosition := (d.Position + rotation.distance) % 100
		if newPosition < d.Position {
			d.TotalZeroCounter++
		}

		d.Position = newPosition
	default:
		return fmt.Errorf("invalid direction: %d", rotation.direction)
	}

	d.TotalZeroCounter += rotation.circles

	if d.Position == 0 {
		d.ZeroCounter++
	}

	return nil
}

func CalculatePasswords(str string) (uint32, uint32, error) {
	dial := NewDial()
	for _, str := range strings.Split(str, "\n") {
		if strings.TrimSpace(str) == "" {
			continue
		}

		if err := dial.Rotate(str); err != nil {
			return 0, 0, err
		}
	}

	return dial.ZeroCounter, dial.TotalZeroCounter, nil
}

type direction int

const (
	directionLeft direction = iota
	directionRight
)

type rotation struct {
	direction direction
	distance  uint8
	circles   uint32
}

func parseRotation(str string) (*rotation, error) {
	var direction direction
	if str[0] == 'L' {
		direction = directionLeft
	} else if str[0] == 'R' {
		direction = directionRight
	} else {
		return nil, fmt.Errorf("invalid direction: %c", str[0])
	}

	distance, err := strconv.Atoi(str[1:])
	if err != nil {
		return nil, err
	}

	return &rotation{
		direction: direction,
		distance:  uint8(distance % 100),
		circles:   uint32(distance / 100),
	}, nil
}
