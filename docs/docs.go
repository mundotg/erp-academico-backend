package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "swagger": "2.0",
  "info": {
    "description": "API REST para gestão acadêmica, recursos humanos e financeiro.",
    "title": "ERP Acadêmico API",
    "contact": {"name": "Suporte API", "email": "suporte@example.com"},
    "license": {"name": "MIT"},
    "version": "1.0"
  },
  "basePath": "/api/v1",
  "paths": {}
}`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "ERP Acadêmico API",
	Description:      "API REST para gestão acadêmica, recursos humanos e financeiro.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
