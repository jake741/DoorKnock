
# 🚪 DoorKnock - Web Directory Brute Forcer 🔓

DoorKnock is a fast and efficient tool for scanning and brute-forcing directories on web servers using a wordlist. 🌐💻 It supports concurrent requests, proxy configurations, and the ability to save results. Perfect for security assessments, discovering hidden resources, and finding forgotten files! 🔍🕵️‍♂️


## 🛠️ Features:


- 🌍 Multi-target Support: Scan a single target or multiple websites at once.
- ⚡ Concurrent Threads: Perform fast scans with multiple threads for quicker results.
- 🕵️‍♂️ Verbose Output: Detailed logs for every request when enabled.
- 💼 Proxy Support: Send requests through a proxy server for anonymity or debugging.
- 📝 Save Results: Export successful URLs to a file.
- 🔒 TLS/SSL Skipping: Skips certificate validation for more flexible requests.

## 📝 Usage:


```terminal
  -u        Single target (e.g., example.com)
  -l        List of targets
  -w        Wordlist for directories
  -burp     Proxy (e.g., 127.0.0.1:8080)
  -threads  Number of concurrent requests
  -v        Verbose output
  -o        Output file to save successful URLs

```
## ⚡ How It Works:

- 🚀 Takes a base URL and attempts to discover directories/files using a wordlist.

- 🏃‍♂️ Executes multiple concurrent requests for faster results.

- 🔍 Logs successful directory hits and optionally exports them to a file.

## 🚧 Requirements:

- Go 1.18+ 🦠

- A wordlist for directories 📄

## ⚙️ Example:

``` golang
go run main.go -u example.com -w /path/to/wordlist.txt -o results.txt
```
