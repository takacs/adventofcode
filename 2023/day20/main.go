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
	p                 PulseType
}

func (p Pulse) print() {
	fmt.Printf("%v %v -> %v\n", p.from, p.p, p.destination)
}

type Node interface {
	handlePulse(Pulse) []Pulse
	inDestination(string) bool
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

func (ff *FlipFlop) handlePulse(p Pulse) []Pulse {
	if p.p == HIGH {
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
		rpt = append(rpt, Pulse{from: ff.name, destination: d, p: rp})
	}
	return rpt
}

func (ff *FlipFlop) flip() {
	ff.state = !ff.state
}

func (ff FlipFlop) inDestination(s string) bool {
	for _, d := range ff.destination {
		if d == s { return true }
	}
	return false
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
	recentPulseType map[string]bool
}

func (c Conjunction) handlePulse(p Pulse) []Pulse {
	t := false
	if p.p == HIGH {
		t = true
	}
	c.recentPulseType[p.from] = t
	if c.allhigh() {
		rpt := []Pulse{}
		for _, d := range c.destination {
			rpt = append(rpt, Pulse{from: c.name, destination: d, p: LOW})
		}
		return rpt
	}
	rpt := []Pulse{}
	for _, d := range c.destination {
		rpt = append(rpt, Pulse{from: c.name, destination: d, p: HIGH})
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

func (c Conjunction) inDestination(s string) bool {
	for _, d := range c.destination {
		if d == s { return true }
	}
	return false
}

type Broadcast struct {
	/*
		There is a single broadcast module (named broadcaster).
		When it receives a pulse, it sends the same pulse to all of its destination modules.
	*/
	name        string
	destination []string
}

func (b Broadcast) handlePulse(p Pulse) []Pulse {
	rpt := []Pulse{}
	for _, d := range b.destination {
		rpt = append(rpt, Pulse{from: b.name, destination: d, p: p.p})
	}
	return rpt
}

func isin(d string, n map[string]Node) bool {
	for k, _ := range n {
		if k == d {
			return true
		}
	}
	return false
}

func isConjunction(t interface{}) bool {
	switch t.(type) {
	case *Conjunction:
		return true
	default:
		return false
	}
}

func (b Broadcast) inDestination(s string) bool {
	for _, d := range b.destination {
		if d == s { return true }
	}
	return false
}

func main() {
	f, err := os.ReadFile("day20.in")
	if err != nil {
		fmt.Print(err)
	}
	fs := string(f)
	instructions := strings.Split(fs, "\n")
	nodeBase := map[string]Node{}
	for _, l := range instructions {
		if len(l) < 2 {
			continue
		}
		i := strings.Split(l, " -> ")
		i[1] = strings.ReplaceAll(i[1], ",", "")
		if strings.HasPrefix(i[0], "broad") {
			nodeBase["broadcast"] = &Broadcast{name: "broadcast", destination: strings.Fields(i[1])}
		} else if strings.HasPrefix(i[0], "%") {
			nodeBase[i[0][1:]] = &FlipFlop{name: i[0][1:], destination: strings.Fields(i[1]), state: false}
		} else if strings.HasPrefix(i[0], "&") {
			nodeBase[i[0][1:]] = &Conjunction{name: i[0][1:], destination: strings.Fields(i[1]), recentPulseType: map[string]bool{}}
		}
	}

	for _, l := range instructions {
		if len(l) < 2 {
			continue
		}
		i := strings.Split(l, " -> ")
		i[0] = strings.ReplaceAll(i[0], "%", "")
		i[0] = strings.ReplaceAll(i[0], "&", "")
		i[1] = strings.ReplaceAll(i[1], ",", "")
		for _, f := range strings.Fields(i[1]) {
			v, _ := nodeBase[f]
			t := isConjunction(v)
			if t {
				v.handlePulse(Pulse{from: i[0], destination: "trash", p: LOW})
			}
		}
	}


	pq := []Pulse{}
	lo, hi := 0, 0
	nodes := nodeBase
	for i := 0; i < 1000; i++ {
		pq = append(pq, Pulse{from: "button", destination: "broadcast", p: LOW})
		for len(pq) > 0 {
			p := pq[0]
			// p.print()
			if p.p == LOW {
				lo++
			}
			if p.p == HIGH {
				hi++
			}
			pq = pq[1:]
			if !isin(p.destination, nodes) {
				continue
			}
			newpulse := nodes[p.destination].handlePulse(p)
			for _, np := range newpulse {
				pq = append(pq, np)
			}
		}
	}

	fmt.Printf("Part 1: %v\n", hi*lo)
	rxFeed := ""
	for k, n := range nodeBase {
		if n.inDestination("rx") {
			rxFeed = k
		}
	}

	feeds := map[string]int{}
	for k, n := range nodeBase {
		if n.inDestination(rxFeed) {
			feeds[k] = -1
		}
	}
	pq = []Pulse{}
	nodes = nodeBase
	// Break and look at output to see cycles, then get lowest common multiple
	for i := 0; true; i++ {
		pq = append(pq, Pulse{from: "button", destination: "broadcast", p: LOW})
		for len(pq) > 0 {
			p := pq[0]
			pq = pq[1:]
			_, ok := feeds[p.from]

			if ok && p.p == HIGH {
				fmt.Println(i, p.from)
			} 

			if !isin(p.destination, nodes) {
				continue
			}
			newpulse := nodes[p.destination].handlePulse(p)
			for _, np := range newpulse {
				pq = append(pq, np)
			}
		}
	}

	fmt.Println(lo, hi, lo*hi)
}
