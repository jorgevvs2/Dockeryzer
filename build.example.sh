#!/bin/bash
API_KEY="your-api-key-here"
go build -ldflags "-X github.com/jorgevvs2/dockeryzer/src/config.APIKey=$API_KEY" . 