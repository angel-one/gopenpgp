package pmcrypto

import (
	"testing"
	"fmt"
				"io/ioutil"
)

const publicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

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
OQTmvM6R
=DvFv
-----END PGP PUBLIC KEY BLOCK-----`
const testMessage = `-----BEGIN PGP MESSAGE-----
Version: ProtonMail
Comment: https://protonmail.com

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
LloBMV1YNGsW45M+oQ==
=BRHT
-----END PGP MESSAGE-----`
const privatekeypassword = "test"
const privatekey = `
-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: OpenPGP.js v3.1.2
Comment: https://openpgpjs.org

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
=xWqS
-----END PGP PRIVATE KEY BLOCK-----
`


// define call back interface
type Callbacks struct {
}

func (t *Callbacks) onBody(body string, mimetype string) {
	fmt.Println(body)
}
func (t Callbacks) onAttachment(headers string, data []byte) {
	fmt.Println(headers)
}
func (t Callbacks) onEncryptedHeaders(headers string) {
	fmt.Println(headers)
}
func (t Callbacks) onVerified(verified int) {
	fmt.Println(verified)
}
func (t Callbacks) onError(err error) {
	fmt.Println(err)
}

func TestDecrypt(t *testing.T) {
	callbacks := Callbacks{}
	o := OpenPGP{}
	block, _ := unArmor(publicKey)
	publicKeyUnarmored, _ := ioutil.ReadAll(block.Body)
	block, _ = unArmor(privatekey)
	privateKeyUnarmored, _ := ioutil.ReadAll(block.Body)
	o.decryptMIMEMessage(testMessage, publicKeyUnarmored, privateKeyUnarmored, privatekeypassword,
		&callbacks, o.GetTime())
}

func TestParse(t *testing.T) {
	testMessage :=
		`Content-Type: multipart/signed; protocol="application/pgp-signature"; micalg=pgp-sha256; boundary="---------------------4a9ea9f4dad3f36079bdb3f1e7b75bd0"; charset=UTF-8
X-Spam-Status: No, score=1.2 required=7.0 tests=ALL_TRUSTED,DKIM_SIGNED, DKIM_VALID,DKIM_VALID_AU,FREEMAIL_ENVFROM_END_DIGIT,FREEMAIL_FROM, FREEMAIL_REPLYTO_END_DIGIT,HTML_IMAGE_ONLY_08,HTML_MESSAGE autolearn=no autolearn_force=no version=3.4.0
X-Spam-Level: *
X-Spam-Checker-Version: SpamAssassin 3.4.0 (2014-02-07) on mail.protonmail.ch

-----------------------4a9ea9f4dad3f36079bdb3f1e7b75bd0
Content-Type: multipart/mixed; boundary="---------------------f0e64db835d0f5c3674df52a164b06bb"

-----------------------f0e64db835d0f5c3674df52a164b06bb
Content-Type: multipart/alternative; boundary="---------------------3ca028eaeffb3ca0fb0dd6461f639c2b"

-----------------------3ca028eaeffb3ca0fb0dd6461f639c2b
Content-Transfer-Encoding: quoted-printable
Content-Type: text/plain;charset=utf-8




[Screenshot from 2018-02-06 17-13-21.png]

--=C2=A0
Julien Palard
https://mdk.fr


-----------------------3ca028eaeffb3ca0fb0dd6461f639c2b
Content-Type: multipart/related; boundary="---------------------ce717c368c2d3981c954ad7c46cd7bf2"

-----------------------ce717c368c2d3981c954ad7c46cd7bf2
Content-Type: text/html;charset=utf-8
Content-Transfer-Encoding: base64

