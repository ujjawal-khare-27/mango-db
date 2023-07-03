/*
Data Access Layer is responsible for reading & writing data to the file system.
*/
package main

import "os"

type pgnum uint64

type page struct {
	num  pgnum
	data []byte
}

type dataAccessLayer struct {
	file     *os.File
	pageSize int
	freelist *freelist
}

// function to create a new data access layer
func newDataAccessLayer(path string, pageSize int) (*dataAccessLayer, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return &dataAccessLayer{file: file, pageSize: pageSize, freelist: newFreelist()}, nil
}

// function to close a file
func (d *dataAccessLayer) close() error {
	err := d.file.Close()
	if err != nil {
		return err
	}
	return nil
}

// allocate a empty page for reading & writing data
func (d *dataAccessLayer) allocateEmptyPage() *page {
	return &page{
		data: make([]byte, d.pageSize),
	}
}

// read a page from the file
func (d *dataAccessLayer) readPage(pageNum pgnum) (*page, error) {
	p := d.allocateEmptyPage()

	// correct offset calculation is performed
	// using the page number and page size
	offset := int(pageNum) * d.pageSize

	// Then we read the data at the correct offset
	_, err := d.file.ReadAt(p.data, int64(offset))
	if err != nil {
		return nil, err
	}
	return p, err
}

// write a page to the file
func (d *dataAccessLayer) writePage(p *page) error {
	offset := int(p.num) * d.pageSize

	_, err := d.file.WriteAt(p.data, int64(offset))
	if err != nil {
		return err
	}
	return nil
}
