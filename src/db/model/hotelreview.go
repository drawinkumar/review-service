package model

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HotelReview struct {
	gorm.Model
	HotelID   int    `gorm:"column:hotel_id" json:"hotelId"`
	Platform  string `gorm:"column:platform" json:"platform"`
	HotelName string `gorm:"column:hotel_name" json:"hotelName"`

	// Comment
	IsShowReviewResponse    bool      `gorm:"column:is_show_review_response" json:"isShowReviewResponse"`
	HotelReviewID           int       `gorm:"column:hotel_review_id;uniqueIndex:idx_review_provider" json:"hotelReviewId"`
	ProviderID              int       `gorm:"column:provider_id;uniqueIndex:idx_review_provider" json:"providerId"`
	Rating                  float64   `gorm:"column:rating" json:"rating"`
	CheckInDateMonthAndYear string    `gorm:"column:check_in_date_month_and_year" json:"checkInDateMonthAndYear"`
	EncryptedReviewData     string    `gorm:"column:encrypted_review_data" json:"encryptedReviewData"`
	FormattedRating         string    `gorm:"column:formatted_rating" json:"formattedRating"`
	FormattedReviewDate     string    `gorm:"column:formatted_review_date" json:"formattedReviewDate"`
	RatingText              string    `gorm:"column:rating_text" json:"ratingText"`
	ResponderName           string    `gorm:"column:responder_name" json:"responderName"`
	ResponseDateText        string    `gorm:"column:response_date_text" json:"responseDateText"`
	ResponseTranslateSource string    `gorm:"column:response_translate_source" json:"responseTranslateSource"`
	ReviewComments          string    `gorm:"column:review_comments" json:"reviewComments"`
	ReviewNegatives         string    `gorm:"column:review_negatives" json:"reviewNegatives"`
	ReviewPositives         string    `gorm:"column:review_positives" json:"reviewPositives"`
	ReviewProviderLogo      string    `gorm:"column:review_provider_logo" json:"reviewProviderLogo"`
	ReviewProviderText      string    `gorm:"column:review_provider_text" json:"reviewProviderText"`
	ReviewTitle             string    `gorm:"column:review_title" json:"reviewTitle"`
	TranslateSource         string    `gorm:"column:translate_source" json:"translateSource"`
	TranslateTarget         string    `gorm:"column:translate_target" json:"translateTarget"`
	ReviewDate              time.Time `gorm:"column:review_date" json:"reviewDate"`
	OriginalTitle           string    `gorm:"column:original_title" json:"originalTitle"`
	OriginalComment         string    `gorm:"column:original_comment" json:"originalComment"`
	FormattedResponseDate   string    `gorm:"column:formatted_response_date" json:"formattedResponseDate"`

	// Reviewer Info
	ReviewerCountryName         string `gorm:"column:reviewer_country_name" json:"reviewerCountryName"`
	ReviewerDisplayMemberName   string `gorm:"column:reviewer_display_member_name" json:"reviewerDisplayMemberName"`
	ReviewerFlagName            string `gorm:"column:reviewer_flag_name" json:"reviewerFlagName"`
	ReviewerReviewGroupName     string `gorm:"column:reviewer_review_group_name" json:"reviewerReviewGroupName"`
	ReviewerRoomTypeName        string `gorm:"column:reviewer_room_type_name" json:"reviewerRoomTypeName"`
	ReviewerCountryID           int    `gorm:"column:reviewer_country_id" json:"reviewerCountryId"`
	ReviewerLengthOfStay        int    `gorm:"column:reviewer_length_of_stay" json:"reviewerLengthOfStay"`
	ReviewerReviewGroupID       int    `gorm:"column:reviewer_review_group_id" json:"reviewerReviewGroupId"`
	ReviewerRoomTypeID          int    `gorm:"column:reviewer_room_type_id" json:"reviewerRoomTypeId"`
	ReviewerReviewedCount       int    `gorm:"column:reviewer_reviewed_count" json:"reviewerReviewedCount"`
	ReviewerIsExpertReviewer    bool   `gorm:"column:reviewer_is_expert_reviewer" json:"reviewerIsExpertReviewer"`
	ReviewerIsShowGlobalIcon    bool   `gorm:"column:reviewer_is_show_global_icon" json:"reviewerIsShowGlobalIcon"`
	ReviewerIsShowReviewedCount bool   `gorm:"column:reviewer_is_show_reviewed_count" json:"reviewerIsShowReviewedCount"`

	// Overall
	OverallScore            float64 `gorm:"column:overall_score" json:"overallScore"`
	OverallReviewCount      int     `gorm:"column:overall_review_count" json:"overallReviewCount"`
	GradeCleanliness        float64 `gorm:"column:grade_cleanliness" json:"overallGradeCleanliness"`
	GradeFacilities         float64 `gorm:"column:grade_facilities" json:"overallGradeFacilities"`
	GradeLocation           float64 `gorm:"column:grade_location" json:"overallGradeLocation"`
	GradeRoomComfortQuality float64 `gorm:"column:grade_room_comfort_quality" json:"overallGradeRoomComfortQuality"`
	GradeService            float64 `gorm:"column:grade_service" json:"overallGradeService"`
	GradeValueForMoney      float64 `gorm:"column:grade_value_for_money" json:"overallGradeValueForMoney"`
}

