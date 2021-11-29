openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout key.pem -out cert.pem \
  -subj "/C=US/ST=CA/L=Irvine/O=Acme Inc./CN=localhost" \
  -reqexts v3_req -reqexts SAN -extensions SAN \
  -config \
  <(echo -e '
    [req]\n
    distinguished_name=req_distinguished_name\n
    [req_distinguished_name]\n
    [SAN]\n
    subjectKeyIdentifier=hash\n
    authorityKeyIdentifier=keyid:always,issuer:always\n
    basicConstraints=CA:TRUE\n
    subjectAltName=DNS:localhost
  ')
# xcrun simctl keychain booted add-root-cert cert.pem