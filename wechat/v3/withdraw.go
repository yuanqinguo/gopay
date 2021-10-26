package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/yuanqinguo/gopay"
	"net/http"
)

// 二级商户余额提现API
//	Code = 0 is success
// 	电商收付通文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_2.shtml
func (c *ClientV3) V3CommerceFundWithDraw(bm gopay.BodyMap) (*CommerceFundWithdrawRsp, error) {
	authorization, err := c.authorization(MethodPost, v3Withdraw, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(bm, v3Withdraw, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &CommerceFundWithdrawRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(CommerceFundWithdraw)
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

// 二级商户查询提现状态API
// 基于微信支付提现单号查询 withdraw_id
//	Code = 0 is success
// 	电商收付通文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_3.shtml
func (c *ClientV3) V3CommerceFundWithdrawQuery(withdrawId string) (*CommerceFundWithdrawQueryRsp, error) {
	uri := fmt.Sprintf(v3WithdrawStatus, withdrawId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &CommerceFundWithdrawQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(CommerceFundWithdrawQuery)
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

// 二级商户查询提现状态API
// 基于商户提现单号查询 out_request_no
//	Code = 0 is success
// 	电商收付通文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_8_3.shtml
func (c *ClientV3) V3CommerceFundWithdrawOutRequestNoQuery(outRequestNo string) (*CommerceFundWithdrawQueryRsp, error) {
	uri := fmt.Sprintf(v3WithdrawOutRequestNoStatus, outRequestNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &CommerceFundWithdrawQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(CommerceFundWithdrawQuery)
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
