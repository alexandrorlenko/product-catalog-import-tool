package fileHandler

import (
	"encoding/csv"
	"os"
)

type CSVHandler struct {
	Reader   *csv.Reader
	Comma    rune
	Bunch    int
	LineChan chan []string
}

func NewCSVHandler() FileHandlerInterface {
	return &CSVHandler{
		Comma:    ',',
		LineChan: make(chan []string),
	}
}

func (c *CSVHandler) InitReader(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = c.Comma
	c.Reader = reader
}

func (c *CSVHandler) ReadLine() ([]string, error) {
	record, e := c.Reader.Read()
	if e != nil {
		return nil, e
	}
	return record, e
}
