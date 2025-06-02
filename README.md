URL Directory Scanner 🔍

A simple yet powerful HTTP directory scanner written in Go. This tool allows you to test and explore different directories and endpoints on a list of target websites using a wordlist. It's useful for discovering hidden pages, paths, or files on a web server. 🚀

Features ✨
Target Options: Scan a single URL or a list of URLs 🌐

Wordlist Integration: Specify a wordlist to append to each target URL for comprehensive directory scanning 📜

Proxy Support: Use an HTTP proxy for routing requests (e.g., Burp Suite) 🕵️‍♂️

Concurrency: Perform concurrent requests with adjustable threads for faster results ⚡

Verbose Mode: Get detailed output on each request to help debug or gather more information 💬

Result Saving: Optionally save successful results to a file for later review 💾

Customizable Timeout: Define a custom timeout for requests ⏳

TLS Verification: Optionally disable SSL certificate verification for testing 🔐

Installation 🛠️
Clone the repository and build the tool:

bash
Copy
Edit
git clone https://github.com/your-username/url-directory-scanner.git
cd url-directory-scanner
go build -o dirscan .
Usage 📚
Basic Usage
bash
Copy
Edit
./dirscan -u <target-url> -w <wordlist> -o <output-file> -threads <num-threads>
-u: Single target URL (e.g., example.com) 🌍

-l: File containing a list of target URLs (one per line) 📄

-w: Path to the wordlist file (e.g., common directories like admin, login) 📑

-burp: Proxy address (e.g., 127.0.0.1:8080) 🌐

-threads: Number of concurrent threads (default is 10) ⚙️

-o: Output file to save successful results 💬

-v: Verbose output to show detailed request status 🧐

-h: Show the help message ❓

Example 1: Scan a Single Target 🧑‍💻
bash
Copy
Edit
./dirscan -u example.com -w wordlist.txt -o results.txt -threads 20 -v

Example 2: Scan Multiple Targets from a List 📋
bash
Copy
Edit
./dirscan -l targets.txt -w wordlist.txt -o results.txt -threads 20

Example 3: Use a Proxy 🕵️‍♀️
bash
Copy
Edit
./dirscan -u example.com -w wordlist.txt -burp 127.0.0.1:8080 -threads 20

Output 📊
For each successful request (non-404 status), the tool will print the URL and its HTTP status code. You can also save the results to a file, which will contain each URL with its respective HTTP status.

Example output:

arduino
Copy
Edit
https://example.com/admin -> 200 OK ✅
https://example.com/login -> 301 Moved Permanently 🔄
Example Output File:
arduino
Copy
Edit
https://example.com/admin -> 200 OK ✅
https://example.com/login -> 301 Moved Permanently 🔄
Contributions 🤝
Feel free to open issues or submit pull requests to improve the tool. Contributions are welcome! 💡

License 📝
This project is licensed under the MIT License - see the LICENSE file for details.

How to Add Images 📸:
Image Hosting: Use a service like Imgur or host the images directly in your GitHub repo under /assets/images/ or a similar folder.

Replace Placeholder URLs: Change the image URLs like https://example.com/banner-image.png to your actual hosted URLs.

Let me know if you'd like more changes or have other questions! 😊
