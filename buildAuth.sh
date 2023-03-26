#! /bin/sh
docker build --tag twiddit_auth_ms .
docker run -d -p 1414:1414 twiddit_auth_ms