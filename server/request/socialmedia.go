package request

type CreateSocialMedia struct {
	Name           string `json:"name" validate:"required" example:""`
	SocialMediaUrl string `json:"social_media_url" validate:"required" example:"https://twitter.com/"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name" validate:"required" example:""`
	SocialMediaUrl string `json:"social_media_url" validate:"required" example:"https://twitter.com/"`
}
