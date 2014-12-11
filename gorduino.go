package gorduino

import (
	"github.com/tarm/goserial"
	"github.com/yanzay/go-firmata"
)

type Gorduino struct {
	pins   map[byte]bool
	client *firmata.FirmataClient
	work   func()
}

func NewGorduino(port string, pins ...byte) (*Gorduino, error) {
	g := new(Gorduino)
	g.pins = make(map[byte]bool)

	c := &serial.Config{Name: port, Baud: 57600}
	conn, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	g.client, _ = firmata.NewClient(conn)

	for _, pin := range pins {
		g.pins[pin] = false
		g.client.SetPinMode(pin, firmata.Output)
	}
	return g, nil
}

func (g *Gorduino) On(p byte) {
	g.client.DigitalWrite(uint(p), true)
}

func (g *Gorduino) Off(p byte) {
	g.client.DigitalWrite(uint(p), false)
}

func (g *Gorduino) Toggle(p byte) {
	g.pins[p] = !g.pins[p]
	g.client.DigitalWrite(uint(p), g.pins[p])
}
