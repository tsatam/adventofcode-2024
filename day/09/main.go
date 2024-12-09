package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/tsatam/adventofcode-2024/common/fp"
)

var (
	//go:embed input
	input string
)

type Block uint16

type File struct {
	b Block
	s int
}

const (
	empty Block = Block(math.MaxUint16)
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	files := parseDiskMap(input)
	blocks := fileToBlocks(files)
	compacted := compactByBlock(blocks)
	return checksum(compacted)
}

func handlePart2(input string) int {
	files := parseDiskMap(input)
	compacted := compactByFile(files)
	compactedBlocks := fileToBlocks(compacted)
	return checksum(compactedBlocks)
}

func parseDiskMap(input string) []File {
	input = strings.TrimSpace(input)
	result := []File{}
	for i, sizeRaw := range []rune(input) {
		s := parseDigit(sizeRaw)
		b := empty
		if i%2 == 0 {
			b = Block(i / 2)
		}
		result = append(result, File{b: b, s: s})
	}

	return fp.Filter(result, func(bs File) bool { return bs.s > 0 })
}

func fileToBlocks(blockSizes []File) []Block {
	result := []Block{}

	for _, bs := range blockSizes {
		result = append(result, bs.toBlocks()...)
	}

	return result
}

func compactByBlock(blocks []Block) []Block {
	b := make([]Block, len(blocks))
	copy(b, blocks)

	freeSpaceIdx := 0

	for i := len(b) - 1; i > freeSpaceIdx+1; i-- {
		if b[i] != empty {
			for b[freeSpaceIdx] != empty {
				freeSpaceIdx++
			}
			b[freeSpaceIdx] = b[i]
			b[i] = empty
		}
	}

	firstFreeSpace := slices.Index(b, empty)
	if firstFreeSpace != -1 {
		return b[:firstFreeSpace]
	} else {
		return b
	}
}

func compactByFile(files []File) []File {
	fs := make([]File, len(files))
	copy(fs, files)
	highestFileId := fs[len(files)-1].b // assume last entry is a file and is not empty space

	for fileId := highestFileId; fileId > 0; fileId-- {
		fileIdx := slices.IndexFunc(fs, func(f File) bool {
			return f.b == fileId
		})
		file := fs[fileIdx]

		validEmptySpaceIdx := slices.IndexFunc(fs[:fileIdx], func(f File) bool {
			return f.b == empty && f.s >= file.s
		})

		if validEmptySpaceIdx == -1 {
			continue
		}

		fs = slices.Replace(fs, fileIdx, fileIdx+1, File{b: empty, s: file.s})
		validEmptySpace := fs[validEmptySpaceIdx]

		if file.s == validEmptySpace.s {
			fs = slices.Replace(fs, validEmptySpaceIdx, validEmptySpaceIdx+1, file)
		} else {
			remaining := validEmptySpace.s - file.s
			fs = slices.Replace(fs, validEmptySpaceIdx, validEmptySpaceIdx+1,
				file,
				File{b: empty, s: remaining},
			)
		}
	}

	return fp.Reduce(fs, []File{}, func(curr []File, next File) []File {
		if len(curr) > 0 && next.b == curr[len(curr)-1].b {
			curr[len(curr)-1].s += next.s
			return curr
		} else {
			return append(curr, next)
		}
	})
}

func checksum(blocks []Block) int {
	checksum := 0

	for i, b := range blocks {
		if b != empty {
			checksum += i * int(b)
		}
	}

	return checksum
}

func parseDigit(r rune) int {
	return int(r - '0')
}

func (bs File) toBlocks() []Block {
	return slices.Repeat([]Block{bs.b}, bs.s)
}

func repeat(b Block, length int) []Block {
	return slices.Repeat([]Block{b}, length)
}
