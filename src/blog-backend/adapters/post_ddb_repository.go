package adapters

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/butterneck/my-blog/src/blog-backend/domain/post"
	"github.com/google/uuid"
)

// FirstYear is the first year that posts were published.
const FirstYear = 2023
const FeedName = "butterneck-blog"

type DDBPost struct {
	Id           string `dynamodbav:"id"`
	Title        string `dynamodbav:"title"`
	Body         string `dynamodbav:"body"`
	Slug         string `dynamodbav:"slug"`
	CreationDate int64  `dynamodbav:"creationDate"`
	FeedName     string `dynamodbav:"feedName"`
	DraftTitle   string `dynamodbav:"draftTitle"`
	DraftBody    string `dynamodbav:"draftBody"`
	DraftSlug    string `dynamodbav:"draftSlug"`
}

func (p *DDBPost) ToPost() (*post.Post, error) {

	postAdapter := &post.PostAdapter{
		Title:        post.Title(p.Title),
		Body:         post.Body(p.Body),
		Slug:         post.Slug(p.Slug),
		CreationDate: p.CreationDate,
		Draft: post.DraftAdapter{
			Title: post.Title(p.DraftTitle),
			Body:  post.Body(p.DraftBody),
			Slug:  post.Slug(p.DraftSlug),
		},
	}

	post, err := post.LoadPostFromAdapter(*postAdapter)
	if err != nil {
		return nil, fmt.Errorf("ToPost - LoadPostFromAdapter - error: %v", err)
	}

	return post, nil
}

type DDBPostRepository struct {
	db                 *dynamodb.Client
	tableName          string
	postsListIndexName string
	slugIndexName      string
}

type DDBPostRepositoryConfig struct {
	TableName          string
	PostsListIndexName string
	SlugIndexName      string
}

type DecodedNextPageToken struct {
	CreationDate int64 `dynamodbav:"creationDate"`
}

func NewDDBPostRepository(db *dynamodb.Client, config DDBPostRepositoryConfig) *DDBPostRepository {
	if db == nil {
		panic("NewDDBPostRepository - db is nil")
	}

	return &DDBPostRepository{
		db:                 db,
		tableName:          config.TableName,
		postsListIndexName: config.PostsListIndexName,
		slugIndexName:      config.SlugIndexName,
	}
}

func (r *DDBPostRepository) GetAnyPost(ctx context.Context, slug string) (*post.Post, error) {
	post, err := r.getAnyPost(ctx, slug)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, nil
	}

	return post.ToPost()
}

func (r *DDBPostRepository) getAnyPost(ctx context.Context, slug string) (*DDBPost, error) {
	resp, err := r.db.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		IndexName:              aws.String(r.slugIndexName),
		KeyConditionExpression: aws.String("slug = :slug"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":slug": &types.AttributeValueMemberS{Value: string(slug)},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("GetPost - db.Query - error: %v", err)
	}

	if resp.Items == nil || len(resp.Items) == 0 {
		return nil, nil
	}

	var post DDBPost
	err = attributevalue.UnmarshalMap(resp.Items[0], &post)
	if err != nil {
		return nil, fmt.Errorf("getPersistedPostBySlug - attributevalue.UnmarshalMap - error: %v", err)
	}

	return &post, nil
}

func (r *DDBPostRepository) GetPublishedPost(ctx context.Context, slug string) (*post.Post, error) {

	resp, err := r.db.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		IndexName:              aws.String(r.slugIndexName),
		KeyConditionExpression: aws.String("slug = :slug"),
		FilterExpression:       aws.String("#title <> :emptyString AND #body <> :emptyString AND #creationDate <> :zero"),
		ExpressionAttributeNames: map[string]string{
			"#title":        "title",
			"#body":         "body",
			"#creationDate": "creationDate",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":slug":        &types.AttributeValueMemberS{Value: string(slug)},
			":emptyString": &types.AttributeValueMemberS{Value: ""},
			":zero":        &types.AttributeValueMemberN{Value: "0"},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("GetPost - db.Query - error: %v", err)
	}

	if resp.Items == nil || len(resp.Items) == 0 {
		return nil, nil
	}

	var post DDBPost
	err = attributevalue.UnmarshalMap(resp.Items[0], &post)
	if err != nil {
		return nil, fmt.Errorf("getPersistedPostBySlug - attributevalue.UnmarshalMap - error: %v", err)
	}

	return post.ToPost()
}

