package main

import (
	"github.com/volatile/i18n"
	"golang.org/x/text/language"
)

var locales = i18n.Locales{

	language.English: {
		"decimalMark":   ".",
		"thousandsMark": ",",
	},

	language.French: {
		"decimalMark":   ",",
		"thousandsMark": " ",
	},
}
