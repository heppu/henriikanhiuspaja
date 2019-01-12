
dist/index.html: data.json index.tmpl
	go run render.go
