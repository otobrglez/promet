.PHONY: clean test-python

promet.so:
	go build \
		-ldflags="-shared" \
		-buildmode=c-shared \
		-o promet.so \
		promet.go

test-python:
	python python/test.py

clean:
	rm -f promet.so promet.h promet.a

