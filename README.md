[![.github/workflows/master.yaml](https://github.com/heppu/henriikanhiuspaja/actions/workflows/master.yaml/badge.svg)](https://github.com/heppu/henriikanhiuspaja/actions/workflows/master.yaml) [![pages-build-deployment](https://github.com/heppu/henriikanhiuspaja/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/heppu/henriikanhiuspaja/actions/workflows/pages/pages-build-deployment)

# Henriikan hiuspaja

Web pages for [henriikanhiuspaja.fi](henriikanhiuspaja.fi)

## Project layout

- `data.json` Web site content that can be edited
- `index.tmpl` Template for index.html
- `render.go` Go script that renders data.json to index.tmpl
- `dist` Static content where render.go inserts index.html

## Deployment

Changes in master will trigger travis which runs `go run render.go` and deploys dist to github pages.
