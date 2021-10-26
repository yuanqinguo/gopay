package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yuanqinguo/gopay"
)

// 申请退款API
//	Code = 0 is success
//	电商收付通申请退款API文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_6_1.shtml
func (c *ClientV3) V3CommerceRefund(bm gopay.BodyMap) (wxRsp *RefundRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3CommerceRefund, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3CommerceRefund, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundOrderResponse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询单笔退款API
//	Code = 0 is success
//	电商收付通查询退款API文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_6_1.shtml
func (c *ClientV3) V3CommerceRefundQuery(outRefundNo string) (wxRsp *RefundQueryRsp, err error) {
	uri := fmt.Sprintf(v3CommerceRefundQuery, outRefundNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundQueryResponse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 申请退款API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_9.shtml
func (c *ClientV3) V3Refund(bm gopay.BodyMap) (wxRsp *RefundRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3DomesticRefund, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3DomesticRefund, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundOrderResponse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询单笔退款API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_10.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_10.shtml
func (c *ClientV3) V3RefundQuery(outRefundNo string) (wxRsp *RefundQueryRsp, err error) {
	uri := fmt.Sprintf(v3DomesticRefundQuery, outRefundNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &RefundQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(RefundQueryResponse)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
