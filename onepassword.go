package onepassword

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"sync"
)

type vaultName string
type itemName string
type documentName string
type documentValue string
type sectionName string
type fieldName string
type fieldValue string
type itemResponse []byte
type fieldMap map[fieldName]fieldValue
type sectionMap map[sectionName]fieldMap

// Client : 1Password client
type Client struct {
	OpPath    string
	Subdomain string
	Email     string
	Password  string
	SecretKey string
	Session   string
	mutex     *sync.Mutex
}

type parsedItem struct {
	UUID    string `json:"uuid"`
	Details struct {
		Sections []struct {
			Title  string `json:"title"`
			Fields []struct {
				Key   string `json:"t"`
				Value string `json:"v"`
			} `json:"fields"`
		} `json:"sections"`
	} `json:"details"`
}

func getArg(key string, value string) string {
	return fmt.Sprintf("--%s=%s", key, value)
}

func (op Client) runCmd(args ...string) ([]byte, error) {
	sessionArg := getArg("session", op.Session)
	args = append(args, sessionArg)
	debugCmd := fmt.Sprintf("op %s", strings.Join(args, " "))
	op.mutex.Lock()
	cmd := exec.Command(string(op.OpPath), args...)
	defer op.mutex.Unlock()
	res, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("error calling 1Password: %s\n%s\n'%s'", err, res, debugCmd)
	}
	return res, err
}

// Calls the `op signin` command and passes in the password via stdin.
// usage: op signin <signinaddress> <emailaddress> <secretkey> [--output=raw]
func (op *Client) authenticate() error {
	signinAddress := fmt.Sprintf("%s.1password.com", op.Subdomain)
	op.mutex.Lock()
	cmd := exec.Command(op.OpPath, "signin", signinAddress, op.Email, op.SecretKey, "--output=raw")
	defer op.mutex.Unlock()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("Cannot attach to stdin: %s", err)
	}
	go func() {
		defer stdin.Close()
		if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", op.Password)); err != nil {
			log.Println("[Error]", err)
		}
	}()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Cannot signin: %s\n%s", err, output)
	}
	op.Session = strings.Trim(string(output), "\n")
	return nil
}

func (res itemResponse) parseResponse() (sectionMap, error) {
	sm := make(sectionMap)
	pItem := parsedItem{}
	if err := json.Unmarshal(res, &pItem); err != nil {
		return sm, err
	}
	for _, section := range pItem.Details.Sections {
		fm := make(fieldMap)
		for _, field := range section.Fields {
			fm[fieldName(field.Key)] = fieldValue(field.Value)
		}
		sm[sectionName(section.Title)] = fm
	}
	return sm, nil
}

// Calls `op get item` command.
// usage: op get item <item> [--vault=<vault>] [--session=<session>]
func (op Client) getItem(vault vaultName, item itemName) (itemResponse, error) {
	vaultArg := getArg("vault", string(vault))
	res, err := op.runCmd("get", "item", string(item), vaultArg)
	if err != nil {
		return itemResponse(""), fmt.Errorf("error getting item: %s", err)
	}
	return itemResponse(res), nil
}

// Calls `op get document` command
// usage: op get document <document> [--vault=<vault>] [--session=<session>]
func (op Client) getDocument(vault vaultName, docName documentName) (documentValue, error) {
	vaultArg := getArg("vault", string(vault))
	res, err := op.runCmd("get", "document", string(docName), vaultArg)
	if err != nil {
		err = fmt.Errorf("error getting document: %s", err)
		return documentValue(""), err
	}
	return documentValue(res), nil
}
