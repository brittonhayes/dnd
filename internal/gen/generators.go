package gen

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"
)

func GenerateMocks() error {
	for i, e := range Endpoints {
		res, err := http.Get(e.URL)
		if err != nil {
			return err
		}

		// read all response body
		data, _ := ioutil.ReadAll(res.Body)
		Endpoints[i] = EndpointDef{
			Name:     e.Name,
			URL:      e.URL,
			Response: string(data),
		}

		// close response body
		res.Body.Close()

		// Artificial rate limit to be a good
		// internet citizen
		time.Sleep(100 * time.Millisecond)
	}

	f, err := os.Create("internal/mocks/mocks.go")
	if err != nil {
		return err
	}
	defer f.Close()

	tpl := template.Must(template.ParseGlob("templates/mock.tmpl"))
	return tpl.ExecuteTemplate(f, "mock.tmpl", Endpoints)
}

func GenerateServices(services []ServiceDefinition) error {
	tpl, err := template.New("service.tmpl").Funcs(template.FuncMap{
		"typeOf": func(i interface{}) string {
			return reflect.TypeOf(i).String()
		},
		"typeName": func(i interface{}) string {
			return reflect.TypeOf(i).Name()
		},
		"toString": func(i interface{}) string {
			return fmt.Sprintf("%#v", i)
		},
		"arg": func(arg string) string {
			if arg != "" {
				return fmt.Sprintf("%s string", strings.ToLower(arg))
			}

			return ""
		},
	}).ParseFiles("internal/gen/service.tmpl")
	if err != nil {
		return err
	}

	for _, s := range services {
		name := strings.ToLower(s.Name) + "_gen" + ".go"
		if fileExists(name) {
			if err := os.Remove(name); err != nil {
				panic(err)
			}
		}

		f, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}

		err = tpl.ExecuteTemplate(f, "service.tmpl", &s)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
