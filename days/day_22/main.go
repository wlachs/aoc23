package day_22

import (
	"fmt"
	"github.com/wlchs/advent_of_code_go_template/types"
	"github.com/wlchs/advent_of_code_go_template/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Brick is a 3D object defined by its 2 corners across one of its diagonals
type Brick struct {
	cornerA     types.Vec3
	cornerB     types.Vec3
	supportedBy []*Brick
	supports    []*Brick
}

// minZ calculates the minimal Z coordinates of the brick
func (b Brick) minZ() int {
	return min(b.cornerA.Z, b.cornerB.Z)
}

// maxZ calculates the maximal Z coordinates of the brick
func (b Brick) maxZ() int {
	return max(b.cornerA.Z, b.cornerB.Z)
}

// intersects checks whether the given brick intersects a specific 2D section of the plane
func (b Brick) intersects(left types.Vec2, right types.Vec2) bool {
	b1, b2 := b.cuttingPlane()
	if left.X <= b2.X && left.Y <= b2.Y && right.X >= b1.X && right.Y >= b1.Y {
		return true
	}
	return false
}

// cuttingPlane returns the 2D plane of the brick from above
func (b Brick) cuttingPlane() (types.Vec2, types.Vec2) {
	return types.Vec2{X: min(b.cornerA.X, b.cornerB.X), Y: min(b.cornerA.Y, b.cornerB.Y)}, types.Vec2{X: max(b.cornerA.X, b.cornerB.X), Y: max(b.cornerA.Y, b.cornerB.Y)}
}

// pushDown finds the lowest possible position of the given brick based on the bricks that are already laying on the ground
func (b Brick) pushDown(bricks []Brick) Brick {
	b1, b2 := b.cuttingPlane()
	minZ := b.minZ()
	highest := highestUnderArea(b1, b2, minZ-1, bricks)
	below := 0
	if highest != nil {
		below = highest.maxZ()
	}
	return Brick{
		cornerA: types.Vec3{
			X: b.cornerA.X,
			Y: b.cornerA.Y,
			Z: b.cornerA.Z - minZ + below + 1,
		},
		cornerB: types.Vec3{
			X: b.cornerB.X,
			Y: b.cornerB.Y,
			Z: b.cornerB.Z - minZ + below + 1,
		},
	}
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
	bricks := prepareBricks(input)
	count := 0
	for _, brick := range bricks {
		canBeRemoved := true
		for _, support := range brick.supports {
			if len(support.supportedBy) < 2 {
				canBeRemoved = false
				break
			}
		}
		if canBeRemoved {
			count++
		}
	}
	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	bricks := prepareBricks(input)
	sum := 0
	for i := range bricks {
		sum += fallingBricksCount(bricks[i+1:], []*Brick{&bricks[i]})
	}
	return strconv.Itoa(sum)
}

// prepareBricks finds each bricks lowest final location and build the brick dependency graph
func prepareBricks(input []string) []Brick {
	bricks := readInput(input)
	finalBricks := make([]Brick, 0, len(bricks))
	for len(bricks) > 0 {
		l, rest := lowest(bricks)
		bricks = rest
		finalBricks = append(finalBricks, l.pushDown(finalBricks))
	}
	buildSupportStructures(finalBricks)
	return finalBricks
}

// readInput reads the input and generates the bricks
func readInput(input []string) []Brick {
	bricks := make([]Brick, 0, len(input))
	for _, row := range input {
		vectors := make([]types.Vec3, 0, 2)
		corners := strings.Split(row, "~")
		for _, corner := range corners {
			coords := strings.Split(corner, ",")
			vectors = append(vectors, types.Vec3{
				X: utils.Atoi(coords[0]),
				Y: utils.Atoi(coords[1]),
				Z: utils.Atoi(coords[2]),
			})
		}
		bricks = append(bricks, Brick{
			cornerA: vectors[0],
			cornerB: vectors[1],
		})
	}
	return bricks
}

// lowest finds the Brick with the lowest Z coordinates and returns the rest of the Bricks.
func lowest(bricks []Brick) (Brick, []Brick) {
	minIndex := 0
	minZ := math.MaxInt
	var minBrick Brick
	for i, brick := range bricks {
		if brick.cornerA.Z < minZ || brick.cornerB.Z < minZ {
			minZ = min(brick.cornerA.Z, brick.cornerB.Z)
			minBrick = brick
			minIndex = i
		}
	}
	bricks[minIndex] = bricks[len(bricks)-1]
	return minBrick, bricks[:len(bricks)-1]
}

// highestUnderArea finds the highest brick that is under the given section of the area seen from above
func highestUnderArea(a types.Vec2, b types.Vec2, maxZ int, bricks []Brick) *Brick {
	topLeft := types.Vec2{X: min(a.X, b.X), Y: min(a.Y, b.Y)}
	bottomRight := types.Vec2{X: max(a.X, b.X), Y: max(a.Y, b.Y)}
	var maxB *Brick
	for i, brick := range bricks {
		if brick.intersects(topLeft, bottomRight) && brick.maxZ() <= maxZ && (maxB == nil || brick.maxZ() > maxB.maxZ()) {
			maxB = &bricks[i]
		}
	}
	return maxB
}

// buildSupportStructures analyzes the bricks and sets the support relationships between them
func buildSupportStructures(bricks []Brick) {
	for i := range bricks {
		buildSupportStructure(&bricks[i], bricks[:i])
	}
}

// buildSupportStructure sets the support relationships originating from and targeted towards the given brick
func buildSupportStructure(brick *Brick, bricks []Brick) {
	bottom := brick.minZ()
	leftCorner, rightCorner := brick.cuttingPlane()
	for i, b := range bricks {
		if b.maxZ() == bottom-1 && b.intersects(leftCorner, rightCorner) {
			bricks[i].supports = append(b.supports, brick)
			brick.supportedBy = append(brick.supportedBy, &bricks[i])
		}
	}
}

// fallingBricksCount recursively calculates the overall number of falling bricks if the provided list of bricks are removed from the system
func fallingBricksCount(bricks []Brick, removedBricks []*Brick) int {
	count := 0
	for i, brick := range bricks {
		willFall := true
		for _, supportedBy := range brick.supportedBy {
			if !slices.Contains(removedBricks, supportedBy) {
				willFall = false
			}
		}
		if willFall && brick.minZ() != 1 {
			removedBricks = append(removedBricks, &bricks[i])
			count++
		}
	}
	return count
}
