package main

import "github.com/volatile/i18n"

var locales = i18n.Locales{
	"en": i18n.Locale{
		"decimalMark":   ".",
		"thousandsMark": ",",
	},
	"fr": i18n.Locale{
		"decimalMark":   ",",
		"thousandsMark": " ",
	},
}
