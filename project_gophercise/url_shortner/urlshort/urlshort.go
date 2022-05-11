package urlshort

import (
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// type HandlerFunc func ( ResponseWriter , *Request)
	

func MapHandler(paths map[string]string, f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, res *http.Request) {
		path := res.URL.Path
		if des, ok := paths[path]; ok {
			http.Redirect(w, res, des, http.StatusFound)
			return
		}
		f.ServeHTTP(w, res)
	}
}

// yaml first comes here as slice of bytes of yaml file
func YAMLHandler(yamlBytes []byte, f http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	fmt.Println(pathUrls)
	paths := buildMap(pathUrls)
	return MapHandler(paths, f), nil
}


// secondly it goes to it and get unmarshal
func parseYaml(data []byte) ([]spathUrl, error) {
	var pathUrls []spathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

// after umarshal it comes here to be mapped in key value pairs
func buildMap(pathUrls []spathUrl) map[string]string {
	paths := make(map[string]string)
	for _, s := range pathUrls {   // iterating through the struct
		paths[s.Path] = s.URL
	}
	return paths
}

type spathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
