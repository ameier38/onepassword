package onepassword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockItemResponse = `
{
  "uuid": "test-item",
  "templateUuid": "102",
  "trashed": "N",
  "createdAt": "2019-05-18T14:58:54Z",
  "updatedAt": "2019-05-18T15:04:56Z",
  "itemVersion": 2,
  "vaultUuid": "test-vault",
  "details": {
    "fields": [],
    "notesPlain": "",
    "sections": [
      {
        "fields": [
          {
            "k": "menu",
            "n": "database_type",
            "t": "type",
            "v": "postgresql"
          },
          {
            "inputTraits": {
              "keyboard": "URL"
            },
            "k": "string",
            "n": "hostname",
            "t": "server",
            "v": "redshift.company.io"
          },
          {
            "inputTraits": {
              "keyboard": "NumberPad"
            },
            "k": "string",
            "n": "port",
            "t": "port",
            "v": "5439"
          },
          {
            "inputTraits": {
              "autocapitalization": "none",
              "autocorrection": "no"
            },
            "k": "string",
            "n": "database",
            "t": "database",
            "v": "test-db"
          },
          {
            "inputTraits": {
              "autocapitalization": "none",
              "autocorrection": "no"
            },
            "k": "string",
            "n": "username",
            "t": "username",
            "v": "test-user"
          },
          {
            "k": "concealed",
            "n": "password",
            "t": "password",
            "v": "test-password"
          },
          {
            "k": "string",
            "n": "sid",
            "t": "SID",
            "v": ""
          },
          {
            "k": "string",
            "n": "alias",
            "t": "alias",
            "v": ""
          },
          {
            "k": "string",
            "n": "options",
            "t": "connection options",
            "v": ""
          },
          {
            "k": "string",
            "n": "custom",
            "t": "schema",
            "v": "development"
          }
        ],
        "name": "",
        "title": "Terraform"
      }
    ]
  },
  "overview": {
    "URLs": [],
    "ainfo": "redshift.company.io",
    "pbe": 0,
    "pgrng": false,
    "ps": 0,
    "tags": [],
    "title": "Redshift",
    "url": ""
  }
}
`

func TestParse(t *testing.T) {
	res := itemResponse(mockItemResponse)
	expectedSectionMap := sectionMap{
		sectionName("Terraform"): fieldMap{
			fieldName("type"):               fieldValue("postgresql"),
			fieldName("server"):             fieldValue("redshift.company.io"),
			fieldName("port"):               fieldValue("5439"),
			fieldName("database"):           fieldValue("test-db"),
			fieldName("username"):           fieldValue("test-user"),
			fieldName("password"):           fieldValue("test-password"),
			fieldName("SID"):                fieldValue(""),
			fieldName("alias"):              fieldValue(""),
			fieldName("connection options"): fieldValue(""),
			fieldName("schema"):             fieldValue("development"),
		},
	}
	actualSectionMap, err := res.parseResponse()
	if assert.Nil(t, err) {
		assert.Equal(t, expectedSectionMap, actualSectionMap, "section maps should equal")
	}
}
