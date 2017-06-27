# Package configuration
PROJECT = core
COMMANDS =

# Including ci Makefile
MAKEFILE = Makefile.main
CI_REPOSITORY = https://github.com/src-d/ci.git
CI_FOLDER = .ci

$(MAKEFILE):
	@git clone --quiet $(CI_REPOSITORY) $(CI_FOLDER); \
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
