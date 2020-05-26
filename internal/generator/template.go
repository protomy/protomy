/*
 * Copyright 2020 Anthony Burns
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import "net/url"

// Template represents a Protomy template.
type Template struct {
	URI  url.URL `json:"uri,string" yaml:"uri,string"`
	Name string  `json:"name,string" yaml:"name,string"`
}

// NewTemplateFromRawURI returns a new Template from the specified URI string.
func NewTemplateFromRawURI(rawURI string) (*Template, error) {
	uri, err := url.Parse(rawURI)
	if err != nil {
		return nil, err
	}

	return NewTemplateFromURI(*uri)
}

// NewTemplateFromURI returns a new Template from the specified url.URL.
func NewTemplateFromURI(uri url.URL) (*Template, error) {
	return &Template{URI: uri}, nil
}

func getSupportedTemplateSchemes() []string {
	return []string{
		// TODO: "bzr",
		// TODO: "bzr+ssh"
		"file",
		// TODO: "git",
		// TODO: "http",
		// TODO: "https",
		// TODO: "scp",
		// TODO: "sftp",
		// TODO: "ssh",
	}
}
