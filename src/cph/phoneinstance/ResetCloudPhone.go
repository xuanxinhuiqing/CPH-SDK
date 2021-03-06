package phoneinstance

import (
	"encoding/json"
	"fmt"
	"global"
	"log"
	"net/http"

	"httphelper"
	"response"
)

type rcpBody struct {
	Phones []struct {
		PhoneID  string `json:"phone_id"`
		Property string `json:"property"`
	} `json:"phones"`
}

func ResetCloudPhone(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResp()

	var projectId string // 必填，项目ID
	var rcp rcpBody
	r.ParseForm()
	if len(r.Form.Get("projectId")) > 0 {
		projectId = r.Form.Get("projectId")
	} else {
		resp.BadReq(w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&rcp)
	if err != nil {
		resp.BadReq(w)
		return
	}

	if len(rcp.Phones) == 0 {
		resp.BadReq(w)
		return
	}

	data, _ := json.Marshal(rcp)
	uri := fmt.Sprintf("%s/%s/cloud-phone/phones/batch-reset", global.BaseUrl, projectId)
	body, err := httphelper.HttpPost(uri, data)
	if err != nil {
		log.Println("ResetCloudPhone err: ", err)
		resp.IntervalServErr(w)
		return
	}

	json.Unmarshal(body, &resp.Data)
	resp.WriteTo(w)
}