PGRpdj48ZGl2Pjxicj48L2Rpdj48ZGl2Pjxicj48L2Rpdj48ZGl2IGNsYXNzPSJwcm90b25tYWls
X3NpZ25hdHVyZV9ibG9jayI+PGRpdiBjbGFzcz0icHJvdG9ubWFpbF9zaWduYXR1cmVfYmxvY2st
dXNlciI+PGRpdj48aW1nIHNyYz0iY2lkOjQ2MDk4ZTliQHByb3Rvbm1haWwuY29tIiBjbGFzcz0i
cHJvdG9uLWVtYmVkZGVkIiBhbHQ9IlNjcmVlbnNob3QgZnJvbSAyMDE4LTAyLTA2IDE3LTEzLTIx
LnBuZyI+PGJyPjwvZGl2PjxkaXY+LS0mbmJzcDs8YnI+PC9kaXY+PGRpdj5KdWxpZW4gUGFsYXJk
PGJyPjwvZGl2PjxkaXY+PGNvZGUgc3R5bGU9ImZvbnQtZmFtaWx5OidTRk1vbm8tUmVndWxhcics
IENvbnNvbGFzLCAnTGliZXJhdGlvbiBNb25vJywgTWVubG8sIENvdXJpZXIsIG1vbm9zcGFjZTtm
b250LXNpemU6MTEuODk5OTk5NjE4NTMwMjczcHg7cGFkZGluZzowcHg7bWFyZ2luOjBweDtiYWNr
Z3JvdW5kLWNvbG9yOnRyYW5zcGFyZW50O3doaXRlLXNwYWNlOnByZTtib3JkZXI6MHB4O2Rpc3Bs
YXk6aW5saW5lO292ZXJmbG93OnZpc2libGU7bGluZS1oZWlnaHQ6aW5oZXJpdDsiPjxhIGhyZWY9
Imh0dHBzOi8vbWRrLmZyIj5odHRwczovL21kay5mcjwvYT48L2NvZGU+PGJyPjwvZGl2PjwvZGl2
PjxkaXYgY2xhc3M9InByb3Rvbm1haWxfc2lnbmF0dXJlX2Jsb2NrLXByb3RvbiBwcm90b25tYWls
X3NpZ25hdHVyZV9ibG9jay1lbXB0eSI+PGJyPjwvZGl2PjwvZGl2PjwvZGl2Pg==
-----------------------ce717c368c2d3981c954ad7c46cd7bf2--
-----------------------3ca028eaeffb3ca0fb0dd6461f639c2b--
-----------------------f0e64db835d0f5c3674df52a164b06bb
Content-Type: image/png; filename="Screenshot from 2018-02-06 17-13-21.png"; name="Screenshot from 2018-02-06 17-13-21.png"
Content-Disposition: inline; filename="Screenshot from 2018-02-06 17-13-21.png"; name="Screenshot from 2018-02-06 17-13-21.png"
Content-ID: <46098e9b@protonmail.com>

wsBcBAEBCAAQBQJbjodcCRDdVS3pfakUGgAAlhoH/2jBRjOSx5EdkiyyNcCT
4DVm+ACoF1KTWE5fLRuDPvSiD934cFoZShJs32r0Wcj/4W4tVhLYzjjtf6xO
Ymqe0p3o4oYxxMXIAd4COrnOPGjeD1ausqT6iUCAadqXoDYfowEg4f0Wd0RK
ElsP/OZaDjsNoRE3WeRcHTr5XWZxhEsIMgnW591iaTliYvbysLoQ08i3G53c
p6q+IANRznx5rDhMdo+shFvhcI2Zszg6X2WuCMhFtUyrqEN8WlZYXMX8PGPO
1kNDRl7B7/r4Ap+FffLpeYw+8rG6lQXGc3RCOAnfMq9X/9Ziqzxr7flYtRJH
RIzX2CG47PuGl/uvImFW/Iw=

-----------------------f0e64db835d0f5c3674df52a164b06bb
Content-Type: application/pgp-keys; filename="publickey - kaykeytest3@protonmail.com - 0xE1DADAE3.asc"; name="publickey - kaykeytest3@protonmail.com - 0xE1DADAE3.asc"
Content-Transfer-Encoding: base64
Content-Disposition: attachment; filename="publickey - kaykeytest3@protonmail.com - 0xE1DADAE3.asc"; name="publickey - kaykeytest3@protonmail.com - 0xE1DADAE3.asc"

LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tDQpWZXJzaW9uOiBPcGVuUEdQLmpz
IHYzLjEuMg0KQ29tbWVudDogaHR0cHM6Ly9vcGVucGdwanMub3JnDQoNCnhzQk5CRnQ2N2I4QkNB
QzBTMlU3VUJ3dnkyUnZqcEpwOHlyMlE0R25oNzk0QjUwVG1TMlViTXJRaEovaQpUWFRTZmJhb21l
NkdkcmJhSjlvVW0xY2hhMU5hVlNnWHE0RVB2WWxmSWVVdkt1U0tnTVdxaVMzT2xSdGYKT3JwMVZp
dXNpZWtldVczL2JtYUZBbzFLQWZvajBMNlF6cVVIY3ZZNnY1RExoWGV2MWlaUW9TN2lrQk9PCmor
Z05lUDJabjVIOXdQTFlxdWlqK0pnWFY4djRRQk5uNU82NUNYRkZIbjl1Mzl1Yng2bDJtY2tqaDVm
VQppU1FiUkh2RzhIdG96Y2xIcWV1bjBpanNjZEtoOTBsaENsWHlDcHlkZ2JwT2pjL2dDQ2R3WmVh
djNNSG0KZnhsd2VReDhmc0xxZ1h5TlNFd3pyODh0MTFWMFQ4TmRjeFdia1lRNllQN2R4UU10RkU0
bWM5akxBQkVCCkFBSE5PU0pyWVhsclpYbDBaWE4wTTBCd2NtOTBiMjV0WVdsc0xtTnZiU0lnUEd0
aGVXdGxlWFJsYzNRegpRSEJ5YjNSdmJtMWhhV3d1WTI5dFBzTEFkUVFRQVFnQUtRVUNXM3J0dndZ
TENRY0lBd0lKRU4xVkxlbDkKcVJRYUJCVUlDZ0lERmdJQkFoa0JBaHNEQWg0QkFBQ0dEZ2Y5SHVq
aG9QUGN5SHpLd25JbFJJaENRL1dLCm13SVVFdEw4eU03aGRvci9zK0kyL1A3WkJHWE1ReXBKb3JD
YTU1NUY2ODRMZnpWdUtBUVFMdXZpNHR6aQo5N2JSNE1ldHY4a2ptSUI1Rk0vNUhpWllWd090SDFK
dU5iM2RQWG9CVjJha09ScnFCeVRjbjEyeUIwOVkKd0tZR0Y4K25XNFhLU1F2V3VBWWVIZ0dQKzMz
ZTdCc0VvODhuRmJJUXl2cmUzUTN0TkJrWTVVVnA2L1M2ClpLZWJwaEFRQVlzM0ZVVlZVeThlZkhr
aEFmZFlSdUFBUUFoWnJGanU0V1J0RGZkUHdLR0UzZDZsUGswSApNSFNTUXNHb2thQTVoTnUycE1S
Mnl5UmtUa3JyNnRXbmlET2ZvSmtKcmc2VllHNVRwaHZnOE43dVdzZHQKZkhDdDBwWW9uYTRzclZx
a0RzN0FUUVJiZXUyL0FRZ0F6b2tCSGZucnFXUXE4RUVSYUVlOUpmZkQ3V0FpCnZyU29lNDdxQ05l
S2I2MmcvLzh0Q1dKMHVlRExTMFRyY0V4cnVzUi8rSFdoYUFTV2srL0lPTGhjRTcrVgp4V2Nab0dp
bzNuc1puTktWdnFhendNSW1vMnFqdUllUnVCVkJqRW5OOERhOXlxNkJ3YVkvSEs2czM5Y3IKQVVV
d2FtWjVuMDNtSFR6MzBMQllYR1VldVc5c3Z4NDIzbTlXb3N2WlQvUWRKWW1qMG44NmpGNGs2Tnh2
ClV4ejR1bmV1YWt1Wld5NmNURHg1WFdWVWF1YzlyV1Z5Uzk3cjJ0NllpRlF6MTFiOU5oZm4rcDVs
UmdZbgpzMzQwT1MyVmppK1BVeUVKZ2YzYi9MQlJ5V3Erbk1Vc2lRV09yYlQxblVCVzNRdDViWWhH
MUhnT1o5eEoKVjd2dndSZHFwQ1p2cVFBUkFRQUJ3c0JmQkJnQkNBQVRCUUpiZXUyL0NSRGRWUzNw
ZmFrVUdnSWJEQUFBClhrTUgvaTBoTnRwRWIrYUZONC9Ba0JCam1GOFJZV243anRyK2UyU05STHdF
RUMxcmdmakhjWXZnanJFTQo2Y3hBcXVzSWJKSEdnSUVYZ0s1YUlNOHBGLzRuamN2VXVxa0x1b1o3
QVJsanpwUTcvaHhyc0FNWnFYR2gKeEU4eXJZVlY0dGhhZWw0Q1NRUTRvRWlpbXB4Z28velE5bzhk
NGNQUTM4Ni9VNEgvMm9OczFUMElCOTZWCkJsTy9pM0w5eGM4K2w4RG8vcm5ieVV1Ym9LUVNKZmtp
RXNLMkRabEZoOVArdVltQ1AzVGJSb2NIUFZtVwpBT3NHVVhJSXROaVNTc05ORUlHeVBSNE8rMXRq
VUJPMFNHRnZoVVZUNnJLYTlaYUVyMzI2ZmU2S0pmZU8KYTl4a21WZEdaQm9SdENmbHhiakdnYjRq
dTJ3Z1E1TE5KWUNWZy9WRkxIRGI3MjQ9DQo9YmNvMw0KLS0tLS1FTkQgUEdQIFBVQkxJQyBLRVkg
QkxPQ0stLS0tLQ0KDQo=
-----------------------f0e64db835d0f5c3674df52a164b06bb--
-----------------------4a9ea9f4dad3f36079bdb3f1e7b75bd0
Content-Type: application/pgp-signature; name="signature.asc"
Content-Description: OpenPGP digital signature
Content-Disposition: attachment; filename="signature.asc"

