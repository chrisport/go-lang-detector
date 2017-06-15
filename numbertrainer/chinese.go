package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const gap string = "    "

var chineseNrs = map[int]string{
	1:  "一",
	2:  "二",
	3:  "三",
	4:  "四",
	5:  "五",
	6:  "六",
	7:  "七",
	8:  "八",
	9:  "九",
	10: "十",
}

var pinyinNrs = map[int]string{
	1:  "yī",
	2:  "èr",
	3:  "sān",
	4:  "sì",
	5:  "wǔ",
	6:  "liù",
	7:  "qī",
	8:  "bā",
	9:  "jiǔ",
	10: "shí",
}

type RandomNumber struct {
	pinyin  string
	chinese string
	intVal  string
}

func GenerateRandomNumber() RandomNumber {
	var pinyin, chinese string

	nr := (rand.Int() % 100) + 1
	if nr >= 20 {
		pinyin = pinyinNrs[nr/10] + " " + pinyinNrs[10] + " "
		chinese = chineseNrs[nr/10] + chineseNrs[10]
	} else if nr >= 10 {
		pinyin = pinyinNrs[10] + " "
		chinese = chineseNrs[10]
	}
	pinyin += pinyinNrs[nr%10]
	chinese += chineseNrs[nr%10]

	return RandomNumber{pinyin: pinyin, chinese: chinese, intVal: strconv.Itoa(nr)}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Chose mode:")
choseMode:
	fmt.Println("  1 ) Active\n  2 ) Passive Chinese letters\n  3 ) Passive Pinyin")
	r, _ := reader.ReadString('\n')
	var mode func(RandomNumber) string
	resultChecker := func(nr RandomNumber, input string) {
		if input == nr.intVal {
			fmt.Println("CORRECT")
		} else {
			fmt.Println("WRONG")
		}
	}
	fmt.Println()
	switch strings.TrimRight(r, "\n") {
	case "1":
		mode = active
		resultChecker = func(RandomNumber, string) {}
		fmt.Println("Please answer for yourself and press <Enter> to see the correct response.")

	case "2":
		mode = passiveChinese
		fmt.Println("Please answer in numbers and confirm with Enter, e.g. 三 --> 3 <Enter>")
	case "3":
		mode = passivePinYin
		fmt.Println("Please answer in numbers and confirm with Enter, e.g. sān --> 3 <Enter>")
	default:
		fmt.Print("Type 1, 2 or 3 and press enter to chose mode. ")
		goto choseMode
	}
	time.Sleep(3 * time.Second)
	playRound(reader, mode, resultChecker)
}

var active = func(number RandomNumber) string {
	return number.intVal
}

var passiveChinese = func(number RandomNumber) string {
	return number.chinese
}

var passivePinYin = func(number RandomNumber) string {
	return number.pinyin
}

func playRound(reader *bufio.Reader, mode func(RandomNumber) string, resultChecker func(RandomNumber, string)) {
	fmt.Println("______________________________")

	for roundCounter := 1; ; roundCounter++ {
		fmt.Println("ROUND " + strconv.Itoa(roundCounter))

		nr := GenerateRandomNumber()
		visiblePart := mode(nr)
		fmt.Println(visiblePart)
		a, _ := reader.ReadString('\n')
		ac := strings.TrimRight(string(a), "\n")
		fmt.Println(pad(string(nr.intVal), 2) + gap + pad(nr.pinyin, 11) + gap + nr.chinese)
		resultChecker(nr, ac)
		fmt.Println("______________________________")
		time.Sleep(2 * time.Second)
	}
}

func pad(s string, length int) string {
	for i := utf8.RuneCount([]byte(s)); i < length; i++ {
		s += " "
	}
	return s
}
