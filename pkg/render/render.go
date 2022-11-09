package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// renders templates using html/template
// http.ResponseWriter writes to web browser
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a temple cache
	tc, err := createTemplateCache()

	if err != nil {
		// close program with log
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		// can;t find template
		// close program with log
		log.Fatal(err)
	}

	// finer grained error checking
	buf := new(bytes.Buffer)

	// parse before writing/execution
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	// same as above
	myCache := map[string]*template.Template{}

	// get all of the files name *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		log.Println("creating cache")
		// filename
		name := filepath.Base(page)

		// parse file named page and store it in a template called name
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// adds the page and all required layouts to myCache
		myCache[name] = ts
	}

	return myCache, nil
}
