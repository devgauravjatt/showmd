package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

func main() {

	args := os.Args
	var fileName string

	if len(args) > 1 {
		fileName = args[1]
	}

	if !strings.Contains(fileName, ".md") {
		fmt.Println("give only md files")
	}

	fmt.Println("md show on http://localhost:4000")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	maybeUnsafeHTML := markdown.ToHTML(bytes, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)

	mainHtml := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Document</title>
   <link
  rel="stylesheet"
  href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.amber.min.css"
>
</head>

<body>
   <main style="padding:6rem">
      %s
   </main>
</body>

</html>`, string(html))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(mainHtml))
	})
	http.ListenAndServe(":4000", nil)

}
