package config

import (
	"log"
	"html/template"

	"example.com/m/v2/internal/models"
	"github.com/alexedwards/scs/v2"
)


type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
