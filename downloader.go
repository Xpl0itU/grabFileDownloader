package grabfiledownloader

import (
	"fmt"

	"github.com/cavaliergopher/grab/v3"
)

func DownloadFile(client *grab.Client, url string, outputPath string) error {
	req, err := grab.NewRequest(outputPath, url)
	if err != nil {
		return err
	}
	resp := client.Do(req)
	if err := resp.Err(); err != nil {
		return err
	}

	fmt.Printf("[Info] Download saved to ./%v \n", resp.Filename)
	return nil
}
