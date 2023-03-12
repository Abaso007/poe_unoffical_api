package main

import (
	"fmt"
	"github.com/isxuelinme/poe_unoffical_api/core"
)

func main() {
	core.SetLogMode(core.LOG_ERROR)
	MutLtiUser := core.NewMutLtiUserGpt(core.GptTypePoeUnofficial)
	ask := &core.AskRequest{
		UserId:           1,
		Question:         "I love forever but you love me",
		CallbackFuncName: "",
		AskResponseCallBack: func(askRequest *core.AskRequest, message *core.CallbackMessageResponse) {
			fmt.Printf("\r answer: %s", message.Data.Text)
		},
	}
	//ask question
	MutLtiUser.Talk(ask)

	select {}

}