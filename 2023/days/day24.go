package days

import (
	"AdventOfCode/models"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
    position models.CoordFloat64
    velocity models.CoordFloat64
}

func (h Hailstone) String() string {
	return fmt.Sprint(h.position) + " @ " + fmt.Sprint(h.velocity)
}

func Day_24_parse_input(use_test_file bool) (hailstones []Hailstone, min, max float64)  {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_24.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		txt := fileScanner.Text()
		line := strings.Split(txt, " @ ")
        pos_text := strings.Split(line[0], ", ")
        vel_text := strings.Split(line[1], ", ")

        pos := [3]float64{}
        vel := [3]float64{}
        for i:=0; i<3; i++ {
            val, err := strconv.Atoi(pos_text[i])
            if err != nil {
                fmt.Println(err)
            }
            pos[i] = float64(val)

            val, err = strconv.Atoi(vel_text[i])
            if err != nil {
                fmt.Println(err)
            }
            vel[i] = float64(val)
        }
        hailstone := Hailstone{position: models.CoordFloat64{}.From(pos),
                               velocity: models.CoordFloat64{}.From(vel)}
        hailstones = append(hailstones, hailstone)
	}
    if use_test_file {
        min, max = 7.0, 27.0
    } else {
        min, max = 200000000000000.0, 400000000000000.0
    }

	file.Close()
	return

}

func algebrafy(hailstone Hailstone) (m, b float64) {
    // Return back y = m x + b
    m = hailstone.velocity.Y / hailstone.velocity.X
    b = hailstone.position.Y - (m * hailstone.position.X)
    return
}

func get_number_of_intersections(hailstones []Hailstone, min, max float64) (num_intersections int) {
    for i := 0; i < len(hailstones); i++ {
        base := hailstones[i]
        m_1, b_1 := algebrafy(base)
        for j := i + 1; j < len(hailstones); j++ {
            comp := hailstones[j]
            m_2, b_2 := algebrafy(comp)

            x_f := (b_2 - b_1) / (m_1 - m_2)
            y_f := (m_1 * x_f) + b_1

            // Check that both Hailstones path's cross in the future
            // Eg: (↑)
            //   |    |    (→)
            // If the interception box is | |, the hailstones vector being the arrows,
            // their paths would intercept in the past and never in the future

            t_1 := (x_f - base.position.X) / base.velocity.X
            t_2 := (x_f - comp.position.X) / comp.velocity.X

            if (min <= x_f) && (x_f <= max) && (min <= y_f) && (y_f <= max) {
                if (t_1 < 0) || (t_2 < 0) {
                    continue
                }
                num_intersections++
            }
        }
    }
    return
}

func Day_24_Part_1() {
	// hailstones, min, max := Day_24_parse_input(true)
	hailstones, min, max := Day_24_parse_input(false)
    num_intersections := get_number_of_intersections(hailstones, min, max)
    fmt.Println("Number of intersections", num_intersections)
}

func Day_24_Part_2() {
    // We have 6 unknowns: The (x, y, z) position and (V_x, V_y, V_z) velocity of the Rock
    // Let the equation of the intersecting rock or Hailstone be given by
    //      P(t, V) = P_i + (V * t)
    // For the rock to intersect a Hailstone, their positions must equal at the same time, giving us
    //      P_r(t, V_r) = P_h(t, V_h)
    //      P_i_r + V_r * t = P_i_h + V_h * t
    // We rearrange this to get
    //      P_i_r - P_i_h = -t(V_r - V_h)
    // On the left, we have a vector. On the right, we have a vector multiplied by a scalar (t). This means the following must be true:
    //      [P_i_r - P_i_h] x [V_r - V_h] = 0  (Cross Product)
    //      [P_i_r x V_r] - [P_i_r x V_h] - [P_i_h x V_r] + [P_i_h x V_h] = 0
    //      - [P_i_r x V_h] + [P_i_h x V_r] + [P_i_h x V_h] = 0  (term [P_i_r x V_r] is common to every pair so we can drop that term)
    //      [V_h x P_i_r] - [P_i_h x V_r] + [P_i_h x V_h] = 0  (Reformatting using A x B = -B x A, to get r terms on rhs)
    // Here we note that the cross product can be calculated using the Skew Assymetric Operator
    //      A_t = <a_1,a_2, a_3>
    //      A_x = | 0    -a_3 a_2  |
    //            | a_3  0    -a_1 |
    //            | -a_2 a_1  0    |
    // With this we can get three equations to sovle for three unknowns. So using two pairs of Hailstones we can solve for all unknowns
    // This also means we avoid having to work with t and thus only have one fewer unknown. We are left with linear system of equations
    // EX let's use pairs [H_0, H_1] and [H_1, H_2]
    // That'll give us
    //      [P_i_r - P_i_h_0] x [V_r - V_h_0] = [P_i_r - P_i_h_1] x [V_r - V_h_1]
    //      [V_h_0 x P_i_r] - [P_i_h_0 x V_r] + [P_i_h_0 x V_h_0] = [V_h_1 x P_i_r] - [P_i_h_1 x V_r] + [P_i_h_1 x V_h_1]  (get Rock variables on the right)
    //      [V_h_0 - V_h_1] x P_i_r + [P_i_h_1 - P_i_h_0] x V_r = [P_i_h_1 x V_h_1] - [P_i_h_0 x V_h_0]
    // All the constants are on the right hand side and the variables are on the left, allowing us to make a vector of the unknowns
    //      | 0 X X 0 X X |   | P_x |   | C_1 |
    //      | X 0 X X 0 X |   | P_y |   | C_1 |
    //      | X X 0 X X 0 | x | P_z | = | C_1 |
    //      | 0 X X 0 X X |   | V_x |   | C_2 |
    //      | X 0 X X 0 X |   | V_y |   | C_2 |
    //      | X X 0 X X 0 |   | V_z |   | C_2 |
    //      M x V = C
    // The matrix M contains all the coefficients, V are the unknowns, and C are the Hailstone constants
    // From there we can do gaussian elimination to find the value of the unknowns

    // Begin work
	// hailstones, _, _ := Day_24_parse_input(true)
	hailstones, _, _ := Day_24_parse_input(false)
    matrix := generate_values_matrix(hailstones)
    fmt.Println(matrix)
    m := matrix.ToArray()
    solution := GaussianElimination(m, 6)
    pos := solution[:3]
    vel := solution[3:]
    fmt.Println(pos)
    fmt.Println(vel)

    total := 0.0
    for _, v := range(pos) {
        total = total + v
    }
    fmt.Printf("%f\n", total)
}

