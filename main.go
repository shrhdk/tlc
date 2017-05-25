package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const version = "0.0.1"

func main() {
	formatOpt := flag.String("f", "yyyy-mm-dd HH:MM:SS.SSS", "date time layout. default is yyyy-mm-dd HH:MM:SS.SSS")
	periodOpt := flag.String("p", "1m", "count period. example: 1ms, 2s, 3m, 4h")
	helpOpt := flag.Bool("h", false, "show help")
	versionOpt := flag.Bool("v", false, "show version")
	flag.Parse()

	if *helpOpt {
		fmt.Println("tlc [options]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *versionOpt {
		fmt.Println(version)
		os.Exit(0)
	}

	period, err := time.ParseDuration(*periodOpt)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unexpected duration format. see https://golang.org/pkg/time/#ParseDuration")
		os.Exit(1)
	}

	layout, re := getLayoutAndPattern(*formatOpt)

	var lastT time.Time
	var ct uint64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		tstr := re.FindString(line)
		if tstr == "" {
			continue
		}

		t, err := time.Parse(layout, tstr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse '%s' as date time at line '%s' : %v", tstr, line, err)
		}

		ct++

		t = t.Truncate(period)
		if lastT.IsZero() {
			lastT = t
		}

		if t != lastT {
			fmt.Printf("%v\t%d\n", lastT, ct)
			lastT = t
			ct = 0
		}
	}

	fmt.Printf("%v\t%d\n", lastT, ct)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, " : %v\n", err)
		os.Exit(1)
	}
}

func getLayoutAndPattern(format string) (string, *regexp.Regexp) {
	l := format
	l = strings.Replace(l, "yyyy", "2006", -1)
	l = strings.Replace(l, "yy", "06", -1)
	l = strings.Replace(l, "mmmm", "January", -1)
	l = strings.Replace(l, "mmm", "Jan", -1)
	l = strings.Replace(l, "mm", "01", -1)
	l = strings.Replace(l, "m", "1", -1)
	l = strings.Replace(l, "dd", "02", -1)
	l = strings.Replace(l, "d", "2", -1)
	l = strings.Replace(l, "HH", "15", -1)
	l = strings.Replace(l, "MM", "04", -1)
	l = strings.Replace(l, "M", "4", -1)
	l = strings.Replace(l, "SS.SSS", "05.000", -1)
	l = strings.Replace(l, "S.SSS", "5.000", -1)
	l = strings.Replace(l, "SS", "05", -1)
	l = strings.Replace(l, "S", "5", -1)

	restr := l
	restr = regexp.MustCompile(`[A-Za-z]`).ReplaceAllString(restr, `[A-Za-z]`)
	restr = regexp.MustCompile(`\d`).ReplaceAllString(restr, `\d`)
	restr = strings.Replace(restr, `.`, `\.`, -1)

	return l, regexp.MustCompile(restr)
}
