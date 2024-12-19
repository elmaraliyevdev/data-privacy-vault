package models

type Token struct {
    ID   string            `json:"id"`
    Data map[string]string `json:"data"`
}