#!/bin/bash

set -eu
set -x

# Test domains
export DNS_DOMAIN_FORWARD="example.com."
export DNS_DOMAIN_REVERSE="1.168.192.in-addr.arpa."

# Run with no authentication

export DNS_SERVER=127.0.0.1
docker run -d -p 53:53/udp \
	-e BIND_DOMAIN_FORWARD=${DNS_DOMAIN_FORWARD} \
	-e BIND_DOMAIN_REVERSE=${DNS_DOMAIN_REVERSE} \
	-e BIND_INSECURE=true \
	--name bind_insecure drebes/bind
make testacc TEST=./builtin/providers/dns
docker stop bind_insecure
docker rm bind_insecure

# Run with authentication

export DNS_KEY_NAME=${DNS_DOMAIN_FORWARD}
export DNS_KEY_ALGORITHM="hmac-md5"
export DNS_KEY_SECRET="c3VwZXJzZWNyZXQ="
docker run -d -p 53:53/udp \
	-e BIND_DOMAIN_FORWARD=${DNS_DOMAIN_FORWARD} \
	-e BIND_DOMAIN_REVERSE=${DNS_DOMAIN_REVERSE} \
	-e BIND_KEY_NAME=${DNS_KEY_NAME} \
	-e BIND_KEY_ALGORITHM=${DNS_KEY_ALGORITHM} \
	-e BIND_KEY_SECRET=${DNS_KEY_SECRET} \
	--name bind_secure drebes/bind
make testacc TEST=./builtin/providers/dns
docker stop bind_secure
docker rm bind_secure
