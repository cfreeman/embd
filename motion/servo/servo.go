// Package servo allows control of servos using a PWM controller.
package servo

import (
	"github.com/golang/glog"
	"github.com/kidoman/embd/util"
)

const (
	minus = 544
	maxus = 2400
)

// A PWM interface implements access to a pwm controller.
type PWM interface {
	SetMicroseconds(channel int, us int) error
}

type Servo struct {
	PWM     PWM
	Channel int

	Minus, Maxus int
}

// New creates a new Servo interface.
func New(pwm PWM, channel int) *Servo {
	return &Servo{
		PWM:     pwm,
		Channel: channel,
		Minus:   minus,
		Maxus:   maxus,
	}
}

// SetAngle sets the servo angle.
func (s *Servo) SetAngle(angle int) error {
	us := util.Map(int64(angle), 0, 180, int64(s.Minus), int64(s.Maxus))

	glog.V(1).Infof("servo: given angle %v calculated %v us", angle, us)

	return s.PWM.SetMicroseconds(s.Channel, int(us))
}
