package util

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func Axis(rowN, colN int) string {
	return excelize.ToAlphaString(colN) + strconv.Itoa(rowN+1)
}
