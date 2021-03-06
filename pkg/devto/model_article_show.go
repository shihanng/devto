/*
 * DEV API (beta)
 *
 * Access DEV articles, comments and other resources via API
 *
 * API version: 0.5.9
 * Contact: yo@dev.to
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package devto

import (
	"time"
)

// ArticleShow struct for ArticleShow
type ArticleShow struct {
	TypeOf                 string    `json:"type_of"`
	Id                     int32     `json:"id"`
	Title                  string    `json:"title"`
	Description            string    `json:"description"`
	CoverImage             string    `json:"cover_image"`
	ReadablePublishDate    string    `json:"readable_publish_date"`
	SocialImage            string    `json:"social_image"`
	TagList                string    `json:"tag_list"`
	Tags                   []string  `json:"tags"`
	Slug                   string    `json:"slug"`
	Path                   string    `json:"path"`
	Url                    string    `json:"url"`
	CanonicalUrl           string    `json:"canonical_url"`
	CommentsCount          int32     `json:"comments_count"`
	PositiveReactionsCount int32     `json:"positive_reactions_count"`
	CreatedAt              time.Time `json:"created_at"`
	EditedAt               time.Time `json:"edited_at"`
	CrosspostedAt          time.Time `json:"crossposted_at"`
	PublishedAt            string    `json:"published_at"`
	LastCommentAt          time.Time `json:"last_comment_at"`
	// Crossposting or published date time
	PublishedTimestamp string              `json:"published_timestamp"`
	BodyHtml           string              `json:"body_html"`
	BodyMarkdown       string              `json:"body_markdown"`
	User               ArticleUser         `json:"user"`
	Organization       ArticleOrganization `json:"organization,omitempty"`
	FlareTag           ArticleFlareTag     `json:"flare_tag,omitempty"`
}
