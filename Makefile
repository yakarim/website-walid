hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	air -c air.toml

runjade:
	jade -pkg=jade -d ./jade -writer ./jade/pages/$(arg)