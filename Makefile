# **************************
# *    Hyperion RestApi    *
# * Étienne "Maiste Marais *
# **************************

EXE=hyperion

all: build

build: src/hyperion/main.go
	go build -o $(EXE) $<

run: build
	./$(EXE)

clear: 
	rm -rf $(EXE)
