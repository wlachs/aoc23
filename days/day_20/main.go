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
}

// messageQueue holding entries to process
type messageQueue struct {
	entries   []MessageQueueEntry
	lowCount  int
	highCount int
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

// Module is the general module interface implemented by all module types
type Module interface {
	pulse(p bool, src Module)
	addInput(m Module)
	addOutput(m Module)
}

// FlipFlop module with an internal state
type FlipFlop struct {
	name         string
	state        bool
	messageQueue MessageQueue
	outputNodes  []Module
}

// pulse handler of a FlipFlop module
func (receiver *FlipFlop) pulse(p bool, _ Module) {
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

// Conjunction module with upstream input memory
type Conjunction struct {
	name         string
	inputMemory  map[Module]bool
	messageQueue MessageQueue
	outputNodes  []Module
}

// pulse handler of a Conjunction module
func (receiver *Conjunction) pulse(p bool, src Module) {
	receiver.inputMemory[src] = p
	out := false
	for _, b := range receiver.inputMemory {
		if !b {
			out = true
			break
		}
	}
	for _, module := range receiver.outputNodes {
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

// Broadcaster module, sends the input to all outputs
type Broadcaster struct {
	messageQueue MessageQueue
	outputNodes  []Module
}

// pulse handler of the Broadcaster module
func (receiver *Broadcaster) pulse(p bool, _ Module) {
	for _, module := range receiver.outputNodes {
		receiver.messageQueue.add(MessageQueueEntry{source: receiver, target: module, pulse: p})
	}
}

// addInput adds the given node to the current node's input queue
func (receiver *Broadcaster) addInput(m Module) {
	//
}

// addOutput adds the given node to the current node's output queue
func (receiver *Broadcaster) addOutput(m Module) {
	receiver.outputNodes = append(receiver.outputNodes, m)
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
			if e.target != nil {
				e.target.pulse(e.pulse, e.source)
			}
		}
	}
	return strconv.Itoa(mq.lowCount * mq.highCount)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

// parseInput sets up the system with the input nodes
func parseInput(input []string) (*messageQueue, *Broadcaster) {
	mq := messageQueue{
		entries:   []MessageQueueEntry{},
		lowCount:  0,
		highCount: 0,
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
				name:         ff[1],
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
				name:         c[1],
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
			src.addOutput(target)
			if target != nil {
				target.addInput(src)
			}
		}
	}
	return &mq, &broadcaster
}
