#!/bin/bash

KEY_FILE=cert.key
CERT_FILE=cert.pem
HOST=grpc.kube.benco.io
CERT_NAME=self-signed

echo "Creating self signed cert"
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${KEY_FILE} -out ${CERT_FILE} -subj "/CN=${HOST}/O=${HOST}"
kubectl create secret tls ${CERT_NAME} --key ${KEY_FILE} --cert ${CERT_FILE}
