package alipay

import (
	"encoding/json"
	"fmt"

	"github.com/yuanqinguo/gopay"
)

// alipay.open.app.qrcode.create(小程序生成推广二维码接口)
//	文档地址：https://opendocs.alipay.com/apis/009zva
func (a *Client) OpenAppQrcodeCreate(bm gopay.BodyMap) (aliRsp *OpenAppQrcodeCreateRsp, err error) {
	err = bm.CheckEmptyError("url_param", "query_param", "describe")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(bm, "alipay.open.app.qrcode.create"); err != nil {
		return nil, err
	}
	aliRsp = new(OpenAppQrcodeCreateRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
