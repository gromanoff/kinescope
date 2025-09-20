package kinescope

import (
    "encoding/json"
    "fmt"
)

// apiError matches Kinescope error envelope: {"error":{"code":...,"message":"...","detail":"..."}}
type apiError struct {
    Err struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
        Detail  string `json:"detail"`
    } `json:"error"`
    raw string
}

func (e *apiError) Error() string {
    if e == nil {
        return "<nil>"
    }
    return fmt.Sprintf("kinescope: %d %s (%s)", e.Err.Code, e.Err.Message, e.Err.Detail)
}

func parseAPIError(body []byte) error {
    var ae apiError
    if err := json.Unmarshal(body, &ae); err == nil && ae.Err.Code != 0 {
        ae.raw = string(body)
        return &ae
    }
    // Fallback: not a structured error
    return fmt.Errorf("kinescope: http error: %s", string(body))
}
