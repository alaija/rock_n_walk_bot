package rnw

import "strings"

type (
	Driver struct {
		Name       string
		Passangers []Passanger
	}

	Passanger struct {
		Name   string
		Weight int
	}
)

func (d *Driver) Description() string {
	var result = "С " + d.Name + " поедут: "
	if len(d.Passangers) > 0 {
		for i, p := range d.Passangers {
			result = result + p.Name
			if i == len(d.Passangers)-1 {
				result += ".\n"
			} else {
				result += ", "
			}
		}
	} else {
		result = d.Name + " поедет один.\n"
	}
	return result
}

func (d *Driver) setPassanger(passanger string) {
	if isGroup(passanger) {
		for _, p := range strings.Split(passanger, "-") {
			d.setPassanger(p)
		}
	} else {
		d.Passangers = append(d.Passangers, Passanger{
			Name:   passanger,
			Weight: 1,
		})
	}
}

func isDriver(driver string) bool {
	return strings.Count(driver, ":") == 1
}

func isPassanger(passanger string) bool {
	return strings.Count(passanger, "@") > 0
}

func isGroup(passanger string) bool {
	return strings.Count(passanger, "@") > 1 && strings.Count(passanger, "-") > 0
}
