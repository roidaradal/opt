package problem

import (
	"fmt"
	"slices"
	"strings"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/number"
	"github.com/roidaradal/fn/str"
	"github.com/roidaradal/opt/discrete"
	"github.com/roidaradal/opt/fn"
)

// Create new Weapon Target Assignment problem
func WeaponTarget(n int) *discrete.Problem {
	name := newName(WEAPON, n)
	cfg := newWeaponTarget(name)
	if cfg == nil {
		return nil
	}
	numWeapons, numTargets := len(cfg.weapons), len(cfg.targets)

	p := discrete.NewProblem(name)
	p.Goal = discrete.Minimize
	p.Type = discrete.Assignment

	// Expand weapons and count into weapons list (uses weapon index)
	weapons := make([]int, 0)
	for i := range numWeapons {
		weapons = append(weapons, slices.Repeat([]int{i}, cfg.count[i])...)
	}

	p.Variables = discrete.Variables(weapons)
	domain := discrete.MapDomain(cfg.targets)
	for _, variable := range p.Variables {
		p.Domain[variable] = domain[:]
	}

	p.ObjectiveFn = func(solution *discrete.Solution) discrete.Score {
		// Compute survival rate of each target, with weapons assigned to attack it
		survival := list.Copy(cfg.value)
		for w, target := range solution.Map {
			weapon := weapons[w]
			survival[target] *= 1 - cfg.chance[weapon][target] // survival = 1 - weaponOnTargetEffectiveness
		}
		total := fmt.Sprintf("%.4f", list.Sum(survival)) // round to 4 decimal places
		return discrete.Score(number.ParseFloat(total))
	}

	weaponTargets := func(solution *discrete.Solution) string {
		// Group the count of weapon => target assignments
		matrix := make([][]int, numWeapons)
		for i := range numWeapons {
			matrix[i] = make([]int, numTargets)
		}
		for w, target := range solution.Map {
			weapon := weapons[w]
			matrix[weapon][target] += 1
		}
		output := make([]string, 0)
		for i, weapon := range cfg.weapons {
			for j, target := range cfg.targets {
				if matrix[i][j] == 0 {
					continue // skip empty count
				}
				line := fmt.Sprintf("%d*%s = %s", matrix[i][j], weapon, target)
				output = append(output, line)
			}
		}
		return str.WrapBraces(output)
	}

	p.SolutionStringFn = weaponTargets
	p.SolutionCoreFn = weaponTargets

	return p
}

type weaponCfg struct {
	weapons []string
	targets []string
	count   []int       // count of each weapon type
	value   []float64   // value of each target
	chance  [][]float64 // effectivity matrix: weapon x value
}

// Load weapon target test case
func newWeaponTarget(name string) *weaponCfg {
	lines, err := fn.LoadProblem(name)
	if err != nil || len(lines) < 5 {
		return nil
	}
	cfg := &weaponCfg{
		weapons: strings.Fields(lines[0]),
		count:   list.Map(strings.Fields(lines[1]), number.ParseInt),
		targets: strings.Fields(lines[2]),
		value:   list.Map(strings.Fields(lines[3]), number.ParseFloat),
	}
	cfg.chance = make([][]float64, len(cfg.weapons))
	for i, line := range lines[4:] {
		cfg.chance[i] = list.Map(strings.Fields(line), number.ParseFloat)
	}
	return cfg
}
