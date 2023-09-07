package api

// ServerStatus is send as part of api response
// this is quite similair how axios responds
type ServerStatus struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

// Response type for json api response
// NOTE! json tags with "" and without spaces
type Response struct {
	ServerStatus
	// empty interface makes it possible
	// to send any data format
	Payload interface{} `json:"payload"`
}

// SetErrorResponse creates standardized api response
func SetErrorResponse(err error, state ServerStatus) Response {
	if state.Status == 0 {
		state.Status = 500
	}
	if state.StatusText == "" {
		state.StatusText = "Internal Server Error"
	}
	var r Response

	r.Status = state.Status
	r.StatusText = state.StatusText
	r.Payload = err.Error()

	return r
}

// SetOKResponse creates standard response with status and payload
func SetOKResponse(data interface{}) Response {
	var r Response

	r.Status = 200
	r.StatusText = "OK"
	r.Payload = data

	return r
}

// func returnError(ctx *fiber.Ctx, err error) {
// 	resp := setErrorResponse(err, ServerStatus{})
// 	ctx.Set("content-type", "application/json")
// 	ctx.Status(500).JSON(resp)
// 	log.Println(ctx.Path(), "ERROR", err)
// 	return
// }

// func returnOK(ctx *fiber.Ctx, payload interface{}) {
// 	resp := SetOKResponse(payload)
// 	// no return needed?
// 	ctx.Status(200).JSON(resp)
// 	return
// }
