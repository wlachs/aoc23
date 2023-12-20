package day_20

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// MessageQueueEntry holds a tuple of Module pointer and input pulse
type MessageQueueEntry struct {
	source Module
	target Module
	pulse  bool
}

// MessageQueue interface
type MessageQueue interface {
	add(entry MessageQueueEntry)
	pop() MessageQueueEntry
	addFinalState(m Module, states int, iter int)
}

// messageQueue holding entries to process
type messageQueue struct {
	entries      []MessageQueueEntry
	lowCount     int
	highCount    int
	lastCount    int
	lastStateMap map[Module]int
}

// add method adds a new MessageQueueEntry to the MessageQueue
func (receiver *messageQueue) add(entry MessageQueueEntry) {
	receiver.entries = append(receiver.entries, entry)
	if entry.pulse {
		receiver.highCount++
	} else {
		receiver.lowCount++
	}
}

// pop removes the first entry of the queue and returns it
func (receiver *messageQueue) pop() MessageQueueEntry {
	e := receiver.entries[0]
	receiver.entries = receiver.entries[1:]
	return e
}

// addFinalState adds a terminating criteria module to the message queue
func (receiver *messageQueue) addFinalState(m Module, states int, iter int) {
	receiver.lastCount = states
	_, ok := receiver.lastStateMap[m]
	if !ok {
		receiver.lastStateMap[m] = iter
	}
}

// finalState gets the final RX state iteration count
func (receiver *messageQueue) finalState() int {
	c := 0
	product := 1
	for _, iter := range receiver.lastStateMap {
		c++
		product = lcm(product, iter)
	}
	if c > 0 && c == receiver.lastCount {
		return product
	}
	return 0
}

// Module is the general module interface implemented by all module types
type Module interface {
	pulse(p bool, src Module, iter int)
	addInput(m Module)
	addOutput(m Module)
	isRX() bool
}

// FlipFlop module with an internal state
type FlipFlop struct {
	state        bool
	messageQueue MessageQueue
	outputNodes  []Module
}

// pulse handler of a FlipFlop module
func (receiver *FlipFlop) pulse(p bool, _ Module, _ int) {
	if p {
		return
	}
	receiver.state = !receiver.state
	for _, module := range receiver.outputNodes {
		receiver.messageQueue.add(MessageQueueEntry{source: receiver, target: module, pulse: receiver.state})
	}
}

// addInput adds the given node to the current node's input queue
func (receiver *FlipFlop) addInput(_ Module) {
	//
}

// addOutput adds the given node to the current node's output queue
func (receiver *FlipFlop) addOutput(m Module) {
	receiver.outputNodes = append(receiver.outputNodes, m)
}

// isRX checks whether the current module is the final one
func (receiver *FlipFlop) isRX() bool {
	return false
}

// Conjunction module with upstream input memory
type Conjunction struct {
	inputMemory  map[Module]bool
	messageQueue MessageQueue
	outputNodes  []Module
}

// pulse handler of a Conjunction module
func (receiver *Conjunction) pulse(p bool, src Module, iter int) {
	receiver.inputMemory[src] = p
	out := false
	for _, b := range receiver.inputMemory {
		if !b {
			out = true
			break
		}
	}
	for _, module := range receiver.outputNodes {
		if module.isRX() {
			for m, b := range receiver.inputMemory {
				if b {
					receiver.messageQueue.addFinalState(m, len(receiver.inputMemory), iter)
				}
			}
		}
		receiver.messageQueue.add(MessageQueueEntry{source: receiver, target: module, pulse: out})
	}
}

// addInput adds the given node to the current node's input queue
func (receiver *Conjunction) addInput(m Module) {
	receiver.inputMemory[m] = false
}

// addOutput adds the given node to the current node's output queue
func (receiver *Conjunction) addOutput(m Module) {
	receiver.outputNodes = append(receiver.outputNodes, m)
}

// isRX checks whether the current module is the final one
func (receiver *Conjunction) isRX() bool {
	return false
}

// Broadcaster module, sends the input to all outputs
type Broadcaster struct {
	messageQueue MessageQueue
	outputNodes  []Module
}

