<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  
</head>
<body>

<h1>HTTP Server in Go 🚀</h1>

<p>
  A simple HTTP server built <b>from scratch in Go</b> to understand how HTTP works behind the scenes.<br>
  This project does not rely on Go's <code>net/http</code> package for request/response handling, but instead manually parses the 
  <b>request line, headers, and body</b>, and constructs HTTP responses.
</p>

<hr>



<h2>⚡ Features</h2>
<ul>
  <li>✅ Parse <b>HTTP request line</b> (method, path, version)</li>
  <li>✅ Parse <b>HTTP headers</b></li>
  <li>✅ Parse <b>HTTP request body</b></li>
  <li>✅ Send back a valid <b>HTTP response</b> to the client</li>
</ul>

<hr>

<h2>🔥 Motivation</h2>
<p>
  I wanted to <b>learn how HTTP actually works</b> behind the scenes.<br>
  Instead of using Go’s built-in <code>net/http</code>, I decided to <b>implement the fundamentals from scratch</b> — 
  manually parsing requests and writing responses — to gain a deeper understanding of networking and the HTTP protocol.
</p>

<hr>



<h2>🛠️ Getting Started</h2>

<h3>1. Clone the Repository</h3>
<pre>
git clone https://github.com/akhand08/http-server-golang.git
cd http-server-golang
</pre>

<h3>2. Run the Server</h3>
<pre>
go run ./cmd/httpserver
</pre>

<p>You should see:</p>
<pre>
Server started on port 42069
</pre>

<h3>3. Test with curl</h3>
<p>Open another terminal and run:</p>
<pre>
curl localhost:42069/home
</pre>

<p>You should get a response like:</p>
<pre>
Hurreh, a warm welcome to Home
</pre>

<hr>

<h2>📂 Project Structure</h2>

<pre>
.
├── cmd
│   ├── httpserver
│   │   └── main.go          # Main entry for the HTTP server
│   └── tcplistener
│       └── main.go          # Basic TCP listener example
├── go.mod
├── go.sum
├── internal
│   ├── handlers             # Request handlers
│   │   ├── coffeHandler.go
│   │   ├── handlers.go
│   │   └── homeHandler.go
│   ├── headers              # HTTP headers parsing
│   │   ├── headers.go
│   │   └── headers_test.go
│   ├── request              # Request parsing logic
│   │   ├── request.go
│   │   └── request_test.go
│   ├── response             # Response building logic
│   │   └── response.go
│   └── server               # Core server implementation
│       └── server.go
├── message.txt
└── trial.txt
</pre>

<hr>
</pre>

<hr>

<h2>🤝 Contributing</h2>
<p>
  This project is mainly for learning purposes, but feel free to fork it, explore, and improve it!
</p>

<hr>



</body>
</html>
