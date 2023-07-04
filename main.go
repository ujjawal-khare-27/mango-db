package main

import (
	"fmt"
	"os"
)

type Document struct {
	ID   string                 `json:"id""`
	Data map[string]interface{} `json:"data""`
}

type DocumentStore struct {
	documents map[string]Document
}

func NewDocumentStore() *DocumentStore {
	return &DocumentStore{
		documents: make(map[string]Document),
	}
}

func (ds *DocumentStore) Create(document Document) error {
	if _, ok := ds.documents[document.ID]; ok {
		return fmt.Errorf("ID %s already exists", document.ID)
	}

	ds.documents[document.ID] = document
	return nil
}

func (ds *DocumentStore) ReadDocument(id string) (*Document, error) {
	document, ok := ds.documents[id]

	if !ok {
		return nil, fmt.Errorf("document with ID '%s' does not exist", id)
	}
	return &document, nil
}

func (ds *DocumentStore) Update(document Document) error {
	if _, ok := ds.documents[document.ID; !ok {
		return fmt.Errorf("document with ID '%s' does not exist", document.ID)
	}

	ds.documents[document.ID] = document
	return nil
}

func main() {
	// initialize db
	dataAccessLayer, _ := newDataAccessLayer("db.db", os.Getpagesize())

	// create a new page
	p := dataAccessLayer.allocateEmptyPage()
	p.num = dataAccessLayer.getNextPage()
	copy(p.data[:], "data")

	// commit it
	_ = dataAccessLayer.writePage(p)
}
