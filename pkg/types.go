package pkg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

type Entry struct {
	Name   string
	Size   string
	Magnet string
}
type Entries []Entry

func (self *Entries) Display() {
	for _, val := range *self {
		fmt.Printf(`
		{
			name:   %v
			size:   %v
			magnet: %v
		}`, val.Name, val.Size, val.Magnet)
	}
}
func (self *Entries) CoppyMagnets() error {
	var magnets []string = make([]string, 0)
	if *self == nil {
		return errors.New("Empty list of entries, therefore, can't coppy any magnets.")
	}

	for _, val := range *self {
		magnets = append(magnets, val.Magnet)
	}
	err := clipboard.WriteAll(strings.Join(magnets, "\n"))
	if err != nil {
		return errors.New(fmt.Sprintf("Error copying to clipboard: %v", err))
	}
	return nil
}
