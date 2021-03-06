// Copyright 2018. bolaxy.org authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package errors

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ErrModel struct {
	Code CodeType    //结合业务的错误code
	Err  error       //error详情
	Data interface{} //不同场景下需要返回的数据
}

func (e *ErrModel) RetErr(ctx *gin.Context) {
	if e.Data == nil {
		e.Data = "{}"
	}

	if e.Code == 0 && e.Err == nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": e.Code, "Data": e.Data, "Msg": "success"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Code": e.Code, "Data": e.Data, "Msg": e.Err.Error()})
	}
}

type ErrJsonModel struct {
	Code int
	Msg  string
	Data interface{}
}

func New(text string) error {
	return errors.New(text)
}
