package main

import (
	"io/ioutil"
)

// model struct holds a list of filenames and the index of the selected file
type ServerConfigs struct {
	filenames     []string // List of filenames
	selectedIndex int      // Index of the currently selected filename
}

// currentSelected returns the currently selected filename
func (m ServerConfigs) currentSelected() string {
	if m.selectedIndex >= 0 && m.selectedIndex < len(m.filenames) {
		return m.filenames[m.selectedIndex]
	}
	return ""
}

// loadFilenames reads the filenames from a given directory and populates the model
func (m *ServerConfigs) loadFilenames(directory string) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	m.filenames = []string{}
	for _, file := range files {
		if !file.IsDir() {
			m.filenames = append(m.filenames, file.Name())
		}
	}

	if len(m.filenames) > 0 {
		m.selectedIndex = 0 // Select the first file by default
	}

	return nil
}
