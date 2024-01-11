cd gosop
echo "replace github.com/angel-one/gopenpgp/v2 => ../gopenpgp" >> go.mod
go get github.com/angel-one/gopenpgp/v2/crypto
go build .
