package days

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	rating     string
	cmp        string
	value      int
	workflow   string
	belongs_to string
}

type Workflow struct {
	steps  []Step
	amount int
	name   string
}

type Part struct {
	ratings map[string]int
}

func Day_19_parse_input(use_test_file bool) (workflows map[string]Workflow, parts []Part) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_19.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	workflows = map[string]Workflow{}

	doing_ratings := false
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if !doing_ratings && txt != "" {
			split := strings.Split(txt, "{")
			name := split[0]
			str_steps := strings.Split(split[1][:len(split[1])-1], ",")

			steps := []Step{}
			for i, s := range str_steps {
				if i == len(str_steps)-1 {
					workflow_or_end := s
					step := Step{workflow: workflow_or_end}
					steps = append(steps, step)
				} else {
					rating := s[0:1]
					cmp := s[1:2]
					colon_ind := utils.FindIndex[string](strings.Split(s, ""), ":")

					num := s[2:colon_ind]
					value, _ := strconv.Atoi(num)
					workflow_or_end := s[colon_ind+1:]
					step := Step{rating: rating, cmp: cmp, value: value, workflow: workflow_or_end, belongs_to: name}
					steps = append(steps, step)
				}
			}
			workflow := Workflow{steps: steps, name: name}
			workflows[name] = workflow
		} else if txt == "" {
			doing_ratings = true
		} else if doing_ratings && txt != "" {
			txt = txt[1 : len(txt)-1]
			split := strings.Split(txt, ",")
			part := Part{}
			part.ratings = map[string]int{}
			for _, s := range split {
				rating := s[0:1]
				num := s[2:]
				value, _ := strconv.Atoi(num)
				part.ratings[rating] = value
			}
			parts = append(parts, part)
		}
	}

	file.Close()
	return
}

func Day_19_Part_1() {
	// workflows, parts := Day_19_parse_input(true)
	workflows, parts := Day_19_parse_input(false)

	accepted := []Part{}

	for _, part := range parts {
		curr_workflow_name := "in"
		workflow := workflows[curr_workflow_name]
		i := 0
		done := 0
		for done < 1 {
			if curr_workflow_name == "A" {
				accepted = append(accepted, part)
				break
			} else if curr_workflow_name == "R" {
				break
			} else {
				step := workflow.steps[i]
				if step.cmp == "<" {
					if part.ratings[step.rating] < step.value {
						curr_workflow_name = step.workflow
						workflow = workflows[curr_workflow_name]
						i = 0
					} else {
						i++
					}
				} else if step.cmp == ">" {
					if part.ratings[step.rating] > step.value {
						curr_workflow_name = step.workflow
						workflow = workflows[curr_workflow_name]
						i = 0
					} else {
						i++
					}
				} else {
					curr_workflow_name = step.workflow
					workflow = workflows[curr_workflow_name]
					i = 0
				}
			}
		}
	}

	total := 0
	for _, part := range accepted {
		total = total + part.ratings["x"] + part.ratings["m"] + part.ratings["a"] + part.ratings["s"]
	}
	fmt.Println(total)
}

type PartRange struct {
	x []int
	m []int
	a []int
	s []int
}

func NewPartRange() PartRange {
	partrange := PartRange{}
	partrange.x = utils.Range(1, 4000+1, 1)
	partrange.m = utils.Range(1, 4000+1, 1)
	partrange.a = utils.Range(1, 4000+1, 1)
	partrange.s = utils.Range(1, 4000+1, 1)
	return partrange
}

func update_range(partrange PartRange, step Step, r []int) PartRange {
	new_range := NewPartRange()
	new_range.x = partrange.x
	new_range.m = partrange.m
	new_range.a = partrange.a
	new_range.s = partrange.s

	switch step.rating {
	case "x":
		new_range.x = r
	case "m":
		new_range.m = r
	case "a":
		new_range.a = r
	case "s":
		new_range.s = r
	}
	return new_range
}

func passes_check(partrange PartRange, step Step) (status string, pass, fail PartRange) {
	r := []int{}
	switch step.rating {
	case "x":
		r = partrange.x
	case "m":
		r = partrange.m
	case "a":
		r = partrange.a
	case "s":
		r = partrange.s
	}

	start, end := r[0], r[len(r)-1]
	rhs := step.value

	if step.cmp == "<" {
		if rhs <= start {
			status = "none"
		} else if end < rhs {
			status = "all"
		} else {
			status = "mix"
			pass = update_range(partrange, step, utils.Range(start, rhs, 1))
			fail = update_range(partrange, step, utils.Range(rhs, end+1, 1))
		}
	} else {
		if start > rhs {
			status = "none"
		} else if end <= rhs {
			status = "all"
		} else {
			status = "mix"
			pass = update_range(partrange, step, utils.Range(rhs+1, end+1, 1))
			fail = update_range(partrange, step, utils.Range(start, rhs+1, 1))
		}
	}
	return status, pass, fail
}

func Combos(partrange PartRange) (total int) {
	total = 1
	total = total * (partrange.x[len(partrange.x)-1] - partrange.x[0] + 1)
	total = total * (partrange.m[len(partrange.m)-1] - partrange.m[0] + 1)
	total = total * (partrange.a[len(partrange.a)-1] - partrange.a[0] + 1)
	total = total * (partrange.s[len(partrange.s)-1] - partrange.s[0] + 1)
	return total
}

func GetAccepted(partrange PartRange, workflow_str string, workflows map[string]Workflow) int {
	count := 0
	workflow := workflows[workflow_str]
	if workflow_str == "A" {
		count = count + Combos(partrange)
		return count
	} else if workflow_str == "R" {
		count = count + 0
	}
	for _, step := range workflow.steps {
		if step.cmp == ">" || step.cmp == "<" {
			status, pass, fail := passes_check(partrange, step)
			if status == "all" {
				count = count + GetAccepted(partrange, step.workflow, workflows)
			} else if status == "none" {
				count = count + 0
			} else {
				partrange = fail
				count = count + GetAccepted(pass, step.workflow, workflows)
			}
		} else {
			count = count + GetAccepted(partrange, step.workflow, workflows)
		}
	}
	return count
}

func Day_19_Part_2() {
	// workflows, _ := Day_19_parse_input(true)
	workflows, _ := Day_19_parse_input(false)

	partrange := NewPartRange()
	total := GetAccepted(partrange, "in", workflows)
	fmt.Println(total)
}
