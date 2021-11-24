openssl req \
    -newkey rsa:2048 \
    -x509 \
    -nodes \
    -keyout key.pem \
    -new \
    -out cert.pem \
    -subj /CN=dev.mycompany.com \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat /System/Library/OpenSSL/openssl.cnf \
        <(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
    -sha256 \
    -days 3650