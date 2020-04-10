INDEX := dist/index.html
INPUT := data.json index.tmpl render.go

all:${INDEX}

${INDEX}: ${INPUT}
	go run render.go

clean:
	rm ${INDEX}
