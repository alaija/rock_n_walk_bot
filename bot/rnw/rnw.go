package rnw

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Seatup(text string) *[]*Driver {
	ds, ps := parseDriversAndPassangers(text)
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Weight > ps[j].Weight
	})
	ps, p := popRandomPassanger(ps)
	for p != nil {
		seatPassanger(&ds, *p)
		ps, p = popRandomPassanger(ps)
	}
	return &ds
}

func parseDriversAndPassangers(text string) ([]*Driver, []Passanger) {
	groups := strings.Fields(strings.ReplaceAll(text, "@rock_n_walk_bot", ""))
	passangers := []Passanger{}
	drivers := []*Driver{}
	for _, group := range groups {
		if isDriver(group) {
			addDriver(&drivers, group)
		} else if isPassanger(group) {
			addPassanger(&passangers, group)
		}
	}
	return drivers, passangers
}

func addDriver(drivers *[]*Driver, driver string) {
	result := strings.Split(driver, ":")
	name := result[0]
	freeSeats, _ := strconv.Atoi(result[1])
	*drivers = append(
		*drivers,
		&Driver{
			Name:       name,
			Passangers: make([]Passanger, 0, freeSeats),
		},
	)
}

func addPassanger(passangers *[]Passanger, passanger string) {
	weight := strings.Count(passanger, "@")
	*passangers = append(
		*passangers,
		Passanger{
			Name:   passanger,
			Weight: weight,
		},
	)
}

func seatPassanger(ds *[]*Driver, p Passanger) {
	min := 10
	dsf := []*Driver{}
	for _, d := range *ds {
		capacity := cap(d.Passangers)
		length := len(d.Passangers)
		free := capacity - length
		if free < p.Weight {
			continue
		}
		if free == capacity {
			dsf = append(dsf, d)
			min = -1
			continue
		}
		if length < min {
			min = length
			dsf = []*Driver{d}
			continue
		}
		if length == min {
			dsf = append(dsf, d)
		}
	}

	rand.Seed(time.Now().Unix())
	dsf[rand.Intn(len(dsf))].setPassanger(p.Name)
}

func popRandomPassanger(ps []Passanger) ([]Passanger, *Passanger) {
	if len(ps) > 0 {
		return ps[1:], &ps[0]
	}
	return []Passanger{}, nil
}
