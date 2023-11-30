package internal

import (
	"github.com/wlchs/advent_of_code_go_template/days/day_0"
	"github.com/wlchs/advent_of_code_go_template/days/day_01"
	"github.com/wlchs/advent_of_code_go_template/days/day_02"
	"github.com/wlchs/advent_of_code_go_template/days/day_03"
	"github.com/wlchs/advent_of_code_go_template/days/day_04"
	"github.com/wlchs/advent_of_code_go_template/days/day_05"
	"github.com/wlchs/advent_of_code_go_template/days/day_06"
	"github.com/wlchs/advent_of_code_go_template/days/day_07"
	"github.com/wlchs/advent_of_code_go_template/days/day_08"
	"github.com/wlchs/advent_of_code_go_template/days/day_09"
	"github.com/wlchs/advent_of_code_go_template/days/day_10"
	"github.com/wlchs/advent_of_code_go_template/days/day_11"
	"github.com/wlchs/advent_of_code_go_template/days/day_12"
	"github.com/wlchs/advent_of_code_go_template/days/day_13"
	"github.com/wlchs/advent_of_code_go_template/days/day_14"
	"github.com/wlchs/advent_of_code_go_template/days/day_15"
	"github.com/wlchs/advent_of_code_go_template/days/day_16"
	"github.com/wlchs/advent_of_code_go_template/days/day_17"
	"github.com/wlchs/advent_of_code_go_template/days/day_18"
	"github.com/wlchs/advent_of_code_go_template/days/day_19"
	"github.com/wlchs/advent_of_code_go_template/days/day_20"
	"github.com/wlchs/advent_of_code_go_template/days/day_21"
	"github.com/wlchs/advent_of_code_go_template/days/day_22"
	"github.com/wlchs/advent_of_code_go_template/days/day_23"
	"github.com/wlchs/advent_of_code_go_template/days/day_24"
)

// RunChallenge executes the challenge of a specific day with the provided input.
func RunChallenge(day int, inputPath string, mode int) {
	input := LoadInputLines(inputPath)
	mapping := map[int]func([]string, int){
		0:  day_0.Run,
		1:  day_01.Run,
		2:  day_02.Run,
		3:  day_03.Run,
		4:  day_04.Run,
		5:  day_05.Run,
		6:  day_06.Run,
		7:  day_07.Run,
		8:  day_08.Run,
		9:  day_09.Run,
		10: day_10.Run,
		11: day_11.Run,
		12: day_12.Run,
		13: day_13.Run,
		14: day_14.Run,
		15: day_15.Run,
		16: day_16.Run,
		17: day_17.Run,
		18: day_18.Run,
		19: day_19.Run,
		20: day_20.Run,
		21: day_21.Run,
		22: day_22.Run,
		23: day_23.Run,
		24: day_24.Run,
	}

	mapping[day](input, mode)
}
