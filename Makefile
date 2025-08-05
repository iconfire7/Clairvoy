all: clean clairvoy

clairvoy:
	go build -o clairvoy main.go

clean:
	rm -rf *.o clairvoy