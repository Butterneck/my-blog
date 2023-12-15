package main

import (
	"net/http"

	"github.com/butterneck/my-blog/src/blog-backend/log"
	"github.com/butterneck/my-blog/src/blog-backend/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	log := log.GetLogger()
	posts, err := models.GetPersistedPosts(c, log, models.PostFilter{})
	if err != nil {
		log.Errorf("GetPosts - GetPersistedPosts - error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostBySlug(c *gin.Context) {
	log := log.GetLogger()

	slug := c.Param("slug")

	post, err := models.GetPersistedPostBySlug(c, log, slug)
	if err != nil {
		log.Errorf("GetPost - GetPersistedPost - error: %v", err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if post == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	log := log.GetLogger()
	var newPostRequest models.NewPostRequest

	err := c.BindJSON(&newPostRequest)
	if err != nil {
		log.Errorf("CreatePost - BindJSON - error: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	post, err := models.NewPost(log, newPostRequest.Title, newPostRequest.Body)
	if err != nil {
		log.Errorf("CreatePost - NewPost - error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = models.PersistNewPost(c, log, post)
	if err != nil {
		log.Errorf("CreatePost - PersistPost - error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, post)
	c.JSON(http.StatusCreated, nil)
}

func UpdatePost(c *gin.Context) {
	log := log.GetLogger()

	id := c.Param("id")
	var updatedPostRequest models.UpdatePostRequest

	err := c.BindJSON(&updatedPostRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Debugf("UpdatePost - updatedPostRequest: %v", updatedPostRequest)

	post, err := models.GetPersistedPostById(c, log, models.PostId(id))
	if err != nil {
		log.Errorf("UpdatePost - GetPersistedPost - error: %v", err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	updatedPost := *post

	// If new title is provided, update title
	if updatedPostRequest.Title != nil {
		updatedPost, err = models.UpdatePostTitle(log, post, *updatedPostRequest.Title)
		if err != nil {
			log.Errorf("UpdatePost - UpdatePostTitle - error: %v", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// If new body is provided, update body
	if updatedPostRequest.Body != nil {
		updatedPost, err = models.UpdatePostBody(log, &updatedPost, *updatedPostRequest.Body)
		if err != nil {
			log.Errorf("UpdatePost - UpdatePostBody - error: %v", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	err = models.PersistUpdatedPost(c, log, &updatedPost)
	if err != nil {
		log.Errorf("UpdatePost - PersistUpdatedPost - error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedPost)
}

func DeletePost(c *gin.Context) {
	log := log.GetLogger()

	id := c.Param("id")

	err := models.DeletePersistedPost(c, log, models.PostId(id))
	if err != nil {
		log.Errorf("DeletePost - DeletePersistedPost - error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
