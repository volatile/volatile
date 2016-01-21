package main

import (
	"github.com/volatile/i18n"
	"golang.org/x/text/language"
)

var locales = i18n.Locales{
	language.English: {
		"decimalMark":   ".",
		"thousandsMark": ",",

		"hello": "Welcome to your new app.",
	},
	language.French: {
		"decimalMark":   ",",
		"thousandsMark": " ",

		"hello": "Bienvenue sur votre nouvelle app.",
	},
}
