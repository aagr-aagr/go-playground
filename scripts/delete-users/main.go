package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type DeleteRequest struct {
	Users []User `json:"users"`
}

type User struct {
	ShortUserID int `json:"short_user_id"`
}

func readIDs(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ids []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		id, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid id %q: %w", line, err)
		}
		ids = append(ids, id)
	}
	return ids, scanner.Err()
}

func deleteUsers(baseURL, apiKey string, ids []int) error {

	req := DeleteRequest{
		Users: make([]User, len(ids)),
	}
	for i, id := range ids {
		req.Users[i] = User{ShortUserID: id}
	}

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal: &w", err)
	}

	httpReq, err := http.NewRequest("POST", baseURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Cmt-Api-Key", apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("OK -- deleted %d users (HTTP %d)\n", len(ids), resp.StatusCode)
		return nil
	}
	body, _ = io.ReadAll(resp.Body)
	return fmt.Errorf("server retuned HTTP %d: %s", resp.StatusCode, string(body))

}

func main() {

	filepath := flag.String("file", "", "path to file with ids (one per line)")
	baseURL := flag.String("url", "", "https://....")
	apiKey := flag.String("apiKey", "", "api key")
	flag.Parse()

	if *filepath == "" || *baseURL == "" || *apiKey == "" {
		fmt.Fprintln(os.Stderr, "usage: bulk-delete -file ids.txt -url https:// -apiKey")
		os.Exit(1)
	}

	ids, err := readIDs(*filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(ids) == 0 {
		fmt.Println("no IDs found in th file")
		return
	}

	fmt.Printf("Deleting %d users...\n", len(ids))

	err = deleteUsers(*baseURL, *apiKey, ids)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}
