package file_store

import (
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	orgId            = "org_id"
	projectId        = "project_id"
	identifier       = "identifier"
	fileContentPath  = "file_content_path"
	name             = "name"
	parentIdentifier = "parent_identifier"
	path             = "path"
	content          = "content"
	description      = "description"
	mimeType         = "mime_type"
	tags             = "tags"
	type_            = "type"
	fileUsage        = "file_usage"
	createdBy        = "created_by"
	lastModifiedBy   = "last_modified_by"
	lastModifiedAt   = "last_modified_at"
	draft            = "draft"
)

func buildTagsJson(tags *schema.Set) string {
	result := "["
	tagMap := make(map[string]string)
	for i := 0; i < tags.Len(); i++ {
		tag := fmt.Sprintf("%v", tags.List()[i])
		if strings.Contains(tag, ":") {
			splitTag := strings.Split(tag, ":")
			key := splitTag[0]
			value := splitTag[1]
			tagMap[key] = value
		} else {
			tagMap[tag] = ""
		}
	}

	first := true
	for key, value := range tagMap {
		if !first {
			result += ","
		} else {
			first = false
		}
		result += fmt.Sprintf(`{"key":"%s","value":"%s"}`, key, value)
	}

	result += "]"
	return result
}

func FlattenTags(tags []nextgen.NgTag) []string {
	var result []string
	for _, tag := range tags {
		result = append(result, tag.Key+":"+tag.Value)
	}
	return result
}

func getOptionalString(str interface{}) optional.String {
	v, ok := str.(string)
	if !ok {
		return optional.String{}
	}

	if len(v) == 0 {
		return optional.EmptyString()
	}

	return optional.NewString(v)
}

func getSafeEmail(user *nextgen.EmbeddedUserDetailsDto) string {
	if user != nil {
		return user.Email
	}
	return ""
}

func getSafeName(user *nextgen.EmbeddedUserDetailsDto) string {
	if user != nil {
		return user.Name
	}
	return ""
}
