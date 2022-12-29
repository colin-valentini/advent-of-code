package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory(t *testing.T) {
	root := newTestDirectory(t)
	assert.Equal(t, 48_381_165, root.size())
}

func newTestDirectory(t *testing.T) *directory {
	// - / (dir)
	//   - a (dir)
	//     - e (dir)
	//       - i (file, size=584)
	//     - f (file, size=29116)
	//     - g (file, size=2557)
	//     - h.lst (file, size=62596)
	//   - b.txt (file, size=14848514)
	//   - c.dat (file, size=8504156)
	//   - d (dir)
	//     - j (file, size=4060174)
	//     - d.log (file, size=8033020)
	//     - d.ext (file, size=5626152)
	//     - k (file, size=7214296)
	t.Helper()

	root := newDirectory(nil, "/")
	root.addFile(newFile("b.txt", 14_848_514))
	root.addFile(newFile("c.dat", 8_504_156))

	a := newDirectory(root, "a")
	a.addFile(newFile("f", 29_116))
	a.addFile(newFile("g", 2_557))
	a.addFile(newFile("h.lst", 62_596))

	e := newDirectory(a, "e")
	e.addFile(newFile("i", 584))

	d := newDirectory(root, "d")
	d.addFile(newFile("j", 4_060_174))
	d.addFile(newFile("d.log", 8_033_020))
	d.addFile(newFile("d.ext", 5_626_152))
	d.addFile(newFile("k", 7_214_296))

	return root
}
