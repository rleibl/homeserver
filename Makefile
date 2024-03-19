
SRVDIR=/srv/homeserver/

all: build

build:
	go build 

install:
	mkdir -p $(SRVDIR)
	cp homeserver $(SRVDIR)
	cp -r assets $(SRVDIR)
	cp -r templates $(SRVDIR)
	cp homeserver.service /etc/systemd/system/
