package utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func ReadFile(fPath string, line func(line string)) {
	f, err := os.Open(fPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadAllFile(fPath string) string {
	buf, err := os.ReadFile(fPath)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsAplpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '-'
}

func IsAplphaAndQuote(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '-' || c == '\''
}

func IsSpace(c rune) bool {
	return c == ' ' || c == '\t'
}

func ReadInt(r *bufio.Reader) int {
	var (
		buf bytes.Buffer
	)
	r.UnreadRune()
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		if !IsDigit(c) {
			r.UnreadRune()
			break
		}
		buf.WriteRune(c)
	}
	return Atoi(buf.String())
}

func ReadStringAndQuote(r *bufio.Reader) string {
	var (
		buf bytes.Buffer
	)
	r.UnreadRune()
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		if !IsAplphaAndQuote(c) {
			r.UnreadRune()
			break
		}
		buf.WriteRune(c)
	}
	return buf.String()
}

func ReadString(r *bufio.Reader) string {
	var (
		buf bytes.Buffer
	)
	r.UnreadRune()
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		if !IsAplpha(c) {
			r.UnreadRune()
			break
		}
		buf.WriteRune(c)
	}
	return buf.String()
}

func Atoi(a string) int {
	r, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return r
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func PowInt(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Great common divisor
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Least common multiplier
func MustLcm(integers []int) int {
	if len(integers) < 2 {
		panic("not enough number to calculate lcm")
	}
	a := integers[0]
	b := integers[1]
	result := a * b / Gcd(a, b)

	for i := 2; i < len(integers); i++ {
		result = MustLcm([]int{result, integers[i]})
	}

	return result
}
