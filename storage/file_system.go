package storage

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jiaqi-yin/go-australian-holidays/domain"
)

var _ Storage = &FileSystem{}

type FileSystem struct {
	Path string
}

func (fs *FileSystem) Save(holidays []domain.Holiday) {
	f, err := os.Create(fs.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, holiday := range holidays {
		_, err = f.WriteString(holiday.ToString() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (fs *FileSystem) Load() []domain.Holiday {
	fmt.Println("Loading holidays from file...")

	file, err := os.Open(fs.Path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	var holidays []domain.Holiday
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		holiday := scanner.Text()
		holidaySlice := strings.Split(holiday, ",")
		holidayObj := domain.Holiday{
			Date:  holidaySlice[0],
			Name:  holidaySlice[1],
			State: holidaySlice[2],
		}
		holidays = append(holidays, holidayObj)
	}

	return holidays
}

func NewFileSystem(path string) *FileSystem {
	return &FileSystem{
		Path: path,
	}
}
