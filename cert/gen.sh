rm *.pem
# 1. generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem

echo "CA's self signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. generate web server's private key and self certificate request
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem

# 3. use CA's private key to sign web-server's CSR to get back signed certificate
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "server's self signed certificate"
openssl x509 -in server-cert.pem -noout -text