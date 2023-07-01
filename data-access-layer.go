/*
Data Access Layer is responsible for reading & writing data to the file system.
*/
package main

import "os"

// function to open a file
func OpenFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// function to close a file
func CloseFile(f *os.File) error {
	err := f.Close()
	if err != nil {
		return err
	}
	return nil
}
