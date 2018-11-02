18: 18.rs
	rustc -o $@ $<

test: 18.rs
	rustc --test -o $@ $<
	./test || true
	rm $@

run: 18
	./18

clean:
	rm -fr 18 test

.PHONY: clean test
