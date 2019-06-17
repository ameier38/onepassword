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

var expectedItemMap = ItemMap{
	SectionName("Terraform"): FieldMap{
		FieldName("type"):               FieldValue("postgresql"),
		FieldName("server"):             FieldValue("redshift.company.io"),
		FieldName("port"):               FieldValue("5439"),
		FieldName("database"):           FieldValue("test-db"),
		FieldName("username"):           FieldValue("test-user"),
		FieldName("password"):           FieldValue("test-password"),
		FieldName("SID"):                FieldValue(""),
		FieldName("alias"):              FieldValue(""),
		FieldName("connection options"): FieldValue(""),
		FieldName("schema"):             FieldValue("development"),
	},
}

func TestParseItemResponse(t *testing.T) {
	actualItemMap, err := parseItemResponse([]byte(mockItemResponse))
	if assert.Nil(t, err) {
		assert.Equal(t, expectedItemMap, actualItemMap, "item maps should equal")
	}
}

func TestGetItem(t *testing.T) {
	opPath, err := buildMockOnePassword()
	if err != nil {
		t.Errorf("failed to build mock 1Password CLI: %s", err)
	}
	client, err := NewClient(opPath, "test-subdomain", "test@subdomain.com", "test-password", "test-secret-key")
	if err != nil {
		t.Errorf("failed to create Client: %s", err)
	}
	assert.Equal(t, "test-session", client.Session)
	actualItemMap, err := client.GetItem(VaultName("test-vault"), ItemName("test-item"))
	if err != nil {
		t.Errorf("error getting item: %s", err)
	}
	assert.Equal(t, expectedItemMap, actualItemMap)
}
