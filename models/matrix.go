package models

import (
	"AdventOfCode/utils"
	"fmt"
)

type Matrix[T any] struct {
	data      [][]T
	rows      int
	cols      int
	nil_value T
}

func (m *Matrix[T]) SetNilValue(value T) {
    m.nil_value = value
}

func (m *Matrix[T]) Rows() int {
	return m.rows
}

func (m *Matrix[T]) Cols() int {
	return m.cols
}

func (m *Matrix[T]) Data() [][]T {
    return m.data
}

func (m *Matrix[T]) AddEmptyRow(at_index, amount int) {
	if m.rows == 0 && m.cols == 0 {
		m.data = append(m.data, []T{})
		m.rows++
		return
	}

	new_row := []T{}
	for i := 0; i < m.cols; i++ {
		new_row = append(new_row, m.nil_value)
	}

	for i := 0; i < amount; i++ {
		if at_index == len(m.data) {
			m.data = append(m.data, utils.Clone(new_row))
		} else {
			m.data = append(m.data[:at_index+1], m.data[at_index:]...)
			m.data[at_index] = utils.Clone(new_row)
		}
	}
    m.rows = m.rows + amount
}

func (m *Matrix[T]) AddRow(at_index int, rows ...[]T) {
	if m.rows == 0 && m.cols == 0 {
        for _, row := range(rows) {
            m.data = append(m.data, row)
        }
		m.rows = m.rows + len(rows)
        m.cols = len(rows[0])
		return
	}

    for _, row := range(rows) {
		if at_index == len(m.data) {
			m.data = append(m.data, row)
		} else {
			m.data = append(m.data[:at_index+1], m.data[at_index:]...)
			m.data[at_index] = row 
            at_index++
		}
    }
    m.rows = m.rows + len(rows) 
}

func (m *Matrix[T]) AddEmptyColumn(at_index, amount int) {
	if m.rows == 0 && m.cols == 0 {
		m.data = append(m.data, []T{[1]T{}[0]})
		m.rows++
		m.cols++
		return
	}

	for row_i := range m.data {
		for i := 0; i < amount; i++ {
			if at_index == m.cols {
				m.data[row_i] = append(m.data[row_i], m.nil_value)
			} else {
				m.data[row_i] = append(m.data[row_i][:at_index+1], m.data[row_i][at_index:]...)
				m.data[row_i][at_index] = m.nil_value
			}
		}
	}
    m.cols = m.cols + amount
}

func (m *Matrix[T]) AddColumn(at_index int, cols ...[]T) {
	if m.rows == 0 && m.cols == 0 {
        for range(cols) {
            m.data = append(m.data, []T{})
        }
        for _, col := range(cols) {
            for row_i, cell := range(col) {
                m.data[row_i] = append(m.data[row_i], cell)
            }
        }
		m.rows = len(cols)
		m.cols = len(cols[0])
		return
	}


    for _, col := range(cols) {
        for row_i, cell := range(col) {
			if at_index == m.cols {
				m.data[row_i] = append(m.data[row_i], cell)
			} else {
				m.data[row_i] = append(m.data[row_i][:at_index+1], m.data[row_i][at_index:]...)
				m.data[row_i][at_index] = cell
			}
        }
    }
}

func (m *Matrix[T]) Get(row_i, col_i int) (v T) {
	if (row_i >= 0 && row_i < len(m.data)) && (col_i >= 0 && col_i < len(m.data[0])) {
		v = m.data[row_i][col_i]
	}
	return
}

func (m *Matrix[T]) GetRow(row_i int) (v []T) {
	if row_i >= 0 && row_i < len(m.data) {
		v = m.data[row_i]
	}
	return
}

func (m *Matrix[T]) GetColumn(col_i int) (v []T) {
	if col_i >= 0 && col_i < len(m.data[0]) {
        for row_i := 0; row_i < m.rows; row_i++ {
            v = append(v, m.Get(row_i, col_i))
        }
	}
	return
}

func (m *Matrix[T]) Set(row_i, col_i int, value T) (success bool) {
	if (row_i >= 0 && row_i < m.rows) && (col_i >= 0 && col_i < m.cols) {
		m.data[row_i][col_i] = value
		success = true
	}
	return
}

func (m Matrix[T]) String() string {
	k := ""
	for i, row := range m.data {
		k = k + fmt.Sprintf("%03d", i) + " - " + fmt.Sprintln(row)
	}
	return k
}

func (m *Matrix[T]) Clone() Matrix[T] {
    cloned := Matrix[T]{}
    for row_i, row := range(m.data) {
        cloned.AddRow(row_i, utils.Clone(row))
    }
    return cloned 
}

func (m *Matrix[T]) AppendVertical(o Matrix[T]) {
    // Both must have same row length
    for i := 0; i < o.Rows(); i++ {
        m.AddRow(m.Rows(), o.GetRow(i))
    }
}

func (m *Matrix[T]) AppendHorizontal(o Matrix[T]) {
    // Both must have same col length
    for i := 0; i < o.Rows(); i++ {
        m.data[i] = append(m.data[i], o.GetRow(i)...)
    }
}

func (m *Matrix[T]) ToArray() ([][]T) {
    return m.data
}


func (m *Matrix[T]) IsValidCell(loc Coord) bool {
    height, width := m.Rows(), m.Cols()
	x_pos, y_pos := loc.X, loc.Y
	if x_pos < 0 || x_pos >= height {
		return false
	}
	if y_pos < 0 || y_pos >= width {
		return false
	}
	return true
}
