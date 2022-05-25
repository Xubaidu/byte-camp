package repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

// hash index
var (
	topicIndexMap map[int64]*Topic
	topicCnt      int64
	postsIndexMap map[int64][]*Post
	postIndexMap  map[int64]*Post
	postCnt       int64
)

func initTopicIndex(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicIndexMap[topic.ID] = &topic
	}
	return nil
}

func QueryTopic(ID int64) (*Topic, error) {
	v, ok := topicIndexMap[ID]
	if !ok {
		return nil, errors.New("not found topic")
	}
	return v, nil
}

func ModifyTopic(ID int64, detail map[string]interface{}) error {
	topic, err := QueryTopic(ID)
	if err != nil {
		return err
	}
	for k, v := range detail {
		switch k {
		case "Title":
			topic.Title = v.(string)
		case "Content":
			topic.Content = v.(string)
		}
	}
	return nil
}

func CreateTopic(title, content string) error {
	topicIndexMap[topicCnt] = &Topic{
		ID:          topicCnt,
		Title:       title,
		Content:     content,
		CreatedTime: time.Now(),
	}
	topicCnt++
	return nil
}

func DeleteTopic(ID int64) error {
	topicIndexMap[topicCnt] = nil
	return nil
}

func initPostsIndex(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	var posts []*Post
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		posts = append(posts, &post)
		postIndexMap[post.ID] = &post
	}
	if len(posts) == 0 {
		return errors.New("no posts here")
	}
	postsIndexMap[posts[0].TopicID] = posts

	return nil
}

func QueryPosts(ID int64) ([]*Post, error) {
	v, ok := postsIndexMap[ID]
	if !ok {
		return nil, errors.New("not found topic")
	}
	return v, nil
}

func QueryPost(ID int64) (*Post, error) {
	v, ok := postIndexMap[ID]
	if !ok {
		return nil, errors.New("not found post")
	}
	return v, nil
}

func ModifyPost(ID int64, detail map[string]interface{}) error {
	post, err := QueryPost(ID)
	if err != nil {
		return err
	}
	for k, v := range detail {
		switch k {
		case "Content":
			post.Content = v.(string)
		}
	}
	return nil
}

func CreatePost(title, content string, topicID int64) error {
	postIndexMap[topicCnt] = &Post{
		ID:          topicCnt,
		TopicID:     topicID,
		Content:     content,
		CreatedTime: time.Now(),
	}
	postCnt++
	return nil
}

func DeletePost(ID int64) error {
	postIndexMap[topicCnt] = nil
	return nil
}
