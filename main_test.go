package main

import (
	"testing"

	"github.com/jo-hoe/ai-cli-assistant/backend"
	"github.com/jo-hoe/ai-cli-assistant/backend/openai"
	"github.com/jo-hoe/ai-cli-assistant/internal"
)

func Test_runCliWithHttpClient(t *testing.T) {
	type args struct {
		aiClient backend.AIClient
		cliArgs  []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				aiClient: openai.NewOpenAIClient("dummy", 0, internal.CreateMockClient(internal.ResponseSummery{
					ResponseCode: 200,
					ResponseBody: sampleResponse,
				})),
				cliArgs: []string{"do", "x"},
			},
			want:    "\n\nThis is a test!",
			wantErr: false,
		}, {
			name: "open ai error",
			args: args{
				aiClient: openai.NewOpenAIClient("dummy", 0, internal.CreateMockClient(internal.ResponseSummery{
					ResponseCode: 500,
					ResponseBody: "dummy",
				})),
				cliArgs: []string{"do", "x"},
			},
			want:    "",
			wantErr: true,
		}, {
			name: "missing argument",
			args: args{
				aiClient: openai.NewOpenAIClient("dummy", 0, internal.CreateMockClient(internal.ResponseSummery{
					ResponseCode: 200,
					ResponseBody: sampleResponse,
				})),
				cliArgs: []string{""},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runCliWithAIClient(tt.args.cliArgs, tt.args.aiClient)
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
