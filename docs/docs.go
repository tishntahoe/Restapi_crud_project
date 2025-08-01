// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add": {
            "post": {
                "description": "Добавляет человека по имени, отчеству, фамилии и email(не обязательно, идет как дополнение), получая возраст, пол и нацию из внешних API",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Добавить нового человека",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Отчество",
                        "name": "thirdname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Человек с именем ... был добавлен",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/addemail": {
            "post": {
                "description": "Добавляет email к человеку по имени и отчеству",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Привязать email к человеку",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Емейл: ... присоединен к человеку: ...",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/addwithparams": {
            "post": {
                "description": "Добавляет человека, используя параметры, переданные напрямую (без внешних API)",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Добавить нового человека с ручными параметрами",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Отчество",
                        "name": "thirdname",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Возраст",
                        "name": "age",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Национальность (напр. US, RU)",
                        "name": "nation",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Пол (male или female, но другой тоже подойдет:) )",
                        "name": "gender",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Человек с именем ... был добавлен с собственными параметрами",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/changeUserInfo": {
            "post": {
                "description": "Обновляет информацию о человеке по имени, отчеству и фамилии",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Обновить информацию о пользователе",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Отчество",
                        "name": "thirdname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Пол",
                        "name": "gender",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Возраст",
                        "name": "age",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Национальность",
                        "name": "nation",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о человеке ... была обновлена!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при обновлении",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/checkfriendship": {
            "post": {
                "description": "Возвращает список друзей пользователя по имени и фамилии",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "friendship"
                ],
                "summary": "Проверить друзей пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя пользователя",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия пользователя",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Список друзей пользователя",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при получении друзей",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/createfriendship": {
            "post": {
                "description": "Создаёт дружбу между пользователями по имени и фамилии",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "friendship"
                ],
                "summary": "Создать дружбу между двумя пользователями",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя первого пользователя",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия первого пользователя",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя второго пользователя",
                        "name": "nameFr",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Фамилия второго пользователя",
                        "name": "secnameFr",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Дружба между ... создана!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при создании дружбы",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getallusers": {
            "get": {
                "description": "Возвращает массив всех людей из базы",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Получить всех пользователей",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AddingHumanStruct"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка при получении пользователей",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getbysecname": {
            "post": {
                "description": "Возвращает структуру человека по фамилии (secname), если он найден",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Получить данные о человеке по фамилии",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Фамилия человека",
                        "name": "secname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/models.AddingHumanStruct"
                        }
                    },
                    "400": {
                        "description": "Ошибка или человек не найден",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddingHumanStruct": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nation": {
                    "type": "string"
                },
                "secname": {
                    "type": "string"
                },
                "thirdname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "redsoft_TEST",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
