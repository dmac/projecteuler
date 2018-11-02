24: 24.rs
	rustc --opt-level 2 -o $@ $<

test: 24.rs
	rustc --test -o $@ $<
	./test || true
	rm $@

run: 24
	./24

clean:
	rm -fr 24 test

.PHONY: clean test
