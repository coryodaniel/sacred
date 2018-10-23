package cmd
import (
  "os"
  "fmt"
  "log"
  "os/exec"
  "runtime"
  "io/ioutil"
  "github.com/spf13/cobra"
)

var PreviewCmd = &cobra.Command{
  Use:   "preview",
  Short: "Preview HTML files in your default browser",
  Run: func(cmd *cobra.Command, args []string) {
    for _, doc := range cfg.Docs {
  		files := Resolve(doc.Files)
  		html := Compile(files, doc.Notice)

      tmpFile, _ := ioutil.TempFile(os.TempDir(), "sacred.*.html")
      tmpFile.Write(html)
      tmpFile.Close()

      openBrowser(fmt.Sprintf("file:///%s", tmpFile.Name()))
    }
  },
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}
