all: build

##@ Development

init: 
	npm install

lint:  ## Lint next files
	npm run lint
	
##@ Build

build: init lint
	npm run export

clean:
	rm -rf .next dist
