package services_test

import (
	"fmt"
	"testing"

	"github.com/Aritiaya50217/GolangByCodeBangkok/services"
	assert "github.com/stretchr/testify/assert"
)

// Unit Testing: ใช้สำหรับตรวจสอบความถูกต้องของฟังก์ชันUnit Testing ใช้ ตรวจสอบว่าฟังก์ชันหรือโมดูลทำงานตามที่คาดหวัง
func TestCheckGrade(t *testing.T) {
	type testCase struct {
		name     string
		score    int
		expected string
	}
	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 0, expected: "F"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			assert.Equal(t, c.expected, grade)
		})
	}
}

// Benchmarking: ใช้สำหรับวัดประสิทธิภาพของฟังก์ชัน
func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

// Example Test: ใช้เพื่อให้ตัวอย่างการใช้งานฟังก์ชันและแสดงผลลัพธ์ที่คาดหวัง
func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
}
