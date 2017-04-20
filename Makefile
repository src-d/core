# Package configuration
PROJECT = core
COMMANDS =
CODECOV_TOKEN = 3738ca87-a09f-4e52-9cf1-d50e15401b72
SRCD_WORKS = true

# Including devops Makefile
MAKEFILE = Makefile.main
DEVOPS_REPOSITORY = https://github.com/src-d/devops.git
DEVOPS_FOLDER = .devops
CI_FOLDER = .ci

$(MAKEFILE):
	@git clone --quiet $(DEVOPS_REPOSITORY) $(DEVOPS_FOLDER); \
	cp -r $(DEVOPS_FOLDER)/ci .ci; \
	rm -rf $(DEVOPS_FOLDER); \
	cp $(CI_FOLDER)/$(MAKEFILE) .;

-include $(MAKEFILE)

ensure-models-generated:
	go get -v -u `go list -f '{{ join .Deps  "\n"}}' . | grep kallax | grep -v types`; \
	go generate ./...; \
	git --no-pager diff; \
	if [ `git status | grep 'Changes not staged for commit' | wc -l` != '0' ]; then \
		echo 'There are differences between the commited kallax.go and the one(s) generated right now'; \
		exit 2; \
	fi; \
