/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: publish.proto

package publish

import (
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/common"
	"mime/multipart"

	_ "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/api"
)

type DouyinPublishActionRequest struct {
	Token string                `json:"token,required" form:"token,required" query:"token,required"` // 用户鉴权token
	Data  *multipart.FileHeader `form:"data,required"`                                               // 视频数据
	Title string                `json:"title,required" form:"title,required" query:"title,required"` // 视频标题
}

type DouyinPublishActionResponse struct {
	StatusCode int32  `json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                     // 返回状态描述
}

type DouyinPublishListRequest struct {
	UserId int64  `json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  string `json:"token,required" form:"token,required" query:"token,required"`       // 用户鉴权token
}

type DouyinPublishListResponse struct {
	StatusCode int32          `json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string         `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                     // 返回状态描述
	VideoList  []common.Video `json:"video_list" form:"video_list" query:"video_list"`                               // 用户发布的视频列表
}
