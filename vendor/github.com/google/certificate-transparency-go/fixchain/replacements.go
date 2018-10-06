// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fixchain

import (
	"encoding/pem"
	"log"

	"github.com/google/certificate-transparency-go/x509"
)

// Go has no PKCS#7 implementation. Rather than fix that, manually
// replace the few PKCS#7 URLs we know of.

var replacements = map[string][]string{
	"http://gca.nat.gov.tw/repository/Certs/IssuedToThisCA.p7b": {
		// subject=/C=TW/O=\xE8\xA1\x8C\xE6\x94\xBF\xE9\x99\xA2/OU=\xE6\x94\xBF\xE5\xBA\x9C\xE6\x86\x91\xE8\xAD\x89\xE7\xAE\xA1\xE7\x90\x86\xE4\xB8\xAD\xE5\xBF\x83
		// issuer=/C=TW/O=Government Root Certification Authority
		`-----BEGIN CERTIFICATE-----
MIIFJTCCAw2gAwIBAgIQCI3SljuLYpwZTjIA2nfOLDANBgkqhkiG9w0BAQsFADA/
MQswCQYDVQQGEwJUVzEwMC4GA1UECgwnR292ZXJubWVudCBSb290IENlcnRpZmlj
YXRpb24gQXV0aG9yaXR5MB4XDTEzMDEzMTAzMjIzNFoXDTMzMDEzMTAzMjIzNFow
RDELMAkGA1UEBhMCVFcxEjAQBgNVBAoMCeihjOaUv+mZojEhMB8GA1UECwwY5pS/
5bqc5oaR6K2J566h55CG5Lit5b+DMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAtX7xPZUtp5iBGQvqYghJUoLeyCJJoaTcc1bcOGHij64WUBYpDu8KKEQK
R1y3zjcDXrLcZX483tmNs92DXSNuBHlx+we1aFyuLpKQVCji97ys3KeMxAEcaXqo
3cZu8nY3g//zkvX80G4RoCyDR86Z420R3mb0GlVw/9TEK8+oduZqAArEdfionpbE
K5zZ/8qaaHafgqMBQGuzfccDKLoWRcTzu3S0IvOpVU6pcB0rJOtc4F7c16tQdXfo
a8sjfcveKKbUQF6AklwugRufHdLqEVpOiGRcDPaHtT6SHJ7D/t+A/rAXMPidcksQ
rea/E+5+lehqEMHSA/gSLa9Ph+/gDQIDAQABo4IBFjCCARIwHwYDVR0jBBgwFoAU
1Wcd4Jx6LJzLxZjnHQcmKobsdM0wHQYDVR0OBBYEFNEYZ8NX/hKakWtfXzHqPsKE
h/u9MA4GA1UdDwEB/wQEAwIBBjAUBgNVHSAEDTALMAkGB2CGdmUAAwMwEgYDVR0T
AQH/BAgwBgEB/wIBADA+BgNVHR8ENzA1MDOgMaAvhi1odHRwOi8vZ3JjYS5uYXQu
Z292LnR3L3JlcG9zaXRvcnkvQ1JMMi9DQS5jcmwwVgYIKwYBBQUHAQEESjBIMEYG
CCsGAQUFBzAChjpodHRwOi8vZ3JjYS5uYXQuZ292LnR3L3JlcG9zaXRvcnkvQ2Vy
dHMvSXNzdWVkVG9UaGlzQ0EucDdiMA0GCSqGSIb3DQEBCwUAA4ICAQBuDj29K1o3
rfT72lhocx0vr18PUI5OEVfiMn+cwE8al5UdPgYAMQL4YIdA1rmL5QResEaC03d7
jFKF1fnGf7rd0k5O47iAa7THDQFtVOks1djLfNecn1l4pdLODWGRNy+DbbqAl87d
at2HSP5OEOl3nt8TxUVRsJx9TDx1IZC+RhUTT8ryalhlJ9UbxORjqbL3C7mMhviY
B9aA4aV6AFa1oAsI+LeXIB9xxmk8V8kzX1VhJ00buIAIjScIhvI39zoeF7z39hzy
Gw9+Av/AnbC4npDvvaLxIhs75LD1Tuh5WY4lk0+/PzdhrK5R0+YaOEoEvpiZljeZ
QuXJVZ08Re6Omb5XYKZ9hjtp+wAIH97k7spxSOFmP76WBy/5o22vxosfvybTxuM3
GFih8XlhoL6UYQ2e29WHW9Mj5yDN00TRp9CYWw7p5sS09PQitGKqYx7AYhJnNBy5
mz4uHLm4nQVI/3jhDb9Xgr+3UHMjz4LM8TQVh2YEDYBYkgH35WhK6pY852yAvIat
usT8CveFOCjr2uJBXaBgmBr4I5/1oJypzZLmP65VtxSMtA5cmgooVRGAe+QrYufJ
lGZiIUjmkqpNzh5q6oShkzqJPpqRviug2oZXQb9q9Qgj4zkr8KA5NkYVG+KNWR5V
LD9SyuP1AJcZmxUKQtDEZJCJtISXfybg0A==
-----END CERTIFICATE-----`,
		// subject=/C=TW/O=\xE8\xA1\x8C\xE6\x94\xBF\xE9\x99\xA2/OU=\xE6\x94\xBF\xE5\xBA\x9C\xE6\x86\x91\xE8\xAD\x89\xE7\xAE\xA1\xE7\x90\x86\xE4\xB8\xAD\xE5\xBF\x83
		//issuer=/C=TW/O=Government Root Certification Authority
		`-----BEGIN CERTIFICATE-----
MIIFJTCCAw2gAwIBAgIRAP+94tm8qUrtFSYcQfB4flUwDQYJKoZIhvcNAQEFBQAw
PzELMAkGA1UEBhMCVFcxMDAuBgNVBAoMJ0dvdmVybm1lbnQgUm9vdCBDZXJ0aWZp
Y2F0aW9uIEF1dGhvcml0eTAeFw0wMzAzMDMwNjUxMjNaFw0yMzAzMDMwNjUxMjNa
MEQxCzAJBgNVBAYTAlRXMRIwEAYDVQQKDAnooYzmlL/pmaIxITAfBgNVBAsMGOaU
v+W6nOaGkeitieeuoeeQhuS4reW/gzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAJXm+CDDPby1h9VhQFkpeEceeJcodEfbO8og0FgZj/cwtbnJbFgpl1KD
GvbBO+GvpxLI18c/T4bMEpDlpugcvlF8tIXT5M35lVTByx7NtT5myohevmppbsnK
1wRC1WwYxWHG3tIRZQdX/mKV0tE1LMDHylTeJWbA+fTPmapiTy7vopSFSmx7zEAI
9x+6p3bWfZ+A/MQh4SC0gUM6sZhlNH13z3VCaK670xHoTryQi5Yv+SuHJz8tIt31
xty/nwrpS8eM2lzonL0zDrHf/wkTIpy+ZU+J3J9gi1l1DTUyEL1thj5udyAnAYBy
aDqT5t3Af7X1DNthPknckzOHz3bOH+0CAwEAAaOCARUwggERMB8GA1UdIwQYMBaA
FMzM78wpYKQ7sZK2PPoyYo+sJRU7MB0GA1UdDgQWBBTk3BdvIqrO+MghGtKrzlOO
TtoYfDAOBgNVHQ8BAf8EBAMCAQYwFAYDVR0gBA0wCzAJBgdghnZlAAMDMBIGA1Ud
EwEB/wQIMAYBAf8CAQAwPQYDVR0fBDYwNDAyoDCgLoYsaHR0cDovL2dyY2EubmF0
Lmdvdi50dy9yZXBvc2l0b3J5L0NSTC9DQS5jcmwwVgYIKwYBBQUHAQEESjBIMEYG
CCsGAQUFBzAChjpodHRwOi8vZ3JjYS5uYXQuZ292LnR3L3JlcG9zaXRvcnkvQ2Vy
dHMvSXNzdWVkVG9UaGlzQ0EucDdiMA0GCSqGSIb3DQEBBQUAA4ICAQBrJPxFGTC3
I2z8AR23BfDjfrQ0Tr8D5ggx5GaPXSe2Re3nHW/mj3TEnpBFMDulZVo8PY8nPx7I
OibTWRAYVkEk33HrAph+FoCAfLnv5BGIwa9KH2FyguebNv6djFnzJf0E6uBOWtN/
Tkl1NOh+Q1PuJ3gGoWt8YVr5UVu+Y84LD6lLf294Y6j76sAlsa4z0HDbVx28qiTP
aPrDkWYIEb5hSW174uazQjiU4yFZMCzwUyGNCAjCoPVPqZ3FIdHdn21KbXJ0Oq87
7bCKtMy/SRUBCeVhDr7DbmELdDgHPjSXqLknMPnrI605jryzpw2nNPxOT/uO9e7n
Cxg1S6wtfVtYUx491RGni1FjARZPh3xGizY+O6UxojdEP3jId4A+LJgn/SrijKTw
m2kiQ3Q7ovMuMKGJdLHeiIOIcyVs7F+Fld2kkfKl7ztQ3mCanuQhvSLaaRSyAiMm
Cfvc6LJHttAQv/8nFQgTvDAWSkGmba7PgBMeACkQE/5e3qd8Hr20gs7df0O1UCgD
ST3oZmF0B2MwtBU4IDNRMmrDlx+ZvfE50guQd1jPSSeNrHyycgILnzLMjwMIEaRf
P34C7KkB1WrJ+IUNL7Wsp4WxEaL8whKzJnaBNMPCDHz4tUuanBtDSuSu2oWYLTNx
58KYKSxTbOGbBt1cRf10CCR1/3YGE9C7rg==
-----END CERTIFICATE-----`},
	"http://crt.usertrust.com/AddTrustExternalCARoot.p7c": {
		// subject=/C=SE/O=AddTrust AB/OU=AddTrust External TTP Network/CN=AddTrust External CA Root
		// issuer=/C=SE/O=AddTrust AB/OU=AddTrust External TTP Network/CN=AddTrust External CA Root
		`-----BEGIN CERTIFICATE-----
MIIENjCCAx6gAwIBAgIBATANBgkqhkiG9w0BAQUFADBvMQswCQYDVQQGEwJTRTEU
MBIGA1UEChMLQWRkVHJ1c3QgQUIxJjAkBgNVBAsTHUFkZFRydXN0IEV4dGVybmFs
IFRUUCBOZXR3b3JrMSIwIAYDVQQDExlBZGRUcnVzdCBFeHRlcm5hbCBDQSBSb290
MB4XDTAwMDUzMDEwNDgzOFoXDTIwMDUzMDEwNDgzOFowbzELMAkGA1UEBhMCU0Ux
FDASBgNVBAoTC0FkZFRydXN0IEFCMSYwJAYDVQQLEx1BZGRUcnVzdCBFeHRlcm5h
bCBUVFAgTmV0d29yazEiMCAGA1UEAxMZQWRkVHJ1c3QgRXh0ZXJuYWwgQ0EgUm9v
dDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALf3GjPm8gAELTngTlvt
H7xsD821+iO2zt6bETOXpClMfZOfvUq8k+0DGuOPz+VtUFrWlymUWoCwSXrbLpX9
uMq/NzgtHj6RQa1wVsfwTz/oMp50ysiQVOnGXw94nZpAPA6sYapeFI+eh6FqUNzX
mk6vBbOmcZSccbNQYArHE504B4YCqOmoaSYYkKtMsE8jqzpPhNjfzp/haW+710LX
a0Tkx63ubUFfclpxCDezeWWkWaCUN/cALw3CknLa0Dhy2xSoRcRdKn23tNbE7qzN
E0S3ySvdQwAl+mG5aWpYIxG3pzOPVnVZ9c0p10a3CitlttNCbxWyuHv77+ldU9U0
WicCAwEAAaOB3DCB2TAdBgNVHQ4EFgQUrb2YejS0Jvf6xCZU7wO94CTLVBowCwYD
VR0PBAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wgZkGA1UdIwSBkTCBjoAUrb2YejS0
Jvf6xCZU7wO94CTLVBqhc6RxMG8xCzAJBgNVBAYTAlNFMRQwEgYDVQQKEwtBZGRU
cnVzdCBBQjEmMCQGA1UECxMdQWRkVHJ1c3QgRXh0ZXJuYWwgVFRQIE5ldHdvcmsx
IjAgBgNVBAMTGUFkZFRydXN0IEV4dGVybmFsIENBIFJvb3SCAQEwDQYJKoZIhvcN
AQEFBQADggEBALCb4IUlwtYj4g+WBpKdQZic2YR5gdkeWxQHIzZlj7DYd7usQWxH
YINRsPkyPef89iYTx4AWpb9a/IfPeHmJIZriTAcKhjW88t5RxNKWt9x+Tu5w/Rw5
6wwCURQtjr0W4MHfRnXnJK3s9EK0hZNwEGe6nQY1ShjTK3rMUUKhemPR5ruhxSvC
Nr4TDea9Y355e6cJDUCrat2PisP29owaQgVR1EX1n6diIWgVIEM8med8vSTYqZEX
c4g/VhsxOBi0cQ+azcgOno4uG+GMmIPLHzHxREzGBHNJdmAPx/i9F4BrLunMTA5a
mnkPIAou1Z5jJh5VkpTYghdae9C8x49OhgQ=
-----END CERTIFICATE-----`,
		// subject=/C=SE/O=AddTrust AB/OU=AddTrust External TTP Network/CN=AddTrust External CA Root
		// issuer=/C=US/ST=UT/L=Salt Lake City/O=The USERTRUST Network/OU=http://www.usertrust.com/CN=UTN - DATACorp SGC
		`-----BEGIN CERTIFICATE-----
MIIEezCCA2OgAwIBAgIQftGpq77jb0bNa04pNJBW8zANBgkqhkiG9w0BAQUFADCB
kzELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAlVUMRcwFQYDVQQHEw5TYWx0IExha2Ug
Q2l0eTEeMBwGA1UEChMVVGhlIFVTRVJUUlVTVCBOZXR3b3JrMSEwHwYDVQQLExho
dHRwOi8vd3d3LnVzZXJ0cnVzdC5jb20xGzAZBgNVBAMTElVUTiAtIERBVEFDb3Jw
IFNHQzAeFw05OTA2MjQxODU3MjFaFw0xOTA2MjQxOTA2MzBaMG8xCzAJBgNVBAYT
AlNFMRQwEgYDVQQKEwtBZGRUcnVzdCBBQjEmMCQGA1UECxMdQWRkVHJ1c3QgRXh0
ZXJuYWwgVFRQIE5ldHdvcmsxIjAgBgNVBAMTGUFkZFRydXN0IEV4dGVybmFsIENB
IFJvb3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC39xoz5vIABC05
4E5b7R+8bA/Ntfojts7emxEzl6QpTH2Tn71KvJPtAxrjj8/lbVBa1pcplFqAsEl6
2y6V/bjKvzc4LR4+kUGtcFbH8E8/6DKedMrIkFTpxl8PeJ2aQDwOrGGqXhSPnoeh
alDc15pOrwWzpnGUnHGzUGAKxxOdOAeGAqjpqGkmGJCrTLBPI6s6T4TY386f4Wlv
u9dC12tE5Met7m1BX3JacQg3s3llpFmglDf3AC8NwpJy2tA4ctsUqEXEXSp9t7TW
xO6szRNEt8kr3UMAJfphuWlqWCMRt6czj1Z1WfXNKddGtworZbbTQm8Vsrh7++/p
XVPVNFonAgMBAAGjge0wgeowHwYDVR0jBBgwFoAUUzLRs89/+uDxoF2FTpLSnkUd
tE8wHQYDVR0OBBYEFK29mHo0tCb3+sQmVO8DveAky1QaMA4GA1UdDwEB/wQEAwIB
BjAPBgNVHRMBAf8EBTADAQH/MBEGA1UdIAQKMAgwBgYEVR0gADA9BgNVHR8ENjA0
MDKgMKAuhixodHRwOi8vY3JsLnVzZXJ0cnVzdC5jb20vVVROLURBVEFDb3JwU0dD
LmNybDA1BggrBgEFBQcBAQQpMCcwJQYIKwYBBQUHMAGGGWh0dHA6Ly9vY3NwLnVz
ZXJ0cnVzdC5jb20wDQYJKoZIhvcNAQEFBQADggEBADwlhyhsmL2dQhxeHmQPVn+W
PPO582kaafSkCNQgTbHyYyfhnwFDN7CxeudxyHoh7qg1wZ3mvGizRoCaPQRyPC9I
/eHMQncOsgU5pAD4NcKseMD9xxO8iyBNWjWvlMoysMZ50ZguO8JSRcGbtyYLywQa
9m6SROF8nMESeKYZAeLvYPt6V/MyKAa1uh2RGyhdZGpfU5wO1erMRb19RguvU0nG
zIAYW1utsWITYE45WVHEpobL8Q1t3t0xC1+jB6D7PkaqSXMEfYoLsC9GYo7hvVBl
KLHIdkr0IgMMVdT8DIdWfgtl74frfPclt80nTNs8CSlpF46LsEfo2mC3p2lm+ws=
-----END CERTIFICATE-----`},
	"http://grca.nat.gov.tw/repository/Certs/IssuedToThisCA.p7b": {
		// subject=/C=TW/O=Government Root Certification Authority
		// issuer=/C=TW/O=Government Root Certification Authority
		`-----BEGIN CERTIFICATE-----
MIIGSTCCBDGgAwIBAgIQMlyJOyY4kQwld2TzSNCtpTANBgkqhkiG9w0BAQsFADA/
MQswCQYDVQQGEwJUVzEwMC4GA1UECgwnR292ZXJubWVudCBSb290IENlcnRpZmlj
YXRpb24gQXV0aG9yaXR5MB4XDTEyMDkyODA5MDcxMloXDTMyMTIwNTEzMjMzM1ow
PzELMAkGA1UEBhMCVFcxMDAuBgNVBAoMJ0dvdmVybm1lbnQgUm9vdCBDZXJ0aWZp
Y2F0aW9uIEF1dGhvcml0eTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
AJoluOzMonWoe/fOW1mKydGGEghU7Jzy50b2iPN86aXfTEc2pBsBHH8eV4qNw8XR
IePaJD9IK/ufLqGU5ywck9G/GwGHU5nOp/UKIXZ3/6m3xnOUT0b3EEk3+qhZSV1q
gQdW8or5BtD3cCJNtLdBuTK4sfCxw5w/cP1T3YGq2GN49thTbqGsaoQkclSGxtKy
yhwOeYHWtXBiCAEuTk8O1RGvqa/lmr/czIdtJuTJV6L7lvnM4T9TjGxMfptTCAts
F/tnyMKtsc2AtJfcdgEWFelq16TheEfOhtX7MfP6Mb40qij7cEwdScevLJ1tZqa2
jWR+tSBqnTuBto9AAGdLiYa4zGX+FVPpBMHWXx1E1wovJ5pGfaENda1UhhXcSTvx
ls4Pm6Dso3pdvtUqdULle96ltqqvKKyskKw4t9VoNSZ63Pc78/1Fm9G7Q3hub/FC
VGqY8A2tl+lSXunVanLeavcbYBT0peS2cWeqH+riTcFCQP5nRhc4L0c/cZyu5SHK
YS1tB6iEfC3uUSXxY5Ce/eFXiGvviiNtsea9P63RPZYLhY3Naye7twWb7LuRqQoH
EgKXTiCQ8P8NHuJBO9NAOueNXdpm5AKwB1KYXA6OM5zCppX7VRluTI6uSw+9wThN
Xo+EHWbNxWCWtFJaBYmOlXqYwZE8lSOyDvR5tMl8wUohAgMBAAGjggE/MIIBOzAf
BgNVHSMEGDAWgBTVZx3gnHosnMvFmOcdByYqhux0zTAdBgNVHQ4EFgQUzMzvzClg
pDuxkrY8+jJij6wlFTswDgYDVR0PAQH/BAQDAgEGMEAGA1UdIAQ5MDcwCQYHYIZ2
ZQADATAJBgdghnZlAAMCMAkGB2CGdmUAAwMwCQYHYIZ2ZQADBDAJBgdghnZlAAMA
MA8GA1UdEwEB/wQFMAMBAf8wPgYDVR0fBDcwNTAzoDGgL4YtaHR0cDovL2dyY2Eu
bmF0Lmdvdi50dy9yZXBvc2l0b3J5L0NSTDIvQ0EuY3JsMFYGCCsGAQUFBwEBBEow
SDBGBggrBgEFBQcwAoY6aHR0cDovL2dyY2EubmF0Lmdvdi50dy9yZXBvc2l0b3J5
L0NlcnRzL0lzc3VlZFRvVGhpc0NBLnA3YjANBgkqhkiG9w0BAQsFAAOCAgEAQYW8
MPfAEZJTO5RgynxIFZVVN1cQCFU6/yF0WS66bEXKVWhz42TXQ9+vXX4R2CPyo1Xx
5Qx+kzOK4jb6LUuAqOYHw5R2QpXox5qjraCoAg9r+cFA3SrzhBe7Mhx+ktCaDaAS
++wxSUJm7Gu8S87grPQT1GKxy7wnCbtOmqmtixhXFu98tAcb5JtWoexD23DdKHlH
tv4Ptn7qhIMd3RflM/fXx3UuiwhtCFWDda7PobuLDXOC9zn96R6Q43EpSgZAq6cU
tKrUc3YB0mRmyRzekCNBcLtgthvK6gFTHBNIotZAkzCEWd6GUVP0V96qxcNP63HB
GmSd5t5INJ4K//iJUVwqnM8Qiugs7eSu2ov9E2korPeUx2IwW7yJogC4GKsvOqY7
ijyWzfq+V9UeHaKBOyVH2BNUFd8yD1nxfUU7pPXTXH/Wlwi3Riuj1N3GegrbmiKN
ILXwLBvwVDyDyksSaOUFLc8ELdzorAFC4/wpV8SFjz0cDg0SQ8KTfrhqnE/49T32
ePMmkuPG6RscGUmt/CHIBELrVuUg9fBsqHhXsEcLwPxpCa3zoz+65l3xC96C7dOs
nILgAleXRJeNnNWztqPMjsmK3bskVfr7JZ/mUfTlkujH33gYl49q+04CcYgeP1zq
Yu1iuTWLXNZzI2QG/3rs8Q9ZYXwyJReCoxC5d+I=
-----END CERTIFICATE-----`,
		// subject=/C=TW/O=Government Root Certification Authority
		// issuer=/C=TW/O=Government Root Certification Authority
		`-----BEGIN CERTIFICATE-----
MIIGSTCCBDGgAwIBAgIRAP7mJmeBQUNcX66p8ttLpm8wDQYJKoZIhvcNAQEFBQAw
PzELMAkGA1UEBhMCVFcxMDAuBgNVBAoMJ0dvdmVybm1lbnQgUm9vdCBDZXJ0aWZp
Y2F0aW9uIEF1dGhvcml0eTAeFw0xMjA5MjgwOTEzMjlaFw0zMjEyMDUxMzIzMzNa
MD8xCzAJBgNVBAYTAlRXMTAwLgYDVQQKDCdHb3Zlcm5tZW50IFJvb3QgQ2VydGlm
aWNhdGlvbiBBdXRob3JpdHkwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoIC
AQC2/5c8gb4BWCQnr44BK9ZykjAyG1+bfNTUf+ihYHMwVxAA+lCWJP5Q5ow6ldFX
eYTVZ1MMKoI+GFy4MCYa1l7GLbIEUQ7v3wxjR+vEEghRK5lxXtVpe+FdyXcdIOxW
juVhYC386RyA3/pqg7sFtR4jEpyCygrzFB0g5AaPQySZn7YKk1pzGxY5vgW28Yyl
ZJKPBeRcdvc5w88tvQ7Yy6gOMZvJRg9nU0MEj8iyyIOAX7ryD6uBNaIgIZfOD4k0
eA/PH07p+4woPN405+2f0mb1xcoxeNLOUNFggmOd4Ez3B66DNJ1JSUPUfr0t4urH
cWWACOQ2nnlwCjyHKenkkpTqBpIpJ3jmrdc96QoLXvTg1oadLXLLi2RW5vSueKWg
OTNYPNyoj420ai39iHPplVBzBN8RiD5C1gJ0+yzEb7xs1uCAb9GGpTJXA9ZN9E4K
mSJ2fkpAgvjJ5E7LUy3Hsbbi08J1J265DnGyNPy/HE7CPfg26QrMWJqhGIZO4uGq
s3NZbl6dtMIIr69c/aQCb/+4DbvVq9dunxpPkUDwH0ZVbaCSw4nNt7H/HLPLo5wK
4/7NqrwB7N1UypHdTxOHpPaY7/1J1lcqPKZc9mA3v9g+fk5oKiMyOr5u5CI9ByTP
isubXVGzMNJxbc5Gim18SjNE2hIvNkvy6fFRCW3bapcOFwIDAQABo4IBPjCCATow
HwYDVR0jBBgwFoAUzMzvzClgpDuxkrY8+jJij6wlFTswHQYDVR0OBBYEFNVnHeCc
eiycy8WY5x0HJiqG7HTNMA4GA1UdDwEB/wQEAwIBBjBABgNVHSAEOTA3MAkGB2CG
dmUAAwEwCQYHYIZ2ZQADAjAJBgdghnZlAAMDMAkGB2CGdmUAAwQwCQYHYIZ2ZQAD
ADAPBgNVHRMBAf8EBTADAQH/MD0GA1UdHwQ2MDQwMqAwoC6GLGh0dHA6Ly9ncmNh
Lm5hdC5nb3YudHcvcmVwb3NpdG9yeS9DUkwvQ0EuY3JsMFYGCCsGAQUFBwEBBEow
SDBGBggrBgEFBQcwAoY6aHR0cDovL2dyY2EubmF0Lmdvdi50dy9yZXBvc2l0b3J5
L0NlcnRzL0lzc3VlZFRvVGhpc0NBLnA3YjANBgkqhkiG9w0BAQUFAAOCAgEAY5IV
CK/CMXYyelsy8bSDPzGiiDG9ZcpXvJdDgV2gnU701Q7uJ52tGOk2pg0CL5WjChVz
Vwmk34jXaBKAQZRF7ruOC3cU2HYfx3IKk8z+CWmVu4PMTmlR5VyPiZoqTfWuiFUP
7p+krL/b97HWZN2G0ein+++/fcdJnTRyxKDBJIpiwXoNjTgN0QYfSRMFAATng3c7
clDtSDYiJlnM0iecu2f1xOWSXwpN9zZQ9KiBKsPS8D5WKBBwLvqq4pwxEHbWJael
UwOzdfW5+P8hzEeuZ7g3BoeoxuPPJACZjgvYUf6Lp1N9HitiwlBYwt1Sk/hggT4r
ykRCghvY2BN0J5Aago8WAtiPH393yLM0PUAPl260C6H4qJCfft+v1LPFRzimukO0
8jZDtJinIKjN4ZNiBO4/wpvpBpAsRZQkbwvMeQKqOhEduH+deDg1LgTidtWapK6D
H4OADQnWfsH96MWrA2OQAU/3n7SGuwDsT3I8oYwXCZ4Za0FMJIcftZuA8soU7bHo
Tvmiar3DZrvPZE6uq0dHboxVt/4Qsogv+3PMRkqV6X8lk18hzkClEvToQh4xUW2R
wnXMUCSjca4A59fi12K6chOo3hv0gIe9OQAkSGWrlAOfCERTio8fW8dC+/or/ZgX
ha+uQ1DoDj7b4KImmT4M6idBYze1/LoKcnizOQs=
-----END CERTIFICATE-----`},
}

func urlReplacement(url string) []*x509.Certificate {
	cs, ok := replacements[url]
	if !ok {
		return nil
	}
	var r []*x509.Certificate
	for _, c := range cs {
		s, _ := pem.Decode([]byte(c))
		if s == nil {
			log.Fatalf("Can't decode built-in: %s", c)
			return nil
		}
		cert, err := x509.ParseCertificate(s.Bytes)
		if err != nil {
			log.Fatalf("Can't parse built-in: %s\n%s", c, err)
			return nil
		}
		if cert == nil {
			log.Fatalf("Parse didn't produce a cert: %s", c)
			return nil
		}
		r = append(r, cert)
	}
	return r
}
