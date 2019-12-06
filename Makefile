# **************************
# *    Hyperion RestApi    *
# * Étienne "Maiste Marais *
# **************************

all: build

build: src/hyperion/main.go
	go build -o hyperion $<

run: build
	./main

clear: 
	rm -rf main
