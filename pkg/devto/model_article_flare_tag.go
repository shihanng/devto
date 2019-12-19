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

// ArticleFlareTag Flare tag of the article
type ArticleFlareTag struct {
	Name string `json:"name,omitempty"`
	// Background color (hexadecimal)
	BgColorHex string `json:"bg_color_hex,omitempty"`
	// Text color (hexadecimal)
	TextColorHex string `json:"text_color_hex,omitempty"`
}
