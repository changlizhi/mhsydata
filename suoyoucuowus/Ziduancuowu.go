package suoyoucuowus

import (
	"encoding/json"
	"time"
)

type Ziduancuowu struct {
	Shijian time.Time
	Wenti   string
}

func (e Ziduancuowu) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
