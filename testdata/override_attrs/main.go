package main

import (
	"time"

	"github.com/swaggo/swag/testdata/override_attrs/common"
	"github.com/swaggo/swag/testdata/override_attrs/model"
)

type Local struct {
	L string `json:"l"`
}

type Payload struct {
	Str   common.Optional[string]         `json:"str"`
	I64   common.Optional[int64]          `json:"i64"`
	Tm    common.Optional[time.Time]      `json:"tm"`
	St    common.Optional[model.Status]   `json:"st"`
	In    common.Optional[model.Inner]    `json:"in"`
	Loc   common.Optional[Local]          `json:"loc"`
	Sl    common.Optional[[]string]       `json:"sl"`
	Ptr   *common.Optional[string]        `json:"ptr"`
	Arr   []common.Optional[string]       `json:"arr"`
	Mp    map[string]common.Optional[int] `json:"mp"`
	Foo   common.Foo                      `json:"foo"`
	Plain string                          `json:"plain"`
}

// @title Swagger Example API
// @version 1.0

// Handler
// @Param body body Payload true "payload"
// @Success 200 {object} Payload
// @Router /payload [post]
func Handler() {}

func main() {}
