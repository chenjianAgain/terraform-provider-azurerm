package parse

import (
	"testing"
)

func TestEventGridDomainTopicId(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *EventGridDomainTopicId
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "Missing Domain",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.EventGrid/topics/topic1",
			Expected: nil,
		},
		{
			Name:  "Domain Topic ID",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.EventGrid/domains/domain1/topics/topic1",
			Expected: &EventGridDomainTopicId{
				Name:          "topic1",
				ResourceGroup: "resGroup1",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := EventGridDomainTopicID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for Resource Group", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
	}
}
