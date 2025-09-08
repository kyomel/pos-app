package response

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccess(t *testing.T) {
	resp := NewSuccessCreated("test",
		WithPayload(map[string]interface{}{
			"foo": "bar",
		}),
		WithStatusCode("2010101100"),
	)

	expected := map[string]interface{}{
		"foo": "bar",
	}

	resJson, _ := json.Marshal(resp)
	fmt.Printf("%s\n", string(resJson))

	require.NotEmpty(t, resp)
	require.Equal(t, expected, resp.Payload)
	require.Equal(t, "2010101100", resp.StatusCode)
}
