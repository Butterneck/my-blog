package models

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/butterneck/my-blog/backend/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PostId string

func NewPostId(log *zap.SugaredLogger) (PostId, error) {
	log.Info("NewPostId")
	id := uuid.New().String()
	return PostId(id), nil
}

type NewPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type UpdatePostRequest struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

type Tag string

type Post struct {
	Id          PostId `json:"id" dynamodbav:"id"`
	Title       string `json:"title" dynamodbav:"title"`
	Body        string `json:"body" dynamodbav:"body"`
	Slug        string `json:"slug" dynamodbav:"slug"`
	CreatedAt   int64  `json:"createdAt" dynamodbav:"createdAt"`
	Description string `json:"description" dynamodbav:"description"`
	IsCompleted bool   `json:"isComplete" dynamodbav:"isComplete"`
	Tags        []Tag  `json:"tags" dynamodbav:"tags"`
}

func NewPost(log *zap.SugaredLogger, title, body string) (*Post, error) {
	log.Info("NewPost")
	log.Debugf("NewPost title: %s", title)
	log.Debugf("NewPost body: %s", body)

	postId, err := NewPostId(log)
	if err != nil {
		return nil, fmt.Errorf("NewPost - NewPostId - error: %v", err)
	}
	log.Debugf("NewPost postId: %s", postId)

	slug, err := slugify(log, title)
	if err != nil {
		return nil, fmt.Errorf("NewPost - error: %v", err)
	}

	return &Post{
		Id:          postId,
		Title:       title,
		Body:        body,
		Description: "description",
		Slug:        slug,
		CreatedAt:   time.Now().Unix(),
		IsCompleted: true,
		Tags: []Tag{
			"tag1",
			"tag2",
		},
	}, nil
}

func UpdatePostTitle(log *zap.SugaredLogger, post *Post, title string) (Post, error) {
	log.Debugf("Update post `%s` title: %s -> %s", post.Id, post.Title, title)
	updatedPost := *post

	// Update title
	updatedPost.Title = title

	// Update slug
	slug, err := slugify(log, title)
	if err != nil {
		return Post{}, fmt.Errorf("UpdatePostTitle - error: %v", err)
	}
	updatedPost.Slug = slug

	return updatedPost, nil
}

func UpdatePostBody(log *zap.SugaredLogger, post *Post, body string) (Post, error) {
	log.Debugf("Update post `%s` body: %s -> %s", post.Id, post.Body, body)
	updatedPost := *post
	updatedPost.Body = body
	return updatedPost, nil
}

func PersistNewPost(ctx context.Context, log *zap.SugaredLogger, post *Post) (err error) {
	log.Debugf("Persist post `%s`", post.Id)
	item, err := attributevalue.MarshalMap(post)
	if err != nil {
		return fmt.Errorf("PersistPost - attributevalue.MarshalMap - error: %v", err)
	}

	db := db.GetDB(log)
	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String("Posts"),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	})
	if err != nil {
		return fmt.Errorf("PersistPost - db.PutItem - error: %v", err)
	}

	return nil
}

func PersistUpdatedPost(ctx context.Context, log *zap.SugaredLogger, post *Post) (err error) {
	log.Debugf("Persist updated post `%s`", post.Id)
	item, err := attributevalue.MarshalMap(post)
	if err != nil {
		return fmt.Errorf("PersistUpdatedPost - attributevalue.MarshalMap - error: %v", err)
	}

	db := db.GetDB(log)
	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String("Posts"),
		Item:                item,
		ConditionExpression: aws.String("attribute_exists(id)"),
	})
	if err != nil {
		return fmt.Errorf("PersistUpdatedPost - db.PutItem - error: %v", err)
	}

	return nil
}

func PersistPost(ctx context.Context, log *zap.SugaredLogger, post *Post) (err error) {
	log.Debugf("Persist post `%s`", post.Id)
	item, err := attributevalue.MarshalMap(post)
	if err != nil {
		return fmt.Errorf("PersistPost - attributevalue.MarshalMap - error: %v", err)
	}

	db := db.GetDB(log)
	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("Posts"),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("PersistPost - db.PutItem - error: %v", err)
	}

	return nil
}

type PostFilter struct{}

func GetPersistedPosts(ctx context.Context, log *zap.SugaredLogger, filter PostFilter) ([]Post, error) {
	log.Debugf("Get persisted posts with filter: %v", filter)
	db := db.GetDB(log)
	resp, err := db.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("Posts"),
	})
	if err != nil {
		return nil, fmt.Errorf("getPersistedPosts - db.Scan - error: %v", err)
	}

	var posts []Post
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &posts)
	if err != nil {
		return nil, fmt.Errorf("getPersistedPosts - attributevalue.UnmarshalListOfMaps - error: %v", err)
	}

	return posts, nil
}

func GetPersistedPost(ctx context.Context, log *zap.SugaredLogger, id PostId) (*Post, error) {
	log.Debugf("Get persisted post with id: %s", id)
	db := db.GetDB(log)
	resp, err := db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("Posts"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: string(id)},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("getPersistedPost - db.GetItem - error: %v", err)
	}

	if resp.Item == nil {
		return nil, nil
	}

	var post Post
	err = attributevalue.UnmarshalMap(resp.Item, &post)
	if err != nil {
		return nil, fmt.Errorf("getPersistedPost - attributevalue.UnmarshalMap - error: %v", err)
	}

	return &post, nil
}

func slugify(log *zap.SugaredLogger, s string) (string, error) {
	log.Debugf("slugify: %s", s)

	// Convert to lowercase
	slug := strings.ToLower(s)

	// Replace spaces with hyphens
	re, err := regexp.Compile(`\s+`)
	if err != nil {
		return "", fmt.Errorf("slugify - regexp.Compile - error: %v", err)
	}
	slug = re.ReplaceAllString(slug, "-")

	// Remove special characters, leaving only letters, numbers, hyphens, and underscores
	re, err = regexp.Compile(`[^\w-]+`)
	if err != nil {
		return "", fmt.Errorf("slugify - regexp.Compile - error: %v", err)
	}
	slug = re.ReplaceAllString(slug, "")

	log.Debugf("slugified: %s", slug)

	return slug, nil
}

func DeletePersistedPost(ctx context.Context, log *zap.SugaredLogger, id PostId) error {
	log.Debugf("Delete post `%s`", id)
	db := db.GetDB(log)
	_, err := db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("Posts"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: string(id)},
		},
	})
	if err != nil {
		return fmt.Errorf("DeletePost - db.DeleteItem - error: %v", err)
	}

	return nil
}
