package day_09

import (
	"fmt"
	"strconv"
)

type memory []int
type filesystem struct {
	memory        memory
	totalFilesize int
	files         int
}

func parseInput(input string) (*filesystem, error) {
	file := true
	fileID := 0
	var filesystem filesystem

	for _, c := range input {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		spaceLayout := make([]int, v)
		for i := range v {
			if file {
				spaceLayout[i] = fileID
			} else {
				spaceLayout[i] = -1
			}
		}
		if file {
			filesystem.files++
			filesystem.totalFilesize += v
		}
		if file {
			fileID++
		}

		file = !file
		filesystem.memory = append(filesystem.memory, spaceLayout...)
	}
	return &filesystem, nil
}

func (f memory) print() {
	for _, v := range f {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}

	fmt.Println()
}

func (f memory) getLastFileBlock() (index, fileID int) {
	for i := len(f) - 1; i >= 0; i-- {
		if f[i] != -1 {
			return i, f[i]
		}
	}

	return -1, -1
}

func (f filesystem) moveLastBlockToFirstAvailableSpot() (anySpaceLeft bool) {
	spotIdx := 0
	for i, v := range f.memory {
		if v == -1 {
			spotIdx = i
			break
		}
	}

	lastBlockIdx, fileID := f.memory.getLastFileBlock()
	if lastBlockIdx == f.totalFilesize-1 {
		return false
	}

	f.memory[spotIdx] = fileID
	f.memory[lastBlockIdx] = -1
	return true
}

func (f filesystem) getNextAvailableSpot(searchFrom int) (startIndex, size int) {
	for i := searchFrom; i < len(f.memory); i++ {
		if f.memory[i] == -1 {
			startIndex = i
			for j := i; j < len(f.memory); j++ {
				if f.memory[j] != -1 {
					size = j - i
					return startIndex, size
				}
			}
		}
	}

	return -1, -1
}

func (f filesystem) getFileByID(fileID int) (startIndex, size int) {
	for i, v := range f.memory {
		if v == fileID {
			startIndex = i
			for j := i; j < len(f.memory); j++ {
				if j == len(f.memory)-1 {
					size = j - i + 1
					return startIndex, size
				}

				if f.memory[j] != fileID {
					size = j - i
					return startIndex, size
				}
			}
		}
	}

	return -1, -1
}

func (f filesystem) moveFileToFirstAvailableSpot(fileID int) {
	fileStartIndex, fileSize := f.getFileByID(fileID)

	var nextSpotIdx, nextSpotSize int
	for i := 0; i < fileStartIndex; i++ {
		nextSpotIdx, nextSpotSize = f.getNextAvailableSpot(i)
		if nextSpotIdx == -1 {
			return
		}
		if nextSpotSize >= fileSize {
			break
		} else {
			i = nextSpotIdx + nextSpotSize
		}
	}

	for i := nextSpotIdx; i < nextSpotIdx+fileSize; i++ {
		f.memory[i] = fileID
	}

	for i := fileStartIndex; i < fileStartIndex+fileSize; i++ {
		f.memory[i] = -1
	}
}

func (f filesystem) calculateChecksum() int {
	var checksum int

	for i, v := range f.memory {
		if v != -1 {
			checksum += i * v
		}
	}

	return checksum
}

func PartOne(input string) (int, error) {
	fs, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for fs.moveLastBlockToFirstAvailableSpot() {

	}

	return fs.calculateChecksum(), nil
}

func PartTwo(input string) (int, error) {
	fs, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for i := fs.files; i > 0; i-- {
		fs.moveFileToFirstAvailableSpot(i - 1)
	}
	return fs.calculateChecksum(), nil
}
