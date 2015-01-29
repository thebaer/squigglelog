package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"strconv"
	"text/template"
	"os"
	"bufio"
	"regexp"
	"flag"
	"encoding/json"
)

var outputPath string = os.Getenv("HOME") + "/public_html/"
const entriesPath = "./entries/"
const templatesPath = "./templates/"
const defaultTemplateFile = "log"
const defaultOutFile = "~log"

type Entry struct {
	Date string
	Body []byte
}

func loadEntry(rawDate string) (*Entry, error) {
	filename := entriesPath + rawDate
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Get raw date parts for formatting
	year := rawDate[:4]
	month, moErr := strconv.Atoi(rawDate[4:6])
	if moErr != nil {
		return nil, moErr
	}
	date, dateErr := strconv.Atoi(rawDate[6:8])
	if dateErr != nil {
		return nil, dateErr
	}

	formattedDate := fmt.Sprintf("%d %s %s", date, time.Month(month).String(), year)

	return &Entry{Date: formattedDate, Body: body}, nil
}

func main() {
	fmt.Println()
	fmt.Println("    ~log generator v1.1")
	fmt.Println()

	// Get any arguments
	templateFilePtr := flag.String("t", defaultTemplateFile, "Squigglelog template file (defined name).")
	outFilePtr := flag.String("o", defaultOutFile, "Squigglelog template file (defined name).")
	flag.Parse()

	entryFiles := getEntries()
	entries := make([]Entry, len(*entryFiles))
	i := 0
	for _, file := range *entryFiles {
		entry, err := loadEntry(file)
		if err != nil {
			fmt.Printf("Error, skipping entry %s: %s\n", file, err)
			continue
		}
		fmt.Printf("Adding entry %s...\n", file)
		entries[i] = *entry
		i++
	}
	
	fmt.Printf("Using template %s...\n", *templateFilePtr)

	generateLog(entries, *templateFilePtr, *outFilePtr)

	fmt.Printf("Finished! Saved to %s%s.html\n", outputPath, *outFilePtr)
}

type Config struct {
	EntriesPath string
	TemplateFile string
}

func configuration() *Config {
	file, err := os.Open(os.Getenv("HOME") + "/.tildelog")
	if err != nil {
		return nil
	}
	d := json.NewDecoder(file)
	c := Config{}
	err = d.Decode(&c)
	if err != nil {
		fmt.Printf("Couldn't decode configuration: %s\n", err)
	}
	return &c
}

var validFileFormat = regexp.MustCompile("^[0-9]{8}$")

func getEntries() *[]string {
	files, _ := ioutil.ReadDir(entriesPath)
	fileList := make([]string, len(files))
	fileCount := 0
	// Traverse file list in reverse, i.e. newest to oldest
	for i := len(files)-1; i >= 0; i-- {
		file := files[i]
		if validFileFormat.Match([]byte(file.Name())) {
			fileList[fileCount] = file.Name()
			fileCount++
		}
	}
	fileList = fileList[:fileCount]
	return &fileList
}

func generateLog(entries []Entry, templateFile string, outFile string) {
	file, err := os.Create(outputPath + outFile + ".html")
	if err != nil {
		panic(err)
	}
	
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	template, err := template.ParseGlob(templatesPath + "*.html")
	if err != nil {
		panic(err)
	}
	template.ExecuteTemplate(writer, templateFile, entries)
	writer.Flush()
}

