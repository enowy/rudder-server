package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/delete/api"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/initialize"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func (d *deleteAPI) handler() http.Handler {
	srvMux := mux.NewRouter()
	srvMux.HandleFunc("/deleteUsers", d.deleteMockServer).Methods("POST")

	return srvMux
}

func TestDelete(t *testing.T) {
	initialize.Init()
	tests := []struct {
		name                 string
		job                  model.Job
		destConfig           map[string]interface{}
		destName             string
		respCode             int
		respBody             string
		expectedDeleteStatus model.JobStatus
		expectedPayload      string
	}{
		{
			name: "test deleter API client with expected status complete",
			job: model.Job{
				ID:            1,
				WorkspaceID:   "1001",
				DestinationID: "1234",
				Status:        model.JobStatusPending,
				UserAttributes: []model.UserAttribute{
					{
						UserID: "Jermaine1473336609491897794707338",
						Phone:  strPtr("6463633841"),
						Email:  strPtr("dorowane8n285680461479465450293436@gmail.com"),
					},
					{
						UserID: "Mercie8221821544021583104106123",
						Email:  strPtr("dshirilad8536019424659691213279980@gmail.com"),
					},
					{
						UserID: "Claiborn443446989226249191822329",
						Phone:  strPtr("8782905113"),
					},
				},
			},
			destConfig: map[string]interface{}{
				"bucketName":  "regulation-test-data",
				"accessKeyID": "abc",
				"accessKey":   "xyz",
				"enableSSE":   false,
				"prefix":      "reg-original",
			},
			destName:             "amplitude",
			respCode:             200,
			respBody:             `[{"status": "complete"}]`,
			expectedDeleteStatus: model.JobStatusComplete,
			expectedPayload:      `[{"jobId":"1","destType":"amplitude","config":{"accessKey":"xyz","accessKeyID":"abc","bucketName":"regulation-test-data","enableSSE":false,"prefix":"reg-original"},"userAttributes":[{"userId":"Jermaine1473336609491897794707338","phone":"6463633841","email":"dorowane8n285680461479465450293436@gmail.com"},{"userId":"Mercie8221821544021583104106123","email":"dshirilad8536019424659691213279980@gmail.com"},{"userId":"Claiborn443446989226249191822329","phone":"8782905113"}]}]`,
		},
		{
			name:                 "test deleter API client with expected status failed: error returned 429",
			destName:             "amplitude",
			respCode:             429,
			respBody:             `[{"status": "failed"}]`,
			expectedDeleteStatus: model.JobStatusFailed,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
		{
			name:                 "test deleter API client with expected status failed-error returned 408",
			destName:             "amplitude",
			respCode:             408,
			respBody:             `[{"status": "failed"}]`,
			expectedDeleteStatus: model.JobStatusFailed,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
		{
			name:                 "test deleter API client with expected status failed: error returned 504",
			destName:             "amplitude",
			respCode:             504,
			respBody:             `[{"status": "failed"}]`,
			expectedDeleteStatus: model.JobStatusFailed,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
		{
			name:                 "test deleter API client with expected status failed: error returned 400",
			destName:             "amplitude",
			respCode:             400,
			respBody:             `[{"status": "failed"}]`,
			expectedDeleteStatus: model.JobStatusAborted,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
		{
			name:                 "test deleter API client with expected status failed: error returned 401",
			destName:             "amplitude",
			respCode:             401,
			respBody:             `[{"status": "failed"}]`,
			expectedDeleteStatus: model.JobStatusAborted,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
		{
			name:                 "test deleter API client with expected status failed: error returned 405",
			destName:             "amplitude",
			respCode:             405,
			respBody:             `[{"status": "failed"}]`,
			expectedDeleteStatus: model.JobStatusNotSupported,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
		{
			name:                 "return 400",
			destName:             "amplitude",
			respCode:             400,
			respBody:             `[{"statusCode":400,"error":"api key/secret for deletion not present"}]`,
			expectedDeleteStatus: model.JobStatusAborted,
			expectedPayload:      `[{"jobId":"0","destType":"amplitude","config":null,"userAttributes":[]}]`,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			d := deleteAPI{
				respStatusCode: tt.respCode,
				respBody:       tt.respBody,
			}
			ctx := context.Background()
			svr := httptest.NewServer(d.handler())

			defer svr.Close()
			os.Setenv("DEST_TRANSFORM_URL", svr.URL)
			api := api.APIManager{
				Client:           &http.Client{},
				DestTransformURL: svr.URL,
			}
			status := api.Delete(ctx, tt.job, tt.destConfig, tt.destName)
			require.Equal(t, tt.expectedDeleteStatus, status)
			require.Equal(t, tt.expectedPayload, d.payload)
		})
	}
}

type deleteAPI struct {
	payload        string
	respStatusCode int
	respBody       string
}

func (d *deleteAPI) deleteMockServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	d.payload = buf.String()

	w.WriteHeader(d.respStatusCode)

	_, err = w.Write([]byte(d.respBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func strPtr(str string) *string {
	return &(str)
}
