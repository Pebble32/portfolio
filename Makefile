.PHONY: dev tw templ

# Run both watchers in parallel
dev:
	@make -j2 tw templ

tw:
	npx tailwindcss -i input.css -o static/css/tw.css --watch

templ:
	templ generate -watch -include-timestamp=false -include-version=false -proxy="http://localhost:3000" -proxyport="8080" -open-browser=true -cmd="go run ."
