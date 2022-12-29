package day7

import (
	"strconv"
	"strings"
)

const (
	cdPrefix  = "$ cd"
	lsPrefix  = "$ ls"
	dirPrefix = "dir"
)

type parser struct {
	stdOut string
	cwd    *directory
	it     *iterator
}

func newParser(stdOut string) *parser {
	lines := strings.Split(stdOut, "\n")
	return &parser{stdOut: stdOut, it: newIterator(lines)}
}

func (p *parser) parse() *directory {
	for p.it.next() {
		line := p.it.value()
		if p.isChangeDir(line) {
			p.parseChangeDir(line)
		} else if p.isListDir(line) {
			// Nothing to do, just know that the next lines will be
			// files. Assumes that the problem inputs are logical.
			continue
		} else if p.isDirEntry(line) {
			p.parseDirEntry(line)
		} else {
			// Must be inside ls output for the current working directory.
			p.parseFileEntry(line)
		}
	}
	return p.cwd.rootDir()
}

func (p *parser) isChangeDir(line string) bool {
	return strings.HasPrefix(line, cdPrefix)
}

func (p *parser) isListDir(line string) bool {
	return strings.HasPrefix(line, lsPrefix)
}

func (p *parser) parseChangeDir(line string) {
	arg := strings.TrimPrefix(line, cdPrefix+" ")
	if arg == "/" {
		if p.cwd == nil {
			p.cwd = newDirectory(nil, "/")
		} else {
			p.cwd = p.cwd.rootDir()
		}
	} else if arg == ".." {
		p.cwd = p.cwd.parentDir()
	} else {
		p.cwd = p.cwd.changeDir(arg)
	}
}

func (p *parser) isDirEntry(line string) bool {
	return strings.HasPrefix(line, dirPrefix)
}

func (p *parser) parseDirEntry(line string) {
	name := strings.TrimPrefix(line, dirPrefix+" ")
	_ = newDirectory(p.cwd, name)
}

func (p *parser) parseFileEntry(line string) {
	fStr := strings.Split(line, " ")
	if len(fStr) != 2 {
		panic("found unrecognizable file line output")
	}
	name := fStr[1]
	size, err := strconv.Atoi(fStr[0])
	if err != nil {
		panic("found unrecognizable file line size")
	}
	p.cwd.addFile(newFile(name, size))
}

type iterator struct {
	idx   int
	lines []string
}

func newIterator(lines []string) *iterator {
	return &iterator{idx: -1, lines: lines}
}

func (it *iterator) next() bool {
	if it.idx+1 >= len(it.lines) {
		return false
	}
	it.idx++
	return true
}

func (it *iterator) value() string {
	return it.lines[it.idx]
}