// pulse handler of the Broadcaster module
func (receiver *Broadcaster) pulse(p bool, _ Module, _ int) {
	for _, module := range receiver.outputNodes {
		receiver.messageQueue.add(MessageQueueEntry{source: receiver, target: module, pulse: p})
	}
}

// addInput adds the given node to the current node's input queue
func (receiver *Broadcaster) addInput(_ Module) {
	//
}

// addOutput adds the given node to the current node's output queue
func (receiver *Broadcaster) addOutput(m Module) {
	receiver.outputNodes = append(receiver.outputNodes, m)
}

// isRX checks whether the current module is the final one
func (receiver *Broadcaster) isRX() bool {
	return false
}

// Dummy module holding noop actions
type Dummy struct {
	name string
}

// pulse handler of the Broadcaster module
func (receiver *Dummy) pulse(_ bool, _ Module, _ int) {
	//
}

// addInput adds the given node to the current node's input queue
func (receiver *Dummy) addInput(_ Module) {
	//
}

// addOutput adds the given node to the current node's output queue
func (receiver *Dummy) addOutput(_ Module) {
	//
}

// isRX checks whether the current module is the final one
func (receiver *Dummy) isRX() bool {
	return receiver.name == "rx"
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	mq, broadcaster := parseInput(input)
	for i := 0; i < 1000; i++ {
		mq.add(MessageQueueEntry{
			target: broadcaster,
			pulse:  false,
		})
		for len(mq.entries) > 0 {
			e := mq.pop()
			e.target.pulse(e.pulse, e.source, i)
		}
	}
	return strconv.Itoa(mq.lowCount * mq.highCount)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	mq, broadcaster := parseInput(input)
	for i := 1; ; i++ {
		fs := mq.finalState()
		if fs > 0 {
			return strconv.Itoa(fs)
		}
		mq.add(MessageQueueEntry{
			target: broadcaster,
			pulse:  false,
		})
		for len(mq.entries) > 0 {
			e := mq.pop()
			if e.target.isRX() && !e.pulse {
				return strconv.Itoa(i)
			}
			e.target.pulse(e.pulse, e.source, i)
		}
	}
}

// parseInput sets up the system with the input nodes
func parseInput(input []string) (*messageQueue, *Broadcaster) {
	mq := messageQueue{
		entries:      []MessageQueueEntry{},
		lowCount:     0,
		highCount:    0,
		lastCount:    0,
		lastStateMap: map[Module]int{},
	}
	moduleMap := map[string]Module{}
	flipFlops := map[string]FlipFlop{}
	conjunctions := map[string]Conjunction{}
	broadcaster := Broadcaster{
		messageQueue: &mq,
		outputNodes:  []Module{},
	}
	moduleMap["broadcaster"] = &broadcaster
	for _, row := range input {
		ff := regexp.MustCompile(`%(?P<a>\w+)`).FindStringSubmatch(row)
		if len(ff) > 0 {
			m := FlipFlop{
				state:        false,
				messageQueue: &mq,
				outputNodes:  []Module{},
			}
			moduleMap[ff[1]] = &m
			flipFlops[ff[1]] = m
			continue
		}
		c := regexp.MustCompile(`&(?P<a>\w+)`).FindStringSubmatch(row)
		if len(c) > 0 {
			m := Conjunction{
				inputMemory:  map[Module]bool{},
				messageQueue: &mq,
				outputNodes:  []Module{},
			}
			moduleMap[c[1]] = &m
			conjunctions[c[1]] = m
			continue
		}
	}
	for _, row := range input {
		r := regexp.MustCompile(`^[%&]*(?P<a>\w+) -> (?P<b>.+$)`).FindStringSubmatch(row)
		src := moduleMap[r[1]]
		for _, t := range strings.Split(r[2], ", ") {
			target := moduleMap[t]
			if target == nil {
				target = &Dummy{name: t}
			}
			src.addOutput(target)
			target.addInput(src)
		}
	}
	return &mq, &broadcaster
}

// gcd calculates the greatest common divisor of a and b.
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// lcm calculates the least common multiple of a and b.
func lcm(a int, b int) int {
	return (a / gcd(a, b)) * b
}
