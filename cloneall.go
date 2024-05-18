package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	apiURL = "https://api.github.com"
)

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cloneall <GitHub org or user URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		fmt.Println("Invalid URL. Please provide a valid GitHub organization or user URL.")
		os.Exit(1)
	}

	usernameOrOrg := parts[len(parts)-1]
	repos := fetchRepos(usernameOrOrg)
	cloneRepos(usernameOrOrg, repos)
	generateMarkdown(usernameOrOrg, repos)
}

func fetchRepos(usernameOrOrg string) []Repo {
	reposURL := fmt.Sprintf("%s/users/%s/repos", apiURL, usernameOrOrg)
	if strings.Contains(usernameOrOrg, "orgs") {
		reposURL = fmt.Sprintf("%s/orgs/%s/repos", apiURL, usernameOrOrg)
	}

	resp, err := http.Get(reposURL)
	if err != nil {
		log.Fatalf("Failed to fetch repos: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch repos: %s", resp.Status)
	}

	var repos []Repo
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	return repos
}

func cloneRepos(usernameOrOrg string, repos []Repo) {
	// Create a directory named after the username or organization
	if err := os.MkdirAll(usernameOrOrg, 0755); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	for _, repo := range repos {
		cloneURL := fmt.Sprintf("https://github.com/%s/%s.git", usernameOrOrg, repo.Name)
		fmt.Printf("Cloning %s...\n", cloneURL)

		cmd := exec.Command("git", "clone", cloneURL, filepath.Join(usernameOrOrg, repo.Name))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Printf("Failed to clone %s: %v", cloneURL, err)
		}
	}
}

func generateMarkdown(usernameOrOrg string, repos []Repo) {
	filename := filepath.Join(usernameOrOrg, "README.md")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create markdown file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("# Repositories of %s\n\n", usernameOrOrg))
	if err != nil {
		log.Fatalf("Failed to write to markdown file: %v", err)
	}

	for _, repo := range repos {
		_, err := file.WriteString(fmt.Sprintf("## [%s](%s)\n\n%s\n\n", repo.Name, repo.HTMLURL, repo.Description))
		if err != nil {
			log.Fatalf("Failed to write to markdown file: %v", err)
		}
	}

	fmt.Printf("Markdown file generated at %s\n", filename)
}
