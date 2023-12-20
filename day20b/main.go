package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Module represents a module with its type and state
type Module struct {
	Type     string
	Name     string
	State    int // 0 for off, 1 for on
	Outs     []string
	Remember map[string]int // Remember the type of the most recent pulse
}

// ProcessPulse processes a pulse for the given module
func (m *Module) ProcessPulse(pulse int, modules map[string]*Module, counters map[int]int, from string, queue []QueueEntry) []QueueEntry {
	switch m.Type {
	case "%":
		if pulse == 0 {
			m.State = 1 - m.State
			if m.State == 1 {
				pulse = 1
			} else {
				pulse = 0
			}
			for _, out := range m.Outs {
				counters[pulse]++
				queue = append(queue, QueueEntry{Module: out, From: m.Name, Pulse: pulse})
			}
		}

	case "&":
		m.Remember[from] = pulse
		allHigh := true
		for _, input := range m.Remember {
			if input != 1 {
				allHigh = false
				break
			}
		}
		pulse = 1
		if allHigh {
			pulse = 0
		}
		for _, out := range m.Outs {
			counters[pulse]++
			queue = append(queue, QueueEntry{Module: out, From: m.Name, Pulse: pulse})
		}
	case "broadcaster":
		for _, out := range m.Outs {
			counters[pulse]++
			queue = append(queue, QueueEntry{Module: out, From: m.Name, Pulse: pulse})
		}
	}

	return queue
}

// QueueEntry represents an entry in the queue containing the destination module and pulse value
type QueueEntry struct {
	Module string
	From   string
	Pulse  int
}

// ParseModule parses a module configuration line and returns a Module
func ParseModule(line string, modules map[string]*Module) *Module {
	parts := strings.Fields(line)
	name := parts[0]
	typ := name
	if strings.HasPrefix(name, "%") || strings.HasPrefix(name, "&") {
		typ = name[0:1]
		name = name[1:]
	}

	module := &Module{
		Type:     typ,
		Name:     name,
		State:    0,
		Outs:     []string{},
		Remember: make(map[string]int),
	}

	for _, dest := range parts[2:] {
		dest = strings.Trim(dest, ",")
		module.Outs = append(module.Outs, dest)
		if modules[dest] == nil {
			modules[dest] = &Module{
				Type:     "",
				Name:     dest,
				State:    0,
				Outs:     []string{},
				Remember: make(map[string]int),
			}
		}
	}

	modules[name] = module
	return module
}

func FindInputs(modules map[string]*Module, dest string) []string {
	inputs := make([]string, 0)
	for k, m := range modules {
		for _, t := range m.Outs {
			if t == dest {
				inputs = append(inputs, k)
			}
		}
	}
	return inputs
}

func main() {
	fmt.Println("show me your input:")
	scanner := bufio.NewScanner(os.Stdin)

	modules := make(map[string]*Module)
	for {
		scanner.Scan()
		line := scanner.Text()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		if line == "" {
			break
		}
		ParseModule(line, modules)
	}
	for _, module := range modules {
		for _, conn := range module.Outs {
			modules[conn].Remember[module.Name] = 0
		}
	}

	inRx := FindInputs(modules, "rx")[0]
	inputs := FindInputs(modules, inRx)

	inputsCounters := map[string]int{}
	pulseCounters := map[int]int{}
	for i := 1; len(inputsCounters) != len(inputs); i++ {
		pulse := 0
		pulseCounters[pulse]++

		// BFS processing
		queue := []QueueEntry{{"broadcaster", "button", 0}}
		for len(queue) > 0 {
			currentModule := queue[0]
			queue = queue[1:]
			module := modules[currentModule.Module]

			queue = module.ProcessPulse(currentModule.Pulse, modules, pulseCounters, currentModule.From, queue)

			for _, k := range inputs {
				_, ok := inputsCounters[k]
				if !ok && modules[inRx].Remember[k] == 1 {
					inputsCounters[k] = i
				}
			}
		}
	}
	products := []int{}
	for _, v := range inputsCounters {
		products = append(products, v)
	}
	fmt.Println("Answer:", lcm(products[0], products[1], products[2:]...))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
