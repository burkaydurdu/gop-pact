#!/bin/bash

pact-broker create-version-tag --pacticipant GopServer --broker-base-url=http://localhost --version 2.5.1 --tag=dev

#pact-broker can-i-deploy --pacticipant=GopServer --broker-base-url=http://localhost --version=2.5.2 --to=master