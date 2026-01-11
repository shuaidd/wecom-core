package meeting

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMeetingRoomsRequest_Marshal(t *testing.T) {
	req := ListMeetingRoomsRequest{
		MeetingRoomName: "ROOMNAME",
		Cursor:          "CURSOR",
		Limit:           50,
	}

	data, err := json.Marshal(req)
	assert.NoError(t, err)

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	assert.NoError(t, err)

	assert.Equal(t, "ROOMNAME", m["meeting_room_name"])
	assert.Equal(t, "CURSOR", m["cursor"])
	v, ok := m["limit"]
	if !ok {
		t.Fatalf("limit not found in marshaled JSON")
	}
	// JSON numbers are unmarshaled as float64
	assert.Equal(t, float64(50), v)
}

func TestGetRoomInfoResponse_Unmarshal(t *testing.T) {
	js := `{"errcode":0,"errmsg":"ok","basic_info":{"rooms_id_list":["200115200039985708"],"meeting_room_name":"会议室测试1","city":"广州","building":"大厦","floor":"10","participant_number":3,"device":"ROOMS","desc":"aGVsbG8=","password":"MzMz"},"account_info":{"account_type":0,"valid_period":"-"},"hardware_info":{"ip":"10.10.10.69","mac":"a1:ee:27:c1:8a:1a","rooms_version":"2.7.2.420","health_status":"ERROR","system_type":"10.13.6","meeting_room_status":1,"active_time":"2021-03-23 15:37:34","camera_model":"FaceTime 高清摄像头（内建）","enable_video_mirror":true,"microphone_info":"内建麦克风 ","speaker_info":"内建输出"},"pmi_info":{"pmi_code":"12345678","pmi_pwd":"XXXXXXX"},"monitor_status":0,"is_allow_call":true,"scheduled_status":1}`

	var resp GetRoomInfoResponse
	err := json.Unmarshal([]byte(js), &resp)
	assert.NoError(t, err)

	if resp.BasicInfo == nil {
		t.Fatalf("basic_info is nil")
	}
	assert.Equal(t, "会议室测试1", resp.BasicInfo.MeetingRoomName)

	if resp.HardwareInfo == nil {
		t.Fatalf("hardware_info is nil")
	}
	assert.Equal(t, "10.10.10.69", resp.HardwareInfo.IP)

	assert.Equal(t, true, resp.IsAllowCall)
	assert.Equal(t, 1, resp.ScheduledStatus)
}

func TestCallRoomResponse_Unmarshal(t *testing.T) {
	js := `{"errcode":0,"errmsg":"ok","invite_id":"INVITEID"}`
	var resp CallRoomResponse
	err := json.Unmarshal([]byte(js), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "INVITEID", resp.InviteID)
}

func TestGetResponseStatusResponse_Unmarshal(t *testing.T) {
	js := `{"errcode":0,"errmsg":"ok","status":2,"response_time":"2022/11/22 14:35:26"}`
	var resp GetResponseStatusResponse
	err := json.Unmarshal([]byte(js), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 2, resp.Status)
	assert.Equal(t, "2022/11/22 14:35:26", resp.ResponseTime)
}

func TestListMeetingsResponse_Roundtrip(t *testing.T) {
	orig := ListMeetingsResponse{
		HasMore:    false,
		NextCursor: "CURSOR",
		MeetingInfoList: []MeetingInfo{
			{
				MeetingID:   "100001",
				MeetingCode: "14512",
				Subject:     "周会",
				Status:      "MEETING_STATE_STARTED",
				MeetingType: 0,
				StartTime:   1679313600,
				EndTime:     1679317200,
			},
		},
	}

	data, err := json.Marshal(orig)
	assert.NoError(t, err)

	var out ListMeetingsResponse
	err = json.Unmarshal(data, &out)
	assert.NoError(t, err)

	if len(out.MeetingInfoList) != 1 {
		t.Fatalf("unexpected meeting info list length: %d", len(out.MeetingInfoList))
	}

	assert.Equal(t, orig.MeetingInfoList[0].Subject, out.MeetingInfoList[0].Subject)
	assert.Equal(t, orig.MeetingInfoList[0].MeetingID, out.MeetingInfoList[0].MeetingID)
	assert.Equal(t, orig.NextCursor, out.NextCursor)
}
