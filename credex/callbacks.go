package credex

import (
	"fmt"

	vcx "github.com/faisal00813/vcx-go/vcx"
)

func cbWithErrorCode(ch chan callbackResult) vcx.Callback_error_code {
	return func(cmd_h vcx.Command_handle_t, error_code vcx.Error_t) {
		fmt.Println("lib Errorcode is : ", error_code)
	}
}
func cbWithIntResponse(ch chan callbackResult) vcx.Callback_int {
	return func(cmd_h vcx.Command_handle_t, error_code vcx.Error_t, int_value uint32) {
		fmt.Println(" lib Errorcode is : ", error_code)
		ch <- callbackResult{uint32(cmd_h), uint32(error_code), uint32(int_value), ""}
	}
}
