package handlers

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	content := `
<!DOCTYPE html>
<html>
	<head>
		<title>Homepage - web</title>
		<link rel="stylesheet" href="static/css/styles.css" />
		<script async src="static/js/index.js"></script>
	</head>
	<body>
		<h1>Echorand Corp. This is the homepage for project web.</h1>
	</body>
</html>
`
	fmt.Fprintf(w, content)
}
