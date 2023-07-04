package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Document struct {
	ID   string                 `json:"id""`
	Data map[string]interface{} `json:"data""`
}

type DocumentStore struct {
	documents map[string]*Document
}

func NewDocumentStore() *DocumentStore {
	return &DocumentStore{
		documents: make(map[string]*Document),
	}
}

func (ds *DocumentStore) Create(document Document) error {
	if _, ok := ds.documents[document.ID]; ok {
		return fmt.Errorf("ID %s already exists", document.ID)
	}

	ds.documents[document.ID] = &document
	return nil
}

func (ds *DocumentStore) ReadDocument(id string) (*Document, error) {
	document, ok := ds.documents[id]

	if !ok {
		return nil, fmt.Errorf("document with ID '%s' does not exist", id)
	}
	return document, nil
}

func (ds *DocumentStore) Update(document Document) error {
	if _, ok := ds.documents[document.ID]; !ok {
		return fmt.Errorf("document with ID '%s' does not exist", document.ID)
	}

	ds.documents[document.ID] = &document
	return nil
}

// SaveToFile -> takes fileName -> Save documentStore to file with given filename
// LoadFromFile -> filename -> Load document to file with given filename

func (ds *DocumentStore) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("invalid file name - %s", err)
	}

	defer file.Close()

	// Creating a list of pointers of document with min length 0 and max length of len(ds.documents)
	documents := make([]*Document, 0, len(ds.documents))

	for _, document := range ds.documents {
		documents = append(documents, document)
	}

	data, err := json.MarshalIndent(documents, "", "	")

	if err != nil {
		return fmt.Errorf("error marshalling documents %s", err)
	}

	_, err = file.Write(data)

	if err != nil {
		return fmt.Errorf("error writing to file %s", err)
	}

	return nil
}

func (ds *DocumentStore) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("error file not exist - %s", err)
		}
		return fmt.Errorf("error loading data from file - %s", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}

	documents := make([]*Document, 0)
	err = json.Unmarshal(data, documents)

	if err != nil {
		return fmt.Errorf("error unmarshalling file into documents: %s", err)
	}

	ds.documents = make(map[string]*Document)
	for _, document := range documents {
		ds.documents[document.ID] = document
	}

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
