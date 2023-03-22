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

// Code generated by hertz generator.

package feed

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	feed "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/basic/feed"
	feed_service "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/service/feed"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/utils"
)

// Feed get a list of recommended videos
//
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, feed.DouyinFeedResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	resp, err := feed_service.NewFeedService(ctx, c).Feed(&req)
	if err != nil {
		bresp := utils.BuildBaseResp(err)
		resp.StatusCode = bresp.StatusCode
		resp.StatusMsg = bresp.StatusMsg
		c.JSON(consts.StatusOK, resp)
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.SuccessMsg

	c.JSON(consts.StatusOK, resp)
}
