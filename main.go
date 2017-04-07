package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	keep = flag.String("keep", "msg", "The message types to keep. Options: join, part, quit, rename, msg")
	path = flag.String("path", ".", "The filepath or directory of logfile(s) to parse.")

	// Holder for what message types to keep
	keeps []string
)

func main() {
	flag.Parse()

	// Verify `path`
	if path == nil || *path == "" {
		fail("Please specify a -path to read and parse logs from")
	}
	if filepath.Dir(*path) == string(filepath.Separator) {
		fail("Reading from the root dir is not supported.")
	}

	// Expand `keep`
	if keep == nil || *keep == "" {
		fail("Please specify message types for -keep")
	}
	keeps = strings.Split(*keep, ",")

	// Walk along `path`
	a, err := filepath.Abs(*path)
	if err != nil {
		fail(err.Error())
	}
	paths, err := filepath.Glob(a)
	if err != nil {
		fail(err.Error())
	}

	// walk each match from the glob
	for i := range paths {
		err = filepath.Walk(paths[i], walk)
		if err != nil {
			fail(err.Error())
		}
	}
}

func walk(path string, info os.FileInfo, err error) error {
	// Ignore SkipDir and directories
	if (err != nil && err != filepath.SkipDir) || info.IsDir() {
		return nil
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	body := string(b)

	kept := parse(body)
	if len(kept) > 0 {
		fmt.Printf("From: %s\n", path)
		for i := range kept {
			fmt.Println(kept[i])
		}
		fmt.Println()
	}

	return nil
}

func parse(body string) []string {
	keepfns := make([]func(string) bool, 0)
	for i := range keeps {
		switch {
		case keeps[i] == "join":
			keepfns = append(keepfns, isJoin)
		case keeps[i] == "part":
			keepfns = append(keepfns, isPart)
		case keeps[i] == "quit":
			keepfns = append(keepfns, isQuit)
		case keeps[i] == "rename":
			keepfns = append(keepfns, isRename)
		}
	}
	keepfns = append(keepfns, isMessage)

	kept := make([]string, 0)
	lines := strings.Split(body, "\n")
	for i := range lines {
		for k := range keepfns {
			if keepfns[k](lines[i]) {
				kept = append(kept, lines[i])
				break
			}
		}
	}

	return kept
}

// regex matching
var (
	tsr = `\[\d\d:\d\d:\d\d\][\s]*`
	msgr = `.+`

	joinr = regexp.MustCompile(tsr + `\*{3} Joins: ` + msgr)
	partr = regexp.MustCompile(tsr + `\*{3} Parts: ` + msgr)
	quitr = regexp.MustCompile(tsr + `\*{3} Quits: ` + msgr)
	renamer = regexp.MustCompile(tsr + `\*{3} \w+ is now known as \w+`)
	messager = regexp.MustCompile(tsr + `<[\w_]+>` + msgr)
)

// Message examples, more in parse_test.go
// [02:33:31] *** Joins: adam (adam@Snoonet-fdl.i3c.1b1g5k.IP)
// [14:36:32] <adamdecaf> They come with a bunch more problems though
// [03:14:57] *** Parts: bo4tdude[penis] (Bo4t@user/bo4tdude) (Leaving)
// [02:33:19] *** Quits: bo4tdude (Bo4t@user/bo4tdude) (Quit: Leaving)
// [02:34:32] *** adam is now known as Snoo60230

func isJoin(line string) bool {
	return joinr.MatchString(line)
}
func isPart(line string) bool {
	return partr.MatchString(line)
}
func isQuit(line string) bool {
	return quitr.MatchString(line)
}
func isRename(line string) bool {
	return renamer.MatchString(line)
}
func isMessage(line string) bool {
	return messager.MatchString(line)
}

func fail(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
