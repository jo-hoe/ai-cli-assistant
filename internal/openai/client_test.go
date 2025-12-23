package openai

import (
	"testing"

	"github.com/jo-hoe/ai-cli-assistant/internal/httpmock"
)

func TestOpenAIClient_GetAnswer(t *testing.T) {
	type args struct {
		prompt string
	}
	tests := []struct {
		name     string
		aiClient *OpenAIClient
		args     args
		want     string
		wantErr  bool
	}{
		{
			name: "positive test",
			aiClient: NewOpenAIClient("", 0, httpmock.CreateMockClient(httpmock.ResponseSummery{
				ResponseCode: 200,
				ResponseBody: sampleResponse,
			}), "", ""),
			args:    args{prompt: "test"},
			want:    "\n\nThis is a test!",
			wantErr: false,
		},
		{
			name: "server failure test",
			aiClient: NewOpenAIClient("", 0, httpmock.CreateMockClient(httpmock.ResponseSummery{
				ResponseCode: 500,
				ResponseBody: sampleResponse,
			}), "", ""),
			args:    args{prompt: "test"},
			want:    "",
			wantErr: true,
		},
		{
			name: "unexpected JSON response",
			aiClient: NewOpenAIClient("", 0, httpmock.CreateMockClient(httpmock.ResponseSummery{
				ResponseCode: 200,
				ResponseBody: "unexpected",
			}), "", ""),
			args:    args{prompt: "test"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.aiClient.GetAnswer(tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenAIClient.GetAnswer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OpenAIClient.GetAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}

const sampleResponse = `{
	"id":"chatcmpl-abc123",
	"object":"chat.completion",
	"created":1677858242,
	"model":"gpt-3.5-turbo-0301",
	"usage":{
	   "prompt_tokens":13,
	   "completion_tokens":7,
	   "total_tokens":20
	},
	"choices":[
	   {
		  "message":{
			 "role":"assistant",
			 "content":"\n\nThis is a test!"
		  },
		  "finish_reason":"stop",
		  "index":0
	   }
	]
 }`
