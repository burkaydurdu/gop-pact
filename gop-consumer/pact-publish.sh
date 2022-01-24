#!/bin/bash

#set -x
#
#VERSION=$1 #like 1.0.0
#
#curl -X PUT \
#    http://localhost/pacts/provider/GopServer/consumer/GopClient/version/${VERSION} \
#    -H "Content-Type: application/json" \
#    -d @/Users/burkay.durdu/Projects/Personal/Learning/Go/gop-pact/gop-consumer/pacts/gopclient-gopserver.json

#pact-broker publish ./pacts/ --consumer-app-version=1.5 --broker-base-url=http://localhost --tag=master

#pact-broker publish ./pacts/ -a=1.2 -b=http://localhost

#pact-broker create-version-tag --pacticipant GopClient --broker-base-url=http://localhost --version 1.2 --tag=dev

#pact-broker can-i-deploy --pacticipant=GopClient --broker-base-url=http://localhost --version=1.5 --to=master