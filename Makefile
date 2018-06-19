

all: server js 

server-mac:
	go build -v ./cmd/bookmarkd
server-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./cmd/bookmarkd
js:
	cd ./web/bookmarks && ng build --prod --aot 
clean:
	go clean
	rm ./bookmarkd

sync-js: 
	rsync -r -a -v -e ssh --delete ./web/bookmarks/dist/bookmarks/ siva@foursigma.io:/var/www/bookmarks.foursigma.io/ 

sync: sync-js 
	scp  ./deploy/bookmarks.service siva@foursigma.io:/home/siva/bookmarks.service  
	ssh -t siva@foursigma.io rm /opt/bookmarks.foursigma.io/bookmarkd 
	scp  ./bookmarkd siva@foursigma.io:/opt/bookmarks.foursigma.io/  
	ssh -t siva@foursigma.io sudo mv /home/siva/bookmarks.service /etc/systemd/system 

deploy: server-linux js sync
	ssh -t siva@foursigma.io systemctl restart bookmarks.service

	

