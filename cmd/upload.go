package cmd

import (
	"fmt"
  "path/filepath"
  "io/ioutil"
	"github.com/spf13/cobra"
  "github.com/coryodaniel/sacred/confluence"
)

var UploadCmd = &cobra.Command{
  Use:   "upload",
  Short: "Uploads to Confluence API",
  Run: func(cmd *cobra.Command, args []string) {
    LoadConfig()

    for _, doc := range cfg.Docs {
  		files := Resolve(doc.Files)
  		html := Compile(files, doc.Notice)

			if outputDir != "" {
				createHtmlFile(doc.Name, html)
			}

      client := confluence.NewClient(cfg.Auth.Domain, cfg.Auth.Token, nil)
  		content, _, _ := client.ContentService.Get(doc.ContentId)

			body := confluence.ContentRequestPayload(content.Space.Key, content.Version.Number, doc.Name, string(html))

			if !dryRun {
				client.ContentService.Update(doc.ContentId, body)
			}
    }
  },
}

func createHtmlFile(name string, output []byte) {
	outputFileName := filepath.Join(outputDir, name)
	htmlFile := fmt.Sprintf("%s.html", outputFileName)
	if verbose {
		fmt.Printf("Creating: %s\n", htmlFile)
	}
	ioutil.WriteFile(htmlFile, output, 0644)
}
