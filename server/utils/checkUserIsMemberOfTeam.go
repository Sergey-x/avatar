package utils

import (
	"avatar.com/avatar/server/conf"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ReqRetries = 3

type isMemberStruct struct {
	TeamMember bool `json:"teamMember"`
}

func CheckUserIsMemberOfTeam(userId uint64, teamId uint64) bool {
	requestURL := fmt.Sprintf("http://%s/schedule/team/%d/%d/isMember", conf.UserService, teamId, userId)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return false
	}

	var res *http.Response
	for i := 0; i < ReqRetries; i++ {
		res, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			continue
		}
		break
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		_ = res.Body.Close()
		fmt.Printf("client: could not read response body: %s\n", err)
		return false
	}
	_ = res.Body.Close()

	var isMember isMemberStruct
	err = json.Unmarshal(resBody, &isMember)
	if err != nil {
		fmt.Printf("Could not unmarshal response body: %s\n", err)
		return false
	}

	return isMember.TeamMember
}