func (h *HotelReview) UnmarshalJSON(data []byte) error {
	var raw struct {
		HotelID   int    `json:"hotelId"`
		Platform  string `json:"platform"`
		HotelName string `json:"hotelName"`
		Comment   struct {
			IsShowReviewResponse    bool    `json:"isShowReviewResponse"`
			HotelReviewID           int     `json:"hotelReviewId"`
			ProviderID              int     `json:"providerId"`
			Rating                  float64 `json:"rating"`
			CheckInDateMonthAndYear string  `json:"checkInDateMonthAndYear"`
			EncryptedReviewData     string  `json:"encryptedReviewData"`
			FormattedRating         string  `json:"formattedRating"`
			FormattedReviewDate     string  `json:"formattedReviewDate"`
			RatingText              string  `json:"ratingText"`
			ResponderName           string  `json:"responderName"`
			ResponseDateText        string  `json:"responseDateText"`
			ResponseTranslateSource string  `json:"responseTranslateSource"`
			ReviewComments          string  `json:"reviewComments"`
			ReviewNegatives         string  `json:"reviewNegatives"`
			ReviewPositives         string  `json:"reviewPositives"`
			ReviewProviderLogo      string  `json:"reviewProviderLogo"`
			ReviewProviderText      string  `json:"reviewProviderText"`
			ReviewTitle             string  `json:"reviewTitle"`
			TranslateSource         string  `json:"translateSource"`
			TranslateTarget         string  `json:"translateTarget"`
			ReviewDate              string  `json:"reviewDate"`
			ReviewerInfo            struct {
				CountryName           string `json:"countryName"`
				DisplayMemberName     string `json:"displayMemberName"`
				FlagName              string `json:"flagName"`
				ReviewGroupName       string `json:"reviewGroupName"`
				RoomTypeName          string `json:"roomTypeName"`
				CountryID             int    `json:"countryId"`
				LengthOfStay          int    `json:"lengthOfStay"`
				ReviewGroupID         int    `json:"reviewGroupId"`
				RoomTypeID            int    `json:"roomTypeId"`
				ReviewerReviewedCount int    `json:"reviewerReviewedCount"`
				IsExpertReviewer      bool   `json:"isExpertReviewer"`
				IsShowGlobalIcon      bool   `json:"isShowGlobalIcon"`
				IsShowReviewedCount   bool   `json:"isShowReviewedCount"`
			} `json:"reviewerInfo"`
			OriginalTitle         string `json:"originalTitle"`
			OriginalComment       string `json:"originalComment"`
			FormattedResponseDate string `json:"formattedResponseDate"`
		} `json:"comment"`
		OverallByProviders []struct {
			ProviderID   int     `json:"providerId"`
			Provider     string  `json:"provider"`
			OverallScore float64 `json:"overallScore"`
			ReviewCount  int     `json:"reviewCount"`
			Grades       struct {
				Cleanliness           float64 `json:"Cleanliness"`
				Facilities            float64 `json:"Facilities"`
				Location              float64 `json:"Location"`
				RoomComfortAndQuality float64 `json:"Room comfort and quality"`
				Service               float64 `json:"Service"`
				ValueForMoney         float64 `json:"Value for money"`
			} `json:"grades"`
		} `json:"overallByProviders"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Assign top-level
	h.HotelID = raw.HotelID
	h.Platform = raw.Platform
	h.HotelName = raw.HotelName

	// Comment
	c := raw.Comment
	h.IsShowReviewResponse = c.IsShowReviewResponse
	h.HotelReviewID = c.HotelReviewID
	h.ProviderID = c.ProviderID
	h.Rating = c.Rating
	h.CheckInDateMonthAndYear = c.CheckInDateMonthAndYear
	h.EncryptedReviewData = c.EncryptedReviewData
	h.FormattedRating = c.FormattedRating
	h.FormattedReviewDate = c.FormattedReviewDate
	h.RatingText = c.RatingText
	h.ResponderName = c.ResponderName
	h.ResponseDateText = c.ResponseDateText
	h.ResponseTranslateSource = c.ResponseTranslateSource
	h.ReviewComments = c.ReviewComments
	h.ReviewNegatives = c.ReviewNegatives
	h.ReviewPositives = c.ReviewPositives
	h.ReviewProviderLogo = c.ReviewProviderLogo
	h.ReviewProviderText = c.ReviewProviderText
	h.ReviewTitle = c.ReviewTitle
	h.TranslateSource = c.TranslateSource
	h.TranslateTarget = c.TranslateTarget
	h.OriginalTitle = c.OriginalTitle
	h.OriginalComment = c.OriginalComment
	h.FormattedResponseDate = c.FormattedResponseDate

	if parsedDate, err := time.Parse(time.RFC3339, c.ReviewDate); err == nil {
		h.ReviewDate = parsedDate
	}

	ri := c.ReviewerInfo
	h.ReviewerCountryName = ri.CountryName
	h.ReviewerDisplayMemberName = ri.DisplayMemberName
	h.ReviewerFlagName = ri.FlagName
	h.ReviewerReviewGroupName = ri.ReviewGroupName
	h.ReviewerRoomTypeName = ri.RoomTypeName
	h.ReviewerCountryID = ri.CountryID
	h.ReviewerLengthOfStay = ri.LengthOfStay
	h.ReviewerReviewGroupID = ri.ReviewGroupID
	h.ReviewerRoomTypeID = ri.RoomTypeID
	h.ReviewerReviewedCount = ri.ReviewerReviewedCount
	h.ReviewerIsExpertReviewer = ri.IsExpertReviewer
	h.ReviewerIsShowGlobalIcon = ri.IsShowGlobalIcon
	h.ReviewerIsShowReviewedCount = ri.IsShowReviewedCount

	// First element of overallByProviders
	if len(raw.OverallByProviders) > 0 {
		op := raw.OverallByProviders[0]
		h.OverallScore = op.OverallScore
		h.OverallReviewCount = op.ReviewCount
		h.GradeCleanliness = op.Grades.Cleanliness
		h.GradeFacilities = op.Grades.Facilities
		h.GradeLocation = op.Grades.Location
		h.GradeRoomComfortQuality = op.Grades.RoomComfortAndQuality
		h.GradeService = op.Grades.Service
		h.GradeValueForMoney = op.Grades.ValueForMoney
	}

	return nil
}

func (h *HotelReview) Validate() error {
	switch {
	case h.HotelReviewID == 0:
		return fmt.Errorf("hotelReviewId is required and cannot be zero")
	case h.ProviderID == 0:
		return fmt.Errorf("providerId is required and cannot be zero")
	case h.HotelID == 0:
		return fmt.Errorf("hotelId is required and cannot be zero")
	case h.FormattedRating == "":
		return fmt.Errorf("formattedRating is required and cannot be empty")
	}

	return nil
}

func (h *HotelReview) CheckDuplicate(dbConn *gorm.DB) (bool, error) {
	var count int64
	err := dbConn.Model(&HotelReview{}).
		Where("hotel_review_id = ? AND provider_id = ?", h.HotelReviewID, h.ProviderID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
