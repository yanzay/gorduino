package gorduino

import (
	"github.com/kraman/go-firmata"
)

type Gorduino struct {
	pins   map[byte]bool
	client *firmata.FirmataClient
	work   func()
}

func NewGorduino(port string, pins ...byte) *Gorduino {
	g := new(Gorduino)
	g.pins = make(map[byte]bool)
	g.client, _ = firmata.NewClient(port, 57600)
	for _, pin := range pins {
		g.pins[pin] = true
		g.client.SetPinMode(pin, firmata.Output)
	}
	return g
}

func (g *Gorduino) On(p byte) {
	if g.pins[p] == true {
		g.client.DigitalWrite(uint(p), true)
	}
}

func (g *Gorduino) Off(p byte) {
	if g.pins[p] == true {
		g.client.DigitalWrite(uint(p), false)
	}
}
