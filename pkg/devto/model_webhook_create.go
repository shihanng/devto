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
// WebhookCreate Webhook creation payload
type WebhookCreate struct {
	WebhookEndpoint WebhookCreateWebhookEndpoint `json:"webhook_endpoint,omitempty"`
}