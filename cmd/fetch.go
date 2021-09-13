/*
Copyright © 2020 wt-l00

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cmd will command and control you.
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch kernel",
	Long:  "fetch kernel",
	Run: func(cmd *cobra.Command, args []string) {
		fetch(args[0])
	},
	Args: cobra.ExactValidArgs(1),
}

var baseURL = "https://kernel.ubuntu.com/~kernel-ppa/mainline/"

func makeVersionID(version string) string {
	number := ""
	rc := ""
	versionID := ""

	parts := strings.Split(version, "-")
	if len(parts) == 2 {
		number = parts[0]
		rc = parts[1]
	} else {
		number = version
	}

	parts = strings.Split(number, ".")
	switch len(parts) {
	case 2:
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		versionID = fmt.Sprintf("%d.%d.0-%02d%02d00", first, second, first, second) + rc
	case 3:
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		third, _ := strconv.Atoi(parts[2])
		versionID = fmt.Sprintf("%d.%d.%d-%02d%02d%02d", first, second, third, first, second, third)
	default:
		log.Fatal("ERROR")
	}

	return versionID
}

func makeURLs(version string) []string {
	// remove "v", e.g. v5.0.0 => 5.0.0
	version = strings.TrimPrefix(version, "v")
	baseVersionURL := baseURL + "v" + version
	res, err := http.Get(baseVersionURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	versionID := makeVersionID(version)
	regexpression := fmt.Sprintf(`(amd64\/)?linux-([\w-])+-%s(-generic)?_%s.([\d])+_(amd64|all).deb`, versionID, versionID)

	r := regexp.MustCompile(regexpression)
	hrefs := []string{}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if r.MatchString(href) {
			hrefs = append(hrefs, href)
		}
	})

	m := make(map[string]bool)
	uniqURLs := []string{}

	for _, href := range hrefs {
		if !m[href] {
			m[href] = true
			uniqURLs = append(uniqURLs, baseVersionURL+"/"+href)
		}
	}

	return uniqURLs
}

func download(url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, fileName := path.Split(url)
	file, err := os.Create(path.Join(".", fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, res.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}

func fetch(version string) {
	urls := makeURLs(version)
	var wg sync.WaitGroup

	for _, url := range urls {
		log.Println("downloading: " + url)
		wg.Add(1)
		go download(url, &wg)
	}

	log.Println("Wait for finishes to download")
	wg.Wait()
	log.Println("Finish!!")
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
