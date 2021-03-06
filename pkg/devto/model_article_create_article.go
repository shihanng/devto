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

// ArticleCreateArticle struct for ArticleCreateArticle
type ArticleCreateArticle struct {
	Title string `json:"title"`
	// The body of the article.  It can contain an optional front matter. For example  ```markdown --- title: Hello, World! published: true tags: discuss, help date: 20190701T10:00Z series: Hello series canonical_url: https://example.com/blog/hello cover_image: article_published_cover_image --- ```  `date`, `series` and `canonical_url` are optional. `date` is the publication date-time `series` is the name of the series the article belongs to `canonical_url` is the canonical URL of the article `cover_image` is the main image of the article  *If the markdown contains a front matter, it will take precedence on the equivalent params given in the JSON payload.*
	BodyMarkdown string `json:"body_markdown,omitempty"`
	// True to create a published article, false otherwise. Defaults to false
	Published bool `json:"published,omitempty"`
	// Article series name.  All articles belonging to the same series need to have the same name in this parameter.
	Series       string   `json:"series,omitempty"`
	MainImage    string   `json:"main_image,omitempty"`
	CanonicalUrl string   `json:"canonical_url,omitempty"`
	Description  string   `json:"description,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	// Only users belonging to an organization can assign the article to it
	OrganizationId int32 `json:"organization_id,omitempty"`
}
