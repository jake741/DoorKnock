package main

import (
    "bufio"
    "crypto/tls"
    "flag"
    "fmt"
    "net/http"
    "net/url"
    "os"
    "path"
    "strings"
    "sync"
    "time"
)

type RequestTask struct {
    URL      string
    ProxyURL string
}

type Result struct {
    URL    string
    Status string
}

var verbose bool

func makeRequest(task RequestTask, wg *sync.WaitGroup, results chan<- Result) {
    defer wg.Done()

    fullURL := task.URL
    if !strings.HasPrefix(fullURL, "http") {
        fullURL = "https://" + fullURL
    }

    transport := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    if task.ProxyURL != "" {
        proxyParsed, err := url.Parse("http://" + task.ProxyURL)
        if err == nil {
            transport.Proxy = http.ProxyURL(proxyParsed)
        } else {
            fmt.Println("Invalid proxy:", err)
            return
        }
    }

    client := &http.Client{
        Transport: transport,
        Timeout:   10 * time.Second,
    }

    resp, err := client.Get(fullURL)
    if err != nil {
        if verbose {
            fmt.Println("[ERROR]", fullURL, "->", err)
        }
        return
    }
    defer resp.Body.Close()

    if verbose || resp.StatusCode != http.StatusNotFound {
        fmt.Println(fullURL, "->", resp.Status)
    }

    if resp.StatusCode != http.StatusNotFound {
        results <- Result{URL: fullURL, Status: resp.Status}
    }
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            lines = append(lines, line)
        }
    }

    return lines, scanner.Err()
}

func worker(tasks <-chan RequestTask, wg *sync.WaitGroup, results chan<- Result) {
    for task := range tasks {
        makeRequest(task, wg, results)
    }
}

func main() {
    // CLI Flags
    target := flag.String("u", "", "Single target URL")
    targetList := flag.String("l", "", "File with list of target URLs")
    wordlist := flag.String("w", "", "Path to wordlist")
    proxy := flag.String("burp", "", "Proxy address (e.g., 127.0.0.1:8080)")
    threads := flag.Int("threads", 10, "Number of concurrent threads")
    outputFile := flag.String("o", "", "Save successful results to file")
    verboseFlag := flag.Bool("v", false, "Verbose output")
    flag.Parse()

    verbose = *verboseFlag

    if *target == "" && *targetList == "" {
        fmt.Println("Usage:")
        fmt.Println("  -u        Single target (e.g., example.com)")
        fmt.Println("  -l        List of targets")
        fmt.Println("  -w        Wordlist for directories")
        fmt.Println("  -burp     Proxy (e.g., 127.0.0.1:8080)")
        fmt.Println("  -threads  Number of concurrent requests")
        fmt.Println("  -v        Verbose output")
        fmt.Println("  -o        Output file to save successful URLs")
        os.Exit(1)
    }

    if *wordlist == "" {
        fmt.Println("Error: You must provide a wordlist with -w")
        os.Exit(1)
    }

    words, err := readLines(*wordlist)
    if err != nil {
        fmt.Println("Failed to read wordlist:", err)
        os.Exit(1)
    }

    var targets []string
    if *target != "" {
        targets = append(targets, *target)
    } else {
        targets, err = readLines(*targetList)
        if err != nil {
            fmt.Println("Failed to read target list:", err)
            os.Exit(1)
        }
    }

    tasks := make(chan RequestTask, *threads*10)
    results := make(chan Result, 100)
    var wg sync.WaitGroup

    // Start workers
    for i := 0; i < *threads; i++ {
        go worker(tasks, &wg, results)
    }

    // Output handling
    var out *os.File
    if *outputFile != "" {
        out, err = os.Create(*outputFile)
        if err != nil {
            fmt.Println("Failed to create output file:", err)
            os.Exit(1)
        }
        defer out.Close()

        go func() {
            for res := range results {
                out.WriteString(fmt.Sprintf("%s -> %s\n", res.URL, res.Status))
            }
        }()
    } else {
        go func() {
            for range results {
            }
        }()
    }

    // Generate request tasks
    for _, base := range targets {
        base = strings.TrimRight(base, "/")
        for _, word := range words {
            fullPath := base + "/" + path.Clean(word)
            wg.Add(1)
            tasks <- RequestTask{
                URL:      fullPath,
                ProxyURL: *proxy,
            }
        }
    }

    go func() {
        wg.Wait()
        close(results)
    }()
    wg.Wait()
}
