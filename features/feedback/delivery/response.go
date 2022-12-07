package delivery

import (
	"be13/project/features/feedback"
)

type FeedbackResponse struct {
	ID       uint   `json:"feedback_id"`
	UserID   uint   `json:"user_id"`
	Status   string `json:"status"`
	MenteeID uint   `json:"mentee_id"`
	Comment  string `json:"comment"`
}

func coreToResponse(core feedback.Core) FeedbackResponse {
	response := FeedbackResponse{
		ID:       core.ID,
		UserID:   core.UserID,
		Status:   core.Status,
		MenteeID: core.MenteeID,
		Comment:  core.Comment,
	}
	return response

}

func responseList(listRes []feedback.Core) []FeedbackResponse {
	var feedList []FeedbackResponse
	for _, v := range listRes {
		feedList = append(feedList, coreToResponse(v))

	}
	return feedList

}
