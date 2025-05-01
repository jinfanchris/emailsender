#!/bin/bash

set -e

CONF_DIR="$(dirname "$0")"
CERT_DIR="runtime"

mkdir -p "$CERT_DIR"

openssl req -x509 -newkey rsa:2048 -nodes \
    -keyout "$CERT_DIR/key.pem" \
    -out "$CERT_DIR/cert.pem" \
    -days 365 \
    -config "$CONF_DIR/cert.conf" \
    -extensions v3_req

echo "✅ Self-signed cert + key generated in $CERT_DIR/"
