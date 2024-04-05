package gcpsp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOldContexts_UnmarshalJSON(t *testing.T) {
	rawContext := `[
						{
							"hash": "https://schema.org/Text",
							"provider": "https://schema.org/Text"
						}
					]`

	ctx := new(Contexts)

	err := json.Unmarshal([]byte(rawContext), ctx)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(*ctx))
	assert.Equal(t, "https://schema.org/Text", (*ctx)[0].Hash)
	assert.Equal(t, "https://schema.org/Text", (*ctx)[0].Provider)

	rawStamp := `{
			"version": "1.0.0",
			"credential": {
				"type": [
					"VerifiableCredential"
				],
				"proof": {
					"jws": "eyJfksjdfjksmFsc2V9..nZflskdfLhBfgN",
					"type": "Ed25519Signature2018",
					"created": "1888-11-15T14:58:55.479Z",
					"proofPurpose": "assertionMethod",
					"verificationMethod": "did:key:z6MkghvfdlskfsdlkfdsqmNU5EfC"
				},
				"issuer": "did:key:z6Maksjndaslkdm;alsd5LC",
				"@context": [
					"https://www.w3.org/2018/credentials/v1"
				],
				"issuanceDate": "1888-11-25T14:58:55.479Z",
				"expirationDate": "1888-01-14T14:58:55.479Z",
				"credentialSubject": {
					"id": "did:pkh:eip155:1:0x3465d06FFb7B08123deadbeef45678979CDb73bf",
					"hash": "v0.0.0:yLhChsdlBcF/NoLJEytqwikmcfxzhlo8Es=",
					"@context": [
						{
							"hash": "https://schema.org/Text",
							"provider": "https://schema.org/Text"
						}
					],
					"provider": "Google"
				}
			}
		}`

	var stamp Stamp

	err = json.Unmarshal([]byte(rawStamp), &stamp)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(stamp.Credential.CredentialSubject.Context))
	assert.Equal(t, "https://schema.org/Text", stamp.Credential.CredentialSubject.Context[0].Hash)
	assert.Equal(t, "https://schema.org/Text", stamp.Credential.CredentialSubject.Context[0].Provider)
}

func TestNewContexts_UnmarshalJSON(t *testing.T) {
	rawContext := `{
						"hash": "https://schema.org/Text",
						"provider": "https://schema.org/Text"
					}`

	ctx := new(Contexts)

	err := json.Unmarshal([]byte(rawContext), ctx)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(*ctx))
	assert.Equal(t, "https://schema.org/Text", (*ctx)[0].Hash)
	assert.Equal(t, "https://schema.org/Text", (*ctx)[0].Provider)

	rawStamp := `{
			"version": "1.0.0",
			"credential": {
				"type": [
					"VerifiableCredential"
				],
				"proof": {
					"type": "EthereumEip712Signature2021",
					"created": "1888-01-01T15:39:55.401Z",
					"@context": "https://w3id.org/security/suites/eip712sig-2021/v1",
					"proofValue": "0x673d192fdeadbeef428a77fa18936fabcdef09876f5435321af60660e1ba",
					"eip712Domain": {
						"types": {
							"Proof": [
								{
									"name": "@context",
									"type": "string"
								},
								{
									"name": "created",
									"type": "string"
								},
								{
									"name": "proofPurpose",
									"type": "string"
								},
								{
									"name": "type",
									"type": "string"
								},
								{
									"name": "verificationMethod",
									"type": "string"
								}
							],
							"@context": [
								{
									"name": "hash",
									"type": "string"
								},
								{
									"name": "provider",
									"type": "string"
								}
							],
							"Document": [
								{
									"name": "@context",
									"type": "string[]"
								},
								{
									"name": "credentialSubject",
									"type": "CredentialSubject"
								},
								{
									"name": "expirationDate",
									"type": "string"
								},
								{
									"name": "issuanceDate",
									"type": "string"
								},
								{
									"name": "issuer",
									"type": "string"
								},
								{
									"name": "proof",
									"type": "Proof"
								},
								{
									"name": "type",
									"type": "string[]"
								}
							],
							"EIP712Domain": [
								{
									"name": "name",
									"type": "string"
								}
							],
							"CredentialSubject": [
								{
									"name": "@context",
									"type": "@context"
								},
								{
									"name": "hash",
									"type": "string"
								},
								{
									"name": "id",
									"type": "string"
								},
								{
									"name": "provider",
									"type": "string"
								}
							]
						},
						"domain": {
							"name": "VerifiableCredential"
						},
						"primaryType": "Document"
					},
					"proofPurpose": "assertionMethod",
					"verificationMethod": "did:ethr:0xa123d6f8d6c8f0987654deadbeef38577e5fb123#controller"
				},
				"issuer": "did:ethr:0xfd46f8d6ca861deadbeef234567891bd8577e5fb",
				"@context": [
					"https://www.w3.org/2018/credentials/v1",
					"https://w3id.org/vc/status-list/2021/v1"
				],
				"issuanceDate": "1888-01-01T15:39:55.400Z",
				"expirationDate": "1777-01-01T15:39:55.400Z",
				"credentialSubject": {
					"id": "did:pkh:eip155:1:0xb12345678900987654321deadbeefaCa9B892f36",
					"hash": "v0.0.0:0PS/3LDUodqidhbqwEyrpvcdv13WNxBU=",
					"@context": {
						"hash": "https://schema.org/Text",
						"provider": "https://schema.org/Text"
					},
					"provider": "Discord"
				}
			}
		}`

	var stamp Stamp

	err = json.Unmarshal([]byte(rawStamp), &stamp)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(stamp.Credential.CredentialSubject.Context))
	assert.Equal(t, "https://schema.org/Text", stamp.Credential.CredentialSubject.Context[0].Hash)
	assert.Equal(t, "https://schema.org/Text", stamp.Credential.CredentialSubject.Context[0].Provider)
}