func (r *DDBPostRepository) GetPublishedPosts(ctx context.Context, pageSize *int, encodedNextPageToken *string) (*post.PaginatedPosts, error) {

	var nextPageToken string

	decodedNextPageToken, err := decodeNextPageToken(encodedNextPageToken)
	if err != nil {
		return nil, fmt.Errorf("GetPublishedPosts - DecodeNextPageToken - error: %v", err)
	}

	var exclusiveStartKey map[string]types.AttributeValue
	if decodedNextPageToken != nil {
		exclusiveStartKey = map[string]types.AttributeValue{
			"feedName":     &types.AttributeValueMemberS{Value: FeedName},
			"creationDate": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", decodedNextPageToken.CreationDate)},
		}
	}

	resp, err := r.db.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		IndexName:              aws.String(r.postsListIndexName),
		KeyConditionExpression: aws.String("feedName = :feedName AND creationDate > :zero"),
		FilterExpression:       aws.String("#title <> :emptyString AND #body <> :emptyString"),
		ExpressionAttributeNames: map[string]string{
			"#title": "title",
			"#body":  "body",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":feedName":    &types.AttributeValueMemberS{Value: FeedName},
			":emptyString": &types.AttributeValueMemberS{Value: ""},
			":zero":        &types.AttributeValueMemberN{Value: "0"},
		},
		ScanIndexForward:  aws.Bool(false),
		Limit:             aws.Int32(int32(getPageSize(pageSize))),
		ExclusiveStartKey: exclusiveStartKey,
	})
	if err != nil {
		return nil, fmt.Errorf("GetPublishedPosts - db.Scan - error: %v", err)
	}

	var ddbPosts []*DDBPost
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &ddbPosts)
	if err != nil {
		return nil, fmt.Errorf("GetPublishedPosts - attributevalue.UnmarshalListOfMaps - error: %v", err)
	}

	var posts []*post.Post
	for _, ddbPost := range ddbPosts {
		post, err := ddbPost.ToPost()
		if err != nil {
			return nil, fmt.Errorf("GetPublishedPosts - ddbPost.ToPost - error: %v", err)
		}

		posts = append(posts, post)
	}

	if len(resp.LastEvaluatedKey) > 0 {
		var decodedNextPageToken DecodedNextPageToken
		err := attributevalue.UnmarshalMap(resp.LastEvaluatedKey, &decodedNextPageToken)
		if err != nil {
			return nil, fmt.Errorf("GetPublishedPosts - attributevalue.UnmarshalMap - error: %v", err)
		}

		nextPageToken, err = encodeNextPageToken(&decodedNextPageToken)
		if err != nil {
			return nil, fmt.Errorf("GetPublishedPosts - EncodeNextPageToken - error: %v", err)
		}
	}

	return &post.PaginatedPosts{
		Posts:         posts,
		NextPageToken: nextPageToken,
	}, nil
}

func (r *DDBPostRepository) GetAllPosts(ctx context.Context, pageSize *int, encodedNextPageToken *string) (*post.PaginatedPosts, error) {

	var nextPageToken string

	decodedNextPageToken, err := decodeNextPageToken(encodedNextPageToken)
	if err != nil {
		return nil, fmt.Errorf("GetAllPosts - DecodeNextPageToken - error: %v", err)
	}

	var exclusiveStartKey map[string]types.AttributeValue
	if decodedNextPageToken != nil {
		exclusiveStartKey = map[string]types.AttributeValue{
			"feedName":     &types.AttributeValueMemberS{Value: FeedName},
			"creationDate": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", decodedNextPageToken.CreationDate)},
		}
	}

	resp, err := r.db.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.tableName),
		IndexName:              aws.String(r.postsListIndexName),
		KeyConditionExpression: aws.String("feedName = :feedName"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":feedName": &types.AttributeValueMemberS{Value: FeedName},
		},
		ScanIndexForward:  aws.Bool(false),
		Limit:             aws.Int32(int32(getPageSize(pageSize))),
		ExclusiveStartKey: exclusiveStartKey,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllPosts - db.Scan - error: %v", err)
	}

	ddbPosts := []*DDBPost{}
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &ddbPosts)
	if err != nil {
		return nil, fmt.Errorf("GetAllPosts - attributevalue.UnmarshalListOfMaps - error: %v", err)
	}

	var posts []*post.Post
	for _, ddbPost := range ddbPosts {
		post, err := ddbPost.ToPost()
		if err != nil {
			return nil, fmt.Errorf("GetAllPosts - ddbPost.ToPost - error: %v", err)
		}

		posts = append(posts, post)
	}

	if len(resp.LastEvaluatedKey) > 0 {
		var decodedNextPageToken DecodedNextPageToken
		err := attributevalue.UnmarshalMap(resp.LastEvaluatedKey, &decodedNextPageToken)
		if err != nil {
			return nil, fmt.Errorf("GetAllPosts - attributevalue.UnmarshalMap - error: %v", err)
		}

		nextPageToken, err = encodeNextPageToken(&decodedNextPageToken)
		if err != nil {
			return nil, fmt.Errorf("GetAllPosts - EncodeNextPageToken - error: %v", err)
		}
	}

	return &post.PaginatedPosts{
		Posts:         posts,
		NextPageToken: nextPageToken,
	}, nil
}

