package main

import (
	"fmt"
	"os"
	"strings"
)

type PulseType int

const (
	LOW PulseType = iota
	HIGH
	NO
)

type Pulse struct {
	from, destination string
	p           PulseType
}

func (p Pulse) print() {
	fmt.Printf("%v %v -> %v\n", p.from, p.p, p.destination)
}

type Node interface {
	handlePulse(PulseType) []Pulse
}

type FlipFlop struct {
	/*
		Flip-flop modules (prefix %) are either on or off; they are initially off.
		If a flip-flop module receives a high pulse, it is ignored and nothing happens.
		However, if a flip-flop module receives a low pulse, it flips between on and off.
		If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.
	*/
	name        string
	destination []string
	state       bool
}

func (ff FlipFlop) handlePulse(p PulseType) []Pulse {
	if p == HIGH {
		return []Pulse{}
	}
	rp := NO
	if ff.state {
		rp = LOW
	} else {
		rp = HIGH
	}
	ff.flip()
	rpt := []Pulse{}
	for _, d := range ff.destination {
		rpt = append(rpt, Pulse{from:ff.name, destination: d, p: rp})
	}
	return rpt
}

func (ff *FlipFlop) flip() {
	ff.state = !ff.state
}

type Conjunction struct {
	/*
		Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules;
		they initially default to remembering a low pulse for each input.
		When a pulse is received, the conjunction module first updates its memory for that input.
		Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.
	*/
	name            string
	destination     []string
	recentPulseType map[Node]bool
}

func (c Conjunction) handlePulse(p PulseType) []Pulse {
	if c.allhigh() {
		rpt := []Pulse{}
		for _, d := range c.destination {
			rpt = append(rpt, Pulse{from: c.name, destination: d, p: LOW})
		}
		return rpt
	}
	rpt := []Pulse{}
	for _, d := range c.destination {
		rpt = append(rpt, Pulse{from: c.name, destination: d, p: LOW})
	}
	return rpt
}

func (c Conjunction) allhigh() bool {
	b := true
	for _, v := range c.recentPulseType {
		b = b && v
	}
	return b
}

type Broadcast struct {
	/*
		There is a single broadcast module (named broadcaster).
		When it receives a pulse, it sends the same pulse to all of its destination modules.
	*/
	name        string
	destination []string
}

func (b Broadcast) handlePulse(p PulseType) []Pulse {
	rpt := []Pulse{}
	for _, d := range b.destination {
		rpt = append(rpt, Pulse{from: b.name, destination: d, p: p})
	}
	return rpt
}

func main() {
	f, err := os.ReadFile("day20.test")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)
	instructions := strings.Split(fs, "\n")
	nodes := map[string]Node{}
	for _, l := range instructions {
		if len(l) < 2 {
			continue
		}
		i := strings.Split(l, " -> ")
		i[1] = strings.ReplaceAll(i[1], ",", "")
		if strings.HasPrefix(i[0], "broad") {
			nodes["broadcast"] = Broadcast{name: "broadcast", destination: strings.Fields(i[1])}
		} else if strings.HasPrefix(i[0], "%") {
			nodes[i[0][1:]] = FlipFlop{name: i[0][1:], destination: strings.Fields(i[1]), state: false}
		} else if strings.HasPrefix(i[0], "&") {
			nodes[i[0][1:]] = Conjunction{name: i[0][1:], destination: strings.Fields(i[1])}
		}
	}

	fmt.Println(nodes)

	pq := []Pulse{}
	lo, hi := 0, 0
	for i:=0;i<1;i++ {
		pq = append(pq, Pulse{from: "button", destination: "broadcast", p: LOW})
		for len(pq) > 0 {
			p := pq[0]
			p.print()
			if p.p == LOW {
				lo++
			}
			if p.p == HIGH {
				hi++
			}
			pq = pq[1:]
			if p.destination == "output" {
				continue
			}	
			newpulse := nodes[p.destination].handlePulse(p.p)
			for _, np := range newpulse {
				pq = append(pq, np)
			}
		}
	}

	fmt.Println(lo, hi)
}
