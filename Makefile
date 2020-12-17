.PHONY: format check test run_file

format:
	bash ./scripts/format.sh
check:
	bash ./scripts/check.sh
test:
	bash ./scripts/test.sh
run_file: format check test
