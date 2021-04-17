package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `<!DOCTYPE html>
    <html>
        <head>
            <meta http-equiv="content-type" content="text/html; charset=utf-8">
        <title>{{ .Title }}</title>
    </head>
    <body>
        {{ .Body }}
    </body>
    </html>
    `
)

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternative template name")
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, tFname string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return err
	}

	// Create temporary file and check for errors
	temp, err := ioutil.TempFile("", "mdp")
	if err != nil {
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name() + ".html"
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outName)

	return preview(outName)
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	// Parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTMl
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// Parse the content of the defaultTemplate const into a new Template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// If user provided alternate template file, replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}

	// Instatiate the content type, adding the title and body
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Execute the template with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	// Write the bytes to the file
	return ioutil.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	// Locate 'open' command in PATH
	openPath, err := exec.LookPath("open")
	if err != nil {
		return err
	}

	// Open the file
	if err := exec.Command(openPath, "-a", "Brave Browser", fname).Start(); err != nil {
		return err
	}

	// Give the browser some time to open the file before it can be deleted
	// by the run function
	time.Sleep(2 * time.Second)

	return nil
}

// content type represents the HTMl content to add into the template
type content struct {
	Title string
	Body  template.HTML
}
