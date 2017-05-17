dist=build
pkg=github.com/kapmahc/fly/web

VERSION=`git rev-parse --short HEAD`
BUILD_TIME=`date -R`
AUTHOR_NAME=`git config --get user.name`
AUTHOR_EMAIL=`git config --get user.email`
COPYRIGHT=`head -n 1 LICENSE`
USAGE=`sed -n '3p' README.md`

build: vue backend
	tar jcvf dist.tar.bz2 $(dist)


backend:
	go build -ldflags "-s -w -X ${pkg}.Version=${VERSION} -X '${pkg}.BuildTime=${BUILD_TIME}' -X '${pkg}.AuthorName=${AUTHOR_NAME}' -X ${pkg}.AuthorEmail=${AUTHOR_EMAIL} -X '${pkg}.Copyright=${COPYRIGHT}' -X '${pkg}.Usage=${USAGE}'" -o ${dist}/fly main.go
	-cp -rv locales db templates $(dist)/

vue:
	mkdir -pv $(dist)
	cd dashboard && npm run build
	cp -rv dashboard/dist $(dist)/public

ng2:
	mkdir -pv $(dist)
	cd ng2-admin && npm run build:prod:aot
	cp -rv ng2-admin/dist $(dist)/public


clean:
	-rm -rv $(dist) dist.tar.bz2
	-rm -rv dashboard/dist
	# -rm -rv ng2-admin/dist

init:
	govendor sync
	cd dashboard && npm install
	# cd ng2-admin && npm install
