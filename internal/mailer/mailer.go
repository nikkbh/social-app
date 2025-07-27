package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"text/template"
	"time"
)

const (
	FromName            = "GopherSocial"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}

func templateParser(templateFile, path string, data any) (*bytes.Buffer, string, error) {
	tmpl, err := template.ParseFS(FS, path+templateFile)
	if err != nil {
		return nil, "", err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return nil, "", err
	}
	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return nil, "", err
	}

	return body, subject.String(), nil
}

// retry executes the provided function up to maxRetries times with exponential backoff.
// It returns the result of the function or the last error encountered.
func retry(maxRetries int, fn func() (int, error)) (int, error) {
	for i := 0; i < maxRetries; i++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}
		log.Printf("Attempt %d/%d failed: %v", i+1, maxRetries, err)
		time.Sleep(time.Second * time.Duration(i+1)) // exponential backoff
	}
	return -1, fmt.Errorf("all %d attempts failed", maxRetries)
}
