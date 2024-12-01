package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Broadcaster
type BroadCaster struct {
	dests []string
}

func (b BroadCaster) String() string {
	return fmt.Sprint("broadcaster -> " + strings.Join(b.dests, ", "))
}

func (b BroadCaster) type_of() string {
	return "BroadCaster"
}

// Flip-Flop
type FlipFlop struct {
	on    bool
	name  string
	dests []string
}

func (ff FlipFlop) type_of() string {
	return "Flip-Flop"
}

func (ff FlipFlop) String() string {
	return fmt.Sprint("%" + ff.name + " -> " + strings.Join(ff.dests, ", "))
}

func PrintFlipFlopMap(m map[string]FlipFlop) {
	for _, flipflop := range m {
		fmt.Println(flipflop)
	}
}

// Conjunction
type Conjunction struct {
	name    string
	sources map[string]Power
	dests   []string
}

func (c Conjunction) type_of() string {
	return "Conjunction"
}

func (c Conjunction) String() string {
	return fmt.Sprint("&" + c.name + " -> " + strings.Join(c.dests, ", "))
}

func PrintConjunctionMap(m map[string]Conjunction) {
	for _, conjunction := range m {
		fmt.Println(conjunction)
	}
}

func Day_20_parse_input(use_test_file bool) (broadcaster BroadCaster, flipflops map[string]FlipFlop, conjunctions map[string]Conjunction) {
	var filename string
	if !use_test_file {
		filename = "2023/inputs/Day_20.txt"
	} else {
		filename = "2023/inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	flipflops = map[string]FlipFlop{}
	conjunctions = map[string]Conjunction{}

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		split := strings.Split(txt, " -> ")

		if strings.Index(txt, "broadcaster") >= 0 {
			// Process broadcaster
			broadcaster.dests = strings.Split(split[1], ", ")
		} else {
			name := split[0][1:]
			dests := strings.Split(split[1], ", ")

			if split[0][0:1] == "%" {
				// Found Flip-Flop
				flipflops[name] = FlipFlop{name: name, dests: dests}

				// If this touches a conjunction omdel, update it it doesn't have it
				conj, prs := conjunctions[split[1]]
				if prs {
					_, mod_prs := conj.sources[name]
					if !mod_prs {
						conj.sources[name] = Low
					}
					conjunctions[split[1]] = conj
				}
			} else {
				// Found Conjunction
				sources := map[string]Power{}
				// Update with sources already in modules list

				for flipflop_name, flipflop := range flipflops {
					for _, dest := range flipflop.dests {
						if dest == name {
							_, prs := sources[flipflop_name]
							if !prs {
								sources[flipflop_name] = Low
							}
						}
					}
				}
				new_conj := Conjunction{name: name, dests: dests, sources: sources}
				conjunctions[name] = new_conj
			}
		}
	}

	file.Close()
	return

}

type Power int

const (
	High Power = iota
	Low
)

func (p Power) String() string {
	if p == High {
		return "-high"
	} else {
		return "-low"
	}
}

type Pulse struct {
	power  Power
	source string
	dest   string
}

func (p Pulse) String() string {
	return p.source + " " + fmt.Sprint(p.power) + "-> " + p.dest
}

func (ff *FlipFlop) GetPower() Power {
	if ff.on {
		return High
	} else {
		return Low
	}
}

func (ff *FlipFlop) SendPulses() (new_pulses []Pulse, high_pulses, low_pulses int) {
	power := ff.GetPower()
	for _, dest := range ff.dests {
		new_pulse := Pulse{source: ff.name, dest: dest, power: power}
		new_pulses = append(new_pulses, new_pulse)
		if power == High {
			high_pulses++
		} else {
			low_pulses++
		}
	}
	return
}

func (c *Conjunction) GetPower() Power {
	for _, source_power := range c.sources {
		if source_power == Low {
			return High
		}
	}
	return Low
}

func (c *Conjunction) SendPulses() (new_pulses []Pulse, high_pulses, low_pulses int) {
	power := c.GetPower()
	for _, dest := range c.dests {
		new_pulse := Pulse{source: c.name, dest: dest, power: power}
		new_pulses = append(new_pulses, new_pulse)
		if power == High {
			high_pulses++
		} else {
			low_pulses++
		}
	}
	return
}

func module_in_initial_state(flipflops map[string]FlipFlop, conjunctions map[string]Conjunction) bool {
	for _, flipflop := range flipflops {
		if flipflop.on {
			return false
		}
	}

	for _, conj := range conjunctions {
		for _, power := range conj.sources {
			if power == High {
				return false
			}
		}
	}
	return true
}

