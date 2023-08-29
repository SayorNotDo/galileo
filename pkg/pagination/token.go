package pagination

import (
	"encoding/base64"
	"encoding/json"
)

func GeneratePageToken(state State) (string, error) {
	/* 序列化分页状态信息为 JSON 字符串 */
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	/* 进行 URL 安全的 Base64 编码 */
	pageToken := base64.URLEncoding.EncodeToString(stateJSON)
	return pageToken, nil
}

func ParsePageToken(pageToken string) (State, error) {
	/* 进行 URL 安全的 Base64 解码 */
	stateJSON, err := base64.URLEncoding.DecodeString(pageToken)
	if err != nil {
		return State{}, err
	}
	/* 解析 JSON 字符串为分页状态信息结构 */
	var state State
	if err := json.Unmarshal(stateJSON, &state); err != nil {
		return State{}, err
	}
	return state, nil
}
