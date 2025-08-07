package builder

import "fmt"

type Lamp struct {
	Material string
	Power    int
	Cost     int
}

func (l *Lamp) On() {
	fmt.Printf("lamp is on; %s %d %d \n", l.Material, l.Power, l.Cost)
}

type LampBuilder interface {
	Material()
	Power(v int)
	Cost()
	GetResult() Lamp
}

type SonyCompany struct {
	lamp Lamp
}

func (c *SonyCompany) Material() {
	c.lamp.Material = "iron"
}
func (c *SonyCompany) Power(v int) {
	c.lamp.Power = v
}
func (c *SonyCompany) Cost() {
	minPower := 1
	if c.lamp.Power > minPower {
		minPower = c.lamp.Power
	}
	c.lamp.Cost = minPower * 3
}
func (c *SonyCompany) GetResult() Lamp {
	return c.lamp
}

type HitachiCompany struct {
	lamp Lamp
}

func (c *HitachiCompany) Material() {
	c.lamp.Material = "plastmass"
}
func (c *HitachiCompany) Power(v int) {
	c.lamp.Power = v
}
func (c *HitachiCompany) Cost() {
	if c.lamp.Power == 0 {
		c.lamp.Cost = 0
	} else {
		c.lamp.Cost = c.lamp.Power * 2
	}
}
func (c *HitachiCompany) GetResult() Lamp {
	return c.lamp
}

type Dir struct {
	builder LampBuilder
}

func (d *Dir) ConstrChip() Lamp {
	d.builder.Material()
	d.builder.Power(3)
	d.builder.Cost()
	return d.builder.GetResult()
}

func (d *Dir) ConstrExpensive() Lamp {
	d.builder.Material()
	d.builder.Power(6)
	d.builder.Cost()
	return d.builder.GetResult()
}

func NewDir(builder LampBuilder) *Dir {
	return &Dir{builder: builder}
}
