package filerw

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
)

func Create(notesDir, category, name string) (string, error) {
	if _, err := os.Stat(filepath.Join(notesDir, category)); err != nil {
		err = os.Mkdir(filepath.Join(notesDir, category), 0755)
		if err != nil {
			return "", err
		}
	}
	if _, err := os.Stat(filepath.Join(notesDir, category, name)); err == nil {
		return "", errors.New("File already exists")
	}
	filePath := filepath.Join(notesDir, category, name)
	_, err := os.Create(filePath)
	return filePath, err
}

func RunInCli(filePath string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	return err
}

func GetLastNotes(notesPath string) ([]noteInfo, error) {
	notes, err := walkDir(notesPath)
	if err != nil {
		return nil, err
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Time.After(notes[j].Time)
	})

	if len(notes) > 10 {
		notes = notes[:10]
	}

	notes = groupNotesByCategory(notes)
	return notes, nil
}

func walkDir(notesPath string) ([]noteInfo, error) {
	entries, err := os.ReadDir(notesPath)
	if err != nil {
		return nil, err
	}

	notes := make([]noteInfo, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			files, err := os.ReadDir(filepath.Join(notesPath, entry.Name()))
			if err != nil {
				return nil, err
			}
			for _, e := range files {
				info, err := e.Info()
				if err != nil {
					return nil, err
				}
				notes = append(notes, noteInfo{
					Name:     e.Name(),
					Category: entry.Name(),
					Time:     info.ModTime(),
					Path:     filepath.Join(notesPath, entry.Name(), e.Name()),
				})
			}
		}
	}

	return notes, nil
}

func groupNotesByCategory(notes []noteInfo) []noteInfo {
	newNotes := make([]noteInfo, len(notes))
	copy(newNotes, notes)
	for i := range newNotes {
		for j := i + 1; j < len(newNotes); j++ {
			if newNotes[i].Category == newNotes[j].Category {
				newNotes[i], newNotes[j] = newNotes[j], newNotes[i]
			}
		}
	}
	return newNotes
}
