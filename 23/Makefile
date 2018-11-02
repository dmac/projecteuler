23: 23.rs
	rustc -o $@ $<

test: 23.rs
	rustc --test -o $@ $<
	./test || true
	rm $@

run: 23
	./23

clean:
	rm -fr 23 test

.PHONY: clean test
