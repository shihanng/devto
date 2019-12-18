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
// ArticleMe struct for ArticleMe
type ArticleMe struct {
	TypeOf string `json:"type_of"`
	Id int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CoverImage string `json:"cover_image"`
	Published bool `json:"published"`
	PublishedAt time.Time `json:"published_at"`
	TagList []string `json:"tag_list"`
	Slug string `json:"slug"`
	Path string `json:"path"`
	Url string `json:"url"`
	CanonicalUrl string `json:"canonical_url"`
	CommentsCount int32 `json:"comments_count"`
	PositiveReactionsCount int32 `json:"positive_reactions_count"`
	PageViewsCount int32 `json:"page_views_count"`
	// Crossposting or published date time
	PublishedTimestamp time.Time `json:"published_timestamp"`
	User ArticleUser `json:"user"`
	Organization ArticleOrganization `json:"organization,omitempty"`
	FlareTag ArticleFlareTag `json:"flare_tag,omitempty"`
}