func (r *DDBPostRepository) UpdatePost(ctx context.Context, slug string, updateFn func(h *post.Post) (*post.Post, error)) error {

	currentDDBPost, err := r.getAnyPost(ctx, slug)
	if err != nil {
		return err
	}

	currentPost, err := currentDDBPost.ToPost()
	if err != nil {
		return err
	}

	updatedPost, err := updateFn(currentPost)
	if err != nil {
		return err
	}

	item := map[string]types.AttributeValue{
		"id":           &types.AttributeValueMemberS{Value: currentDDBPost.Id},
		"title":        &types.AttributeValueMemberS{Value: updatedPost.Title()},
		"body":         &types.AttributeValueMemberS{Value: updatedPost.Body()},
		"slug":         &types.AttributeValueMemberS{Value: updatedPost.Slug()},
		"creationDate": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", updatedPost.CreationDate())},
		"draftTitle":   &types.AttributeValueMemberS{Value: updatedPost.Draft().Title()},
		"draftBody":    &types.AttributeValueMemberS{Value: updatedPost.Draft().Body()},
		"draftSlug":    &types.AttributeValueMemberS{Value: updatedPost.Draft().Slug()},
		"feedName":     &types.AttributeValueMemberS{Value: fmt.Sprintf("%d", FeedName)},
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("PersistPost - db.PutItem - error: %v", err)
	}

	return nil
}

func (r *DDBPostRepository) CreatePost(ctx context.Context, p *post.Post) error {
	newId, err := newPostId()
	if err != nil {
		return fmt.Errorf("CreatePost - newPostId - error: %v", err)
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberS{Value: newId},
			"title":        &types.AttributeValueMemberS{Value: p.Title()},
			"body":         &types.AttributeValueMemberS{Value: p.Body()},
			"slug":         &types.AttributeValueMemberS{Value: p.Slug()},
			"creationDate": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", p.CreationDate())},
			"draftTitle":   &types.AttributeValueMemberS{Value: p.Draft().Title()},
			"draftBody":    &types.AttributeValueMemberS{Value: p.Draft().Body()},
			"draftSlug":    &types.AttributeValueMemberS{Value: p.Draft().Slug()},
			"feedName":     &types.AttributeValueMemberS{Value: FeedName},
		},
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	})
	if err != nil {
		return fmt.Errorf("PersistPost - db.PutItem - error: %v", err)
	}

	return nil
}

func newPostId() (string, error) {
	return uuid.New().String(), nil
}

func decodeNextPageToken(encodedNextPageToken *string) (*DecodedNextPageToken, error) {
	if encodedNextPageToken == nil {
		return nil, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(*encodedNextPageToken)
	if err != nil {
		return nil, err
	}

	var decodedNextPageToken DecodedNextPageToken
	err = json.Unmarshal(decoded, &decodedNextPageToken)
	if err != nil {
		return nil, err
	}

	return &decodedNextPageToken, nil
}

func encodeNextPageToken(nextPageToken *DecodedNextPageToken) (string, error) {
	encoded, err := json.Marshal(nextPageToken)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(encoded)), nil
}

func getPageSize(pageSize *int) int {
	if pageSize == nil {
		return 10
	}

	return *pageSize
}
