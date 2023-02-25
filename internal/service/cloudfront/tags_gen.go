// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package cloudfront

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists cloudfront service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn cloudfrontiface.CloudFrontAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &cloudfront.ListTagsForResourceInput{
		Resource: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags.Items), nil
}

// []*SERVICE.Tag handling

// Tags returns cloudfront service tags.
func Tags(tags tftags.KeyValueTags) []*cloudfront.Tag {
	result := make([]*cloudfront.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &cloudfront.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from cloudfront service tags.
func KeyValueTags(ctx context.Context, tags []*cloudfront.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// UpdateTags updates cloudfront service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn cloudfrontiface.CloudFrontAPI, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &cloudfront.UntagResourceInput{
			Resource: aws.String(identifier),
			TagKeys:  &cloudfront.TagKeys{Items: aws.StringSlice(removedTags.IgnoreAWS().Keys())},
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &cloudfront.TagResourceInput{
			Resource: aws.String(identifier),
			Tags:     &cloudfront.Tags{Items: Tags(updatedTags.IgnoreAWS())},
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
