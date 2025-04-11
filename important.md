
For generate the Generating a self-signed TLS certificate.
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost


For integration test. 

CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'test_web'@'localhost';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';
ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';


For run test

go test ./... -v

