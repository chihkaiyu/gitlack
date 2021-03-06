package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func (g *gitlab) GetSingleCommit(id int, sha string) (*Commit, error) {
	url := g.GitLabAPI + fmt.Sprintf("/projects/%v/repository/commits/%v", id, sha)
	params := map[string]string{
		"private_token": g.GitLabToken,
	}
	res, err := g.client.Get(url, nil, params, nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err := fmt.Errorf("Invalid GitLab API error: %v", string(body))
		logrus.Errorln(err)
		return nil, err
	}
	var commit Commit
	err = json.Unmarshal(body, &commit)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return &commit, nil
}
