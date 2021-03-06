package bandwidth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"httphelper"
	"response"
)

func QueryBandwidth(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	uri := "https://cph.cn-east-3.myhuaweicloud.com/v1/09402bad5e80f3902fc1c0188cab3cd5/cloud-phone/bandwidths"

	body, err := httphelper.HttpGet(uri)
	if err != nil {
		log.Println("QueryBandwidth err: ", err)
		resp.IntervalServErr(w)
		return
	}
	fmt.Println("test UpdateBandwidth: ", string(body))

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
