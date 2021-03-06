package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// GitLabUser is the response of getting GitLab user list
type GitLabUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (g *gitlab) GetUser() ([]*GitLabUser, error) {
	url := g.GitLabAPI + "/users"
	params := map[string]string{
		"private_token": g.GitLabToken,
		"per_page":      "100",
		"active":        "true",
	}

	var allUsers []*GitLabUser
	// run at most 100 times for preventing from infinite loop
	for i := 0; i < 100; i++ {
		res, err := g.client.Get(url, nil, params, nil)
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}

		if res.StatusCode != 200 {
			err := fmt.Errorf("Invalid GitLab API error: %v", string(body))
			logrus.Errorln(err)
			return nil, err
		}

		var user []*GitLabUser
		err = json.Unmarshal(body, &user)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		allUsers = append(allUsers, user...)

		// check next page
		if res.Header["X-Next-Page"][0] == "" {
			break
		}
		params["page"] = res.Header["X-Next-Page"][0]
	}
	return allUsers, nil
}
