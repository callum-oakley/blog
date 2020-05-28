publishdir = callum-oakley.github.io
publishrepo = git@github.com:callum-oakley/$(publishdir).git

build/.dirstamp: $(shell fd . content templates)
	anne && \
	touch build/.dirstamp

build/.publish: build/.dirstamp
	cd build && \
	git init && \
	git add . && \
	git commit -m publish && \
	git remote add origin $(publishrepo) && \
	git push -f && \
	touch .publish

.PHONY: publish
publish: build/.publish

.PHONY: serve
serve: build/.dirstamp
	http-server build

.PHONY: pull-private
pull-private:
	git clone $(publishrepo) && \
	cp -r $(publishdir)/fonts content/ && \
	cp -r $(publishdir)/slides/private content/slides/ && \
	rm -rf $(publishdir)

.PHONY: clean
clean:
	-rm -rf build
