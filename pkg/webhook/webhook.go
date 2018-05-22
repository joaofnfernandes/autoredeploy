package webhook

import (
	"encoding/json"
	"reflect"
)

func Unmarshal(str string) Webhook {
	var webhook Webhook
	err := json.Unmarshal([]byte(str), &webhook)
	if err != nil {
		return Webhook{}
	}
	return webhook
}

type image struct {
	namespace string
	name      string
	tag       string
}

var (
	validImages = []image{
		image{"joaofnfernandes", "test", "latest"},
	}
)

func (w *Webhook) isDefault() bool {
	defaultWebhook := Webhook{}
	return reflect.DeepEqual(*w, defaultWebhook)
}

func (w *Webhook) IsValid() bool {
	if w.isDefault() {
		return false
	}
	for _, v := range validImages {
		if v.namespace == w.Repository.Namespace &&
			v.name == w.Repository.Name &&
			v.tag == w.PushData.Tag {
			return true
		}
	}
	return false
}
