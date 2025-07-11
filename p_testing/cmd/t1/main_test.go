package main

import "testing"

//простой тестик
func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
	}
}

//простой табличный тестик (хз зачем, но пускай будет)
func TestSumTable(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{10, 10, 20},
		{0, 0, 0},
		{-5, 5, 0},
	}

	//тесты выполн послед-о, в выводе все ошибки
	//будут отображаться как часть одного теста
	for _, tt := range tests {
		result := Sum(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Sum(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

//Если нужна более детализированная отчетность (чтобы видеть, какой именно подтест упал),
//используем t.Run(). Можно так же сделать с t.Parallel(),
// но с захватом переменной "tt := tt" чтобы не было race condition
// пример:
// for _, tt := range tests {
//     tt := tt
//     t.Run(tt.name, func(t *testing.T) {
//         t.Parallel()
//          ...
//     })
// }
func TestSumWithSubtests(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "zero and negative",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{
			name:     "all zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Sum(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