func generate_values_matrix(hailstones []Hailstone) (matrix models.Matrix[float64]) {
    consts := []float64{}
    h_0 := hailstones[0]
    h_1 := hailstones[1]
    h_2 := hailstones[10]

    pairs := [][]Hailstone{
        {h_0, h_1},
        {h_1, h_2},
    }

    for _, pair := range(pairs) {
        h_0 := pair[0]
        h_1 := pair[1]
        temp_matrix := models.Matrix[float64]{}

        term_1 := h_0.velocity.Sub(h_1.velocity)
        term_1_skew_matrix := term_1.SkewMatrixOperator()
        term_1_matrix := SkewMatrixToMatrix(term_1_skew_matrix)
        temp_matrix.AppendVertical(term_1_matrix)

        term_2 := h_1.position.Sub(h_0.position)
        term_2_skew_matrix := term_2.SkewMatrixOperator()
        term_2_matrix := SkewMatrixToMatrix(term_2_skew_matrix)
        temp_matrix.AppendHorizontal(term_2_matrix)

        const_1 := h_1.position.CrossProduct(h_1.velocity)
        const_2 := h_0.position.CrossProduct(h_0.velocity)
        constants := const_1.Sub(const_2)
        consts = append(consts, constants.ToArray()...)

        matrix.AppendVertical(temp_matrix)
    }
    matrix.AddColumn(matrix.Cols(), consts)
    return
}

func SkewMatrixToMatrix(m []models.CoordFloat64) (matrix models.Matrix[float64]) {
    for _, coord := range(m) {
        matrix.AddRow(matrix.Rows(), coord.ToArray())
    }

    return
}

// Swap: Swap two rows
func Swap(m [][]float64, i, j int) [][]float64 {
    for k := 0; k < len(m[0]); k++ {
        temp := m[i][k]
        m[i][k] = m[j][k]
        m[j][k] = temp
    }
    return m
}

func GaussianElimination(m [][]float64, n int) ([]float64)  {
    var singular_flag bool
    singular_flag, m = forward_elim(m, n)
    if singular_flag {
        fmt.Println("unsolvable")
    }
    sol := back_sub(m, n)
    return sol
}

func forward_elim(m [][]float64, n int) (bool, [][]float64) {
    for k := 0; k < n; k++ {
        i_max := k
        v_max := m[i_max][k]

        for i := k + 1; i < n; i++ {
            if math.Abs(m[i][k]) > v_max {
                v_max = m[i][k]
                i_max = i
            }
        }

        // if m[k][i_max] == 0.0 {
        //     return true, m
        // }

        if i_max != k {
            m = Swap(m, k, i_max)
        }

        for i := k + 1; i < n; i++ {
            ratio := m[i][k] / m[k][k]
            if m[i][k] == 0 || m[k][k] == 0 {
                ratio = 0
            }
            for j := k + 1; j < n + 1; j++ {
                m[i][j] = m[i][j] - (ratio * m[k][j])
            }
            m[i][k] = 0
        }
    }

    return false, m
}

func back_sub(m [][]float64, n int) []float64 {
    sol := []float64{}
    for i := 0; i < n; i++ {
        sol = append(sol, 0)
    }
    for i := n-1; i >= 0; i-- {
        sol[i] = m[i][n]

        for j := i + 1; j < n; j++ {
            sol[i] = sol[i] - (m[i][j] * sol[j])
        }
        sol[i] = sol[i] / m[i][i]
    }
    return sol
}
