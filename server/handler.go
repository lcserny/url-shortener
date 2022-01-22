package server

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(rw, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw, r)
	}
}

func YAMLHandler(ymlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYAML(ymlBytes)
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(pathUrls), fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathUrlMap := make(map[string]string)
	for _, pu := range pathUrls {
		pathUrlMap[pu.Path] = pu.URL
	}
	return pathUrlMap
}

func parseYAML(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	if err := yaml.Unmarshal(data, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
