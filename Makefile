test:
	@go test ./... -v | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/'' | GREP_COLOR="01;33;40" egrep --color=always '\s*[a-zA-Z0-9\-_.]+[:][0-9]+[:]|^'

server:
	go run .
