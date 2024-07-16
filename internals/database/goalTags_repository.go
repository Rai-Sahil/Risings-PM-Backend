package database

import (
	"errors"
	"fmt"
)

// Tag represents a tag with a name and color
type Tag struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// TagManager manages a collection of tags
type TagManager struct {
	tags map[string]Tag // Use a map for easy lookup by tag name
}

// NewTagManager creates a new TagManager
func NewTagManager() *TagManager {
	return &TagManager{
		tags: make(map[string]Tag),
	}
}

// AddTag adds a new tag to the manager
func (tm *TagManager) AddTag(name, color string) error {
	if _, exists := tm.tags[name]; exists {
		return errors.New("tag already exists")
	}
	tm.tags[name] = Tag{Name: name, Color: color}
	return nil
}

// DeleteTag removes a tag from the manager by name
func (tm *TagManager) DeleteTag(name string) error {
	if _, exists := tm.tags[name]; !exists {
		return errors.New("tag not found")
	}
	delete(tm.tags, name)
	return nil
}

// ChangeTagName changes the name of an existing tag
func (tm *TagManager) ChangeTagName(oldName, newName string) error {
	if _, exists := tm.tags[oldName]; !exists {
		return errors.New("tag not found")
	}
	if _, exists := tm.tags[newName]; exists {
		return errors.New("new tag name already exists")
	}
	tag := tm.tags[oldName]
	tag.Name = newName
	tm.tags[newName] = tag
	delete(tm.tags, oldName)
	return nil
}

// PrintTags prints all the tags for debugging purposes
func (tm *TagManager) PrintTags() {
	for _, tag := range tm.tags {
		fmt.Printf("Tag: %s, Color: %s\n", tag.Name, tag.Color)
	}
}

func (tm *TagManager) GetTags() []Tag {
	tags := []Tag{}
	for _, tag := range tm.tags {
		tags = append(tags, tag)
	}
	return tags
}

var NewGoalTagManager = NewTagManager()
