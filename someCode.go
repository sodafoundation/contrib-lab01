// Copyright 2023 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//Adding a comment testing

package main

import "fmt"

// this is sample function
func getLanguage(project string) string {
	var projectLang map[string]string
	projectLang = map[string]string{
		"strato": "goLang",
	"delfin": "python"
	"kahu": "goLang"
	"dashboard": "nodeJs"
	"installer": "ansible"
	}

	return projectLang[project]
}

func main() {
	project := "delfin"
	lang := getLanguage(project)
	fmt.Printf("This ode will print the project name and the language the project is implemented on ")
	fmt.Printf("the language used for project: %s is %s\n", project, lang)

}
