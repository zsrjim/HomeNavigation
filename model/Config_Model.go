// -*- coding: utf-8 -*-
// @Time    : 2021/11/15 17:10
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : Config_Model.go

package model

// Config 配置项
type Config struct {
	Index       conf
	SoftWare    software
	Data        []Data
	Footer      []footer
	FooterStyle style
}

type conf struct {
	Logo    string `json:"logo"`
	Title   string `json:"title"`
	Favicon string `json:"favicon"`
}

type software struct {
	Port     int    `json:"port"`
	Password string `json:"password"`
	Mode     string `json:"mode"`
	Debug    string `json:"Debug"`
	LogLeveL string `json:"LogLevel"`
}

type Data struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Ico   string `json:"ico"`
	WLink string `json:"w_link"`
	NLink string `json:"n_link"`
}

type footer struct {
	Text     string `json:"Text"`
	TextLink string `json:"TextLink"`
}

type style struct {
	Background       string `json:"background"`
	BackgroundColorA string `json:"BackgroundColorA"`
	BackgroundColorB string `json:"BackgroundColorB"`
	LColor           string `json:"LColor"`
	SColor           string `json:"SColor"`
	FColor           string `json:"FColor"`
}
