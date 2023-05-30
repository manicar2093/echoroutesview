test:
	@ ginkgo -v ./...

mocking:
	@ rm -r mocks
	@ mockery
