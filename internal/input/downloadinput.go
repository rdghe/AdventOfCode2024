package input

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"AdventOfCode2024/internal/xerrors"
)

var AocRootURL = "https://adventofcode.com/2024"

func Download(session, day string) error {
	if session == "" {
		return xerrors.SessionTokenMissingError
	}
	if day == "" {
		day = "1"
	}

	log.Printf("downloading input for day %s\n", day)

	url := AocRootURL + "/day/" + day + "/input"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set(
		"Cookie",
		fmt.Sprintf("session=%s", session),
	)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	sb := string(body)

	f, err := os.Create(fmt.Sprintf("inputs/day%s.txt", day))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(sb)
	if err != nil {
		return err
	}
	return nil
}
