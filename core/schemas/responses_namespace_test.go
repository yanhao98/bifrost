package schemas

import "testing"

func TestResponsesMessage_FunctionCallNamespaceRoundTrip(t *testing.T) {
	raw := []byte(`{
	  "type": "function_call",
	  "status": "completed",
	  "call_id": "call_123",
	  "name": "get_app_state",
	  "namespace": "mcp__computer_use__",
	  "arguments": "{\"app\":\"Google Chrome\"}"
	}`)

	var msg ResponsesMessage
	if err := Unmarshal(raw, &msg); err != nil {
		t.Fatalf("failed to unmarshal message: %v", err)
	}

	if msg.ResponsesToolMessage == nil {
		t.Fatal("expected tool message to be present")
	}
	if msg.Namespace == nil || *msg.Namespace != "mcp__computer_use__" {
		t.Fatalf("expected namespace to survive unmarshal, got %+v", msg.Namespace)
	}

	marshaled, err := MarshalSorted(msg)
	if err != nil {
		t.Fatalf("failed to marshal message: %v", err)
	}

	var roundTrip ResponsesMessage
	if err := Unmarshal(marshaled, &roundTrip); err != nil {
		t.Fatalf("failed to unmarshal round-tripped message: %v", err)
	}
	if roundTrip.Namespace == nil || *roundTrip.Namespace != "mcp__computer_use__" {
		t.Fatalf("expected namespace to survive round trip, got %+v; raw=%s", roundTrip.Namespace, marshaled)
	}
}

func TestBifrostResponsesStreamResponse_FunctionCallItemNamespaceRoundTrip(t *testing.T) {
	raw := []byte(`{
	  "type": "response.output_item.done",
	  "sequence_number": 43,
	  "output_index": 2,
	  "item": {
	    "id": "fc_123",
	    "type": "function_call",
	    "status": "completed",
	    "arguments": "{\"app\":\"Google Chrome\"}",
	    "call_id": "call_123",
	    "name": "get_app_state",
	    "namespace": "mcp__computer_use__"
	  },
	  "extra_fields": {}
	}`)

	var event BifrostResponsesStreamResponse
	if err := Unmarshal(raw, &event); err != nil {
		t.Fatalf("failed to unmarshal stream event: %v", err)
	}
	if event.Item == nil || event.Item.ResponsesToolMessage == nil {
		t.Fatal("expected function_call item to be present")
	}
	if event.Item.Namespace == nil || *event.Item.Namespace != "mcp__computer_use__" {
		t.Fatalf("expected namespace to survive unmarshal, got %+v", event.Item.Namespace)
	}

	marshaled, err := MarshalSorted(event)
	if err != nil {
		t.Fatalf("failed to marshal stream event: %v", err)
	}

	var roundTrip BifrostResponsesStreamResponse
	if err := Unmarshal(marshaled, &roundTrip); err != nil {
		t.Fatalf("failed to unmarshal marshaled event: %v", err)
	}
	if roundTrip.Item == nil || roundTrip.Item.Namespace == nil || *roundTrip.Item.Namespace != "mcp__computer_use__" {
		t.Fatalf("expected namespace to survive stream-event round trip, got %+v; raw=%s", roundTrip.Item, marshaled)
	}
}
