package folders

import "testing"

// customerSingleFolderPayload reflects a response shape seen in the wild where
// "resource" is a single root folder object (with nested "sub_folders") rather
// than an array of folders.
const customerSingleFolderPayload = `{
	"items": 50,
	"resource": {
		"Children": [
			{"id": "22", "name": "Folder A"},
			{"id": "25", "name": "Folder B"}
		],
		"child_count": 2,
		"created_at": "2023-11-09T16:48:17.785000+00:00",
		"id": "shared",
		"name": "Shared Folder",
		"permission": "core_dashboards_edit",
		"sub_folders": [
			{
				"Children": [{"id": "461", "name": "Folder C"}],
				"child_count": 1,
				"created_at": "2026-01-23T17:59:52.493000+00:00",
				"id": "189",
				"name": "Nested Folder",
				"permission": "core_dashboards_edit",
				"sub_folders": [],
				"type": "ACCOUNT"
			}
		],
		"type": "ACCOUNT"
	}
}`

func TestParseListFoldersResponse_SingleRootFolderObject(t *testing.T) {
	folders, err := parseListFoldersResponse([]byte(customerSingleFolderPayload))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(folders) != 1 {
		t.Fatalf("expected 1 root folder, got %d", len(folders))
	}
	root := folders[0]
	if root.Id != "shared" || root.Name != "Shared Folder" {
		t.Fatalf("unexpected root folder: %+v", root)
	}
	if len(root.SubFolders) != 1 || root.SubFolders[0].Id != "189" {
		t.Fatalf("expected root to retain sub_folders, got: %+v", root.SubFolders)
	}

	flattened := flattenFolders(folders)
	if len(flattened) != 2 {
		t.Fatalf("expected root + 1 nested sub-folder after flatten, got %d: %+v", len(flattened), flattened)
	}
}

func TestParseListFoldersResponse_ArrayEnvelope(t *testing.T) {
	payload := `{"resource": [{"id": "1", "name": "Folder A"}, {"id": "2", "name": "Folder B"}]}`
	folders, err := parseListFoldersResponse([]byte(payload))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(folders) != 2 {
		t.Fatalf("expected 2 folders, got %d", len(folders))
	}
}

func TestParseListFoldersResponse_BareArray(t *testing.T) {
	payload := `[{"id": "1", "name": "Folder A"}]`
	folders, err := parseListFoldersResponse([]byte(payload))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(folders) != 1 {
		t.Fatalf("expected 1 folder, got %d", len(folders))
	}
}

func TestParseListFoldersResponse_UnexpectedShape(t *testing.T) {
	payload := `{"resource": "not a folder"}`
	_, err := parseListFoldersResponse([]byte(payload))
	if err == nil {
		t.Fatal("expected an error for an unrecognized response shape, got nil")
	}
}