func send_pulses(broadcaster BroadCaster, flipflops map[string]FlipFlop, conjunctions map[string]Conjunction, amount int, is_part_two bool, stop_at string) int {
	high_pulses, low_pulses := 0, 0
	button_pressed_count := 0
	found_cycle := false
	for button_pressed_count < amount {
		// fmt.Println("at button press", button_pressed_count)
		queue := []Pulse{}

		// Print out the pulses like the instructions
		button_pushes_catalog := "button -low-> broadcaster\n"

		new_high_pulses, new_low_pulses := 0, 1 // Start at one for button module
		for _, dest := range broadcaster.dests {
			new_pulse := Pulse{power: Low, source: "broadcaster", dest: dest}
			queue = append(queue, new_pulse)
			// button_pushes_catalog = button_pushes_catalog + fmt.Sprintln(new_pulse)
			new_low_pulses++
		}

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			if is_part_two && pulse.source == stop_at && pulse.dest == "gf" && pulse.power == High {
				fmt.Println(pulse.source, button_pressed_count+1)
				return button_pressed_count + 1
			}

			button_pushes_catalog = button_pushes_catalog + fmt.Sprintln(pulse)

			flipflop, prs_ff := flipflops[pulse.dest]
			conj, prs_c := conjunctions[pulse.dest]
			if prs_ff {
				if pulse.power == High {
					// Ignored
				} else {
					flipflop.on = !flipflop.on
					new_pulses, high_count, low_count := flipflop.SendPulses()
					new_high_pulses = new_high_pulses + high_count
					new_low_pulses = new_low_pulses + low_count
					queue = append(queue, new_pulses...)
					flipflops[pulse.dest] = flipflop
				}
			} else if prs_c {
				conj.sources[pulse.source] = pulse.power
				new_pulses, high_count, low_count := conj.SendPulses()
				new_high_pulses = new_high_pulses + high_count
				new_low_pulses = new_low_pulses + low_count
				queue = append(queue, new_pulses...)
				conjunctions[pulse.dest] = conj
			} else {
				// Module with no receiver, skip
			}
		}

		button_pressed_count++
		high_pulses = high_pulses + new_high_pulses
		low_pulses = low_pulses + new_low_pulses

		in_initial_state := module_in_initial_state(flipflops, conjunctions)
		if in_initial_state && !found_cycle && !is_part_two {
			cycle_length := button_pressed_count

			left := amount - button_pressed_count
			to_mul := int(math.Floor(float64(left / cycle_length)))

			high_pulses = high_pulses + (high_pulses * to_mul)
			low_pulses = low_pulses + (low_pulses * to_mul)

			button_pressed_count = button_pressed_count + (cycle_length * to_mul)
			found_cycle = true
		}
	}
	return high_pulses * low_pulses
}

func reset(flipflops map[string]FlipFlop, conjunctions map[string]Conjunction) (map[string]FlipFlop, map[string]Conjunction) {
	for name, flipflop := range flipflops {
		flipflop.on = false
		flipflops[name] = flipflop
	}

	for name, conjunction := range conjunctions {
		sources := map[string]Power{}
		for source := range conjunction.sources {
			sources[source] = Low
		}
		conjunction.sources = sources
		conjunctions[name] = conjunction
	}
	return flipflops, conjunctions
}

func Day_20_Part_1() {
	// broadcaster, flipflops, conjunctions := Day_20_parse_input(true)
	broadcaster, flipflops, conjunctions := Day_20_parse_input(false)
	// fmt.Println(broadcaster)
	// PrintFlipFlopMap(flipflops)
	// PrintConjunctionMap(conjunctions)

	count := send_pulses(broadcaster, flipflops, conjunctions, 1000, false, "")
	fmt.Println(count)
}

func Day_20_Part_2() {
	// broadcaster, flipflops, conjunctions := Day_20_parse_input(true)
	broadcaster, flipflops, conjunctions := Day_20_parse_input(false)

	zs := send_pulses(broadcaster, flipflops, conjunctions, 1e9, true, "zs")
	flipflops, conjunctions = reset(flipflops, conjunctions)
	kr := send_pulses(broadcaster, flipflops, conjunctions, 1e9, true, "kr")
	flipflops, conjunctions = reset(flipflops, conjunctions)
	kf := send_pulses(broadcaster, flipflops, conjunctions, 1e9, true, "kf")
	flipflops, conjunctions = reset(flipflops, conjunctions)
	qk := send_pulses(broadcaster, flipflops, conjunctions, 1e9, true, "qk")
	flipflops, conjunctions = reset(flipflops, conjunctions)

	fmt.Println("zs", zs, "kr", kr, "kf", kf, "qk", qk)

	lcm := utils.LCMMultiple([]int{zs, kr, kf, qk})
	fmt.Println(lcm)
}