-----BEGIN PGP SIGNATURE-----
		Version: ProtonMail
Comment: https://protonmail.com

wsBcBAEBCAAQBQJbjodcCRDdVS3pfakUGgAAlhoH/2jBRjOSx5EdkiyyNcCT
4DVm+ACoF1KTWE5fLRuDPvSiD934cFoZShJs32r0Wcj/4W4tVhLYzjjtf6xO
Ymqe0p3o4oYxxMXIAd4COrnOPGjeD1ausqT6iUCAadqXoDYfowEg4f0Wd0RK
ElsP/OZaDjsNoRE3WeRcHTr5XWZxhEsIMgnW591iaTliYvbysLoQ08i3G53c
p6q+IANRznx5rDhMdo+shFvhcI2Zszg6X2WuCMhFtUyrqEN8WlZYXMX8PGPO
1kNDRl7B7/r4Ap+FffLpeYw+8rG6lQXGc3RCOAnfMq9X/9Ziqzxr7flYtRJH
RIzX2CG47PuGl/uvImFW/Iw=
=t7ak
-----END PGP SIGNATURE-----


-----------------------4a9ea9f4dad3f36079bdb3f1e7b75bd0--
`
	o := OpenPGP{}
	body, _, atts, attHeaders, err := o.parseMIME(testMessage, nil)
	if err != nil {
		t.Error(err)
	}

	fmt.Println()
	fmt.Println("==BODY:")
	fmt.Println(body.GetBody())
	fmt.Println("==BODY HEADERS:")
	fmt.Println(body.GetHeaders())

	fmt.Println("==ATTACHMENTS:")
	fmt.Println(len(atts))
	for i, attachment := range atts {
		fmt.Println("==ATTACHMENT HEADERS:")
		fmt.Println(attHeaders[i])
		fmt.Println("==ATTACHMENT:")
		fmt.Println(attachment)
	}
}

