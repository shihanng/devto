# ArticleUpdateArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | [optional] 
**BodyMarkdown** | **string** | The body of the article.  It can contain an optional front matter. For example  &#x60;&#x60;&#x60;markdown --- title: Hello, World! published: true tags: discuss, help date: 20190701T10:00Z series: Hello series canonical_url: https://example.com/blog/hello cover_image: article_published_cover_image --- &#x60;&#x60;&#x60;  &#x60;date&#x60;, &#x60;series&#x60; and &#x60;canonical_url&#x60; are optional. &#x60;date&#x60; is the publication date-time &#x60;series&#x60; is the name of the series the article belongs to &#x60;canonical_url&#x60; is the canonical URL of the article &#x60;cover_image&#x60; is the main image of the article  *If the markdown contains a front matter, it will take precedence on the equivalent params given in the JSON payload.*  | [optional] 
**Published** | **bool** | True to create a published article, false otherwise. Defaults to false  | [optional] 
**Series** | **string** | Article series name.  All articles belonging to the same series need to have the same name in this parameter.  To remove an article from a series, the &#x60;null&#x60; value can be used.  | [optional] 
**MainImage** | **string** |  | [optional] 
**CanonicalUrl** | **string** |  | [optional] 
**Description** | **string** |  | [optional] 
**Tags** | **[]string** |  | [optional] 
**OrganizationId** | **int32** | Only users belonging to an organization can assign the article to it  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


